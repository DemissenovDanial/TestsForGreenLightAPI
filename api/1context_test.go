package main

import (
	"net/http/httptest"
	"testing"

	"github.com/DataDavD/snippetbox/greenlight/internal/data"
)

func TestContextSetUser(t *testing.T) {
	app := &application{}

	user := &data.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	req := httptest.NewRequest("GET", "/", nil)

	req = app.contextSetUser(req, user)

	ctxUser := app.contextGetUser(req)

	if ctxUser.ID != user.ID || ctxUser.Name != user.Name || ctxUser.Email != user.Email {
		t.Errorf("got %+v, want %+v", ctxUser, user)
	}
}

func TestContextGetUser(t *testing.T) {
	app := &application{}

	user := &data.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	req := httptest.NewRequest("GET", "/", nil)

	req = app.contextSetUser(req, user)

	ctxUser := app.contextGetUser(req)

	if ctxUser.ID != user.ID || ctxUser.Name != user.Name || ctxUser.Email != user.Email {
		t.Errorf("got %+v, want %+v", ctxUser, user)
	}
}
