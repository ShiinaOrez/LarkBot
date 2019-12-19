package githubbot

import "time"

type GithubBot interface {
	Do()
	Run(duration time.Duration)
}

type GBS []GithubBot

func (gbs GBS) Append(bot GithubBot) {
	if gbs == nil {
		gbs = GBS{}
	}
	gbs = append(gbs, bot)
}
