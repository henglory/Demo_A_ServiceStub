package spec

type AReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ARes struct {
	StatusCode string `json:"statusCode"`
	Result     int    `json:"result"`
}
