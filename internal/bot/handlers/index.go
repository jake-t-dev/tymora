package handlers

import "github.com/bwmarrin/discordgo"

func AddHandlers(dg *discordgo.Session) (*discordgo.Session, error) {
	dg.AddHandler(test)
	return dg, nil
}
