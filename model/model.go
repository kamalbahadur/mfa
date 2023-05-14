package model

type User struct {
	id        string
	firstName string
	lastName  string
}

type UserMFASharedSecret struct {
	userId       string
	sharedSecret string
}
