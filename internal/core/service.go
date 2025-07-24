package core

type UserService interface {
	GetWelcomeMessage(name string) string
}
