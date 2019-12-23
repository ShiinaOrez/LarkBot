package main

import (
	"github.com/ShiinaOrez/LarkBot/bot/githubbot/trending"
	_ "github.com/ShiinaOrez/LarkBot/conf"
	"github.com/ShiinaOrez/LarkBot/timeTable"
)

var githubBotTimeTable = timeTable.NewTimeTable()

func main() {
	// githubBot := event.NewBot("backend")
	// githubBotTimeTable.Append(githubBot, 20)

	goRepoBot := trending.NewBot("go")
	githubBotTimeTable.Append(goRepoBot, 10)

	javaRepoBot := trending.NewBot("java")
	githubBotTimeTable.Append(javaRepoBot, 10)

	kotRepoBot := trending.NewBot("kotlin")
	githubBotTimeTable.Append(kotRepoBot, 10)

	githubBotTimeTable.Run()
	defer githubBotTimeTable.Close()
}
