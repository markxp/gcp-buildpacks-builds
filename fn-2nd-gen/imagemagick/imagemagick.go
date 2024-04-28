// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START functions_imagemagick_setup]

// Package imagemagick contains an example of using ImageMagick to process a
// file uploaded to Cloud Storage.
package imagemagick

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Global API clients used across function invocations.
var (
	storageClient *storage.Client
)

func init() {
	// Declare a separate err variable to avoid shadowing the client variables.
	var err error

	bgctx := context.Background()
	storageClient, err = storage.NewClient(bgctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	functions.CloudEvent("blur-images", blurEverything)
}

// [END functions_imagemagick_setup]

// [START functions_imagemagick_analyze]

// GCSEvent is the payload of a GCS event.
// additional fields are documented at
// https://cloud.google.com/storage/docs/json_api/v1/objects#resource
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// blurEverything blurs an image uploaded to Cloud Storage
func blurEverything(ctx context.Context, e cloudevents.Event) error {
	outputBucket := os.Getenv("BLURRED_BUCKET_NAME")
	if outputBucket == "" {
		return errors.New("environment variable BLURRED_BUCKET_NAME must be set")
	}

	gcsEvent := &GCSEvent{}
	if err := e.DataAs(gcsEvent); err != nil {
		return fmt.Errorf("e.DataAs: failed to decode event data: %v", err)
	}

	

	return blur(ctx, gcsEvent.Bucket, outputBucket, gcsEvent.Name)
}

// [START functions_imagemagick_blur]

// blur blurs the image stored at gs://inputBucket/name and stores the result in
// gs://outputBucket/name.
func blur(ctx context.Context, inputBucket, outputBucket, name string) error {
	inputBlob := storageClient.Bucket(inputBucket).Object(name)
	r, err := inputBlob.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("inputBlob.NewReader: %w. gs://%s/%s", err, inputBucket, name)
	}

	outputBlob := storageClient.Bucket(outputBucket).Object(name)
	w := outputBlob.NewWriter(ctx)
	defer w.Close()

	// Use - as input and output to use stdin and stdout.
	cmd := exec.Command("convert", "-", "-blur", "0x8", "-")
	cmd.Stdin = r
	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cmd.Run: %v", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}
	log.Printf("Blurred image uploaded to gs://%s/%s", outputBlob.BucketName(), outputBlob.ObjectName())

	return nil
}

// [END functions_imagemagick_blur]
