using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;
using Azure.Storage.Sas;

namespace ManimCompilerService.Services;

public class UploadService : IUploadService
{
    public async Task<string> UploadBlobAsync(string localFilePath)
    {
        const string storageAccountName = "ltplotter";
        const string containerName = "ltplotter";

        var blobServiceClient = new BlobServiceClient(
            new Uri($"https://{storageAccountName}.blob.core.windows.net"),
            new DefaultAzureCredential());

        var containerClient = blobServiceClient.GetBlobContainerClient(containerName);

        await containerClient.CreateIfNotExistsAsync(PublicAccessType.None);

        var blobName = Path.GetFileName(localFilePath);

        var blobClient = containerClient.GetBlobClient(blobName);

        await blobClient.UploadAsync(localFilePath, true);
         
        var sasBuilder = new BlobSasBuilder
        {
            BlobContainerName = containerName,
            BlobName = blobName,
            Resource = "ltplotterrg",
            StartsOn = DateTimeOffset.UtcNow,
            ExpiresOn = DateTimeOffset.UtcNow.AddHours(1)
        };
        sasBuilder.SetPermissions(BlobSasPermissions.Read);

        var sasUri = blobClient.GenerateSasUri(sasBuilder);

        return sasUri.AbsoluteUri;
    }
}