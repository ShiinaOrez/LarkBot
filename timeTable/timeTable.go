package timeTable

import (
	"github.com/ShiinaOrez/LarkBot/bot/githubbot"
	"log"
	"time"
)

type Worker struct {
	Bot githubbot.GithubBot
	Fd  time.Duration
}

type Workers []Worker

type TimeTable struct {
	Map     map[int]*githubbot.GBS
	Workers *Workers
}

func getDuration(startHour int) time.Duration {
	timeNow := time.Now()
	// 计算当前时间和每个下一个需要进行任务的时间的差值
	dHour := time.Hour * time.Duration(startHour-timeNow.Hour()-1)
	dMinute := time.Minute * time.Duration(60-timeNow.Minute()-1)
	dSecond := time.Second * time.Duration(60-timeNow.Second())
	// 当前时间为整分钟时，second为0
	if dSecond == 60*time.Second {
		// 将多减去的一分钟加回来
		dMinute += time.Minute
		dSecond -= 60 * time.Second
		// 当前时间为整小时时
		if dMinute == 60*time.Minute {
			// 将多减的一个小时加回来
			dHour += time.Hour
			dMinute -= 60 * time.Minute
		}
	}
	if dHour < 0 {
		dHour += 24 * time.Hour
	}
	log.Printf("Bot will firstly run after %d hours %d mins %d seconds\n", dHour/time.Hour, dMinute/time.Minute, dSecond/time.Second)
	return dHour + dMinute + dSecond
}

func (w Worker) Run() {
	w.Bot.Run(w.Fd)
	for {
		w.Bot.Run(time.Hour * 24)
	}
}

func (ws *Workers) Append(worker Worker) {
	*ws = append(*ws, worker)
}

func NewTimeTable() TimeTable {
	return TimeTable{Map: make(map[int]*githubbot.GBS), Workers: &Workers{}}
}

func (tt *TimeTable) Append(bot githubbot.GithubBot, startHour int) {
	if _, ok := tt.Map[startHour]; !ok {
		tt.Map[startHour] = &githubbot.GBS{bot}
	} else {
		tt.Map[startHour].Append(bot)
	}
}

func (tt TimeTable) Register() {
	log.Printf("%v\n", tt.Map)
	for startHour, gbs := range tt.Map {
		duration := getDuration(startHour)
		for _, githubBot := range *gbs {
			log.Println("[TimeTable] [Register] [Bot]")
			tt.Workers.Append(Worker{
				Bot: githubBot,
				Fd:  duration,
			})
		}
	}
}

func (tt TimeTable) Run() {
	for _, worker := range *tt.Workers {
		log.Println("[Worker] [Start]")
		go worker.Run()
	}
	// 睡一年
	time.Sleep(time.Hour * 24 * 30 * 12)
}
