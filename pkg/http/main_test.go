package http

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stovenn/dataimpact/internal"
	"github.com/stovenn/dataimpact/pkg/util"
)

func newTestServer(t *testing.T, store internal.Store) *Server {
	config := util.Config{
		SymmetricKey:  util.RandomString(32),
		TokenDuration: time.Minute,
	}
	server := NewServer(store, log.Default(), log.Default(), config)

	return server
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
