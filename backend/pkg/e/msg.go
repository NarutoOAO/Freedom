package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "success",
	NotExistInentifier:    "1",
	ERROR:                 "fail",
	InvalidParams:         "failed",

	ErrorExistNick:          "failed",
	ErrorExistUser:          "failed",
	ErrorNotExistUser:       "failed",
	ErrorNotCompare:         "failed",
	ErrorNotComparePassword: "failed",
	ErrorFailEncryption:     "failed",
	ErrorNotExistProduct:    "failed",
	ErrorNotExistAddress:    "failed",
	ErrorExistFavorite:      "failed",
	ErrorUserNotFound:       "failed",

	ErrorBossCheckTokenFail:        "failed",
	ErrorBossCheckTokenTimeout:     "failed",
	ErrorBossToken:                 "failed",
	ErrorBoss:                      "failed",
	ErrorBossInsufficientAuthority: "failed",
	ErrorBossProduct:               "failed",

	ErrorProductExistCart: "failed，数量+1",
	ErrorProductMoreCart:  "failed",

	ErrorAuthCheckTokenFail:        "failed",
	ErrorAuthCheckTokenTimeout:     "failed",
	ErrorAuthToken:                 "failed",
	ErrorAuth:                      "failed",
	ErrorAuthInsufficientAuthority: "failed",
	ErrorReadFile:                  "failed",
	ErrorSendEmail:                 "failed",
	ErrorCallApi:                   "failed",
	ErrorUnmarshalJson:             "failed",

	ErrorUploadFile:    "failed",
	ErrorAdminFindUser: "failed",

	ErrorDatabase: "failed",

	ErrorOss: "failed",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
