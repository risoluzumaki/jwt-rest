import express from "express"
import { authMiddleware } from "../../middleware/auth.middleware";
import { UserController } from "./user.controller";
import {Request, Response, NextFunction} from "express"

export const userRouter = express.Router();

// Middleware
userRouter.use(authMiddleware)

// Routes
userRouter.get("/profile" ,(req: Request, res: Response, next: NextFunction) => UserController.profile(req, res, next))