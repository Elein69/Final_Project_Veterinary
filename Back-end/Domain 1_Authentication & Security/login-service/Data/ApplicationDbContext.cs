using login_service.Models;
using Microsoft.EntityFrameworkCore;

namespace login_service.Data
{
    public class ApplicationDbContext : DbContext
    {
        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
            : base(options)
        {
        }

        public DbSet<User> Users { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            
            modelBuilder.Entity<User>().ToTable("users");

            base.OnModelCreating(modelBuilder);
        }
    }
}
