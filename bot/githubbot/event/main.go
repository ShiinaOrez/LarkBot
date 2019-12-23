package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	lark "github.com/ShiinaOrez/LarkBot/bot"
	"github.com/ShiinaOrez/LarkBot/bot/githubbot"
	"github.com/ShiinaOrez/LarkBot/conf"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type EventBot struct {
	TargetUserList []string
	Client         *http.Client
	WebHookList    []string
}

type Event struct {
	Type      string `json:"type"`
	Repo      Repo   `json:"repo"`
	CreatedAt string `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type Result struct {
	Name    string
	PushTot int
	RepoTot int
}

type Results struct {
	Data []Result
}

func (r Results) Len() int {
	return len(r.Data)
}

func (r Results) Less(i, j int) bool {
	if r.Data[i].PushTot > r.Data[j].PushTot {
		return true
	} else if r.Data[i].PushTot == r.Data[j].PushTot {
		return r.Data[i].RepoTot > r.Data[j].RepoTot
	}
	return false
}

func (r Results) Swap(i, j int) {
	r.Data[i], r.Data[j] = r.Data[j], r.Data[i]
	return
}

func (bot EventBot) getUserGithubEvents(username string) []Event {
	resp, err := bot.Client.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		log.Println("[Github] [Request]", err.Error())
		return nil
	}
	if resp.Body == nil {
		log.Println("[Github] [Response]", "Response body is nil")
		return nil
	}
	events := make([]Event, 0)
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		log.Println("[Github] [Json]", err.Error())
		return nil
	}
	return events
}

func countEvents(events []Event, today string) (pushTot, repoTot int) {
	repoMap := make(map[string]bool)
	for _, event := range events {
		if strings.HasPrefix(event.CreatedAt, today) {
			if event.Type == "PushEvent" {
				pushTot += 1
				repoMap[event.Repo.Name] = true
			}
		} else {
			break
		}
	}
	repoTot = len(repoMap)
	return
}

func (bot EventBot) Do() {
	failed := 0
	failLocker := sync.Mutex{}

	// actually `today` is yesterday
	today := time.Now().AddDate(0, 0, -1).String()[:10]

	results := Results{Data: make([]Result, 0)}
	resultsLocker := sync.Mutex{}

	waitGroup := sync.WaitGroup{}
	for _, user := range bot.TargetUserList {
		waitGroup.Add(1)
		go func(username string) {
			events := bot.getUserGithubEvents(username)
			if events == nil {
				failLocker.Lock()
				failed += 1
				failLocker.Unlock()
			}
			pushTot, repoTot := countEvents(events, today)

			resultsLocker.Lock()
			results.Data = append(results.Data, Result{
				Name:    username,
				PushTot: pushTot,
				RepoTot: repoTot,
			})
			resultsLocker.Unlock()

			waitGroup.Done()
		}(user)
	}
	waitGroup.Wait()

	sort.Sort(results)
	newMsg := lark.Message{
		Title: today + " | Github每日Push量统计",
		Text: fmt.Sprintf("Tot: %d, Success: %d, fail: %d\n🎉",
			len(bot.TargetUserList),
			len(bot.TargetUserList)-failed,
			failed),
	}
	for _, result := range results.Data {
		newMsg.Text += fmt.Sprintf("[%s] push %d time(s), to %d repo(s).\n",
			result.Name,
			result.PushTot,
			result.RepoTot)
	}
	// fmt.Printf("%v\n", newMsg)
	bs, err := json.Marshal(newMsg)
	if err != nil {
		log.Println("[Github] [Json]", err.Error())
		return
	}
	for _, wh := range bot.WebHookList {
		_, err = bot.Client.Post(wh, "application/json", bytes.NewReader(bs))
		if err != nil {
			log.Println("[Lark] [WebHook]", err.Error())
			return
		}
	}
	return
}

func (bot EventBot) Run(duration time.Duration) {
	time.Sleep(duration)
	log.Println("[Github] [Push] [Bot] [TODO]")
	bot.Do()
	log.Println("[Github] [Push] [Bot] [Done]")
}

func NewBot(groupName string) githubbot.GithubBot {
	return EventBot{
		TargetUserList: conf.C.GroupUsersMap[groupName],
		Client:         &http.Client{},
		WebHookList:    conf.C.Webhooks.Github.Event.Push,
	}
}
