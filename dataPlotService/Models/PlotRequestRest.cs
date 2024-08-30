namespace PlotService.Models;

public class PlotRequestRest
{
    public List<double> X { get; set; }
    public List<List<double>> Y { get; set; }
    public Metadata Metadata { get; set; }
}