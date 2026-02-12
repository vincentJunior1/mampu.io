package usecase

import "errors"

func InitUsecase(cfg UsecaseConfig) (UsecaseInterface, error) {
	if cfg.Repo == nil {
		return nil, errors.New("Repo cannot be nil")
	}
	return &usecase{
		repo: cfg.Repo,
	}, nil
}
