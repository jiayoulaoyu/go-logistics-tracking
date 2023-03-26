package service

import (
	"go-logistics-tracking/api/ups"
	"go-logistics-tracking/model"
	"sort"
	"strings"
)

var API_MAP = make(map[string]func(request []model.TrackRequest) []model.TrackResponse)

func init() {
	API_MAP["API_UPS_CRAWL"] = ups.Track()
}

func Execute(request []model.TrackRequest, api string) []model.TrackResponse {
	res := API_MAP[api](request)
	for _, v := range res {
		eventsEn := v.EventListEn
		eventsZh := v.EventListZh
		if len(eventsEn) > 0 {
			sort.Slice(eventsEn, func(i, j int) bool {
				return strings.Compare(eventsEn[i].Date, eventsEn[j].Date) < 0
			})
		}

		if len(eventsZh) > 0 {
			sort.Slice(eventsZh, func(i, j int) bool {
				return strings.Compare(eventsZh[i].Date, eventsZh[j].Date) < 0
			})
		}
	}
	return res
}
