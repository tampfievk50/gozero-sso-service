package model

type VT8x79 struct {
	ID            uint    `json:"id" gorm:"column:id;primaryKey"`
	Operator      string  `json:"operator" gorm:"column:operator"`
	ServiceID     string  `json:"service_id" gorm:"column:service_id"`
	UserID        string  `json:"user_id" gorm:"column:user_id"`
	SubmitDate    string  `json:"submit_date" gorm:"submit_date"`
	DoneDate      string  `json:"done_date" gorm:"done_date"`
	SendTime      string  `json:"send_time" gorm:"column:send_time"`
	DeliverTime   string  `json:"deliver_time" gorm:"column:deliver_time"`
	Keyword       string  `json:"keyword" gorm:"column:keyword"`
	Info          string  `json:"info" gorm:"column:info"`
	ContentType   string  `json:"content_type" gorm:"column:content_type"`
	MessageType   string  `json:"message_type" gorm:"column:message_type"`
	NumberMessage int     `json:"number_message" gorm:"column:number_message"`
	RequestID     string  `json:"request_id" gorm:"column:request_id"`
	Thirdparty    string  `json:"thirdparty" gorm:"column:thirdparty"`
	ProcessResult string  `json:"process_result" gorm:"column:process_result"`
	Money         float64 `json:"money" gorm:"column:money"`
	AutoTimestamp string  `json:"autotimestamp" gorm:"column:autotimestamp"`
	LoggingTime   string  `json:"loggingTime" gorm:"column:loggingTime"`
}

func (VT8x79) TableName() string {
	return "mo202508_vt"
}
