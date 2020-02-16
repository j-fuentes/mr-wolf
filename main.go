package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/j-fuentes/mr-wolf/internal/config"
	"github.com/j-fuentes/mr-wolf/pkg/auth"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	token, err := readAPIToken()
	if err != nil {
		log.Fatalf("Error reading API token: %+v", err)
	}

	cfg, err := config.Read(expandPath("./config.yaml"))
	if err != nil {
		log.Fatalf("Error reading config file: %+v", err)
	}

	a, err := auth.NewAuth(cfg.AllowedUsers)
	if err != nil {
		log.Fatalf("Error creating auth: %+v", err)
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Cannot start Telegram bot: %+v", err)
	}

	commands := []string{"/hello", "/help"}

	// hello
	bot.Handle(commands[0], authorizedOnly(a, bot, hello))
	// help
	bot.Handle(commands[1], authorizedOnly(a, bot, func(bot *tb.Bot, m *tb.Message) {
		bot.Send(m.Sender, fmt.Sprintf("This is the list of available commands:\n%s", strings.Join(commands, "\n")))
	}))

	log.Println("Starting bot...")
	bot.Start()
}

type handler func(bot *tb.Bot, m *tb.Message)

func hello(bot *tb.Bot, m *tb.Message) {
	bot.Send(m.Sender, fmt.Sprintf("hello Mr %s", m.Sender.FirstName))
}

func deny(bot *tb.Bot, m *tb.Message) {
	bot.Send(m.Sender, "you are not welcomed here")
}

func authorizedOnly(a *auth.Auth, bot *tb.Bot, fn handler) func(m *tb.Message) {
	return func(m *tb.Message) {
		msg := "-> received %q from %d. allowed: %v \n"
		if a.UserAllowed(m.Sender) {
			log.Printf(msg, m.Text, m.Sender.ID, true)
			fn(bot, m)
		} else {
			log.Printf(msg, m.Text, m.Sender.ID, false)
			deny(bot, m)
		}
	}
}

func readAPIToken() (string, error) {
	dat, err := ioutil.ReadFile(expandPath("./.api_token"))
	if err != nil {
		return "", nil
	}

	return strings.Trim(string(dat), "\t \n"), nil
}

func expandPath(path string) string {
	if len(path) == 0 {
		log.Fatalf("Unexpected empty path")
	}

	if path[:2] == "~/" {
		u, err := user.Current()
		if err != nil {
			log.Fatalf("Cannot expand path because it was impossible to get current user: %+v", err)
		}
		return filepath.Join(u.HomeDir, path[2:])
	}

	return path
}
