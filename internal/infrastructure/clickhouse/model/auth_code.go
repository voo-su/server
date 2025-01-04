package model

type AuthCode struct {
	Email        string
	Code         string
	Token        string
	ErrorMessage string
}
