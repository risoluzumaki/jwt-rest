import { UserRepositoryInterface, User } from "./user.repository";
import { userRepository } from "./user.repository";
import { ApiError } from "../../utils/err.utils";

export interface UserServiceInterface {
  getUserID(id: number): User;
  getUserEmail(email: string, password: string): User;
}

class UserService implements UserServiceInterface {
  constructor(private userRepository: UserRepositoryInterface) {}

  getUserID(id: number): User {
    const user = this.userRepository.getUserByID(id);
    if (!user) {
      throw new Error("User not found");
    }
    return user;
  }

  getUserEmail(email: string, password: string): User {
    const user = this.userRepository.getUserByEmail(email);
    if (!user) {
      throw new ApiError(404, "User Not Found");
    }
    const compare = user.password === password;
    if (!user || !compare) {
      throw new ApiError(401, "Invalid email or password");
    }
    return user;
  }
}

export const userService = new UserService(userRepository)