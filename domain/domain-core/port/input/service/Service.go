package service

import "gozero-sso-service/domain/domain-core/port/output/repository"

type Service interface {
	InitService(repository repository.Repository) error
}
