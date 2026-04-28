package auth

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"proyecto-monolito/internal/store"
)

type contextKey string

const (
	sessionName     = "app-session"
	contextUserID   = contextKey("user_id")
	flashSessionKey = "flash"
)

type Auth struct {
	store        store.Store
	sessionStore *sessions.CookieStore
}

func init() {
	gob.Register(int64(0))
}

func NewAuth(store store.Store, sessionStore *sessions.CookieStore) *Auth {
	return &Auth{store: store, sessionStore: sessionStore}
}

func (a *Auth) Register(ctx context.Context, email, password string) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user, err := a.store.CreateUser(ctx, email, string(hash))
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (a *Auth) Login(ctx context.Context, email, password string) (int64, error) {
	user, err := a.store.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 0, errors.New("credenciales inválidas")
	}
	return user.ID, nil
}

func (a *Auth) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := a.sessionStore.Get(r, sessionName)
		userValue, ok := session.Values["user_id"]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		userID, ok := userValue.(int64)
		if !ok || userID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), contextUserID, userID))
		next.ServeHTTP(w, r)
	})
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := a.sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return sessions.Save(r, w)
}

func (a *Auth) SetSessionUser(w http.ResponseWriter, r *http.Request, userID int64) error {
	session, err := a.sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user_id"] = userID
	return sessions.Save(r, w)
}

func GetCurrentUserID(r *http.Request) int64 {
	value := r.Context().Value(contextUserID)
	if value == nil {
		return 0
	}
	userID, ok := value.(int64)
	if !ok {
		return 0
	}
	return userID
}

func (a *Auth) SetFlash(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := a.sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.AddFlash(message, flashSessionKey)
	return sessions.Save(r, w)
}

func (a *Auth) GetFlash(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := a.sessionStore.Get(r, sessionName)
	if err != nil {
		return "", err
	}
	flashes := session.Flashes(flashSessionKey)
	if len(flashes) == 0 {
		return "", sessions.Save(r, w)
	}
	if err := sessions.Save(r, w); err != nil {
		return "", err
	}
	if msg, ok := flashes[0].(string); ok {
		return msg, nil
	}
	return "", nil
}
