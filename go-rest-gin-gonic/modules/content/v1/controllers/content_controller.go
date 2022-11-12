//////////////////////////////////////////////
//
// Copyright (C) 2022 Thanawat (Q) Chan-in,
// All rights reserved.
//
// Author: Thanawat (Q) Chan-in
// Year: 2022
//
//

package controllers

///////////////////////////////////////////
//
//  Libraries
//

import (
	"fmt"
	"net/http"
	"strconv"

	"example/go-rest-gin-gonic/modules/content/v1/models"
	"example/go-rest-gin-gonic/modules/content/v1/services"

	"github.com/gin-gonic/gin"
)

///////////////////////////////////////////
//
//  Helper functions
//

///////////////////////////////////////////
//
//  Struct
//

type IContentController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	AddContent(ctx *gin.Context)
}

type ContentController struct {
	Service services.IContentService
}

///////////////////////////////////////////
//
//  Struct functions
//

func (c *ContentController) GetAll(ctx *gin.Context) {

	contents, err := c.Service.GetAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, contents)
}

func (c *ContentController) GetById(ctx *gin.Context) {

	// Get id param
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)

	// Response error
	if err != nil {
		errMsg := fmt.Sprintf("Invalid id: %s", idStr)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	content, err := c.Service.GetById(int32(idInt))

	ctx.IndentedJSON(http.StatusOK, content)
}

func (c *ContentController) AddContent(ctx *gin.Context) {

	var newContent models.Content

	// Bind 'Content' struct with JSON input (from http request)
	if err := ctx.BindJSON(&newContent); err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, newContent)

}
