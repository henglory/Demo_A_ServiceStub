package spec

type AReq struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type ARes struct {
	StatusCode string `json:"statusCode"`
	Result     int    `json:"result"`
}
