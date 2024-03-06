package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
