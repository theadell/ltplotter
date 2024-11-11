using Grpc.Core;
using ManimCompilerService.Services;

namespace ManimCompilerService.Protobuf;

public class ManimCompilerServiceGrpc(IManimService manimService, IUploadService uploadService)
    : ManimCompilerService.ManimCompilerServiceBase
{
    public override async Task<ManimResponse> GenerateVideo(ManimRequest request, ServerCallContext context)
    {
        var path = manimService.CreateVideo(request.PythonSource);
        var mediaDir = Path.Combine(path, "media");
        var file = Directory.GetFiles(mediaDir, "*.mp4", SearchOption.AllDirectories).FirstOrDefault();
        if (file != null)
        {
            var url = await uploadService.UploadBlobAsync(file);
            Directory.Delete(path, true);
            return new ManimResponse
            {
                ManimVideoUrl = url.ToString()
            };
        }

        return new ManimResponse
        {
            ManimVideoUrl = string.Empty
        };
    }
}