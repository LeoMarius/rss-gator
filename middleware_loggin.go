package main

import (
	"context"
	"fmt"

	"github.com/LeoMarius/rss-gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Récupérer le nom d'utilisateur actuellement connecté
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("no user logged in: %w", err)
		}

		// Exécuter le handler en lui passant l'utilisateur
		return handler(s, cmd, user)
	}
}
