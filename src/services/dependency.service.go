package services

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"trouble-ticket-ms/src/config"
)

type AppDependencies struct {
	KeycloakClient *gocloak.GoCloak
	Context        context.Context
	KeycloakCfg    config.KeyCloakConfig
}

func InitAppDependencies() *AppDependencies {
	// Initialize the config
	cfg := config.New()

	// Initialize the Keycloak client
	client := gocloak.NewClient(cfg.KEYCLOAK.Host)

	// Initialize the context
	ctx := context.Background()

	return &AppDependencies{
		KeycloakClient: client,
		Context:        ctx,
		KeycloakCfg:    cfg.KEYCLOAK,
	}
}
