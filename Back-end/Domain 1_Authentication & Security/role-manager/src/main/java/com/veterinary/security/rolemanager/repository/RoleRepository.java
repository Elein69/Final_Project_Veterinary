package com.veterinary.security.rolemanager.repository;

import com.veterinary.security.rolemanager.model.Role;
import org.springframework.data.jpa.repository.JpaRepository;

public interface RoleRepository extends JpaRepository<Role, Long> {
}