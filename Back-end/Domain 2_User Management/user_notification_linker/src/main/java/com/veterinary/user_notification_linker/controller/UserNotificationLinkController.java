package com.veterinary.user_notification_linker.controller;

import com.veterinary.user_notification_linker.model.UserNotificationLink;
import com.veterinary.user_notification_linker.service.UserNotificationLinkService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/notification-links")
public class UserNotificationLinkController {

    private final UserNotificationLinkService service;

    @Autowired
    public UserNotificationLinkController(UserNotificationLinkService service) {
        this.service = service;
    }

    @PostMapping
    public UserNotificationLink createLink(@RequestBody UserNotificationLink link) {
        return service.saveLink(link);
    }

    @GetMapping
    public List<UserNotificationLink> getAllLinks() {
        return service.getAllLinks();
    }

    @GetMapping("/{id}")
    public UserNotificationLink getLink(@PathVariable Long id) {
        return service.getLinkById(id);
    }

    @DeleteMapping("/{id}")
    public void deleteLink(@PathVariable Long id) {
        service.deleteLink(id);
    }
}

