package main

import (
	"os"
)

var (
	Token     string
	JuanBotID string
)

func init() {
	Token = os.Getenv("TOKEN")
	JuanBotID = os.Getenv("JUANBOT_ID")
}
