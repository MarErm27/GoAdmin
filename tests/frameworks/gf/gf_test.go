package gf

import (
	"net/http"
	"testing"

	"github.com/MarErm27/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestGf(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}