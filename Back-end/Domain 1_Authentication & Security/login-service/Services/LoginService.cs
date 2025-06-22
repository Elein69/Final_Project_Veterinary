using login_service.DTOs;
using login_service.Models;
using login_service.Repositories;
using login_service.Utils;


namespace login_service.Services
{
    public class LoginService : ILoginService
    {
        private readonly IUserRepository _userRepository;

        public LoginService(IUserRepository userRepository)
        {
            _userRepository = userRepository;
        }

        public async Task<LoginResponse?> LoginAsync(LoginRequest request)
        {
            var user = await _userRepository.GetByEmailAsync(request.Email);

            if (user == null || user.Password != request.Password)
            {
                return null;
            }

            var token = JwtHelper.GenerateToken(user);

            return new LoginResponse
            {
                Token = token,
                Email = user.Email,
                Role = user.Role
            };
        }
    }
}
