package quote

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service { return &Service{repo: r} }
func (s *Service) Random(ctx context.Context) (Quote, error) {
	return s.repo.Random(ctx)
}
