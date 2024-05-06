package session

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionManager struct {
	store *session.Store
}

func NewSessionManager() *SessionManager {

	store := session.New(session.Config{
		CookiePath:        "/",
		CookieHTTPOnly:    true,
		CookieSessionOnly: true,
	})
	return &SessionManager{
		store: store,
	}
}

const USER_ID = "USER_ID"

func (manager *SessionManager) HandleCreateSession(c *fiber.Ctx, userId int64) (*session.Session, error) {
	sess, err := manager.store.Get(c)
	if err != nil {
		return nil, err
	}

	defer sess.Save()
	sess.SetExpiry(time.Hour * 24)
	if sess.Fresh() {
		sess.Set(USER_ID, userId)
	} else {
		log.Println("Regen sess")
		err = sess.Regenerate()
		return sess, err
	}

	return sess, nil
}
