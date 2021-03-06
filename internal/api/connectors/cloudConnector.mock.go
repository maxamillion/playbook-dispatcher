package connectors

import (
	"context"

	"github.com/google/uuid"
)

type cloudConnectorClientMock struct {
}

func NewConnectorClientMock() CloudConnectorClient {
	return &cloudConnectorClientMock{}
}

func (this *cloudConnectorClientMock) SendCloudConnectorRequest(
	ctx context.Context,
	account string,
	recipient uuid.UUID,
	correlationId uuid.UUID,
	url string,
) (*string, bool, error) {
	id := uuid.New().String()
	return &id, false, nil
}
