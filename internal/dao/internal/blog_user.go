// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogUserDao is the data access object for table blog_user.
type BlogUserDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns BlogUserColumns // columns contains all the column names of Table for convenient usage.
}

// BlogUserColumns defines and stores column names for table blog_user.
type BlogUserColumns struct {
	UserId      string //
	UserName    string //
	Password    string //
	PhoneNumber string //
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	Avatar      string //
	Role        string //
}

// blogUserColumns holds the columns for table blog_user.
var blogUserColumns = BlogUserColumns{
	UserId:      "user_id",
	UserName:    "user_name",
	Password:    "password",
	PhoneNumber: "phone_number",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	Avatar:      "avatar",
	Role:        "role",
}

// NewBlogUserDao creates and returns a new DAO object for table data access.
func NewBlogUserDao() *BlogUserDao {
	return &BlogUserDao{
		group:   "default",
		table:   "blog_user",
		columns: blogUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogUserDao) Columns() BlogUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
