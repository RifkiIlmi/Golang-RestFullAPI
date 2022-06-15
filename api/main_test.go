package api

import (
	"os"
	"testing"

	db "github.com/Rifkiilmi/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func newTestServer(t *testing.T, store db.IStore) *Server {

	server := NewServer(store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
