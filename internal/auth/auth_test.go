package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
)

// A simple mock for store to test Auth
type mockStore struct {
	err error
}

func (m *mockStore) CreateUser(ctx context.Context, email, passwordHash string) (interface{}, error) {
	return nil, m.err
}

func TestAuth_SetFlashAndGetFlash(t *testing.T) {
	sessionStore := sessions.NewCookieStore([]byte("secret"))
	a := NewAuth(nil, sessionStore)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	err := a.SetFlash(w, req, "test flash")
	if err != nil {
		t.Fatal(err)
	}

	res := w.Result()
	req2, _ := http.NewRequest("GET", "/", nil)
	for _, c := range res.Cookies() {
		req2.AddCookie(c)
	}

	w2 := httptest.NewRecorder()
	msg, err := a.GetFlash(w2, req2)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "test flash" {
		t.Fatalf("expected 'test flash', got '%s'", msg)
	}
}
