using Google.Protobuf;
using Grpc.Core;
using PlotService.Services;

namespace Plot
{
    public class PgfPlotServiceGrpc : PlotService.PlotServiceBase
    {
        private readonly ILatexService _latexService;

        public PgfPlotServiceGrpc(ILatexService latexService)
        {
            _latexService = latexService;
        }

        public override Task<PlotResponse> GeneratePlot(PlotRequest request, ServerCallContext context)
        {
            var response = new PlotResponse();
            var latexString = _latexService.GenerateLatex(request);
            var pdf = _latexService.CompileLatex(latexString);
            response.Pdf = ByteString.CopyFrom(pdf);
            return Task.FromResult(response);
        }
    }
}