using ManimCompilerService.Extensions;
using ManimCompilerService.Utils;
using Serilog;

namespace ManimCompilerService.Services;

public class ManimService : IManimService
{
    public string CreateVideo(string pythonSource)
    {
        try
        {
            var tempDir = Path.Combine(Path.GetTempPath(), Guid.NewGuid().ToString());
            Directory.CreateDirectory(tempDir);

            var pythonFilePath = Path.Combine(tempDir, "manimScene.py");
            File.WriteAllText(pythonFilePath, pythonSource);

            var mediaDir = Path.Combine(tempDir, "media");
            var arguments = $"-ql --media_dir {mediaDir} {pythonFilePath}".Pipe(Log.Information);
            
            ProcessUtils.RunProcess("manim", arguments);
            
            var file = Directory.GetFiles(mediaDir, "*.mp4", SearchOption.AllDirectories).FirstOrDefault().Pipe(Log.Information);
            
            return tempDir;
        }
        catch (Exception ex)
        {
            Log.Error("Exception while compiling manim: {@ex}", ex);
            return null;
        }
    }
 
}