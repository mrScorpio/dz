package orderapi

type AuthService struct {
	repo *UserRepository
}

func NewAuthService(userRepo *UserRepository) *AuthService {
	return &AuthService{repo: userRepo}
}

func (as *AuthService) ChkPhone(phone string) (string, error) {
	user, err := as.repo.GetUserByPhone(phone)
	if err == nil {
		user.GenSessionId()
		as.repo.Update(user)
		return user.SessionID, nil
	}

	newUser := NewUser(phone)
	user, err = as.repo.Create(newUser)
	if err != nil {
		return "", err
	}

	return user.SessionID, nil
}

func (as *AuthService) ChkCode(sessionId, code string) (string, error) {
	user, err := as.repo.GetBySession(sessionId)
	if err != nil {
		return "", err
	}
	if user.Code == code {
		newJwt := NewJWT(user.Secret)
		token, err := newJwt.Create(user.Phone)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", nil
}
