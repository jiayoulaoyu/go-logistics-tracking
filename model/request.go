package model

type InnerTrackRequest struct {
	ApiName      string         `json:"apiName"`
	Language     string         `json:"language"`
	PassWord     string         `json:"passWord"`
	Priority     int            `json:"priority"`
	TrackRequest []TrackRequest `json:"trackObjects"`
}

type TrackRequest struct {
	LogisticsNo string `json:"logisticsNo"`
	PostCode    string `json:"postCode"`
	Des         string `json:"des"`
}
