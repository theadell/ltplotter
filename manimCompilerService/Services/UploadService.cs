using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;

namespace ManimCompilerService.Services;

public class UploadService : IUploadService
{
    // TODO: Use DI
    private BlobServiceClient GetBlobServiceClient(string accountName)
    {
        BlobServiceClient client = new(
            new Uri($"https://{accountName}.blob.core.windows.net"),
            new DefaultAzureCredential());

        return client;
    } 
    
    public async Task UploadFromFileAsync(
        string localFilePath)
    {
        var blobServiceClient = GetBlobServiceClient("ltplotter");
        var response = await blobServiceClient.CreateBlobContainerAsync("container");
        var blobClient = response.Value.GetBlobClient("test");

        var fileName = Path.GetFileName(localFilePath);

        await blobClient.UploadAsync(localFilePath, true);
    }
}