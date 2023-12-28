package gorm

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/db"
)

/** 会员审核表 **/
func (r *gormRepository) userAuditHelper() *db.Helper[repository.UserAudit] {
	return db.NewHelper[repository.UserAudit](r.db)
}

func (r *gormRepository) CreateUserAudit(ctx context.Context, audit *repository.UserAudit) error {
	return r.userAuditHelper().Create(ctx, audit)
}

func (r *gormRepository) PassUserAudit(ctx context.Context, id uint64, isPass uint8) error {
	var data repository.UserAudit
	return r.db.Model(&data).Where("id = ?", id).Update("status", isPass).Error
}

func (r *gormRepository) GetUserAuditById(ctx context.Context, id uint64) (*repository.UserAudit, error) {
	return r.userAuditHelper().Where("id = ?", id).First(ctx)
}

/** 会员表 **/
func (r *gormRepository) userHelper() *db.Helper[repository.User] {
	return db.NewHelper[repository.User](r.db)
}

func (r *gormRepository) CreateUser(ctx context.Context, user *repository.User) error {
	return r.userHelper().Create(ctx, user)
}

func (r *gormRepository) UpdateUser(ctx context.Context, user *repository.User) error {
	return r.db.Model(&user).Where("id = ?", user.Id).Updates(user).Error
}

func (r *gormRepository) GetUserById(ctx context.Context, id uint64) (*repository.User, error) {
	return r.userHelper().Where("is_del = 1 AND id = ?", id).First(ctx)
}

func (r *gormRepository) GetUsersByMoney(ctx context.Context, classify uint8) ([]*repository.User, error) {
	return r.userHelper().Order("money desc").Limit(10).Where("is_type = 1 AND is_del = 1").
		Where("is_classify = ?", classify).Find(ctx)
}

func (r *gormRepository) GetUserByMobile(ctx context.Context, mobile string) (*repository.User, error) {
	return r.userHelper().Where("is_del = 1 AND mobile = ?", mobile).First(ctx)
}

/** 用户token表 **/
func (r *gormRepository) userTokenHelper() *db.Helper[repository.UserToken] {
	return db.NewHelper[repository.UserToken](r.db)
}

func (r *gormRepository) GetUserTokenByUserId(ctx context.Context, userId uint64) (*repository.UserToken, error) {
	return r.userTokenHelper().Where("user_id = ?", userId).First(ctx)
}

/* 用户关注表 */
func (r *gormRepository) userCollectHelper() *db.Helper[repository.UserCollect] {
	return db.NewHelper[repository.UserCollect](r.db)
}

func (r *gormRepository) GetUserCollectByUserId(ctx context.Context, userId uint64) ([]*repository.UserCollect, error) {
	return r.userCollectHelper().Where("user_id = ? AND status = ?", userId, 1).Find(ctx)
}
