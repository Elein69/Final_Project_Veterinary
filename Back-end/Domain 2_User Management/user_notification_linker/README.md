# 🧩 User Notification Linker Microservice

Microservicio del sistema veterinario distribuido encargado de **vincular usuarios con sus preferencias de notificación**. Pertenece al **Dominio 2 - Gestión de Usuario** del sistema.

---

## 🚀 Funcionalidad

Permite realizar operaciones CRUD sobre las asociaciones entre:

- `userId`: ID del usuario.
- `preferenceId`: ID de la preferencia de notificación asociada.

---

## 🛠️ Tecnologías

- **Lenguaje**: Java 17
- **Framework**: Spring Boot 3.5.3
- **Base de datos**: PostgreSQL
- **ORM**: Spring Data JPA (Hibernate)
- **Build tool**: Maven
- **Testing**: JUnit 5, Mockito
- **Contenedores**: Docker

---

## 📦 Estructura del Proyecto

```
src/
└── main/
    ├── java/com/veterinary/user_notification_linker/
    │   ├── controller/             # Controladores REST
    │   ├── service/                # Lógica de negocio
    │   ├── repository/             # Acceso a datos JPA
    │   ├── model/                  # Entidades JPA
    │   └── UserNotificationLinkerApplication.java
    └── resources/
        └── application.properties
```

---

## 📄 Endpoints

| Método | Endpoint                      | Descripción                        |
|--------|-------------------------------|------------------------------------|
| GET    | `/notification-links`         | Listar todos los vínculos         |
| GET    | `/notification-links/{id}`    | Obtener un vínculo específico     |
| POST   | `/notification-links`         | Crear un nuevo vínculo            |
| DELETE | `/notification-links/{id}`    | Eliminar un vínculo existente     |

---

## 🧪 Pruebas

Las pruebas unitarias se encuentran en:
```
src/test/java/com/veterinary/user_notification_linker/service/
```

Y pueden ejecutarse con:
```bash
mvn test
```

---

## 🐳 Docker

### Crear imagen
```bash
docker build -t user-notification-linker .
```

### Ejecutar contenedor
```bash
docker run -p 8080:8080 user-notification-linker
```

---

## 🔐 Seguridad

Este microservicio puede requerir un JWT para acceder a los endpoints. Verifica si la configuración de seguridad está habilitada en tu entorno antes de realizar peticiones con Postman.

---

## 📫 Autor

Proyecto académico desarrollado por Steven Alexander para la materia **Programación Distribuida**.
