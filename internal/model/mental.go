package model

type Answer struct {
	QuestionId int    `v:"required" json:"question_id"`
	OptionId   string `v:"required" json:"option_id"`
}

type MentalAnswerInput struct {
	ScaleId int `v:"required" json:"scale_id"`
	Answers []struct {
		QuestionId int    `v:"required" json:"question_id"`
		OptionId   string `v:"required" json:"option_id"`
	}
}
type MentalAnswerOutput struct {
	Score  int    `json:"score"`
	Report string `json:"report"`
	Advice string `json:"advice"`
}
