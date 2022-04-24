package roles

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var reactChannelId = "965312058079985774"
var reactMessageId = "967877389189263390"

type ReactionRole struct {
	Name   string
	RoleId string
	Emoji  string
}

var reactionRoles []ReactionRole

func InitReactionRoles(session *discordgo.Session) {
	reactionRoles = append(reactionRoles, createReactionRole("Apex Legends", "965313077262290956", "🅰️"))
	reactionRoles = append(reactionRoles, createReactionRole("League of Legends", "965313152789143602", "🇱"))
	reactionRoles = append(reactionRoles, createReactionRole("Rust", "965313241259593748", "🇷"))

	//initReactChannel(session)
}

func createReactionRole(name string, roleId string, emoji string) ReactionRole {
	role := ReactionRole{
		Name:   name,
		RoleId: roleId,
		Emoji:  emoji,
	}
	return role
}

func initReactChannel(session *discordgo.Session) {
	channel, _ := session.Channel(reactChannelId)
	if channel.MessageCount <= 0 {
		message, _ := session.ChannelMessageSend(reactChannelId, "React to the role that suits you accordingly."+
			"\n\nGames:"+
			"\n\n:a: for Apex Legends"+
			"\n\n:regional_indicator_l: for League of Legends"+
			"\n\n:regional_indicator_r: for Rust")

		for _, role := range reactionRoles {
			session.MessageReactionAdd(reactChannelId, message.ID, role.Emoji)
		}

		reactMessageId = message.ID
	}
}

func MessageReactAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.UserID == s.State.User.ID {
		return
	}
	if m.MessageID == reactMessageId {
		for _, role := range reactionRoles {
			if m.Emoji.Name == role.Emoji {
				s.GuildMemberRoleAdd(m.GuildID, m.UserID, role.RoleId)
				log.Printf("Added role %s to %s\n", role.Name, m.UserID)
			}
		}

	}
}

func MessageReactRemove(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	if m.UserID == s.State.User.ID {
		return
	}
	if m.MessageID == reactMessageId {
		for _, role := range reactionRoles {
			if m.Emoji.Name == role.Emoji {
				s.GuildMemberRoleRemove(m.GuildID, m.UserID, role.RoleId)
				log.Printf("Removed role %s from %s\n", role.Name, m.UserID)
			}
		}
	}
}
