package domain

type UserRepositoryStub struct {
	users []User
}

func (s UserRepositoryStub) FindAll() ([]User, error) {
	return s.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	jsonMock := map[string]interface{}{
		"test": "1",
	}

	users := []User{
		{"1", "Rob", "123", "18/02/2019 11:15:45", "admin", "rob@test.com", jsonMock, "1"},
		{"2", "Rob", "123", "18/02/2019 11:15:45", "user", "rob@test.com", jsonMock, "1"},
		{"3", "Rob", "123", "18/02/2019 11:15:45", "admin", "rob@test.com", jsonMock, "1"},
		{"4", "Rob", "123", "18/02/2019 11:15:45", "user", "rob@test.com", jsonMock, "1"},
		{"5", "Rob", "123", "18/02/2019 11:15:45", "user", "rob@test.com", jsonMock, "1"},
		{"6", "Rob", "123", "18/02/2019 11:15:45", "user", "rob@test.com", jsonMock, "1"},
		{"7", "Rob", "123", "18/02/2019 11:15:45", "user", "rob@test.com", jsonMock, "1"},
	}
	return UserRepositoryStub{users}
}
