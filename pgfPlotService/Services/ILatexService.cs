using PlotService.Models;
using Plot;

namespace PlotService.Services;

public interface ILatexService
{
    string GenerateLatex(PlotRequest plotRequest);
    string GenerateLatex(PlotRequestRest requestRest);
    byte[] CompileLatex(string latex);
}
