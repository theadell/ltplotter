using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using System.Threading.Tasks;
using Serilog;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddSingleton<PlotService.Services.PlotService>();

var app = builder.Build();

Log.Logger = new LoggerConfiguration()
    .WriteTo.Console()
    .CreateLogger();

app.MapPost("/generate-plot", async (HttpContext context, PlotRequest request, PlotService.Services.PlotService plotService) =>
{
    Log.Information("Received PlotRequest {@plotRequest}", request);
    var latexString = plotService.GenerateLatex(request);
    var pdf = plotService.CompileLatex(latexString);
    context.Response.ContentType = "application/pdf";
    await context.Response.Body.WriteAsync(pdf);
});

app.Run();

