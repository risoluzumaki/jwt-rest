package user

type UserServiceInterface interface {
	GetUserProfile(id int) (*User, error)
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserProfile(id int) (*User, error) {
	if user, err := s.repo.GetByID(id); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
