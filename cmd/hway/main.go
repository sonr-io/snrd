package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/pkg/common/clients"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/common/signer"
	"github.com/onsonr/sonr/pkg/gateway"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/gateway/database"
)

//go:embed config.pkl
var configBz []byte

func main() {
	env, err := loadConfig()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.HwayMiddleware(env))
	e.Use(signer.Middleware())
	e.Use(database.Middleware(env))
	e.Use(clients.Middleware(env))
	gateway.RegisterRoutes(e)

	if err := e.Start(fmt.Sprintf(":%d", env.GetServePort())); err != http.ErrServerClosed {
		log.Fatal(err)
		os.Exit(1)
		return
	}
}

func loadConfig() (config.Env, error) {
	return config.LoadFromBytes(configBz)
}
