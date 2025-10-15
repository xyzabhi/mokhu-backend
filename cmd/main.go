package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mail.v2"
)

func main() {
	fmt.Println("Starting Mokhu Backend...")

	// Test email functionality with MailHog
	if err := sendTestEmail(); err != nil {
		log.Printf("Failed to send test email: %v", err)
		log.Println("Make sure MailHog is running: make mailhog")
		return
	}

	fmt.Println("Email sent successfully to MailHog!")
	fmt.Println("Check MailHog web interface at: http://localhost:8025")
}

func sendTestEmail() error {
	m := mail.NewMessage()
	m.SetHeader("From", "noreply@mokhu.com")
	m.SetHeader("To", "abhinavkumar8083@gmail.com")
	m.SetHeader("Subject", "MailHog Test - Mokhu Backend")
	m.SetBody("text/html", `
		<h1>Hello from Mokhu Backend!</h1>
		<p>This is a test email sent via MailHog.</p>
		<p>Time: `+time.Now().Format("2006-01-02 15:04:05")+`</p>
	`)

	// Configure dialer for MailHog (localhost:1025, no auth)
	d := mail.NewDialer("localhost", 1025, "", "")
	d.Timeout = 10 * time.Second

	// Attempt to send email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
