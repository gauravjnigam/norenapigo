package norenapigo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

// UserSession represents the response after a successful authentication.
type UserSession struct {
	UserProfile
}

// UserSessionTokens represents response after renew access token.
type UserSessionTokens struct {
	SuserToken string `json:"susertoken"`
}

// UserProfile represents a user's personal and financial profile.
type UserProfile struct {
	RequestTime    string `json:"request_time"`
	Actid          string `json:"actid"`
	Uname          string `json:"uname"`
	Susertoken     string `json:"susertoken"`
	Email          string `json:"email"`
	UID            string `json:"uid"`
	Brkname        string `json:"brkname"`
	Lastaccesstime string `json:"lastaccesstime"`
}

// GenerateSession gets a user session details in exchange of username and password.
// Access token is automatically set if the session is retrieved successfully.
// Do the token exchange with the `requestToken` obtained after the login flow,
// and retrieve the `accessToken` required for all subsequent requests. The
// response contains not just the `accessToken`, but metadata for the user who has authenticated.
// totp used is required for 2 factor authentication
func (c *Client) GenerateSession(totp string) (UserSession, error) {
	// fmt.Printf("Generating sessions for user - %s, appkey - %s", c.clientCode, c.accessToken)
	u_app_key := fmt.Sprintf("%s|%s", c.clientCode, c.apiKey)
	passwordBytes := []byte(c.password)
	uAppKeyBytes := []byte(u_app_key)
	// Calculate the SHA256 hash
	hash := sha256.Sum256(passwordBytes)
	hashAppKey := sha256.Sum256(uAppKeyBytes)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash[:])
	hashAppKeyString := hex.EncodeToString(hashAppKey[:])

	// construct url values
	params := make(map[string]interface{})
	params["apkversion"] = "1.0.0"
	params["uid"] = c.clientCode
	params["pwd"] = hashString
	params["factor2"] = totp
	params["imei"] = "abc1234"
	params["source"] = "API"
	params["vc"] = "FA87226_U"
	params["appkey"] = hashAppKeyString

	// url := c.baseURI + URILogin
	// fmt.Printf("URL - %s \nParam : %v", url, params)
	var session UserSession
	err := c.doEnvelope(http.MethodPost, URILogin, params, nil, &session)
	// Set accessToken on successful session retrieve
	if err == nil && session.Susertoken != "" {
		c.SetAccessToken(session.Susertoken)
	}
	return session, err
}

// RenewAccessToken renews expired access token using valid refresh token.
func (c *Client) RenewAccessToken(refreshToken string) (string, error) {

	params := map[string]interface{}{}
	params["refreshToken"] = refreshToken

	var session UserSessionTokens
	err := c.doEnvelope(http.MethodPost, URIUserSessionRenew, params, nil, &session, true)

	// Set accessToken on successful session retrieve
	if err == nil && session.SuserToken != "" {
		c.SetAccessToken(session.SuserToken)
	}

	return session.SuserToken, err
}

// GetUserProfile gets user profile.
func (c *Client) GetUserProfile() (UserProfile, error) {
	var userProfile UserProfile
	params := make(map[string]interface{})
	params["uid"] = c.clientCode
	err := c.doEnvelope(http.MethodPost, URIUserProfile, params, nil, &userProfile, true)
	return userProfile, err
}

// Logout from User Session.
func (c *Client) Logout() (bool, error) {
	var status bool
	params := map[string]interface{}{}
	params["clientcode"] = c.clientCode
	err := c.doEnvelope(http.MethodPost, URILogout, params, nil, nil, true)
	if err == nil {
		status = true
	}
	return status, err
}
