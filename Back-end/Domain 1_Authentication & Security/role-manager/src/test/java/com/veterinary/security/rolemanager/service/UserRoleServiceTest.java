package com.veterinary.security.rolemanager.service;

import com.veterinary.security.rolemanager.model.Role;
import com.veterinary.security.rolemanager.model.UserRole;
import com.veterinary.security.rolemanager.repository.RoleRepository;
import com.veterinary.security.rolemanager.repository.UserRoleRepository;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.junit.jupiter.api.Assertions.assertEquals;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Optional;

@ExtendWith(MockitoExtension.class)
public class UserRoleServiceTest {

    @Mock
    private UserRoleRepository userRoleRepository;

    @Mock
    private RoleRepository roleRepository;

    @InjectMocks
    private UserRoleService userRoleService;

    @Test
    public void assignRole_ShouldReturnSavedUserRole() {
        // Arrange
        Long userId = 1L;
        Long roleId = 1L;
        Role role = new Role(1L, "ADMIN");

        when(roleRepository.findById(anyLong())).thenReturn(Optional.of(role));
        when(userRoleRepository.save(any(UserRole.class)))
                .thenAnswer(invocation -> invocation.getArgument(0));

        // Act
        UserRole result = userRoleService.assignRole(userId, roleId);

        // Assert
        assertEquals(userId, result.getUserId());
        assertEquals(role.getId(), result.getRole().getId());
        assertEquals(role.getName(), result.getRole().getName());
    }
}
