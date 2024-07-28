using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using System.Threading.Tasks;
using Serilog;
using Plot;
using PlotService.Models;
using PlotService.Services;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddGrpc();

builder.Services.AddScoped<ILatexService, LatexService>();

var app = builder.Build();

Log.Logger = new LoggerConfiguration()
    .WriteTo.Console()
    .CreateLogger();

app.UseRouting();


app.UseEndpoints(endpoints =>
{
    endpoints.MapGrpcService<PgfPlotServiceGrpc>();
});

app.MapPost("/generate-plot", async (HttpContext context, PlotRequestRest request, PlotService.Services.ILatexService latexService) =>
{
    Log.Information("Received PlotRequestRest {@plotRequest}", request);
    var latexString = latexService.GenerateLatex(request);
    var pdf = latexService.CompileLatex(latexString);
    context.Response.ContentType = "application/pdf";
    await context.Response.Body.WriteAsync(pdf);
});

app.Run();

