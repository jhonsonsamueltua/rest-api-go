package models

type User struct {
	UserID   int64  `json:userID,omitempty`
	Username string `json:username,omitempty`
	Password string `json:password,omitempty`
	Name     string `json:name,omitempty"`
}

const (
	StatusSucces   = "success"
	StatusFailed   = "failed"
	MessageSucces  = "Berhasil"
	MeassageFailed = "Tidak Berhasil"
)

type Responses struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
