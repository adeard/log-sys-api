package domain

type LogRequest struct {
	StatusCode int    `json:"status_code" form:"status_code"`
	Source     string `json:"source" form:"source"`
	Request    string `json:"request" form:"request"`
	Response   string `json:"response" form:"response"`
	Recipients string `json:"recipients" form:"recipients"`
	CreatedAt  string `json:"created_at" form:"created_at"`
	UpdatedAt  string `json:"updated_at" form:"updated_at"`
	Method     string `json:"method" form:"method"`
	ClientIp   string `json:"client_ip" form:"client_ip"`
	UserAgent  string `json:"user_agent" form:"user_agent"`
	Duration   string `json:"duration" form:"duration"`
}

type LogData struct {
	Id int `json:"id" form:"id"`
	LogRequest
}

func (LogRequest) TableName() string {
	return "logs"
}

type LogFilterRequest struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	LogRequest
	FilterRequest
}

type LogTotalData struct {
	LogDate  string `json:"log_date" form:"log_date"`
	LogTotal int    `json:"log_total" form:"log_total"`
}

type LogTopErrorData struct {
	Source        string `json:"source" form:"source"`
	LogTotal      int    `json:"log_total" form:"log_total"`
	LastCreatedAt string `json:"last_created_at" form:"last_created_at"`
	Response      string `json:"response" form:"response"`
	LastId        string `json:"last_id" form:"last_id"`
}
