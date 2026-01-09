package handlers

import "github.com/bwmarrin/discordgo"

func test(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!test" {

		s.ChannelMessageSend(m.ChannelID, "I am live! See my repository at https://github.com/jake-t-dev/tymora")
	}
}
