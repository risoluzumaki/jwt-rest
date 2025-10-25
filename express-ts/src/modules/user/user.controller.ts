import { Request, Response, NextFunction } from "express";
import { userService } from "./user.service";
import { UserResponse } from "./user.dto";

export class UserController{

  static profile(req: Request, res: Response , next : NextFunction) {
    
    const id: number = res.locals.id as number
    // const email: string = res.locals.email as string
  
    try {
      const result = userService.getUserID(id)
      const response: UserResponse = {
        id: result.id,
        email: result.email,
        name: result.name,
        username: result.username
      }
      res.status(200).json(response)
    } catch (err){
      next(err)
    }
  }
} 