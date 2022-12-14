package command

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/devproje/kuma-engine/utils"
	"github.com/devproje/kuma-engine/utils/emoji"
	"github.com/devproje/plog/log"
)

type GuildCommand struct {
	Commands []Command
	GuildId  string
}

// Build create guild command handler
func (g *GuildCommand) BuildHandler(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, e *discordgo.InteractionCreate) {
		if e.GuildID != g.GuildId || e.Member.User.Bot {
			return
		}

		if e.Interaction.Type == discordgo.InteractionApplicationCommand {
			for _, j := range g.Commands {
				if j.Data.Name == e.ApplicationCommandData().Name {
					debug(e)
					err := j.Execute(s, e)
					if err != nil {
						s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Embeds: []*discordgo.MessageEmbed{
									utils.ErrorEmbed(e.Member.User, emoji.NoEntry, "An error occurred while executing the code"),
								},
							},
						})
						return
					}
				}
			}
		}
	})
}

// AddGuildData add target guild application commands data
func (g *GuildCommand) AddGuildData(session *discordgo.Session) error {
	if g.GuildId == "" {
		return errors.New("you must type spectified guild id")
	}

	for i, j := range g.Commands {
		guild, err := session.State.Guild(g.GuildId)
		if err != nil {
			return err
		}

		log.Infof("Register '%s' guild command %s data (%d/%d)", guild.Name, j.Data.Name, i+1, len(g.Commands))
		_, err = session.ApplicationCommandCreate(session.State.User.ID, g.GuildId, j.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

// DropGuildData delete target guild application commands data
func (g *GuildCommand) DropGuildData(session *discordgo.Session) error {
	commands, err := session.ApplicationCommands(session.State.User.ID, g.GuildId)
	if err != nil {
		return err
	}

	for _, i := range commands {
		guild, err := session.State.Guild(g.GuildId)
		if err != nil {
			return err
		}

		log.Infof("Remove '%s' guild command %s data", guild.Name, i.Name)
		err = session.ApplicationCommandDelete(session.State.User.ID, g.GuildId, i.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
