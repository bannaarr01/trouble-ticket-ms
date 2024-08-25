package services

import (
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"trouble-ticket-ms/src/models"
)

type AuthService interface {
	SignIn(models.Auth) (*gocloak.JWT, error)
	SignUp(models.SignUpDTO) (*gocloak.User, error)
}

type authService struct {
	deps AppDependencies
}

func (auth *authService) SignIn(authM models.Auth) (*gocloak.JWT, error) {
	// login with user credentials to get jwt Token payload
	jwtPayload, err := auth.deps.KeycloakClient.Login(
		auth.deps.Context,
		auth.deps.KeycloakCfg.ClientID,
		"", // not used actually
		auth.deps.KeycloakCfg.Realm,
		authM.Username,
		authM.Password,
	)

	if err != nil {
		return nil, err
	}

	return jwtPayload, nil
}

func (auth *authService) SignUp(signUpM models.SignUpDTO) (*gocloak.User, error) {
	token, err := auth.loginToKeycloak()
	if err != nil {
		return nil, err
	}
	user, err := auth.createUser(token, signUpM)
	if err != nil {
		return nil, err
	}

	userId := *user.ID

	err = auth.setPassword(token, userId, signUpM.Password)
	if err != nil {
		return nil, auth.handleError(token, userId, err)
	}

	err = auth.assignRoles(token, userId, signUpM.RealmRoles)
	if err != nil {
		return nil, auth.handleError(token, userId, err)
	}

	return auth.getUserByID(token, userId)
}

// loginToKeycloak to use Service Account Token for user management activities
func (auth *authService) loginToKeycloak() (string, error) {
	token, err := auth.deps.KeycloakClient.LoginClient(
		auth.deps.Context,
		auth.deps.KeycloakCfg.ClientServiceActID,
		auth.deps.KeycloakCfg.ClientSecret,
		auth.deps.KeycloakCfg.Realm,
	)
	if err != nil {
		return "", fmt.Errorf("invalid service account credentials: %v", err)
	}
	return token.AccessToken, nil
}

func (auth *authService) createUser(token string, signUpM models.SignUpDTO) (*gocloak.User, error) {
	user := gocloak.User{
		Username:      gocloak.StringP(signUpM.Username),
		Email:         gocloak.StringP(signUpM.Email),
		FirstName:     gocloak.StringP(signUpM.FirstName),
		LastName:      gocloak.StringP(signUpM.LastName),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(true),
	}

	newUserId, err := auth.deps.KeycloakClient.CreateUser(
		auth.deps.Context,
		token,
		auth.deps.KeycloakCfg.Realm,
		user,
	)
	if err != nil {
		return nil, err
	}

	return &gocloak.User{ID: gocloak.StringP(newUserId)}, nil
}

func (auth *authService) setPassword(token, userId, password string) error {
	return auth.deps.KeycloakClient.SetPassword(
		auth.deps.Context,
		token,
		userId,
		auth.deps.KeycloakCfg.Realm,
		password,
		false,
	)
}

func (auth *authService) assignRoles(token, userId string, realmRoles []string) error {
	roles, err := auth.getRealmRoles(token, realmRoles)
	if err != nil {
		return err
	}

	return auth.deps.KeycloakClient.AddRealmRoleToUser(
		auth.deps.Context,
		token,
		auth.deps.KeycloakCfg.Realm,
		userId,
		roles,
	)
}

func (auth *authService) getRealmRoles(token string, realmRoles []string) ([]gocloak.Role, error) {
	param := gocloak.GetRoleParams{}

	roles, err := auth.deps.KeycloakClient.GetRealmRoles(
		auth.deps.Context,
		token,
		auth.deps.KeycloakCfg.Realm,
		param,
	)

	if err != nil {
		return nil, err
	}

	foundRoles := make([]gocloak.Role, len(realmRoles))
	for i, roleName := range realmRoles {
		for _, role := range roles {
			if *role.Name == roleName {
				foundRoles[i] = *role
				break
			}
		}
		if foundRoles[i].ID == nil {
			return nil, fmt.Errorf("invalid role. %s does not exist", roleName)
		}
	}

	return foundRoles, nil
}

func (auth *authService) getUserByID(token, userId string) (*gocloak.User, error) {
	return auth.deps.KeycloakClient.GetUserByID(
		auth.deps.Context,
		token,
		auth.deps.KeycloakCfg.Realm,
		userId,
	)
}

func (auth *authService) handleError(token, userId string, err error) error {
	auth.deleteUserWithRetry(token, userId)
	return err
}

// deleteUserWithRetry attempts to delete the user and retries once if it fails
func (auth *authService) deleteUserWithRetry(token, userId string) error {
	const maxRetries = 2
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		lastErr = auth.deps.KeycloakClient.DeleteUser(
			auth.deps.Context,
			token,
			auth.deps.KeycloakCfg.Realm,
			userId,
		)
		if lastErr == nil {
			return nil
		}
	}

	return lastErr
}

func NewAuthService(dependencies AppDependencies) AuthService {
	return &authService{dependencies}
}
