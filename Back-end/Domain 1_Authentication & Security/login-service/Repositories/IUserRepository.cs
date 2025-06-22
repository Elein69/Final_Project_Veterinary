using System.Threading.Tasks;
using login_service.Models;

namespace login_service.Repositories
{
    public interface IUserRepository
    {
        Task<User?> GetByEmailAsync(string email);
    }
}
