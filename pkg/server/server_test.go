package server

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestWellKnown(t *testing.T) {
	app := setup()

	const expectedBody = `{"modules.v1":"/v1/modules/","providers.v1":"/v1/providers/"}`

	req := httptest.NewRequest("GET", "/.well-known/terraform.json", nil)
	res, _ := app.Test(req, -1)
	body, _ := ioutil.ReadAll(res.Body)

	assert.Equalf(t, expectedBody, string(body), "wellknown test")
}
