package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
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
	// Load environment variables from .env file
	godotenv.Load()
	// SES SMTP credentials
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	// Set the sender and recipient
	// "from" address must be verified with Amazon SES
	// If your account is still in the sandbox, "to" address must also be verified
	from := os.Getenv("FROM_EMAIL")
	to := os.Getenv("TO_EMAIL")

	host := "email-smtp.us-east-2.amazonaws.com"
	port := "587" // STARTTLS Ports: 25, 587, or 2587

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	// Connect to the SMTP server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting to SMTP server:", err)
		return
	}

	// Create new SMTP client
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		return
	}

	// Set TLS config
	if err = client.StartTLS(tlsConfig); err != nil {
		fmt.Println("Error starting TLS:", err)
		return
	}

	// Authenticate
	auth := smtp.PlainAuth("", username, password, host)
	if err = client.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		return
	}

	// Set email headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = "Test Email from Amazon SES"

	// Set email body
	body := "This is a test email sent from Amazon SES using Go."

	// Compose the message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Send the email
	if err = client.Mail(from); err != nil {
		fmt.Println("Error setting sender:", err)
		return
	}
	if err = client.Rcpt(to); err != nil {
		fmt.Println("Error setting recipient:", err)
		return
	}
	w, err := client.Data()
	if err != nil {
		fmt.Println("Error getting data writer:", err)
		return
	}
	_, err = w.Write([]byte(message)) // Transmit email content to SMTP server
	if err != nil {
		fmt.Println("Error writing message:", err)
		return
	}
	err = w.Close() // Finalize transmission
	if err != nil {
		fmt.Println("Error closing data writer:", err)
		return
	}

	client.Quit()

	fmt.Println("Email sent successfully!")
}
