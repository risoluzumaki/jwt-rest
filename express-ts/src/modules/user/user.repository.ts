export type User = {
  id: number;
  username: string;
  name: string;
  email: string;
  password: string;
}

export interface UserRepositoryInterface {
  getUserByID(id: number): User | undefined;
  getUserByEmail(email: string): User | undefined;
}

class UserRepository implements UserRepositoryInterface {
  private users: User[] = [{
    id: 1,
    username: 'testuser',
    name: 'Test User',
    email: 'user1@test.com',
    password: 'passworduser1'
  }];

  getUserByID(id: number): User | undefined {
    return this.users.find(user => user.id === id);
  }

  getUserByEmail(email: string): User | undefined {
    return this.users.find(user => user.email.toLowerCase() === email.toLowerCase());
  }
}

export const userRepository = new UserRepository();