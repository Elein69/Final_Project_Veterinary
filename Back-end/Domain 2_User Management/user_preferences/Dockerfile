
FROM openjdk:17-jdk-slim


WORKDIR /app


COPY target/user_preferences-0.0.1-SNAPSHOT.jar app.jar


EXPOSE 8082


ENTRYPOINT ["java", "-jar", "app.jar"]
