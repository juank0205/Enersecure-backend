package mock

type UserServiceMock struct {
	WelcomeMsg string
}

func (m *UserServiceMock) GetWelcomeMessage(name string) string {
	return m.WelcomeMsg
}
