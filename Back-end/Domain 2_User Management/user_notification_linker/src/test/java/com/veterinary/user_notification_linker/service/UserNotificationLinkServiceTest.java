package com.veterinary.user_notification_linker.service;

import com.veterinary.user_notification_linker.model.UserNotificationLink;
import com.veterinary.user_notification_linker.repository.UserNotificationLinkRepository;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class UserNotificationLinkServiceTest {

    private final UserNotificationLinkRepository repository = mock(UserNotificationLinkRepository.class);
    private final UserNotificationLinkService service = new UserNotificationLinkService(repository);

    @Test
    void saveLink_shouldReturnSavedLink() {
        UserNotificationLink link = new UserNotificationLink();
        link.setUserId(1L);
        link.setPreferenceId(2L);

        when(repository.save(link)).thenReturn(link);

        UserNotificationLink result = service.saveLink(link);

        assertNotNull(result);
        assertEquals(1L, result.getUserId());
        assertEquals(2L, result.getPreferenceId());
        verify(repository, times(1)).save(link);
    }

    @Test
    void getAllLinks_shouldReturnListOfLinks() {
        UserNotificationLink link1 = new UserNotificationLink();
        UserNotificationLink link2 = new UserNotificationLink();

        when(repository.findAll()).thenReturn(List.of(link1, link2));

        List<UserNotificationLink> result = service.getAllLinks();

        assertEquals(2, result.size());
        verify(repository, times(1)).findAll();
    }
}
