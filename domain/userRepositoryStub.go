package domain

type UserRepositoryStub struct {
	users []User
}

func (s UserRepositoryStub) FindAll() ([]User, error) {
	return s.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	users := []User{
		{1, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
		{2, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
		{3, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
		{4, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
		{5, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
		{6, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
		{7, "Rob", "18/02/2019 11:15:45", "rob@test.com", "dfsdf,dsfds", "", 1},
	}
	return UserRepositoryStub{users}
}
