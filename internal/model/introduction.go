package model

type GetIntroductionByUserIdInput struct {
	UserId string
}

type GetIntroductionByUserIdOutput struct {
	ProfileId  string
	UserId     string
	Content    string
	CreateTime string
	UpdateTime string
}
