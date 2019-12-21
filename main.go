package main

import (
	"github.com/ShiinaOrez/LarkBot/bot/githubbot/trending"
	"github.com/ShiinaOrez/LarkBot/timeTable"
)

var githubBotTimeTable = timeTable.NewTimeTable()

func main() {
	// githubBot := event.NewBot("backend")
	// githubBotTimeTable.Append(githubBot, 20)

	repoBot := trending.NewBot("go")
	githubBotTimeTable.Append(repoBot, 10)

	githubBotTimeTable.Run()
	defer githubBotTimeTable.Close()
}
