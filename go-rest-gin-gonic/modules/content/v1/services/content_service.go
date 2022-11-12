//////////////////////////////////////////////
//
// Copyright (C) 2022 Thanawat (Q) Chan-in,
// All rights reserved.
//
// Author: Thanawat (Q) Chan-in
// Year: 2022
//
//

package services

///////////////////////////////////////////
//
//  Libraries
//

import (
	"example/go-rest-gin-gonic/modules/content/v1/models"
	repo "example/go-rest-gin-gonic/modules/content/v1/repositories"
)

///////////////////////////////////////////
//
//  Helper functions
//

///////////////////////////////////////////
//
//  Struct
//

type IContentService interface {
	GetAll() ([]models.Content, error)
	GetById(id int32) (*models.Content, error)
}

type ContentService struct {
	Repo repo.IContentRepository
}

///////////////////////////////////////////
//
//  Struct functions
//

func (s *ContentService) GetAll() ([]models.Content, error) {
	return s.Repo.GetAll()
}

func (s *ContentService) GetById(id int32) (*models.Content, error) {
	return s.Repo.GetById(id)
}
