package com.jwt.app.auth;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;
import io.javalin.http.Context;
import io.javalin.http.Handler;
import org.jetbrains.annotations.NotNull;

public class JwtMiddleware implements Handler {
    private final String secret = "your-secret"; // Replace with a real secret

    @Override
    public void handle(@NotNull Context ctx) throws Exception {
        String token = ctx.header("Authorization");
        if (token == null || !token.startsWith("Bearer ")) {
            ctx.status(401).json("Unauthorized");
            return;
        }

        String jwt = token.substring(7);

        try {
            Algorithm algorithm = Algorithm.HMAC256(secret);
            DecodedJWT decodedJWT = JWT.require(algorithm).build().verify(jwt);
            ctx.attribute("userId", decodedJWT.getClaim("userId").asInt());
        } catch (Exception e) {
            ctx.status(401).json("Unauthorized");
        }
    }
}
