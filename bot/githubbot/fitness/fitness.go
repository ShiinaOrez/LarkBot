package fitness

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	lark "github.com/ShiinaOrez/LarkBot/bot"
	"github.com/ShiinaOrez/LarkBot/conf"
)

type FitnessBot struct {
	Client      *http.Client
	WebHookList []string
}

func (bot FitnessBot) Do() {
	today := time.Now().AddDate(0, 0, -1).String()[:10]
	newMsg := lark.Message{
		Title: fmt.Sprintf("%s | 今日头疼健身提醒", today),
		Text:  "今天你健身了吗？打工人！",
	}
	bs, err := json.Marshal(newMsg)
	if err != nil {
		log.Println("[Github] [Json]", err.Error())
		return
	}
	log.Println("[Fitness] [Sending] ... ")
	for _, wh := range bot.WebHookList {
		log.Println("[Fitness] [SendTo]", wh)
		_, err := bot.Client.Post(wh, "application/json", bytes.NewReader(bs))
		if err != nil {
			log.Println("[Lark] [WebHook]", err.Error())
			return
		}
	}
	return
}

func (bot FitnessBot) Run(duration time.Duration) {
	time.Sleep(duration)
	log.Println("[Fitness] [Bot] [TODO]")
	bot.Do()
	log.Println("[Fitness] [Push] [Bot] [Done]")
}

func NewBot() FitnessBot {
	return FitnessBot{
		Client:      &http.Client{},
		WebHookList: conf.C.Webhooks.Github.Fitness,
	}
}
