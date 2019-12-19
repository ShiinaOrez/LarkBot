package main

import (
	"github.com/ShiinaOrez/LarkBot/bot/githubbot/event"
	"github.com/ShiinaOrez/LarkBot/bot/githubbot/trending"
	"github.com/ShiinaOrez/LarkBot/timeTable"
)

var githubBotTimeTable = timeTable.NewTimeTable()

func main() {
	githubBot := event.NewBot("backend")
	githubBotTimeTable.Hour(20).Append(githubBot)

	repoBot := trending.NewBot("go")
	githubBotTimeTable.Hour(10).Append(repoBot)

	githubBotTimeTable.Register()
	githubBotTimeTable.Run()
}
