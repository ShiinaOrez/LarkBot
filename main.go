package main

import (
	"github.com/ShiinaOrez/LarkBot/bot/githubbot/event"
	"github.com/ShiinaOrez/LarkBot/bot/githubbot/trending"
	"time"
)

func main() {
	githubBot := event.NewBot("backend")
	githubBot.Run(time.Duration(time.Hour * 24))

	repoBot := trending.NewBot("go")
	repoBot.Run(time.Duration(time.Hour * 24))
	repoBot.Do()
}
