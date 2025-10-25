import { Request, Response, NextFunction } from "express";

import { authService } from "./auth.service";
import { AuthRequest, AuthResponse } from "./auth.dto";

export class AuthController {
  
  static login (req: Request, res: Response, next : NextFunction) {
    try {
      const { email, password } = req.body as AuthRequest;
      const token = authService.LoginUser(email, password);
      const response: AuthResponse = { token };
      res.json(response);
    } catch (err) {
      next(err);
    }
  }
}