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
	Topics  []*pubsub.Topic
	Context context.Context
}

func main() {
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
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pub/Sub emulator toolkit")
	fmt.Println("------------------------")

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

			app.Topics = append(app.Topics, t)
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
		default:
			fmt.Printf("%q is not a valid input\n", input)
		}
	}
}
