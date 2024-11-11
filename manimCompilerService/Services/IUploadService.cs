namespace ManimCompilerService.Services;

public interface IUploadService
{

    Task<Uri> UploadBlobAsync(string localFilePath);
}