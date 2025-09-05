package middleware

import (
	"context"
	"job-finder/internal/models"
)

type contextKey string

const userContextKey = contextKey("user")

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	return user, ok
}
