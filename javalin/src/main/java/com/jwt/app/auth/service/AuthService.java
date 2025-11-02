package com.jwt.app.auth.service;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.jwt.app.user.dto.User;
import com.jwt.app.user.repository.UserRepository;

public class AuthService {
    private final UserRepository userRepository;
    private final String secret = "your-secret"; // Replace with a real secret

    public AuthService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public String login(String email, String password) {
        User user = userRepository.findByEmail(email);
        if (user == null || !user.getPassword().equals(password)) {
            return null;
        }

        Algorithm algorithm = Algorithm.HMAC256(secret);
        return JWT.create()
                .withClaim("userId", user.getId())
                .sign(algorithm);
    }
}
