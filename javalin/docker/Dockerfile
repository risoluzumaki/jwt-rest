# Stage 1: Build the application
FROM maven:3.9.6-eclipse-temurin-21 AS builder
WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn package

# Stage 2: Create the final image
FROM eclipse-temurin:21-jre-jammy
WORKDIR /app
COPY --from=builder /app/target/classes ./classes
COPY --from=builder /app/target/dependency ./dependency
EXPOSE 3000
CMD ["java", "-cp", "classes:dependency/*", "com.jwt.app.Main"]
