package main

import (
    "github.com/bwmarrin/discordgo"
    "os"
    "os/signal"
    "syscall"
    "fmt"
    "./handlers"
)


func main () {
    token := os.Getenv("TOKEN")
    dg, err := discordgo.New("Bot " + token)

    if err != nil {
        fmt.Println(err)
        return
    }

    dg.AddHandler(handlers.MessageCreate)

    err = dg.Open()
    if err != nil {
        fmt.Println(err)
        return
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    dg.Close()
}
