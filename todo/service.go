package todo

type TodoService interface {
	Create(title, description string) (Todo, error)
	FindById(id int) (Todo, error)
	FindAll(limit, offset int) ([]Todo, error)
	Delete(id int) (Todo, error)
	Update(id int, title, description string, status bool) (Todo, error)
	Done(id int, status bool) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(title, description string) (Todo, error) {
	if len(title) <= 0 {
		return Todo{}, TitleEmptyErr
	}
	if len(title) >= 50 {
		return Todo{}, TitleSizeErr
	}

	return s.repository.Create(title, description)
}

func (s *service) FindById(id int) (Todo, error) {
	return s.repository.FindById(id)
}

func (s *service) FindAll(limit, offset int) ([]Todo, error) {
	return s.repository.FindAll(limit, offset)
}

func (s *service) Delete(id int) (Todo, error) {
	return s.repository.Delete(id)
}

func (s *service) Update(id int, title, description string, status bool) (Todo, error) {
	if len(title) <= 0 {
		return Todo{}, TitleEmptyErr
	}
	if len(title) >= 50 {
		return Todo{}, TitleSizeErr
	}

	return s.repository.Update(id, title, description, status)
}

func (s *service) Done(id int, status bool) (Todo, error) {
	return s.repository.Done(id, status)
}
