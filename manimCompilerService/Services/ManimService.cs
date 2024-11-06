using ManimCompilerService.Extensions;
using ManimCompilerService.Utils;
using Serilog;

namespace ManimCompilerService.Services;

public class ManimService : IManimService
{
    public bool CreateVideo(string pythonSource)
    {
        try
        {
            var tempDir = Path.Combine(Path.GetTempPath(), Guid.NewGuid().ToString());
            Directory.CreateDirectory(tempDir);

            var pythonFilePath = Path.Combine(tempDir, "manimScene.py");
            File.WriteAllText(pythonFilePath, pythonSource);

            var arguments = $"-ql --media_dir {Path.Combine(tempDir, "media")} {pythonFilePath}".Pipe(Log.Information);
            
            ProcessUtils.RunProcess("manim", arguments);
            
            // TODO: Return location of file and implement method

            Directory.Delete(tempDir, true);

            return true;
        }
        catch (Exception ex)
        {
            Log.Error("Exception while compiling latex: {@ex}", ex);
            return false;
        }
    }
 
}