package model

type TrackResponse struct {
	LogisticsNo string  `json:"cNo"`
	Status      int     `json:"nState"`
	Message     string  `json:"cMess"`
	PodDate     string  `json:"dPodDate"`
	Sign        string  `json:"cSign"`
	EventListZh []Event `json:"xEventList"`
	EventListEn []Event `json:"xEventListEn"`
}

type Event struct {
	Date     string `json:"xDate"`
	Location string `json:"xPlace"`
	Message  string `json:"xInfo"`
}

type TrackRequest struct {
	LogisticsNo string `json:"logisticsNo"`
	PostCode    string `json:"postCode"`
	Des         string `json:"des"`
}

// 构建空的返回结果
func BuildEmptyTrackResponse(logisticsNo, message string, status int) (res TrackResponse) {
	res.LogisticsNo = logisticsNo
	res.Message = message
	res.Status = status
	return
}
