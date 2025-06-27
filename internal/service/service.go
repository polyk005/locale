package service

type Service struct {

}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}