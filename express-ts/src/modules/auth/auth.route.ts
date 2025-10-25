import express from "express"
import { Request, Response , NextFunction} from "express"
import { AuthController } from "./auth.controller"

export const authRouter = express.Router()

authRouter.post("/login", function (req :Request, res : Response, next : NextFunction) {
  return AuthController.login(req, res, next)
})
