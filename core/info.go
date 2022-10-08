package core

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/devproje/kuma-engine/command"
	"github.com/devproje/kuma-engine/emoji"
	"github.com/devproje/kuma-engine/utils"
	"math/rand"
	"os"
	"runtime"
)

var KumaInfo = command.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "kumainfo",
		Description: "KumaEngine system information",
	},
	Usage: "/kumainfo",
	Execute: func(session *discordgo.Session, event *discordgo.InteractionCreate) {
		embed := utils.Embed{
			Title:       fmt.Sprintf("%s **KumaInfo**", emoji.Dart),
			Description: "KumaEngine system information",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   fmt.Sprintf("%s **ENGINE VERSION**", emoji.ElectricPlug),
					Value:  fmt.Sprintf("`%s`", KUMA_ENGINE_VERSION),
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%s **GO VERSION**", emoji.PageFacingUp),
					Value:  fmt.Sprintf("`%s`", runtime.Version()),
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%s **API LATENCY**", emoji.PingPong),
					Value:  fmt.Sprintf("`%dms`", session.HeartbeatLatency().Milliseconds()),
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%s **OS**", emoji.Desktop),
					Value:  fmt.Sprintf("`%s/%s`", runtime.GOOS, runtime.GOARCH),
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%s **BOT SERVERS**", emoji.Satellite),
					Value:  fmt.Sprintf("`%d`", len(session.State.Guilds)),
					Inline: true,
				},
				{
					Name:   fmt.Sprintf("%s **SYSTEM PID**", emoji.FileFolder),
					Value:  fmt.Sprintf("`%d`", os.Getpid()),
					Inline: true,
				},
			},
			Color: rand.Intn(0xFFFFFF),
			Footer: &discordgo.MessageEmbedFooter{
				Text:    event.Member.User.String(),
				IconURL: event.Member.User.AvatarURL("512x512"),
			},
		}.Build()

		_ = session.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
				Flags:  1 << 6,
			},
		})
	},
}