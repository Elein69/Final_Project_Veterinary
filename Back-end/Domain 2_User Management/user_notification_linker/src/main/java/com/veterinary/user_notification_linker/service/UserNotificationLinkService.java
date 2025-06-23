package com.veterinary.user_notification_linker.service;

import com.veterinary.user_notification_linker.model.UserNotificationLink;
import com.veterinary.user_notification_linker.repository.UserNotificationLinkRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserNotificationLinkService {

    private final UserNotificationLinkRepository repository;

    @Autowired
    public UserNotificationLinkService(UserNotificationLinkRepository repository) {
        this.repository = repository;
    }

    public UserNotificationLink saveLink(UserNotificationLink link) {
        return repository.save(link);
    }

    public List<UserNotificationLink> getAllLinks() {
        return repository.findAll();
    }

    public UserNotificationLink getLinkById(Long id) {
        return repository.findById(id).orElse(null);
    }

    public void deleteLink(Long id) {
        repository.deleteById(id);
    }
}
