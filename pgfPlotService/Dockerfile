FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /source
COPY . .

RUN rm -rf ./Generated


RUN dotnet restore "./TimeSeriesPlotter.csproj" --disable-parallel
RUN dotnet publish "./TimeSeriesPlotter.csproj" -c release -o /app --no-restore

FROM mcr.microsoft.com/dotnet/sdk:8.0 
WORKDIR /app
COPY --from=build /app ./

EXPOSE 8080 5001

RUN apt-get update && apt-get install -y texlive-latex-base
    
    
ENTRYPOINT ["dotnet", "TimeSeriesPlotter.dll"]
