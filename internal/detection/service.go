package detection

import (
	"context"

	"go.temporal.io/sdk/client"
)

type Service interface {
	Start(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error
	End(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error
}

type DetectionService struct {
	repository Repository
}

func NewDetectionService(repository Repository) *DetectionService {
	return &DetectionService{
		repository: repository,
	}
}

func (s *DetectionService) Start(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error {
	err := s.repository.StartDetection(ctx, temporalClient, alarmScheduleID, userEmail)
	if err != nil {
		return err
	}
	return nil
}

func (s *DetectionService) End(ctx context.Context, temporalClient client.Client, alarmScheduleID string, userEmail string) error {
	err := s.repository.EndDetection(ctx, temporalClient, alarmScheduleID, userEmail)
	if err != nil {
		return err
	}
	return nil
}
