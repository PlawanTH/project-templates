//////////////////////////////////////////////
//
// Copyright (C) 2022 Thanawat (Q) Chan-in,
// All rights reserved.
//
// Author: Thanawat (Q) Chan-in
// Year: 2022
//
//

package content

///////////////////////////////////////////
//
//  Libraries
//

import (
	"example/go-rest-gin-gonic/modules/content/v1/controllers"
	repo "example/go-rest-gin-gonic/modules/content/v1/repositories"
	"example/go-rest-gin-gonic/modules/content/v1/services"
)

///////////////////////////////////////////
//
//  Helper functions
//

///////////////////////////////////////////
//
//  Struct
//

///////////////////////////////////////////
//
//  Struct functions
//

var ContentRepository repo.IContentRepository = &repo.ContentMockRepository{}
var ContentService services.IContentService = &services.ContentService{Repo: ContentRepository}
var ContentController controllers.IContentController = &controllers.ContentController{Service: ContentService}
