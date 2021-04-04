package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

type Config struct {
	Client  *pubsub.Client
	Topics  map[string]*pubsub.Topic
	Context context.Context
}

func main() {
	if len(os.Args) < 1 {
		log.Fatal("You must pass the emulator's projectID as argument")
	} else if len(os.Args) > 2 {
		log.Fatal("You can pass only the projectID as argument")
	}
	projectID := os.Args[1]

	ctx := context.Background()

	c, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Println(err)
		return
	}

	app := Config{
		Client:  c,
		Context: ctx,
		Topics:  map[string]*pubsub.Topic{},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pub/Sub emulator toolkit")
	fmt.Println("------------------------")
	fmt.Println("[!] newtopic <topic-name>")
	fmt.Println("[!] publish <topic> <data> <attributes> ('data' field can have only oneliners)")
	fmt.Println("[!] showtopics")
	fmt.Println("[!] exit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		commands := strings.Split(input, " ")
		switch commands[0] {
		case "newtopic":
			topicID := commands[1]
			t, err := app.Client.CreateTopic(ctx, topicID)
			if err != nil {
				fmt.Println(err)
			}

			app.Topics[topicID] = t
			fmt.Println(t.String())
		case "showtopics":
			it := app.Client.Topics(app.Context)
			for {
				t, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Println(err)
					continue
				}
				fmt.Println(t)
			}
		case "publish":
			topic := app.Topics[commands[1]]
			res := topic.Publish(app.Context, &pubsub.Message{
				Data: []byte(commands[2]),
				Attributes: map[string]string{
					"event": commands[3],
				},
			})

			msgID, err := res.Get(ctx)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("Message published with ID: %q\n", msgID)
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%q is not a valid input\n", input)
		}
	}
}
