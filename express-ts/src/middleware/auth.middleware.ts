import { verifyToken } from "../utils/jwt.utils";
import { Response, Request, NextFunction } from "express";
import { ApiError } from "../utils/err.utils";


export function authMiddleware (req: Request, res: Response, next: NextFunction) :void{
  const authHeader = req.headers.authorization;
  try {
    if (!authHeader) {
      throw new ApiError(401, "Unauthorized");
    }
    const token = authHeader.split(" ")[1];
    console.log(token)
    const decoded = verifyToken(token);
    if (!decoded) {
      throw new ApiError(401, "Invalid Token");
    }
    res.locals.id = decoded.id;
    res.locals.email = decoded.email;
    next();
  } catch (err) {
    next(err);
  }
}