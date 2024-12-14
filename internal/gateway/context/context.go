package context

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/segmentio/ksuid"
	"golang.org/x/exp/rand"
)

// isUnavailableDevice returns true if the device is unavailable
func IsUnavailableDevice(c echo.Context) bool {
	s, err := Get(c)
	if err != nil {
		return true
	}
	return s.IsBot() || s.IsTV()
}

// initSession initializes or loads an existing session
func (s *HTTPContext) initSession() error {
	sessionID := s.getOrCreateSessionID()
	f := rand.Intn(5) + 1
	l := rand.Intn(4) + 1
	challenge, err := protocol.CreateChallenge()
	if err != nil {
		return err
	}

	// Try to load existing session
	var sess models.Session
	result := s.db.Where("id = ?", sessionID).First(&sess)
	if result.Error != nil {
		// Create new session if not found
		sess = models.Session{
			ID:             sessionID,
			BrowserName:    s.GetBrowser(),
			BrowserVersion: s.GetMajorVersion(),
			Platform:       s.GetOS(),
			IsMobile:       s.IsMobile(),
			IsTablet:       s.IsTablet(),
			IsDesktop:      s.IsDesktop(),
			IsBot:          s.IsBot(),
			IsTV:           s.IsTV(),
			IsHumanFirst:   f,
			IsHumanLast:    l,
			Challenge:      challenge.String(),
		}
		if err := s.db.Create(&sess).Error; err != nil {
			return err
		}
	}
	s.sess = &sess
	return nil
}

func (s *HTTPContext) getOrCreateSessionID() string {
	if ok := common.CookieExists(s.Context, common.SessionID); !ok {
		sessionID := ksuid.New().String()
		common.WriteCookie(s.Context, common.SessionID, sessionID)
		return sessionID
	}

	sessionID, err := common.ReadCookie(s.Context, common.SessionID)
	if err != nil {
		sessionID = ksuid.New().String()
		common.WriteCookie(s.Context, common.SessionID, sessionID)
	}
	return sessionID
}
