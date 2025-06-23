package com.veterinary.user_preferences.controller;

import com.veterinary.user_preferences.model.UserPreferences;
import com.veterinary.user_preferences.repository.UserPreferencesRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/preferences")
public class UserPreferencesController {

    @Autowired
    private UserPreferencesRepository preferencesRepository;

    @PostMapping
    public ResponseEntity<UserPreferences> createPreferences(@RequestBody UserPreferences preferences) {
        UserPreferences savedPreferences = preferencesRepository.save(preferences);
        return new ResponseEntity<>(savedPreferences, HttpStatus.CREATED);
    }

    @GetMapping
    public ResponseEntity<List<UserPreferences>> getAllPreferences() {
        List<UserPreferences> preferences = preferencesRepository.findAll();
        return ResponseEntity.ok(preferences);
    }
}
