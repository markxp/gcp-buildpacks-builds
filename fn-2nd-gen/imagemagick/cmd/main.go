package main

import (
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var (
	dest        string = "http://127.0.0.1:8080"
	client      cloudevents.Client
	inputBucket string = "blur-input-001"
)

type gcsEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func main() {
	ctx := context.Background()
	var err error
	client, err = cloudevents.NewClientHTTP()
	if err != nil {
		panic(fmt.Errorf("could not create cloudevent.Client: %w", err))
	}

	ev := cloudevents.NewEvent()
	ev.SetSource("example/uri")
	ev.SetType("example.type")
	ev.SetData(cloudevents.ApplicationJSON, &gcsEvent{
		Bucket: inputBucket,
		Name:   "1.png",
	})

	ctx = cloudevents.ContextWithTarget(ctx, dest)
	if result := client.Send(ctx, ev); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send: %v", result)
	}

}
