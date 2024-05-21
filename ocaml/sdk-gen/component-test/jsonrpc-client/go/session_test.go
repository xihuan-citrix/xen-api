package componenttest

import (
	"testing"
	"xenapi"
)

func TestLoginSuccess(t *testing.T) {
	session = xenapi.NewSession(&xenapi.ClientOpts{
		URL: ServerURL,
		Headers: map[string]string{
			"Test-ID": "xapi-24/session_login_01",
		},
	})
	_, err := session.LoginWithPassword(*USERNAME_FLAG, *PASSWORD_FLAG, "1.0", "Go SDK component test")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedXapiVersion := "24.11"
	getXapiVersion := session.XAPIVersion
	if expectedXapiVersion != getXapiVersion {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expectedXapiVersion, getXapiVersion)
	}
	var expectedAPIVersion xenapi.APIVersion = xenapi.APIVersion2_21
	getAPIVersion := session.APIVersion
	if expectedAPIVersion != getAPIVersion {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expectedAPIVersion, getAPIVersion)
	}

	err = session.Logout()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}
