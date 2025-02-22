package server

import (

)

func NewRouter() *gin.Engine {
    r := gin.Default()

	r.GET("/", auth.HelloWorldHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}


