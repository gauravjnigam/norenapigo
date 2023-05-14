package norenapigo

import (
	"testing"
)

func (ts *TestSuite) TestGetRMS(t *testing.T) {
	t.Parallel()
	rms, err := ts.TestConnect.GetRMS()
	if err != nil {
		t.Errorf("Error while fetching RMS. %v", err)
	}

	if rms.Net == "" {
		t.Errorf("Error while fetching Net from RMS. %v", err)
	}

}
