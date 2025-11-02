package com.jwt.app;

import com.jwt.app.auth.JwtMiddleware;
import com.jwt.app.auth.controller.AuthController;
import com.jwt.app.auth.service.AuthService;
import com.jwt.app.user.controller.UserController;
import com.jwt.app.user.repository.UserRepository;
import com.jwt.app.user.service.UserService;
import io.javalin.Javalin;

public class Main {
    
    public static void main(String[] args) {
        UserRepository userRepository = new UserRepository();
        UserService userService = new UserService(userRepository);
        AuthService authService = new AuthService(userRepository);

        UserController userController = new UserController(userService);
        AuthController authController = new AuthController(authService);

        Javalin app = Javalin.create().start(3000);

        app.exception(Exception.class, (e, ctx) -> {
            ctx.status(500).json("Internal Server Error: " + e.getMessage());
        });

        app.post("/api/v1/auth/login", authController::login);
        app.before("/api/v1/user/profile", new JwtMiddleware());
        app.get("/api/v1/user/profile", userController::getProfile);
    }
}