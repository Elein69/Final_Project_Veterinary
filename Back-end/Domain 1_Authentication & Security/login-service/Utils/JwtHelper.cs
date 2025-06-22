using System;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using login_service.Models;
using Microsoft.Extensions.Configuration;
using Microsoft.IdentityModel.Tokens;

namespace login_service.Utils
{
    public static class JwtHelper
    {
        public static string GenerateToken(User user)
        {
            // Carga configuración desde appsettings.json
            var config = new ConfigurationBuilder()
                .AddJsonFile("appsettings.json")
                .Build();

            var secretKey = config["JwtSettings:Secret"];

            if (string.IsNullOrEmpty(secretKey) || Encoding.UTF8.GetByteCount(secretKey) < 32)
            {
                throw new Exception("❌ JWT secret is missing or too short (must be at least 256 bits = 32 characters).");
            }

            var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(secretKey));
            var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

            var claims = new[]
            {
                new Claim(ClaimTypes.NameIdentifier, user.Id.ToString()),
                new Claim(ClaimTypes.Email, user.Email),
                new Claim(ClaimTypes.Role, user.Role)
            };

            var token = new JwtSecurityToken(
                issuer: "login-service",
                audience: "veterinary-app",
                claims: claims,
                expires: DateTime.UtcNow.AddHours(1),
                signingCredentials: creds
            );

            return new JwtSecurityTokenHandler().WriteToken(token);
        }
    }
}
