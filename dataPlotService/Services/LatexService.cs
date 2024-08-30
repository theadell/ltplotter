using System.Diagnostics;
using System.Text.RegularExpressions;
using PlotService.Extensions;
using PlotService.Protobuf;
using Serilog;

namespace PlotService.Services;

public partial class LatexService : ILatexService
{

    [GeneratedRegex(@"([\\&%$#_{}~^<>])", RegexOptions.Compiled)]
    private static partial Regex EscapeRegex();
    
    private static readonly Regex EscapePattern = EscapeRegex();

    private static readonly Dictionary<string, string> ConversionMap = new()
    {
        { @"\", @"\textbackslash{}" },
        { "&",  @"\&" },
        { "%",  @"\%" },
        { "$",  @"\$" },
        { "#",  @"\#" },
        { "_",  @"\_" },
        { "{",  @"\{" },
        { "}",  @"\}" },
        { "~",  @"\textasciitilde{}" },
        { "^",  @"\^{} " },
        { "<",  @"\textless{}" },
        { ">",  @"\textgreater{}" }
    };
    
    public string GenerateLatex(PlotRequest plotRequest)
    {
        try
        {
            var sb = new System.Text.StringBuilder();
            sb.AppendLine(@"\documentclass{standalone}");
            sb.AppendLine(@"\usepackage{pgfplots}");
            sb.AppendLine(@"\begin{document}");
            sb.AppendLine(@"\begin{tikzpicture}");
            sb.AppendLine(@$"\begin{{axis}}[title={{{Escape(plotRequest.Metadata.Title)}}}, xlabel={{{Escape(plotRequest.Metadata.Labels.X)}}}, ylabel={{{Escape(plotRequest.Metadata.Labels.Y)}}}, legend pos=outer north east]");
            
            for (var i = 0; i < plotRequest.Y.Count; i++)
            {
                sb.Append(@"\addplot coordinates {");
                sb.Append(string.Join(' ', plotRequest.X.Zip(plotRequest.Y[i].Values).Select(pair => $"({pair.Item1}, {pair.Item2})")));
                sb.AppendLine("};");
                sb.AppendLine(@$"\addlegendentry{{{Escape(plotRequest.Metadata.Legends[i])}}}");
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

            var pdfBytes = Path.Combine(tempDir, "plot.pdf")
                .Pipe(File.ReadAllBytes);

            Directory.Delete(tempDir, true);

            return pdfBytes;
        }
        catch (Exception ex)
        {
            Log.Error("Exception while compiling latex: {@ex}", ex);
            return [];
        }
    }
    
    private static string Escape(string s)
    {
        return EscapePattern.Replace(s, match => ConversionMap[match.Value]);
    }
}
