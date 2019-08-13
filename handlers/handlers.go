package handlers

import (
    "bytes"
    "time"
    "math/rand"
    "strings"
    "github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }

    if strings.HasPrefix(m.Content, "!twos") {

        HandleShuffle(s, m, 4)
    } else if strings.HasPrefix(m.Content, "!threes") {
        HandleShuffle(s, m, 6)
    } else if strings.HasPrefix(m.Content, "!fours") {
        HandleShuffle(s, m, 8)
    }
}


func HandleShuffle(s *discordgo.Session, m *discordgo.MessageCreate, a int) {
    if len(m.Mentions) != a {
        return
    }

    users := Shuffle(m.Mentions)

    Idx := a / 2

    leftTeam := users[Idx:]
    rightTeam := users[:Idx]

    var desc string

    switch Idx {
    case 4:
        desc = "Teams of four"
    case 3:
        desc = "Teams of three"
    case 2:
        desc = "Teams of two"
    }

    embed := makeEmbed(desc, leftTeam, rightTeam)

    s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func MentionString (team [] *discordgo.User) string {
    var result bytes.Buffer

    for _, user := range team {
        result.WriteString("<@" + user.ID +"> (")
        result.WriteString(user.Username + "#" + user.Discriminator + ")\n")
    }

    return result.String()
}

func Shuffle (users [] *discordgo.User) ([] *discordgo.User) {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    uarray := make([] *discordgo.User, len(users))
    perm := r.Perm(len(users))

    for i, randIdx := range perm {
        uarray[i] = users[randIdx]
    }

    return uarray
}


func makeEmbed(desc string, team1 [] *discordgo.User, team2 [] *discordgo.User) *discordgo.MessageEmbed {
    fields := []*discordgo.MessageEmbedField{
        &discordgo.MessageEmbedField{
            Name: "Team One:",
            Value: MentionString(team1),
            Inline: false,
        },
        &discordgo.MessageEmbedField{
            Name: "Team Two:",
            Value: MentionString(team2),
            Inline: false,
        },
    }



    embed := &discordgo.MessageEmbed{
        Author: &discordgo.MessageEmbedAuthor{},
        Color: 0x0066ff,
        Description: desc,
        Fields: fields,
        Timestamp: time.Now().Format(time.RFC3339),
        Title: "Scrim Team Generator",
    }

    return embed
}
