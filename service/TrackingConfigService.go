package service

import (
	"go-logistics-tracking/api/ups"
	"go-logistics-tracking/model"
)

var API_MAP = make(map[string]func(request []model.TrackRequest) []model.TrackResponse)

func init() {
	API_MAP["API_UPS_CRAWL"] = ups.Track()
}

func Execute(request []model.TrackRequest, api string) []model.TrackResponse {
	return API_MAP[api](request)
}
