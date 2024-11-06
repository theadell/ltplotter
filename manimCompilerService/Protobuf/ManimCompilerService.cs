using Grpc.Core;
using ManimCompilerService.Services;

namespace ManimCompilerService.Protobuf;

public class ManimCompilerServiceGrpc(IManimService manimService) : ManimCompilerService.ManimCompilerServiceBase
{
    public override async Task<ManimResponse> GenerateVideo(ManimRequest request, ServerCallContext context)
    {
        await Task.Yield();
        manimService.CreateVideo(request.PythonSource);
        return new ManimResponse()
        {
            ManimVideoUrl = "SUCCESS"
        };
    }
}