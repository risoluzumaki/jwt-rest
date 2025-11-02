package com.jwt.app.auth.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.jwt.app.auth.dto.LoginRequest;
import com.jwt.app.auth.dto.LoginResponse;
import com.jwt.app.auth.service.AuthService;
import io.javalin.http.Context;

import java.io.IOException;

public class AuthController {
    private final AuthService authService;
    private final ObjectMapper objectMapper = new ObjectMapper();

    public AuthController(AuthService authService) {
        this.authService = authService;
    }

    public void login(Context ctx) throws IOException {
        LoginRequest loginRequest = objectMapper.readValue(ctx.body(), LoginRequest.class);
        String token = authService.login(loginRequest.getEmail(), loginRequest.getPassword());

        if (token != null) {
            ctx.json(new LoginResponse(token));
        } else {
            ctx.status(401).json("Invalid credentials");
        }
    }
}
