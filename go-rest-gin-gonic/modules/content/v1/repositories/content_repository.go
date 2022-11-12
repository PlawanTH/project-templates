//////////////////////////////////////////////
//
// Copyright (C) 2022 Thanawat (Q) Chan-in,
// All rights reserved.
//
// Author: Thanawat (Q) Chan-in
// Year: 2022
//
//

package repositories

///////////////////////////////////////////
//
//  Libraries
//

import (
	"encoding/json"
	"example/go-rest-gin-gonic/modules/content/v1/models"
)

///////////////////////////////////////////
//
//  Helper functions
//

///////////////////////////////////////////
//
//  Struct
//

type IContentRepository interface {
	GetAll() ([]models.Content, error)
	GetById(id int32) (*models.Content, error)
}

type ContentMockRepository struct {
	Name string
}

///////////////////////////////////////////
//
//  Struct functions
//

func (r *ContentMockRepository) GetAll() ([]models.Content, error) {

	jsonResp := `
[
	{
		"id": 1,
		"title": "My Title 1",
		"description": "My description 1",
		"thumbnailUrl": ""
	},
	{
		"id": 2,
		"title": "My Title 2",
		"description": "My description 2",
		"thumbnailUrl": ""
	},
	{
		"id": 3,
		"title": "My Title 3",
		"description": "My description 3",
		"thumbnailUrl": ""
	}
]
	`

	var contents []models.Content

	// Unmarshal json
	err := json.Unmarshal([]byte(jsonResp), &contents)

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (r *ContentMockRepository) GetById(id int32) (*models.Content, error) {

	contents, err := r.GetAll()

	if err != nil {
		return nil, err
	}

	for _, content := range contents {
		if content.Id == id {
			return &content, nil
		}
	}

	return nil, nil

}
