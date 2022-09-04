package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"nhooyr.io/websocket"
)

func main() {
	log.Println(Token)
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}
	s.AddHandler(messageCreate)

	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"pi00:8000"},
		})
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close(websocket.StatusInternalError, "connection unexpectedly died")

		for {
			ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
			defer cancel()
			fmt.Println("Trying to write")
			conn.Write(ctx, websocket.MessageText, []byte("test"))
			time.Sleep(5 * time.Second)
		}
	})

	httpServ := &http.Server{
		Addr:    ":5000",
		Handler: handler,
	}

	go func() {
		fmt.Println("Serving...")
		if err := httpServ.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_ = httpServ.Shutdown(ctx)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == JuanBotID {
		if strings.HasPrefix(m.Content, "Bryant the type of guy to") {
			log.Println(fmt.Sprintf("Juan Said a Newtype in %s", m.ChannelID))
		}
	}
}
