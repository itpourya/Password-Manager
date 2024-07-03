package serilizers

type SaveRequest struct {
	Username  string `json:"username" form:"username" binding:"required,min=4"`
	Password  string `json:"password" form:"password" binding:"required,min=4"`
	Website   string `json:"website" form:"website" binding:"required,min=4"`
}