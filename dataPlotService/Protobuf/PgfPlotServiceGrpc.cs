using Google.Protobuf;
using Grpc.Core;
using PlotService.Services;
using PlotService.Extensions;
using Serilog;

namespace PlotService.Protobuf
{
    public class PgfPlotServiceGrpc(ILatexService latexService) : PlotService.PlotServiceBase
    {
        public override Task<PlotResponse> GeneratePlot(PlotRequest request, ServerCallContext context)
        {
            var latex = string.Empty;
            return request
                .Pipe(latexService.GenerateLatex)
                .Pipe(Log.Information)
                .Pipe(l =>  latex = l)
                .Pipe(latexService.CompileLatex)
                .Pipe(ByteString.CopyFrom)
                .Pipe(pdf => new PlotResponse { Pdf = pdf, Latex = latex })
                .Pipe(Task.FromResult);
        }
    }
}