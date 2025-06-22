using System.Threading.Tasks;
using login_service.DTOs;

namespace login_service.Services
{
    public interface ILoginService
    {
        Task<LoginResponse?> LoginAsync(LoginRequest request);
    }
}
