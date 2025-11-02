package com.jwt.app.user.controller;

import com.jwt.app.user.dto.User;
import com.jwt.app.user.service.UserService;
import io.javalin.http.Context;

public class UserController {
    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    public void getProfile(Context ctx) {
        Integer userId = ctx.attribute("userId");
        if (userId == null) {
            ctx.status(401).json("Unauthorized");
            return;
        }
        User user = userService.getUserById(userId);
        if (user != null) {
            ctx.json(user);
        } else {
            ctx.status(404).json("User not found");
        }
    }
}
