import jwt from "jsonwebtoken";

type payload = {
  id: number;
  email: string;
}

const password = process.env.JWT_SECRET || "secret";

export function generateToken(payload: payload): string {
  const token = jwt.sign(payload, password, { expiresIn: "1h" })
  return token;
}

export function verifyToken(token: string): payload | null {
  try {
    const decoded = jwt.verify(token, password);
    return decoded as payload;
  } catch (err) {
    return null;
  }
}