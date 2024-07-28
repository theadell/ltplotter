using System.Diagnostics;
using Google.Protobuf.Collections;
using Plot;
using PlotService.Models;
using Serilog;

namespace PlotService.Services;

public class LatexService : ILatexService
{
    public string GenerateLatex(PlotRequest plotRequest)
    {
        try
        {
            var sb = new System.Text.StringBuilder();
            sb.AppendLine(@"\documentclass{standalone}");
            sb.AppendLine(@"\usepackage{pgfplots}");
            sb.AppendLine(@"\begin{document}");
            sb.AppendLine(@"\begin{tikzpicture}");
            sb.AppendLine(@$"\begin{{axis}}[title={{{plotRequest.Metadata.Title}}}, xlabel={{{plotRequest.Metadata.Labels.X}}}, ylabel={{{plotRequest.Metadata.Labels.Y}}}, legend pos=outer north east]");
            
            for (var i = 0; i < plotRequest.Y.Count; i++)
            {
                sb.Append(@"\addplot coordinates {");
                sb.Append(string.Join(' ', plotRequest.X.Zip(plotRequest.Y[i].Values).Select(pair => $"({pair.Item1}, {pair.Item2})")));
                sb.AppendLine("};");
                sb.AppendLine(@$"\addlegendentry{{{plotRequest.Metadata.Legends[i]}}}");
            }

            sb.AppendLine(@"\end{axis}");
            sb.AppendLine(@"\end{tikzpicture}");
            sb.AppendLine(@"\end{document}");
            return sb.ToString();
        }
        catch (Exception ex)
        {
            Log.Error("Exception while generating latex string: {@ex}", ex);
            return string.Empty;
        }
    }

    public string GenerateLatex(PlotRequestRest requestRest)
    {
        try
        {
            var sb = new System.Text.StringBuilder();
            sb.AppendLine(@"\documentclass{standalone}");
            sb.AppendLine(@"\usepackage{pgfplots}");
            sb.AppendLine(@"\begin{document}");
            sb.AppendLine(@"\begin{tikzpicture}");
            sb.AppendLine(@$"\begin{{axis}}[title={{{requestRest.Metadata.Title}}}, xlabel={{{requestRest.Metadata.Labels.X}}}, ylabel={{{requestRest.Metadata.Labels.Y}}}, legend pos=outer north east]");
            
            for (var i = 0; i < requestRest.Y.Count; i++)
            {
                sb.Append(@"\addplot coordinates {");
                sb.Append(string.Join(' ', requestRest.X.Zip(requestRest.Y[i]).Select(pair => $"({pair.Item1}, {pair.Item2})")));
                sb.AppendLine("};");
                sb.AppendLine(@$"\addlegendentry{{{requestRest.Metadata.Legends[i]}}}");
            }

            sb.AppendLine(@"\end{axis}");
            sb.AppendLine(@"\end{tikzpicture}");
            sb.AppendLine(@"\end{document}");
            return sb.ToString();
        }
        catch (Exception ex)
        {
            Log.Error("Exception while generating latex string: {@ex}", ex);
            return string.Empty;
        }
    }

    public byte[] CompileLatex(string latex)
    {
        try
        {
            var tempDir = Path.Combine(Path.GetTempPath(), Guid.NewGuid().ToString());
            Directory.CreateDirectory(tempDir);

            var texFile = Path.Combine(tempDir, "plot.tex");
            File.WriteAllText(texFile, latex);

            var psi = new ProcessStartInfo
            {
                FileName = "/usr/bin/pdflatex",
                Arguments = $"-output-directory {tempDir} {texFile}",
                RedirectStandardOutput = true,
                RedirectStandardError = true,
                UseShellExecute = false,
                CreateNoWindow = true
            };

            using (var process = Process.Start(psi))
            {
                process?.WaitForExit();
            }

            var pdfFile = Path.Combine(tempDir, "plot.pdf");
            var pdfBytes = File.ReadAllBytes(pdfFile);

            Directory.Delete(tempDir, true);

            return pdfBytes;
        }
        catch (Exception ex)
        {
            Log.Error("Exception while compiling latex: {@ex}", ex);
            return [];
        }
    }
}