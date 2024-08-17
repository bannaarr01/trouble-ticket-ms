package services

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"trouble-ticket-ms/src/config"
	"trouble-ticket-ms/src/models"
)

type AuthService interface {
	SignIn(models.Auth) (*gocloak.JWT, error)
	SignUp() error
}

type authService struct {
	keycloakClient *gocloak.GoCloak
	context        context.Context
	keycloakCfg    config.KeyCloakConfig
}

func (auth *authService) SignIn(authM models.Auth) (*gocloak.JWT, error) {
	token, err := auth.keycloakClient.Login(
		auth.context,
		auth.keycloakCfg.ClientID,
		auth.keycloakCfg.ClientSecret,
		auth.keycloakCfg.Realm,
		authM.Username,
		authM.Password,
	)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (auth *authService) SignUp() error {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(
	keycloakClient *gocloak.GoCloak,
	ctx context.Context,
	keycloakCfg config.KeyCloakConfig,
) AuthService {
	return &authService{keycloakClient, ctx, keycloakCfg}
}
