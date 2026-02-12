package server

import "github.com/gin-gonic/gin"

type ServerConfig struct {
	Port         string
	Handler      *gin.Engine
	IdleTimeout  int // in second
	ReadTimeout  int // in second
	WriteTimeout int // in second
}
