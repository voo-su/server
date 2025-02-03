package model

import "time"

type AccessLog struct {
	RequestId       string    `ch:"request_id"`
	RequestMethod   string    `ch:"request_method"`
	RequestHeader   string    `ch:"request_header"`
	RequestURI      string    `ch:"request_uri"`
	RequestBody     string    `ch:"request_body"`
	RequestTime     time.Time `ch:"request_time"`
	RequestQuery    string    `ch:"request_query"`
	ResponseTime    time.Time `ch:"response_time"`
	HttpUserAgent   string    `ch:"http_user_agent"`
	HttpStatus      int       `ch:"http_status"`
	ServerName      string    `ch:"server_name"`
	HostName        string    `ch:"host_name"`
	RemoteAddr      string    `ch:"remote_addr"`
	ResponseHeader  string    `ch:"response_header"`
	ResponseBody    string    `ch:"response_body"`
	ResponseBodyRaw string    `ch:"response_body_raw"`
	RequestDuration string    `ch:"request_duration"`
	RequestBodyRaw  string    `ch:"request_body_raw"`
	CreatedAt       time.Time `ch:"created_at"`
}

func (AccessLog) TableName() string {
	return "access_logs"
}
