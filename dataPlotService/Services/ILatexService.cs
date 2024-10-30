using Google.Protobuf;
using PlotService.Protobuf;
namespace PlotService.Services;

public interface ILatexService
{
    string GenerateLatex(PlotRequest plotRequest);
    (ByteString, string) CompileLatex(string latex);
}
