package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/P-H-Pancholi/Golang-Projects/gator/internal/config"
	"github.com/P-H-Pancholi/Golang-Projects/gator/internal/database"
)

type state struct {
	db *database.Queries
	c  *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	list_of_commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list_of_commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.list_of_commands[cmd.name]
	if !ok {
		return fmt.Errorf("command does not exists, please register it")
	}
	if err := f(s, cmd); err != nil {
		return err
	}
	return nil
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rss := RSSFeed{}

	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, err
	}

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for _, item := range rss.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
	return &rss, nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("expected one argument, got zero")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == sql.ErrNoRows {
		return fmt.Errorf("no user with given username")
	}
	if err := s.c.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("The user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("expected one argument, got zero")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		return fmt.Errorf("user already exists")
	}
	if err != sql.ErrNoRows {
		return err
	}
	arg := database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	fmt.Printf("%s user is created with id %d\n", user.Name, user.ID)
	if err := s.c.SetUser(cmd.args[0]); err != nil {
		return err
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUser(context.Background()); err != nil {
		return err
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user == s.c.User {
			fmt.Printf("%s (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Printf("RSSFeed: %+v\n", *rss)
	return nil
}

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	db, err := sql.Open("postgres", c.DbURL)
	dbQueries := database.New(db)
	s := state{
		c:  &c,
		db: dbQueries,
	}

	comm := commands{
		make(map[string]func(*state, command) error),
	}
	comm.register("login", handlerLogin)
	comm.register("register", handlerRegister)
	comm.register("reset", handlerReset)
	comm.register("users", handlerUsers)
	comm.register("agg", handlerAgg)
	if len(os.Args) < 2 {
		fmt.Println("Expected more than 1 args")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := comm.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
