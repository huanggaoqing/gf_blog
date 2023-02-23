// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogTagAssociateDao is the data access object for table blog_tag_associate.
type BlogTagAssociateDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns BlogTagAssociateColumns // columns contains all the column names of Table for convenient usage.
}

// BlogTagAssociateColumns defines and stores column names for table blog_tag_associate.
type BlogTagAssociateColumns struct {
	ArticleId string //
	TagId     string //
}

// blogTagAssociateColumns holds the columns for table blog_tag_associate.
var blogTagAssociateColumns = BlogTagAssociateColumns{
	ArticleId: "article_id",
	TagId:     "tag_id",
}

// NewBlogTagAssociateDao creates and returns a new DAO object for table data access.
func NewBlogTagAssociateDao() *BlogTagAssociateDao {
	return &BlogTagAssociateDao{
		group:   "default",
		table:   "blog_tag_associate",
		columns: blogTagAssociateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogTagAssociateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogTagAssociateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogTagAssociateDao) Columns() BlogTagAssociateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogTagAssociateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogTagAssociateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogTagAssociateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
