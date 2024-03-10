package norenapigo

/*
import (
	"testing"
)

// MockSession is a mock implementation of norenapigo.Client interface
type MockSession struct {
	Susertoken    string
	GenerateErr   error
	RenewToken    string
	RenewErr      error
	UserProfile   UserProfile
	GetProfileErr error
	LogoutStatus  bool
	LogoutErr     error
}

// Implement GenerateSession method for the MockSession
func (m *MockSession) GenerateSession(totp string) (UserSession, error) {
	if m.GenerateErr != nil {
		return UserSession{}, m.GenerateErr
	}
	return UserSession{UserProfile: UserProfile{Susertoken: m.Susertoken}}, nil
}

// Implement RenewAccessToken method for the MockSession
func (m *MockSession) RenewAccessToken(refreshToken string) (string, error) {
	if m.RenewErr != nil {
		return "", m.RenewErr
	}
	return m.RenewToken, nil
}

// Implement GetUserProfile method for the MockSession
func (m *MockSession) GetUserProfile() (UserProfile, error) {
	if m.GetProfileErr != nil {
		return UserProfile{}, m.GetProfileErr
	}
	return m.UserProfile, nil
}

// Implement Logout method for the MockSession
func (m *MockSession) Logout() (bool, error) {
	if m.LogoutErr != nil {
		return false, m.LogoutErr
	}
	return m.LogoutStatus, nil
}

func (ts *TestSuite) TestGenerateSession(t *testing.T) {
	t.Parallel()
	mock := &MockSession{
		Susertoken:  "mockedToken",
		GenerateErr: nil,
	}

	totp := "123456" // Mock TOTP for testing

	session, err := mock.GenerateSession(totp)
	if err != nil {
		t.Errorf("GenerateSession error: %v", err)
	}

	if session.UserProfile.Susertoken != "mockedToken" {
		t.Error("Expected mockedToken, got different token")
	}

}

func (ts *TestSuite) TestRenewAccessToken(t *testing.T) {
	t.Parallel()

	mock := &MockSession{
		Susertoken:  "mockedToken",
		GenerateErr: nil,
	}

	totp := "123456" // Mock TOTP for testing

	session, err := mock.GenerateSession(totp)
	if err != nil {
		t.Errorf("GenerateSession error: %v", err)
	}

	if session.Susertoken == "" {
		t.Errorf("Error while fetching new access token. %v", err)
	}

}

func (ts *TestSuite) TestUserProfile(t *testing.T) {
	t.Parallel()
	session, err := ts.TestConnect.GetUserProfile()
	if err != nil {
		t.Errorf("Error while fetching user profile. %v", err)
	}

	if session.UID == "" {
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
*/
