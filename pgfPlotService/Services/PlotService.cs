using System.Diagnostics;

namespace PlotService.Services;

public class PlotService
{
    public string GenerateLatex(PlotRequest request)
    {
        var sb = new System.Text.StringBuilder();
        sb.AppendLine(@"\documentclass{standalone}");
        sb.AppendLine(@"\usepackage{pgfplots}");
        sb.AppendLine(@"\begin{document}");
        sb.AppendLine(@"\begin{tikzpicture}");
        sb.AppendLine(@$"\begin{{axis}}[title={{{request.Metadata.Title}}}, xlabel={{{request.Metadata.Labels.X}}}, ylabel={{{request.Metadata.Labels.Y}}}, legend pos=outer north east]");
        for (var i = 0; i < request.Y.Count; i++)
        {
            sb.Append(@"\addplot coordinates {");
            sb.Append(string.Join(' ', request.X.Zip(request.Y[i]).Select(pair => $"({pair.Item1}, {pair.Item2})")));
            sb.AppendLine("};");
            sb.AppendLine(@$"\addlegendentry{{{request.Metadata.Legends[i]}}}");
        }
        sb.AppendLine(@"\end{axis}");
        sb.AppendLine(@"\end{tikzpicture}");
        sb.AppendLine(@"\end{document}");
        return sb.ToString();
    }

    public byte[] CompileLatex(string latex)
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
}