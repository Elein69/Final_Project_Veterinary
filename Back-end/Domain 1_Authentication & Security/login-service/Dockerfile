
FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app
EXPOSE 5046


FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["login-service.csproj", "./"]
RUN dotnet restore "./login-service.csproj"
COPY . .
RUN dotnet build "login-service.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "login-service.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "login-service.dll"]
