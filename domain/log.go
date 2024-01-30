package domain

type LogRequest struct {
	StatusCode int    `json:"status_code" form:"status_code"`
	Source     string `json:"source" form:"source"`
	Request    string `json:"request" form:"request"`
	Response   string `json:"response" form:"response"`
	Recipients string `json:"recipients" form:"recipients"`
	CreatedAt  string `json:"created_at" form:"created_at"`
	UpdatedAt  string `json:"updated_at" form:"updated_at"`
}

type LogData struct {
	Id int `json:"id" form:"id"`
	LogRequest
}

func (LogData) TableName() string {
	return "logs"
}

type LogFilterRequest struct {
	LogRequest
	FilterRequest
}
