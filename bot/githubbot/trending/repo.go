package trending

import (
	"bytes"
	"encoding/json"
	"fmt"
	lark "github.com/ShiinaOrez/LarkBot/bot"
	"github.com/ShiinaOrez/LarkBot/constvar"
	"log"
	"net/http"
	"time"
)

type RepoBot struct {
	Language    string
	Client      *http.Client
	WebHookList []string
}

type Repo struct {
	Author string `json:"author"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Stars  int    `json:"stars"`
	Forks  int    `json:"forks"`
	DStars int    `json:"currentPeriodStars"`
}

type Repos struct {
	Data []Repo
}

func (r Repos) Len() int {
	return len(r.Data)
}

func (r Repos) Less(i, j int) bool {
	return r.Data[i].DStars > r.Data[j].DStars
}

func (r Repos) Swap(i, j int) {
	r.Data[i], r.Data[j] = r.Data[j], r.Data[i]
	return
}

func (bot RepoBot) getTop5Repos() []Repo {
	resp, err := bot.Client.Get(
		fmt.Sprintf("https://github-trending-api.now.sh/repositories?language=%s&since=daily", bot.Language))
	if err != nil {
		log.Println("[Github] [Request]", err.Error())
		return nil
	}
	if resp.Body == nil {
		log.Println("[Github] [Response]", "response body is nil.")
		return nil
	}
	repos := Repos{Data: make([]Repo, 0)}
	err = json.NewDecoder(resp.Body).Decode(&repos.Data)
	if err != nil {
		log.Println("[Github] [Json]", err.Error())
		return nil
	}
	// sort.Sort(repos)
	if len(repos.Data) > 5 {
		repos.Data = repos.Data[:5]
	}
	return repos.Data
}

func (bot RepoBot) Do() {
	repos := bot.getTop5Repos()
	if repos == nil {
		log.Printf("[Github] [Repo] [%s] Get Top 5 Repo Failed.", bot.Language)
		return
	}
	today := time.Now().AddDate(0, 0, -1).String()[:10]
	newMsg := lark.Message{
		Title: fmt.Sprintf("%s | Github昨日%s语言TOP%d趋势库", today, bot.Language, len(repos)),
		Text:  "今天本BOT也运行正常\n",
	}
	for index, repo := range repos {
		appendText := fmt.Sprintf(
			"%s[TOP%d]: %s\n    作者:%s\n    增长Star数:%d\n    总Star数:%d\n    总Fork数:%d\n    传送门:%s\n",
			constvar.NumberToEmoji[index+1],
			index+1,
			repo.Name,
			repo.Author,
			repo.DStars,
			repo.Stars,
			repo.Forks,
			repo.URL,
		)
		newMsg.Text += appendText
	}
	bs, err := json.Marshal(newMsg)
	if err != nil {
		log.Println("[Github] [Json]", err.Error())
		return
	}
	for _, wh := range bot.WebHookList {
		_, err := bot.Client.Post(wh, "application/json", bytes.NewReader(bs))
		if err != nil {
			log.Println("[Lark] [WebHook]", err.Error())
			return
		}
	}
	// fmt.Printf("%v\n", newMsg)
	return
}

func (bot RepoBot) Run(duration time.Duration) {
	time.Sleep(duration)
	log.Println("[Github] [Repo] [Bot] [TODO]")
	bot.Do()
	log.Println("[Github] [Push] [Bot] [Done]")
}

func NewBot(language string) RepoBot {
	return RepoBot{
		Language:    language,
		Client:      &http.Client{},
		WebHookList: constvar.WebHooks["github"][language],
	}
}
