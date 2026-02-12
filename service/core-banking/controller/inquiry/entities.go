package inquiry

import (
	entitiesCore "core-system/core/entities/core-banking"
	usecase "core-system/service/core-banking/usecase/inquiry"
)

type ControllerInterface interface {
	entitiesCore.CoreBankingInterface
}

type controller struct {
	usecase usecase.UsecaseInterface
}

type ControllerConfig struct {
	Usecase usecase.UsecaseInterface
}
