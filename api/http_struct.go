package api

type CommonResponse struct {
	Msg                string `json:"msg"`                  //
	Ret                int    `json:"ret"`                  //
	Data               string `json:":data"`                //
	ServerExecutedTime string `json:"server_executed_time"` //
}
