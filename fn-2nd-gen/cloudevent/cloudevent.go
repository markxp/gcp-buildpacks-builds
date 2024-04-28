package cloudevent

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func init() {
	functions.CloudEvent("pubsub", pubsubHandler)
}

type pubsubPayload struct {
	Message pubsubMessage
}

type pubsubMessage struct {
	Data []byte `json:"data"`
}

func pubsubHandler(ctx context.Context, e cloudevents.Event) error {
	var msg pubsubPayload
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	s := string(msg.Message.Data)
	fmt.Printf("got message: %s", s)
	return nil
}
