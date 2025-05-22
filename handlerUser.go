package main

import (
	"context"
	"fmt"
	"time"

	"github.com/arishimam/blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Not enough arguments were provided")
	}
	username := cmd.Args[0]
	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Couldn't set current user")
	}

	fmt.Printf("User has been set to %v\n", username)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Not enough arguments were provided")
	}

	username := cmd.Args[0]
	err := s.Db.GetUser(context.Background(), username)

	if err == nil {

	}

	userParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username
	}

	s.Db.CreateUser(context.Background(), userParams)


	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Couldn't set current user")
	}

	fmt.Printf("User has been set to %v\n", username)

	return nil

}
