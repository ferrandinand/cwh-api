package domain

type ProjectRepositoryStub struct {
	projects []Project
}

func (s ProjectRepositoryStub) FindAll() ([]Project, error) {
	return s.projects, nil
}

func NewProjectRepositoryStub() ProjectRepositoryStub {
	projects := []Project{
		{2, "test-2", "stan", "01/01/2021", 1, "http://www.bictbucket.com/opda/test", "{}", "{}", 1},
		{3, "test-2", "stan", "01/01/2021", 1, "http://www.bictbucket.com/opda/test", "{}", "{}", 1},
		{4, "test-2", "stan", "01/01/2021", 1, "http://www.bictbucket.com/opda/test", "{}", "{}", 1},
		{5, "test-2", "stan", "01/01/2021", 1, "http://www.bictbucket.com/opda/test", "{}", "{}", 1},
	}
	return ProjectRepositoryStub{projects}
}
