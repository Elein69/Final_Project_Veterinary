package com.veterinary.user_registry.repository;

import com.veterinary.user_registry.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Long> {
}

