package main

import "fmt"

// SMS Email

type INotificationFactory interface {
	SendNotification()
	GetSender() Isender
}

type Isender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification Via SMS")

}

func (SMSNotification) GetSender() Isender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
