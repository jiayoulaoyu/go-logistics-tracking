package model

import "go-logistics-tracking/consts"

type Event struct {
	Date     string `json:"xDate"`
	Location string `json:"xPlace"`
	Message  string `json:"xInfo"`
}
type TrackResponse struct {
	LogisticsNo string  `json:"cNo"`
	Status      int     `json:"nState"`
	Message     string  `json:"cMess"`
	PodDate     string  `json:"dPodDate"`
	Sign        string  `json:"cSign"`
	EventListZh []Event `json:"xEventList"`
	EventListEn []Event `json:"xEventListEn"`
}

type WebTrackingResponse struct {
	ReturnValue    int             `json:"ReturnValue"`
	Message        string          `json:"cMess"`
	TrackResponses []TrackResponse `json:"NoList"`
}

func NewErrorResponse(returnValue int, message string) WebTrackingResponse {
	return WebTrackingResponse{returnValue, message, nil}
}

func NewSuccessResponse(res []TrackResponse) WebTrackingResponse {
	return WebTrackingResponse{consts.SUCCESS, "", res}
}

// 构建空的返回结果
func BuildEmptyTrackResponse(logisticsNo, message string, status int) (res TrackResponse) {
	res.LogisticsNo = logisticsNo
	res.Message = message
	res.Status = status
	return
}
