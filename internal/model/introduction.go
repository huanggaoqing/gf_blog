package model

type GetIntroductionByUserIdInput struct {
	UserId string
}

type GetIntroductionByUserIdOutput struct {
	IntroductionId string
	UserId         string
	Content        string
	CreateTime     string
	UpdateTime     string
}
