package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type config struct {
	port         string
	handler      *gin.Engine
	idleTimeout  int // in second
	readTimeout  int // in second
	writeTimeout int // in second
	server       *http.Server
}
