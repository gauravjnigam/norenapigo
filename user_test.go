package norenapigo

import (
	"testing"
)

func (ts *TestSuite) TestGenerateSession(t *testing.T) {
	t.Parallel()
	session, err := ts.TestConnect.GenerateSession()
	if err != nil {
		t.Errorf("Error while generating session. %v", err)
	}

	if session.AccessToken == "" {
		t.Errorf("Error while fetching access token. %v", err)
	}

}

func (ts *TestSuite) TestRenewAccessToken(t *testing.T) {
	t.Parallel()
	session, err := ts.TestConnect.RenewAccessToken("test")
	if err != nil {
		t.Errorf("Error while regenerating session. %v", err)
	}

	if session.AccessToken == "" {
		t.Errorf("Error while fetching new access token. %v", err)
	}

}

func (ts *TestSuite) TestUserProfile(t *testing.T) {
	t.Parallel()
	session, err := ts.TestConnect.GetUserProfile()
	if err != nil {
		t.Errorf("Error while fetching user profile. %v", err)
	}

	if session.ClientCode == "" {
		t.Errorf("Error while fetching client code. %v", err)
	}

}

func (ts *TestSuite) TestLogout(t *testing.T) {
	t.Parallel()
	resp, err := ts.TestConnect.Logout()
	if err != nil {
		t.Errorf("Error while calling log out api. %v", err)
	}

	if !resp {
		t.Errorf("Error while logging out. %v", err)
	}

}
