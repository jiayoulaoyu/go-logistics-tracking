package service

import (
	"go-logistics-tracking/api/ups"
	"go-logistics-tracking/model"
)

var SERVICE_MAP = make(map[string]func(request []model.TrackRequest) []model.TrackResponse)

func init() {
	SERVICE_MAP["UPS"] = ups.Track()
}

func Execute(request []model.TrackRequest, api string) []model.TrackResponse {
	return SERVICE_MAP[api](request)
}
