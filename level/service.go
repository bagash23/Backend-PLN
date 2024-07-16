package level

type Service interface {
	CreateLevel(input LevelInput)(Level, error)
	GetLevel()([]Level, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateLevel(input LevelInput)(Level, error) {
	level := Level{}
	level.IDLevel = input.IDLevel
	level.Level = input.Level

	newLevel, err := s.repository.InputLevel(level)
	if err != nil {
		return newLevel, err
	}
	return newLevel, nil
}

func (s *service) GetLevel()([]Level, error) {
	level, err := s.repository.FindAll()
	if err != nil {
		return level, err
	}
	return level, nil
}