package githubbot

import "time"

type GithubBot interface {
	Do()
	Run(duration time.Duration)
}
