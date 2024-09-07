using System.Diagnostics;
using PlotService.Extensions;
using Serilog;

namespace PlotService.Utils;

public static class ProcessUtils
{
   private static ProcessStartInfo GetProcessStartInfo(string process, string args)
   {
         return new ProcessStartInfo
         {
             FileName = process,
             Arguments = args,
             RedirectStandardOutput = true,
             RedirectStandardError = true,
             UseShellExecute = false,
             CreateNoWindow = true
         };
   }
   
   public static void RunProcess(string process, string args)
   {
         GetProcessStartInfo(process, args)
            .Pipe(Process.Start)?
            .WaitForExit();
   }

   public static string GetProcessOutput(string process, string args)
   {
      try
      {
         return GetProcessStartInfo(process, args)
            .Pipe(Process.Start)?
            .StandardOutput.ReadToEnd() ?? string.Empty;
      }
      catch (Exception ex)
      {
         Log.Error("Running the process failed with @{ex}", ex);
         return string.Empty;
      }
   }
}