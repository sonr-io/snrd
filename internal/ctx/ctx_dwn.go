package ctx

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
)

type DWNContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request
}

func (s *DWNContext) ID() string {
	return s.id
}

func (s *DWNContext) Address() string {
	cnfg, _ := GetConfig(s.Context)
	return cnfg.Motr.Address
}

func (s *DWNContext) ChainID() string {
	cnfg, _ := GetConfig(s.Context)
	return cnfg.Sonr.ChainId
}

func (s *DWNContext) Schema() *dwngen.Schema {
	cnfg, _ := GetConfig(s.Context)
	return cnfg.Schema
}

func GetDWNContext(c echo.Context) (*DWNContext, error) {
	ctx, ok := c.(*DWNContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return ctx, nil
}

// HighwaySessionMiddleware establishes a Session Cookie.
func DWNSessionMiddleware(config *dwngen.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionID := GetSessionID(c)
			SetConfig(c, config)
			cc := &DWNContext{
				Context: c,
				id:      sessionID,
			}
			return next(cc)
		}
	}
}
