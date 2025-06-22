package com.veterinary.security.rolemanager.service;

import com.veterinary.security.rolemanager.model.Role;
import com.veterinary.security.rolemanager.model.UserRole;
import com.veterinary.security.rolemanager.repository.RoleRepository;
import com.veterinary.security.rolemanager.repository.UserRoleRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserRoleService {

    @Autowired
    private UserRoleRepository userRoleRepository;

    @Autowired
    private RoleRepository roleRepository;

    public UserRole assignRole(Long userId, Long roleId) {
        Role role = roleRepository.findById(roleId)
                .orElseThrow(() -> new RuntimeException("Role not found with id " + roleId));

        UserRole userRole = new UserRole();
        userRole.setUserId(userId);
        userRole.setRole(role);

        return userRoleRepository.save(userRole);
    }
}

