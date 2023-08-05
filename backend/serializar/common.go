package serializar

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type TokenResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
