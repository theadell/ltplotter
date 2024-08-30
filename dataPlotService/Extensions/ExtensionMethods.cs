namespace PlotService.Extensions;

public static class ExtensionMethods
{
    public static TResult Pipe<TInput, TResult>(this TInput input, Func<TInput, TResult> func)
    {
        return func(input);
    }

    public static T Pipe<T>(this T input, Action<T> action)
    {
        action(input);
        return input;
    }
}