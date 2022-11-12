//////////////////////////////////////////////
//
// Copyright (C) 2022 Thanawat (Q) Chan-in,
// All rights reserved.
//
// Author: Thanawat (Q) Chan-in
// Year: 2022
//
//

package main

///////////////////////////////////////////
//
//  Libraries
//

import (
	"github.com/gin-gonic/gin"

	"example/go-rest-gin-gonic/middleware"
	content "example/go-rest-gin-gonic/modules/content"
)

///////////////////////////////////////////
//
//  Helper functions
//

///////////////////////////////////////////
//
//  Body functions
//

func main() {

	router := gin.Default()

	router.Use(middleware.MiddlewareFunc())

	router.GET("/contents", content.ContentController.GetAll)
	router.GET("/contents/:id", content.ContentController.GetById)

	router.Run("localhost:8080")

}
