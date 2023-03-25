package service

import (
	"go-logistics-tracking/model"
)

type TrackingService interface {
	track(request model.TrackRequest) []model.TrackResponse
}
