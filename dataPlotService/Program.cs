using System.Security.Cryptography.X509Certificates;
using PlotService.Protobuf;
using Serilog;
using PlotService.Services;
using PlotService.Settings;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddGrpc();
builder.Services.AddGrpcReflection();

builder.Services.Configure<LatexCompilerSettings>(
    builder.Configuration.GetSection(nameof(LatexCompilerSettings)));

builder.Services.Configure<CertificateSettings>(
    builder.Configuration.GetSection(nameof(CertificateSettings)));

builder.Services.AddScoped<ILatexService, LatexService>();

builder.WebHost.ConfigureKestrel(options =>
{
    var certificateSettings = new CertificateSettings();
    builder.Configuration.GetSection(nameof(CertificateSettings)).Bind(certificateSettings);

    var certificate = X509Certificate2.CreateFromPem(
        File.ReadAllText(certificateSettings.ServerCertificatePath),
        File.ReadAllText(certificateSettings.ServerKeyPath)); 
    
    options.ListenAnyIP(5001, listenOptions =>
    {
        listenOptions.Protocols = Microsoft.AspNetCore.Server.Kestrel.Core.HttpProtocols.Http2;
        listenOptions.UseHttps(certificate);
    });
    
});

var app = builder.Build();

Log.Logger = new LoggerConfiguration()
    .WriteTo.Console()
    .CreateLogger();

app.UseRouting();

app.MapGrpcService<PgfPlotServiceGrpc>();

app.MapGrpcReflectionService();

app.Run();

