package com.veterinary.user_notification_linker.model;

import jakarta.persistence.*;
import lombok.Data;

@Entity
@Data
@Table(name = "user_notification_links")
public class UserNotificationLink {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private Long userId;
    private Long preferenceId;
}
