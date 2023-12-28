package repository

import "context"

type Repository interface {
	UserAuditRepository
	UserRepository
	UserTokenRepository
	ArticleRepository
	ClassifyRepository
	SetMealRepository
	UserCollectRepository
}

type UserAuditRepository interface {
	CreateUserAudit(ctx context.Context, audit *UserAudit) error
	PassUserAudit(ctx context.Context, id uint64, isPass uint8) error
	GetUserAuditById(ctx context.Context, id uint64) (*UserAudit, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, id uint64) (*User, error)
	GetUsersByMoney(ctx context.Context, classify uint8) ([]*User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*User, error)
}

type UserCollectRepository interface {
	GetUserCollectByUserId(ctx context.Context, userId uint64) ([]*UserCollect, error)
}

type UserTokenRepository interface {
	GetUserTokenByUserId(ctx context.Context, userId uint64) (*UserToken, error)
}

type ArticleRepository interface {
	GetArticleByTitle(ctx context.Context, title string) (*Article, error)
}

type ClassifyRepository interface {
	GetClassify(ctx context.Context, types uint8) ([]*Classify, error)
}

type SetMealRepository interface {
	GetSetMeal(ctx context.Context, types uint8) ([]*SetMeal, error)
}
