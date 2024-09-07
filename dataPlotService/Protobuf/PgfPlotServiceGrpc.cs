using Google.Protobuf;
using Grpc.Core;
using PlotService.Services;
using PlotService.Extensions;

namespace PlotService.Protobuf
{
    public class PgfPlotServiceGrpc(ILatexService latexService) : PlotService.PlotServiceBase
    {
        public override Task<PlotResponse> GeneratePlot(PlotRequest request, ServerCallContext context)
        {
            return request
                .Pipe(latexService.GenerateLatex)
                .Pipe(Console.WriteLine)
                .Pipe(latexService.CompileLatex)
                .Pipe(ByteString.CopyFrom)
                .Pipe(pdf => new PlotResponse { Pdf = pdf })
                .Pipe(Task.FromResult);
        }
    }
}