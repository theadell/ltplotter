namespace ManimCompilerService.Services;

public interface IUploadService
{

    Task<string> UploadBlobAsync(string localFilePath);
}