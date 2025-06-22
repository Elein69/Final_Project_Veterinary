using System;
using System.Threading.Tasks;
using login_service.DTOs;
using login_service.Models;
using login_service.Repositories;
using login_service.Services;
using Moq;
using Xunit;

namespace login_service.Tests
{
    public class LoginServiceTests
    {
        [Fact]
        public async Task LoginAsync_ReturnsToken_WhenCredentialsAreValid()
        {
            // Arrange
            var user = new User
            {
                Id = Guid.NewGuid(),
                Email = "test@example.com",
                Password = "123456",
                Role = "cliente",
                Name = "Test User",
                CreatedAt = DateTime.UtcNow
            };

            var mockRepo = new Mock<IUserRepository>();
            mockRepo.Setup(repo => repo.GetByEmailAsync(user.Email)).ReturnsAsync(user);

            var loginService = new LoginService(mockRepo.Object);

            var request = new LoginRequest
            {
                Email = user.Email,
                Password = "123456"
            };

            // Act
            var response = await loginService.LoginAsync(request);

            // Assert
            Assert.NotNull(response);
            Assert.Equal(user.Email, response.Email);
            Assert.Equal(user.Role, response.Role);
            Assert.False(string.IsNullOrEmpty(response.Token));
        }
    }
}
