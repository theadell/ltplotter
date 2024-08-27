package main

import (
	"fmt"
	"ltplotter/gen/pb"

	"github.com/bufbuild/protovalidate-go"
)

func main() {
	msg := &pb.ExprPlotRequest{
		XLabel:    "",
		YLabel:    "",
		XMin:      0,
		XMax:      0,
		YMin:      0,
		YMax:      0,
		AxisLines: "none",
		Border:    5,
		Grid:      "major",
		Title:     "",
		Plots: []*pb.PlotExpression{
			{Expression: "sdsd", LineStyle: "solid", LineWidth: "1.5pt",
				Samples: 12, Domain: "1.1:1.1"}},
	}

	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate(msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
