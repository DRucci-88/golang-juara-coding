package main

import (
	"errors"
	"fmt"
)

/*
Latihan 2: Sistem Notifikasi Penyiaran E-Commerce (Slice of Interfaces)
*/

type NotificationChannel interface {
	Send(userID string, message string) error
	GetChannelName() string
}

type EmailNotifier struct {
	EmailAddress string
}

func (notifier EmailNotifier) GetChannelName() string { return "EmailNotifier" }
func (notifier EmailNotifier) Send(userID string, message string) error {
	fmt.Printf("EmailAddress [%s]", notifier.EmailAddress)
	fmt.Printf("UserID [%s] Message [%s]\n", userID, message)
	return nil
}

type SmsNotifier struct {
	PhoneNumber string
}

func (notifier SmsNotifier) GetChannelName() string { return "SmsNotifier" }
func (notifier SmsNotifier) Send(userID string, message string) error {
	fmt.Printf("PhoneNumber [%s]", notifier.PhoneNumber)
	fmt.Printf("UserID [%s] Message [%s]\n", userID, message)
	return nil
}

type PushNotifier struct {
	DeviceToken string
}

func (notifier PushNotifier) GetChannelName() string { return "PushNotifier" }
func (notifier PushNotifier) Send(userID string, message string) error {
	fmt.Printf("DeviceToken [%s]", notifier.DeviceToken)
	fmt.Printf("UserID [%s] Message [%s]\n", userID, message)
	return nil
}

type NotificationDispatcher struct {
	Channels []NotificationChannel
}

func (dispatcher *NotificationDispatcher) AddChannel(channel NotificationChannel) error {
	if dispatcher == nil || dispatcher.Channels == nil {
		return errors.New("Dispatcher not initialized")
	}

	dispatcher.Channels = append(dispatcher.Channels, channel)
	// fmt.Println(dispatcher.Channels)
	return nil
}

func (dispatcher *NotificationDispatcher) BroadcastNotification(
	userID string,
	msg string,
) error {
	for _, channel := range dispatcher.Channels {
		fmt.Println(channel)
		fmt.Printf("Channel Name [%s]\n", channel.GetChannelName())
		if err := channel.Send(userID, msg); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	email := EmailNotifier{EmailAddress: "le@gmail.com"}
	sms := SmsNotifier{PhoneNumber: "08123456789"}
	push := PushNotifier{DeviceToken: "805020"}

	dispatcher := NotificationDispatcher{Channels: []NotificationChannel{}}
	if err := dispatcher.AddChannel(email); err != nil {
		fmt.Println(err)
	}
	if err := dispatcher.AddChannel(sms); err != nil {
		fmt.Println(err)
	}
	if err := dispatcher.AddChannel(push); err != nil {
		fmt.Println(err)
	}

	dispatcher.BroadcastNotification("User-1", "Hello World")
}
