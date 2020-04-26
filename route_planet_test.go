package main

import (
	"context"
	"fmt"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func withTestServer() *httptest.Server {
	data.WithTestDatabase()
	initSessionManager("sqlite3")
	ts := httptest.NewServer(routes(false))
	if err := os.Setenv("SERVER_URL", ts.URL); err != nil {
		panic(err)
	}
	return ts
}

func newTestClient() *http.Client {
	return &http.Client{
		Jar: &TestJar{},
	}
}

// Copied from net/http/jar.go
type TestJar struct {
	m      sync.Mutex
	perURL map[string][]*http.Cookie
}

func (j *TestJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.m.Lock()
	defer j.m.Unlock()
	if j.perURL == nil {
		j.perURL = make(map[string][]*http.Cookie)
	}
	j.perURL[u.Host] = cookies
}

func (j *TestJar) Cookies(u *url.URL) []*http.Cookie {
	j.m.Lock()
	defer j.m.Unlock()
	return j.perURL[u.Host]
}

func withTestServerAuthenticated() (*httptest.Server, *ent.User) {
	ts := withTestServer()
	u := data.NewUserFactory().MustCreate()
	ts.Config.ConnContext = func(ctx context.Context, c net.Conn) context.Context {
		var err error
		ctx, err = sessionManager.Load(ctx, "")
		if err != nil {
			panic(err)
		}
		sessionManager.Put(ctx, userIDKey, u.ID)
		return ctx
	}
	return ts, u
}

func TestRoutePublicPages(t *testing.T) {
	ts := withTestServer()
	defer ts.Close()

	routes := []string{
		"/",
		"/signup",
	}

	for _, r := range routes {
		t.Run(r, func(t *testing.T) {
			res, err := http.Get(ts.URL + r)
			assert.NoError(t, err)
			assert.Equal(t, 200, res.StatusCode)
		})
	}
}

func TestRouteDashboardPages(t *testing.T) {
	ts, _ := withTestServerAuthenticated()
	defer ts.Close()

	routes := []string{
		"/dashboard",
		"/dashboard/fleetcontrol",
		"/dashboard/technology",
		"/dashboard/diplomacy",
		"/dashboard/story",
		"/dashboard/wiki",
		"/dashboard/news",
	}

	for _, r := range routes {
		t.Run(r, func(t *testing.T) {
			res, err := http.Get(ts.URL + r)
			assert.NoError(t, err)
			assert.Equal(t, 200, res.StatusCode)
		})
	}
}

func TestRoutePlanetPages(t *testing.T) {
	ts, u := withTestServerAuthenticated()
	defer ts.Close()

	p := data.NewPlanetFactory().ForOwner(u).MustCreate()

	routes := []string{
		fmt.Sprintf("/planet/%d/", p.ID),
		fmt.Sprintf("/planet/%d/constructions", p.ID),
		fmt.Sprintf("/planet/%d/factories", p.ID),
		fmt.Sprintf("/planet/%d/research", p.ID),
		fmt.Sprintf("/planet/%d/fleets", p.ID),
		fmt.Sprintf("/planet/%d/defenses", p.ID),
	}

	for _, r := range routes {
		t.Run(r, func(t *testing.T) {
			res, err := http.Get(ts.URL + r)
			assert.NoError(t, err)
			assert.Equal(t, 200, res.StatusCode)
		})
	}
}
