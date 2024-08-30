using System.Security.Cryptography.X509Certificates;
using PlotService.Protobuf;
using Serilog;
using PlotService.Services;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddGrpc();
builder.Services.AddGrpcReflection();

builder.Services.AddScoped<ILatexService, LatexService>();

builder.WebHost.ConfigureKestrel(options =>
{
    var certificate = X509Certificate2.CreateFromPem(
        File.ReadAllText("/app/certs/data_plot_server.crt"),
        File.ReadAllText("/app/certs/data_plot_server.key")); 
    
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

