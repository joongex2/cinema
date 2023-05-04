package main

import (
	"fmt"

	"github.com/joongex2/cinema/movie"
	"github.com/joongex2/cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
