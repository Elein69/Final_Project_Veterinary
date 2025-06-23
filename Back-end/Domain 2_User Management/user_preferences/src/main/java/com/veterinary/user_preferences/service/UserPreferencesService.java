package com.veterinary.user_preferences.service;

import com.veterinary.user_preferences.model.UserPreferences;
import com.veterinary.user_preferences.repository.UserPreferencesRepository;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
public class UserPreferencesService {

    private final UserPreferencesRepository repository;

    public UserPreferencesService(UserPreferencesRepository repository) {
        this.repository = repository;
    }

    public UserPreferences savePreferences(UserPreferences preferences) {
        return repository.save(preferences);
    }

    public Optional<UserPreferences> getPreferencesByUserId(Long userId) {
        return repository.findByUserId(userId);
    }
}
