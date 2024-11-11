using Azure.Identity;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Blobs.Specialized;
using Azure.Storage.Sas;

namespace ManimCompilerService.Services;

public class UploadService : IUploadService
{
    public async Task<Uri> UploadBlobAsync(string localFilePath)
    {
        const string storageAccountName = "ltplotter";
        const string containerName = "ltplotter";

        var blobServiceClient = new BlobServiceClient(
            new Uri($"https://{storageAccountName}.blob.core.windows.net"),
            new DefaultAzureCredential());

        var containerClient = blobServiceClient.GetBlobContainerClient(containerName);
        await containerClient.CreateIfNotExistsAsync();
        var blobName = Path.GetFileName(localFilePath);
        var blobClient = containerClient.GetBlobClient(blobName);


        await blobClient.UploadAsync(localFilePath, true);
        var userDelegationKey = await RequestUserDelegationKeyAsync(blobServiceClient);
        var uri = CreateUserDelegationSasUri(blobClient, userDelegationKey);

        return uri;
    }

    private async Task<UserDelegationKey> RequestUserDelegationKeyAsync(
        BlobServiceClient blobServiceClient)
    {
        UserDelegationKey userDelegationKey =
            await blobServiceClient.GetUserDelegationKeyAsync(
                DateTimeOffset.UtcNow,
                DateTimeOffset.UtcNow.AddDays(1));

        return userDelegationKey;
    }

    private Uri CreateUserDelegationSasUri(
        BlobBaseClient blobClient,
        UserDelegationKey userDelegationKey)
    {
        var sasBuilder = new BlobSasBuilder()
        {
            BlobContainerName = blobClient.BlobContainerName,
            BlobName = blobClient.Name,
            Resource = "b",
            StartsOn = DateTimeOffset.UtcNow,
            ExpiresOn = DateTimeOffset.UtcNow.AddDays(1)
        };

        sasBuilder.SetPermissions(BlobSasPermissions.Read);

        var uriBuilder = new BlobUriBuilder(blobClient.Uri)
        {
            Sas = sasBuilder.ToSasQueryParameters(
                userDelegationKey,
                blobClient
                    .GetParentBlobContainerClient()
                    .GetParentBlobServiceClient().AccountName)
        };

        return uriBuilder.ToUri();
    }
}