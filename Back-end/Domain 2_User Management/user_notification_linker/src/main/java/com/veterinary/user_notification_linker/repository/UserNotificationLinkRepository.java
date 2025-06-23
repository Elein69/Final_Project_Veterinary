package com.veterinary.user_notification_linker.repository;

import com.veterinary.user_notification_linker.model.UserNotificationLink;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserNotificationLinkRepository extends JpaRepository<UserNotificationLink, Long> {
}
