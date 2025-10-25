import express, { Express, Router } from "express";
import cors from "cors";
import { errorHandler } from "./middleware/err.midleware";
import { authRouter } from "./modules/auth/auth.route";
import { userRouter } from "./modules/user/user.route";


export default function bootstrap() {
  const app: Express = express();
  const router: Router = express.Router();

  app.use(express.json());
  app.use(cors());

  // 
  router.use('/auth', authRouter);
  router.use('/user', userRouter);

  // 
  app.use('/api/v1', router);

  app.use(errorHandler)
  app.listen(3000, () => {
    console.log("Server running on http://localhost:3000/api/v1");
  });
}
