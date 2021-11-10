package dto

import "time"

type AgentLink struct {
	ID           string
	Port         int
	UserHome     string // the user-home-dir
	PublicKey    string // the user-public-key-file
	PublicKeySum string // sha256sum
	Timeout      int
	Online       bool
}

type AgentMessage struct {
	ID      int
	Time    time.Time
	Name    string
	Title   string
	Content string
}
