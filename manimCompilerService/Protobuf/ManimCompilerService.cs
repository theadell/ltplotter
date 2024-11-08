using Grpc.Core;
using ManimCompilerService.Services;

namespace ManimCompilerService.Protobuf;

public class ManimCompilerServiceGrpc(IManimService manimService, IUploadService uploadService) : ManimCompilerService.ManimCompilerServiceBase
{
    public override async Task<ManimResponse> GenerateVideo(ManimRequest request, ServerCallContext context)
    {
        var path = manimService.CreateVideo(request.PythonSource);
        var url = await uploadService.UploadBlobAsync(path); 
        Directory.Delete(path, true);
        return new ManimResponse()
        {
            ManimVideoUrl = url
        };
    }
}