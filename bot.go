package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

const (
	ServerName = "Memes of Our Lives"
	MyName     = "Adjutant"
)

// Welcomer sends a welcome message regarding new members.
func Welcomer(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	msg := fmt.Sprintf("Hi %s, welcome to %s 👏👏👏.", m.User.Username, ServerName)
	s.ChannelMessageSend(AnnouncementsChannelID, msg)
}

// WhoAmI responds to questions about my existence.
func WhoAmI(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	content := strings.ToLower(m.Content)
	if (strings.Contains(content, "wtf") || strings.Contains(content, "what") ||
		strings.Contains(content, "who")) && strings.Contains(content, "is") &&
		strings.Contains(content, strings.ToLower(MyName)) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`
			@%s I am a sentient AI 🤖. I am here to help manage '%s'.
		`, m.Author.Username, ServerName))
	}
}

var (
	ShitPostKeywords = []string{
		"i.4cdn.org",
		"trump",
		"pol ",
	}
)

func ShitPostManagement(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	content := strings.ToLower(m.Content)
	for _, word := range ShitPostKeywords {
		if strings.Contains(content, word) {
			msg := fmt.Sprintf(
				"@%s Y'all best not be shit-postin' round these here parts 🤠.",
				m.Author.Username,
			)
			s.ChannelMessageSend(
				m.ChannelID,
				msg,
			)
			return
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

var (
	AnnouncementsChannelID string
	Token                  string
)

func initConfig() error {
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	cid, ok := (viper.Get("announcements_channel_id")).(string)
	if !ok {
		return errors.New("Error reading announcements_channel_id")
	}
	AnnouncementsChannelID = cid
	tok, ok := (viper.Get("token")).(string)
	if !ok {
		return errors.New("Error reading token")
	}
	Token = tok
	return nil
}

/*func Soundboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	msg := strings.Trim(strings.ToLower(m.Content), " 	")
	switch (msg) {
		case "!wololo"
	}
}*/

func main() {
	checkErr(initConfig())
	discord, err := discordgo.New("Bot " + Token)
	checkErr(err)
	discord.AddHandler(Welcomer)
	discord.AddHandler(ShitPostManagement)
	discord.AddHandler(WhoAmI)
	err = discord.Open()
	checkErr(err)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}
