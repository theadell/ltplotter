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
            return request
                .Pipe(latexService.GenerateLatex)
                .Pipe(Log.Information)
                .Pipe(latexService.CompileLatex)
                .Pipe(value => new PlotResponse { Pdf = value.Item1, Latex = value.Item2 })
                .Pipe(Task.FromResult);
        }
    }
}