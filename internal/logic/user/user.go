package user

import (
	"context"
	"database/sql"
	"gf_blog/internal/dao"
	"gf_blog/internal/model"
	"gf_blog/internal/model/do"
	"gf_blog/internal/service"
	"gf_blog/utility"
	"github.com/gogf/gf/v2/errors/gerror"
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

type sUser struct{}

// Register 用户注册
func (s *sUser) Register(ctx context.Context, in *model.RegisterInput) (out *model.RegisterOutput, err error) {
	if in.Password != in.Password2 {
		return nil, gerror.New("两次输入密码不匹配")
	}
	_, err = s.checkPhoneNumberIsExits(ctx, in.PhoneNumber)
	if err != nil {
		return nil, err
	}
	in.UserId = utility.IdBySnowflake.Generate()
	_, err = dao.BlogUser.Ctx(ctx).Data(in).Insert()
	if err != nil {
		return nil, err
	}
	return &model.RegisterOutput{UserId: in.UserId}, nil
}

// Login 用户登录
func (s *sUser) Login(ctx context.Context, in *model.LoginInput) (out *model.LoginOutput, err error) {
	return
}

// checkPhoneNumberExits 通过手机号查找用户是否存在
func (s *sUser) checkPhoneNumberIsExits(ctx context.Context, phone string) (out *do.BlogUser, err error) {
	user := &do.BlogUser{}
	m := dao.BlogUser.Ctx(ctx).Where("phone_number = ?", phone)
	if err := m.Scan(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, gerror.New("用户已存在")
}
