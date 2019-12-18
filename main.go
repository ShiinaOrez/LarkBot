package main

import (
	"github.com/ShiinaOrez/LarkBot/githubbot"
	"time"
)

func main() {
	githubBot := githubbot.NewBot("backend")

	d := time.Duration(time.Hour * 24)
	t := time.NewTicker(d)
	defer t.Stop()

	for {
		<-t.C
		githubBot.Run()
	}
}
