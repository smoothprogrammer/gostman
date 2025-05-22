// Package postman_test contains examples of gostman using Postman Echo.
// The following list are comparison of hierarchy between gostman and gostman.
//
//   - The package postman_test is equal to collection in the postman
//   - All tests in the form of TextXxx equal to folder in the postman (gostman doesn't support recursive folder)
//   - Request itself is running in the subtests
package postman_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/smoothprogrammer/gostman"
	"github.com/smoothprogrammer/testr"
)

type postmanResponse struct {
	Args          map[string]string `json:"args"`
	Authenticated bool              `json:"authenticated"`
	Headers       map[string]string `json:"headers"`
	Data          string            `json:"data"`
}

func TestMain(m *testing.M) {
	os.Exit(gostman.Run(m))
}

// TestRequest tests request to postman-echo.
// go test -run Request
func TestRequest(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	gm := gostman.New(t)

	// go test -run Request/Params
	gm.GET("Params", "https://postman-echo.com/get", func(r *gostman.Request) {
		r.Params(func(v url.Values) {
			v.Set("foo1", "bar1")
			v.Set("foo2", "bar2")
		})

		r.Send(func(t *testing.T, req *http.Request, res *http.Response) {
			defer res.Body.Close()

			assert := testr.New(t)
			assert.Equal(res.StatusCode, http.StatusOK)

			var resp postmanResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			assert.ErrorIs(err, nil)
			assert.Equal(resp.Args["foo1"], "bar1")
			assert.Equal(resp.Args["foo2"], "bar2")
		})
	})

	// go test -run Request/Authorization
	gm.GET("Authorization", "https://postman-echo.com/basic-auth", func(r *gostman.Request) {
		r.Authorization(gostman.AuthBasic(
			"postman",
			"password",
		))

		r.Send(func(t *testing.T, req *http.Request, res *http.Response) {
			defer res.Body.Close()

			assert := testr.New(t)
			assert.Equal(res.StatusCode, http.StatusOK)

			var resp postmanResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			assert.ErrorIs(err, nil)
			assert.Equal(resp.Authenticated, true)
		})
	})

	// go test -run Request/Headers
	gm.GET("Headers", "https://postman-echo.com/headers", func(r *gostman.Request) {
		r.Headers(func(h http.Header) {
			h.Set("foo1", "bar1")
			h.Set("foo2", "bar2")
		})

		r.Send(func(t *testing.T, req *http.Request, res *http.Response) {
			defer res.Body.Close()

			assert := testr.New(t)
			assert.Equal(res.StatusCode, http.StatusOK)

			var resp postmanResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			assert.ErrorIs(err, nil)
			assert.Equal(resp.Headers["foo1"], "bar1")
			assert.Equal(resp.Headers["foo2"], "bar2")
		})
	})

	// go test -run Request/Body
	gm.POST("Body", "https://postman-echo.com/post", func(r *gostman.Request) {
		text := "This is expected to be sent back as part of response body."

		r.Body(gostman.BodyText(text))

		r.Send(func(t *testing.T, req *http.Request, res *http.Response) {
			defer res.Body.Close()

			assert := testr.New(t)
			assert.Equal(res.StatusCode, http.StatusOK)

			var resp postmanResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			assert.ErrorIs(err, nil)
			assert.Equal(resp.Data, text)
		})
	})
}

// TestVariable tests request to postman-echo with variable.
//
//	Set env: go test -run Variable -env postman
//	Set env for the future request too: go test -run Variable -setenv postman
func TestVariable(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	gm := gostman.New(t)

	// go test -run Variable/Authorization -env postman
	gm.GET("Authorization", "https://postman-echo.com/basic-auth", func(r *gostman.Request) {
		r.Authorization(gostman.AuthBasic(
			gm.V("username"),
			gm.V("password"),
		))

		r.Send(func(t *testing.T, req *http.Request, res *http.Response) {
			defer res.Body.Close()

			assert := testr.New(t)
			assert.Equal(res.StatusCode, http.StatusOK)

			var resp postmanResponse
			err := json.NewDecoder(res.Body).Decode(&resp)
			assert.ErrorIs(err, nil)
			assert.Equal(resp.Authenticated, true)
		})
	})
}
