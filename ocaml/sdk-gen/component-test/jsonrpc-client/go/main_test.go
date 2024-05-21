package componenttest

import (
	"encoding/json"
	"flag"
	"os"
	"testing"
	"xenapi"
)

const ServerURL = "http://localhost:5000"

var session *xenapi.Session

var USERNAME_FLAG = flag.String("root", "", "the username of the host (e.g. root)")
var PASSWORD_FLAG = flag.String("secret", "", "the password of the host")

func GetSession(testId string) (*xenapi.Session, error) {
	var newSession *xenapi.Session
	newSession = xenapi.NewSession(&xenapi.ClientOpts{
		URL: ServerURL,
		Headers: map[string]string{
			"Test-ID": testId,
		},
	})

	return newSession, nil
}

func TestMain(m *testing.M) {
	flag.Parse()
	exitVal := m.Run()
	os.Exit(exitVal)
}
