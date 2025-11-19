package orderapi

type PhoneReq struct {
	Phone string `json:"phone" validate:"required"`
}

type SessionResp struct {
	SessionId string `json:"session_id"`
}

type CodeReq struct {
	SessionId string `json:"session_id"`
	Code      string `json:"code"`
}

type TokenResp struct {
	Token string `json:"token"`
}
