import { ApiError } from "../utils/err.utils";
import { Request, Response, NextFunction } from "express";

export function errorHandler(err: Error, req_: Request, res: Response, next_: NextFunction){
  if (err instanceof ApiError) {
    return res.status(err.status).json({ message: err.message });
  }
  return res.status(500).json({ message: "Internal server error" });
}