package auth

var _ AuthUseCase = (*auth)(nil)

//TODO определить, какие еще нужны интерфейсы

// AuthUseCase - интерфейс для аутентификации
type AuthUseCase interface {
	Login(email string, password string) (accessToken, refreshToken string, err error)
	Register(email string, password string) (userId string, err error)
	RefreshToken(refreshToken string) (accessToken string, err error)
	ValidateToken(accessToken string) (valid bool, err error)
	Logout(accessToken string) (err error)
}

// Auth - структура для аутентификации, пока пустая
type auth struct {
}

func NewAuth() AuthUseCase {
	return &auth{}
}

// Login - логика аутентификации
func (uc *auth) Login(email string, password string) (accessToken, refreshToken string, err error) {
	//TODO определить логику
	return "", "", nil
}

// Register - логика регистрации
func (uc *auth) Register(email string, password string) (userId string, err error) {
	//TODO определить логику
	return "", nil
}

// RefreshToken - логика обновления токена
func (uc *auth) RefreshToken(refreshToken string) (accessToken string, err error) {
	//TODO определить логику
	return "", nil
}

// ValidateToken - логика проверки токена
func (uc *auth) ValidateToken(accessToken string) (valid bool, err error) {
	//TODO определить логику
	return false, nil
}

// Logout - логика выхода из системы
func (uc *auth) Logout(accessToken string) error {
	//TODO определить логику
	return nil
}
