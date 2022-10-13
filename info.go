package kuma

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/devproje/kuma-engine/command"
	"github.com/devproje/kuma-engine/log"
	"github.com/devproje/kuma-engine/utils"
	"github.com/devproje/kuma-engine/utils/emoji"
)

const logo = "https://github.com/devproje/kuma-engine/raw/master/assets/kuma-engine-logo.png"

var kumaInfo = command.Command{
	Data: &discordgo.ApplicationCommand{
		Name:        "kumainfo",
		Description: "KumaEngine system information",
	},
	Usage: "/kumainfo",
	Execute: func(session *discordgo.Session, event *discordgo.InteractionCreate) {
		embed := utils.Embed{
			Title:       fmt.Sprintf("%s **KumaInfo**", emoji.Dart),
			Description: "KumaEngine system information",
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    logo,
				Width:  512,
				Height: 512,
			},
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

		data := &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		}

		if infoEphemeral {
			data.Flags = 1 << 6
		}

		_ = session.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: data,
		})
	},
}

func (k *KumaEngine) SetEphemeralKumaInfo(e bool) {
	if !engineStarted {
		infoEphemeral = e
		return
	}

	log.Logger.Errorln("You cannot use this method, Please try to engine enabled before")
}

func (k *KumaEngine) DisableKumaInfo() {
	if !engineStarted {
		go func() {
			count := 0
			for !engineStarted {
				time.Sleep(time.Second * 1)
				count++
			}

			if engineStarted {
				err := command.DropDataManual(k.Session, kumaInfo)
				if err != nil {
					log.Logger.Errorln(err)
				}

				log.Logger.Infof("KumaInfo disabled: %ds", count)
			}
		}()

		command.DropCommand(kumaInfo)
		return
	}

	log.Logger.Errorln("You cannot use this method, Please try to engine enabled before")
}
