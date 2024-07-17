package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mmcdole/gofeed"
)

func main() {

	sess, err := discordgo.New("Bot <key>")

	if err != nil {
		log.Fatal(err)
	}

	starcitizenChannelId := "your discord channel id"

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Enter a poll duration like 1 s, 1 m, or 1 hr.")
		os.Exit(1)
	}

	pollDuration := args[0]
	var durationUnit time.Duration
	if args[1] == "hr" || args[1] == "h" || args[1] == "H" || args[1] == "HR" {
		durationUnit = time.Hour
	} else if args[1] == "m" || args[1] == "min" || args[1] == "M" || args[1] == "MIN" {
		durationUnit = time.Minute
	} else if args[1] == "s" || args[1] == "sec" || args[1] == "S" || args[1] == "SEC" {
		durationUnit = time.Second
	} else {
		fmt.Println("Second argument should be a time duration unit like s, sec, m, min, h, or hr.")
		os.Exit(1)
	}
	pollDurationInt, _ := strconv.Atoi(pollDuration)

	// Get first item to compare
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://status.robertsspaceindustries.com/index.xml")
	lastLatestItem := feed.Items[0]
	fmt.Println("Latest post:")
	fmt.Println()
	fmt.Println(lastLatestItem.Title, "\n", lastLatestItem.Description)
	fmt.Println()
	fmt.Println()

	for {
		fmt.Println("Waiting ", time.Duration(pollDurationInt)*durationUnit, ". Current time:", time.Now().Format("15:04:05"))
		<-time.After(time.Duration(pollDurationInt) * durationUnit)

		// get next item to compare
		feed, _ := fp.ParseURL("https://status.robertsspaceindustries.com/index.xml")
		latestItem := feed.Items[0]
		if latestItem.PublishedParsed.After(*lastLatestItem.PublishedParsed) {
			post := "New post!!" + "\n" + latestItem.Title + "\n" + latestItem.Description
			post = strip.StripTags(post)
			fmt.Println("New post!!")
			fmt.Println()
			fmt.Println(post)
			fmt.Println()
			fmt.Println("Sending to Discord...")

			sess.ChannelMessageSend(starcitizenChannelId, post)
			fmt.Println("Sent!")
		} else {
			fmt.Println("No new post...")
		}
		lastLatestItem = latestItem
	}
}
