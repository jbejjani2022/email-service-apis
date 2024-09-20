package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	godotenv.Load()
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	// Construct the path to the .env file (one directory up)
	envPath := filepath.Join(currentDir, "..", ".env")

	// Load the .env file
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	from := os.Getenv("FROM_EMAIL")
	to := os.Getenv("TO_EMAIL")

	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	// Create an Amazon SES service client
	client := sesv2.NewFromConfig(cfg)

	// Compose the email
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from),
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String("Test email from Amazon SES"),
				},
				Body: &types.Body{
					Text: &types.Content{
						Data: aws.String("This is a test email sent from Amazon SES using the AWS SDK for Go"),
					},
				},
			},
		},
	}

	// Send the email
	result, err := client.SendEmail(context.TODO(), input)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	fmt.Printf("Email sent! Message ID: %s\n", *result.MessageId)
}
