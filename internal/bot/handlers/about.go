package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func about(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!about" {
		msg := "**Tymora Dice Bot Commands:**\n" +
			"- `!roll <num>d<sides>[+/-modifier]` (alias: `!r`): Rolls dice (e.g., `!roll 5d20+10`).\n" +
			"- `!test`: Checks if the bot is online.\n" +
			"- `!about`: Shows this command list."
		s.ChannelMessageSend(m.ChannelID, msg)
	}
}
