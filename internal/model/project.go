package model

import "gf_blog/internal/model/entity"

type GetProjectListByIdInput struct {
	Page   int
	Size   int
	UserId string
}

type GetProjectListByUserIdOutput struct {
	Page  int
	Size  int
	List  []*entity.BlogProject
	Total int
}
