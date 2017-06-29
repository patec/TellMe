package main

import (
    "github.com/SlyMarbo/rss"
    "fmt"
    "os"
)

func main() {
    feed, err := rss.Fetch("http://rss.nytimes.com/services/xml/rss/nyt/US.xml")
    if err != nil {
        fmt.Printf("coudnt fetch feed\n")
        os.Exit(1)
    }

    fmt.Printf("%v", feed)

    var first_feed = feed.Items[0]
    fmt.Printf("first title is: %v\n", first_feed.Title)
    fmt.Printf("content: %v\n", first_feed.Link)
}


