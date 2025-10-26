package utility

import (
	"context"

	"go-jwt-auth/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

var applicationConfig = config.GetApplicationConfig()

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     applicationConfig.GOOGLE_CLIENT_ID,
	ClientSecret: applicationConfig.GOOGLE_CLIENT_SECRET,
	RedirectURL:  applicationConfig.GOOGLE_REDIRECT_URL,
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint:     google.Endpoint,
}

func VerifyGoogleIDToken(ctx context.Context, idToken string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, idToken, applicationConfig.GOOGLE_CLIENT_ID)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
