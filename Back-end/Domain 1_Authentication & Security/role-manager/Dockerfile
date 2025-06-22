# Usa una imagen base oficial de OpenJDK 17
FROM eclipse-temurin:17-jdk-alpine

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el jar compilado (verifica que target/role-manager-0.0.1-SNAPSHOT.jar exista)
COPY target/role-manager-0.0.1-SNAPSHOT.jar app.jar

# Expone el puerto en el que corre tu microservicio
EXPOSE 8084

# Comando para ejecutar la aplicación
ENTRYPOINT ["java", "-jar", "app.jar"]
