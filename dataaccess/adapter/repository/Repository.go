package repository

import oport "gozero-sso-service/domain/domain-core/port/output/repository"

type Repository struct {
	UserRepository oport.UserRepository
	RoleRepository oport.RoleRepository
}
