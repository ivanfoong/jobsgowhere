package message

import "context"

// Service - message service interface
type Service interface {
	SendMessage(ctx context.Context, toPersonID string,
		fromPersonID string, subject string, body string) error
}

type messageService struct {
	repo Repository
}

func (m *messageService) SendMessage(ctx context.Context, toPersonID string,
	fromPersonID string, subject string, body string) error {

	err := m.repo.SendMessage(ctx, toPersonID, fromPersonID, subject, body)

	if err != nil {
		return err
	}

	return nil
}
