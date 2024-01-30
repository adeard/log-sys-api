package domain

type ApplicationRequest struct {
	Name      string `json:"name" form:"name"`
	Url       string `json:"url" form:"url"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}

type ApplicationData struct {
	Id int `json:"id" form:"id"`
	ApplicationRequest
}

func (ApplicationData) TableName() string {
	return "applications"
}

type ApplicationFilterRequest struct {
	ApplicationRequest
	FilterRequest
}
