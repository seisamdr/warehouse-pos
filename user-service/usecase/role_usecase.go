package usecase

import (
	"context"
	"micro-warehouse/user-service/models"
	"micro-warehouse/user-service/repositories"
)

type RoleUsecaseInterface interface {
	CreateRole(ctx context.Context, role models.Role) error
	UpdateRole(ctx context.Context, role models.Role) error
	DeleteRole(ctx context.Context, id uint) error
	GetRoleByID(ctx context.Context, id uint) (*models.Role, error)
	GetAllRoles(ctx context.Context) ([]models.Role, error)
}

type roleUsecase struct {
	roleRepo repositories.RoleRepositoryInterface
}

// CreateRole implements [RoleUsecaseInterface].
func (r *roleUsecase) CreateRole(ctx context.Context, role models.Role) error {
	return r.roleRepo.CreateRole(ctx, role)
}

// DeleteRole implements [RoleUsecaseInterface].
func (r *roleUsecase) DeleteRole(ctx context.Context, id uint) error {
	return r.roleRepo.DeleteRole(ctx, id)
}

// GetAllRoles implements [RoleUsecaseInterface].
func (r *roleUsecase) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	return r.roleRepo.GetAllRoles(ctx)
}

// GetRoleByID implements [RoleUsecaseInterface].
func (r *roleUsecase) GetRoleByID(ctx context.Context, id uint) (*models.Role, error) {
	return r.roleRepo.GetRoleByID(ctx, id)
}

// UpdateRole implements [RoleUsecaseInterface].
func (r *roleUsecase) UpdateRole(ctx context.Context, role models.Role) error {
	return r.roleRepo.UpdateRole(ctx, role)
}

func NewRoleUsecase(roleRepo repositories.RoleRepositoryInterface) RoleUsecaseInterface {
	return &roleUsecase{roleRepo: roleRepo}
}
