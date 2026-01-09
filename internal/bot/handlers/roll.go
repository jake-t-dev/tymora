package handlers

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var rollRegex = regexp.MustCompile(`^!roll\s+(\d+)d(\d+)$`)

func roll(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "!roll") {
		return
	}

	matches := rollRegex.FindStringSubmatch(m.Content)
	if len(matches) != 3 {
		s.ChannelMessageSend(m.ChannelID, "Usage: !roll <number_of_dice>d<sides> (e.g., !roll 5d20)")
		return
	}

	numDice, err := strconv.Atoi(matches[1])
	if err != nil || numDice <= 0 {
		s.ChannelMessageSend(m.ChannelID, "Invalid number of dice.")
		return
	}

	sides, err := strconv.Atoi(matches[2])
	if err != nil || sides <= 0 {
		s.ChannelMessageSend(m.ChannelID, "Invalid number of sides.")
		return
	}

	if numDice > 100 || sides > 1000 {
		s.ChannelMessageSend(m.ChannelID, "Limit is 100d1000.")
		return
	}

	var rolls []string
	total := 0
	for i := 0; i < numDice; i++ {
		roll := rand.Intn(sides) + 1
		rolls = append(rolls, strconv.Itoa(roll))
		total += roll
	}

	response := fmt.Sprintf("%s rolled %dd%d: [%s] (Total: %d)",
		m.Author.Mention(), numDice, sides, strings.Join(rolls, ", "), total)

	if len(response) > 2000 {
		response = fmt.Sprintf("%s rolled %dd%d: (Total: %d) (Rolls were too long to display)",
			m.Author.Mention(), numDice, sides, total)
	}

	s.ChannelMessageSend(m.ChannelID, response)
}
