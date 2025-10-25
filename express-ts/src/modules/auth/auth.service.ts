import { UserRepositoryInterface , userRepository} from "../user/user.repository"
import { ApiError } from "../../utils/err.utils"
import { generateToken } from "../../utils/jwt.utils"

export interface AuthServiceInterface {
  LoginUser(email: string, password: string): string ; 
}

export class AuthService implements AuthServiceInterface {
  constructor (private userRepository: UserRepositoryInterface) {}

  LoginUser(email: string, password: string): string  {
    console.log(email, password)
    const user = this.userRepository.getUserByEmail(email)
    if (!user) {
      throw new ApiError(404, "User Not Found")
    }
    const compare = user.password === password
    if (!compare) {
      throw new ApiError(401, "Invalid email or password")
    }
    const token = generateToken({ id: user.id, email: user.email })
    return token
  }
}

export const authService = new AuthService(userRepository)