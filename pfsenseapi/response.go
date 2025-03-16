package pfsenseapi

type apiResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Return  int         `json:"return"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
