package com.veterinary.security.rolemanager.controller;

import com.veterinary.security.rolemanager.model.UserRole;
import com.veterinary.security.rolemanager.service.UserRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/user-roles")
public class UserRoleController {

    @Autowired
    private UserRoleService userRoleService;

    @PostMapping("/assign")
    public ResponseEntity<UserRole> assignRole(@RequestBody AssignRoleRequest request) {
        UserRole userRole = userRoleService.assignRole(request.getUserId(), request.getRoleId());
        return ResponseEntity.ok(userRole);
    }

    // Clase interna estática para recibir el JSON del POST
    public static class AssignRoleRequest {
        private Long userId;
        private Long roleId;

        public Long getUserId() {
            return userId;
        }

        public void setUserId(Long userId) {
            this.userId = userId;
        }

        public Long getRoleId() {
            return roleId;
        }

        public void setRoleId(Long roleId) {
            this.roleId = roleId;
        }
    }
}
