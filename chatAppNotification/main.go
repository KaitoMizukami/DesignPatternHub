/*

問題: チャットアプリの通知機能の実装

あなたは、チャットアプリの通知機能をGo言語で実装する必要があります。
このチャットアプリは、ユーザーがメッセージを送信すると、他のユーザーに対して通知が送信されるものです。
通知は、アプリ内の通知バナーやメールなど、複数の通知方法をサポートします。

このシステムの要件は以下の通りです：

ユーザーがメッセージを送信すると、通知を受け取る対象のユーザーに通知が送信される。
通知は複数の方法で送信されることがあり、将来的に新しい通知方法を追加することができるようにする。
各通知方法は個別に設定可能であり、ユーザーは通知を有効/無効にしたり、通知方法を選択したりできる。
上記の要件を満たすために、デザインパターンを使用してシステムを実装してください。

*/

package main

import "fmt"

type Message struct {
	From    string
	To      string
	Content string
}

type NotificationStrategy interface {
	SendNotification(Message)
}

type MessageSender struct {
	observers     []NotificationStrategy
	messageBuffer []Message
}

func NewMessageSender() *MessageSender {
	return &MessageSender{
		observers:     []NotificationStrategy{},
		messageBuffer: []Message{},
	}
}

func (ms *MessageSender) AddObserver(observer NotificationStrategy) {
	ms.observers = append(ms.observers, observer)
}

func (ms *MessageSender) SendMessage(message Message) {
	ms.messageBuffer = append(ms.messageBuffer, message)
	ms.NotifyObservers()
}

func (ms *MessageSender) NotifyObservers() {
	for _, observer := range ms.observers {
		for _, message := range ms.messageBuffer {
			observer.SendNotification(message)
		}
	}
	ms.messageBuffer = []Message{}
}

type BannerNotification struct{}

func (bn *BannerNotification) SendNotification(message Message) {
	fmt.Printf("[通知バナー] 送信者: %s, 受信者: %s, メッセージ: %s\n", message.From, message.To, message.Content)
}

// メール通知
type EmailNotification struct{}

func (en *EmailNotification) SendNotification(message Message) {
	fmt.Printf("[メール通知] 送信者: %s, 受信者: %s, メッセージ: %s\n", message.From, message.To, message.Content)
}

func NewNotificationStrategy(notificationType string) NotificationStrategy {
	switch notificationType {
	case "banner":
		return &BannerNotification{}
	case "email":
		return &EmailNotification{}
	default:
		return nil
	}
}

func main() {
	// メッセージ送信者を作成
	messageSender := NewMessageSender()

	// 通知方法を登録
	messageSender.AddObserver(NewNotificationStrategy("banner"))
	messageSender.AddObserver(NewNotificationStrategy("email"))

	// メッセージを送信
	message1 := Message{From: "UserA", To: "UserB", Content: "こんにちは！"}
	message2 := Message{From: "UserC", To: "UserA", Content: "おはよう！"}
	messageSender.SendMessage(message1)
	messageSender.SendMessage(message2)
}
