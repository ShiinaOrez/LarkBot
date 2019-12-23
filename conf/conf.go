package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Event struct {
	Push []string `yaml:"push"`
}

type Github struct {
	Trending map[string][]string `yaml:"trending"`
	Event    Event               `yaml:"event"`
}

type Webhooks struct {
	Github    Github `yaml:"github"`
	WorkBench string `yaml:"workbench"`
}

type Config struct {
	Webhooks      Webhooks            `yaml:"webhooks"`
	NumberToEmoji map[int]string      `yaml:"number_to_emoji"`
	GroupUsersMap map[string][]string `yaml:"group_users_map"`
}

var C Config

func init() {
	content, _ := ioutil.ReadFile("./conf/conf.yaml")
	err := yaml.Unmarshal(content, &C)
	if err != nil {
		log.Println("[config] [yaml]", err.Error())
	}
	fmt.Printf("%v\n", C)
}
