package com.veterinary.security.rolemanager.model;

import jakarta.persistence.*;
import lombok.Data;
@Entity
@Table(name = "roles")
@Data
public class Role {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, unique = true)
    private String name;

    // Constructor por defecto (requerido por JPA)
    public Role() {}

    // Constructor adicional para pruebas o uso manual
    public Role(Long id, String name) {
        this.id = id;
        this.name = name;
    }
}

