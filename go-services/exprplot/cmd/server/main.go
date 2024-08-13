package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"ltplotter/common/middleware"
	"ltplotter/common/utils"
	"ltplotter/exprplot/internal/config"
	"ltplotter/gen/pb"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedExpressionPlotServiceServer
	tmpl *template.Template
}

func newServer() (*server, error) {
	funcMap := template.FuncMap{
		"axisLinesToLatex": axisLinesToLatex,
	}

	tmpl, err := template.New("expr_plot_template.go.tex").Funcs(funcMap).ParseFiles("expr_plot_template.go.tex")
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %v", err)
	}

	return &server{
		tmpl: tmpl,
	}, nil
}

func (s *server) GeneratePlot(ctx context.Context, req *pb.ExprPlotRequest) (*pb.ExprPlotResponse, error) {
	requestID := uuid.New().String()
	timestamp := time.Now().Format("20060102_150405")
	baseFilename := fmt.Sprintf("plot_%s_%s", timestamp, requestID)

	latexFileName := fmt.Sprintf("%s.tex", baseFilename)
	pdfFileName := fmt.Sprintf("%s.pdf", baseFilename)

	latexFilePath := filepath.Join(os.TempDir(), latexFileName)
	latexFile, err := os.Create(latexFilePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create LaTeX file: %v", err)
	}
	defer latexFile.Close()

	err = s.tmpl.Execute(latexFile, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to execute LaTeX template: %v", err)
	}

	latexFileContent, err := os.ReadFile(latexFilePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to read LaTeX file: %v", err)
	}

	pdfFilePath := filepath.Join(os.TempDir(), pdfFileName)
	cmd := exec.Command(
		"pdflatex",
		"-output-directory", os.TempDir(),
		"-no-shell-escape",
		"-interaction=nonstopmode",
		"-halt-on-error",
		latexFilePath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("pdflatex failed to compile", "error", err.Error(), "output", string(output))
		return nil, status.Errorf(codes.Internal, "Failed to generate PDF: %v", err)
	}

	pdfData, err := os.ReadFile(pdfFilePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to read generated PDF: %v", err)
	}

	// Clean up temporary files
	defer func() {
		_ = os.Remove(latexFilePath)
		_ = os.Remove(pdfFilePath)
	}()

	return &pb.ExprPlotResponse{
		LatexSrcCode: string(latexFileContent),
		PdfPlot:      pdfData,
	}, nil
}

func main() {

	config := config.LoadConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv, err := newServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	creds, err := utils.LoadCerts(config.CertsPath, "expression_plot_server")
	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			middleware.UnaryServerLogger(slog.Default()),
			middleware.UnaryRecoverer(),
		),
	)

	pb.RegisterExpressionPlotServiceServer(grpcServer, srv)

	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func axisLinesToLatex(axisLines pb.AxisLines) string {
	switch axisLines {
	case pb.AxisLines_AXIS_LINES_BOX:
		return "box"
	case pb.AxisLines_AXIS_LINES_LEFT:
		return "left"
	case pb.AxisLines_AXIS_LINES_RIGHT:
		return "right"
	case pb.AxisLines_AXIS_LINES_MIDDLE:
		return "middle"
	case pb.AxisLines_AXIS_LINES_CENTER:
		return "center"
	case pb.AxisLines_AXIS_LINES_NONE:
		return "none"
	default:
		return "box"
	}
}
