// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf_blog/internal/dao/internal"
)

// internalBlogUserDao is internal type for wrapping internal DAO implements.
type internalBlogUserDao = *internal.BlogUserDao

// blogUserDao is the data access object for table blog_user.
// You can define custom methods on it to extend its functionality as you wish.
type blogUserDao struct {
	internalBlogUserDao
}

var (
	// BlogUser is globally public accessible object for table blog_user operations.
	BlogUser = blogUserDao{
		internal.NewBlogUserDao(),
	}
)

// Fill with you ideas below.