package models

type Bridge struct {
	FromChatID   string `json:"first_chat_id" gorm:"primaryKey"`
	FirstURL     string `json:"first_url"`
	SecondChatID string `json:"second_chat_id" gorm:"primaryKey"`
	SecondURL    string `json:"second_url"`
}