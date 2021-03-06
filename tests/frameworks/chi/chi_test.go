package chi

import (
	"net/http"
	"testing"

	"github.com/MarErm27/GoAdmin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestChi(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
