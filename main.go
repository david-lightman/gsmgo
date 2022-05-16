package main

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

func main() {
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		log.Fatalf("Environment variable SECRET is required")
	}
	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		return
	}
	defer c.Close()
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", credentials.ProjectID, secret),
	}
	res, err := c.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("unable to access secret-version: %v", err)
	}
	fmt.Println(string(res.Payload.Data))
}
