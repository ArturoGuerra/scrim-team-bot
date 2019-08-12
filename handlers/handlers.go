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

    s.ChannelMessageSend(m.ChannelID, "First Team is: " + MentionString(leftTeam))
    s.ChannelMessageSend(m.ChannelID, "Second Team is: " + MentionString(rightTeam))
}

func MentionString (team [] *discordgo.User) string {
    var result bytes.Buffer

    for _, user := range team {
        result.WriteString("<@" + user.ID +"> ")
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
