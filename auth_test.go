package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
)

func TestServeHTTP(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	ah := &authHandler{
		next: http.HandlerFunc(nil),
	}

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	ah.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusPermanentRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}

func TestLoginHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/auth/login/github", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up gomniauth
	gomniauth.SetSecurityKey("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	gomniauth.WithProviders(
		github.New("60a4c05cdc1199396d32", "3dfc10b50c17fed83bc354b268e10366a9ce46d9", "/auth/callback/github"),
	)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}
