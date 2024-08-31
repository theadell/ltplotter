using PlotService.Protobuf;
namespace PlotService.Services;

public interface ILatexService
{
    string GenerateLatex(PlotRequest plotRequest);
    byte[] CompileLatex(string latex);
}
