package restapi

type APIResponse struct {
	Code    int32  `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
