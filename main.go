package main

import (
	"Gator/internal"
	"Gator/internal/command"
	"Gator/internal/config"
	"Gator/internal/database"
	"Gator/internal/state"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config - %v", err)
		os.Exit(1)
	}

	appState := state.New(&conf)
	db, err := sql.Open("postgres", conf.DBUrl)
	if err != nil {
		fmt.Printf("Error opening database - %v", err)
		os.Exit(1)
	}
	defer db.Close()
	dbQueries := database.New(db)
	appState.DB = dbQueries

	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1)
	}
	var commands command.Commands
	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerUsers)
	commands.Register("agg", command.HandlerAgg)
	commands.Register("addfeed", internal.MiddlewareLoggedIN(command.HandlerAddFeed))
	commands.Register("feeds", command.HandlerGetAllFeeds)
	commands.Register("follow", internal.MiddlewareLoggedIN(command.HandlerFollow))
	commands.Register("following", internal.MiddlewareLoggedIN(command.HandlerListFeedFollows))
	commands.Register("unfollow", internal.MiddlewareLoggedIN(command.HandlerUnfollow))
	commands.Register("browse", internal.MiddlewareLoggedIN(command.HandlerBrowse))

	cmd := command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(appState, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}
