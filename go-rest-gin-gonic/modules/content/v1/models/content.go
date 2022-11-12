//////////////////////////////////////////////
//
// Copyright (C) 2022 Thanawat (Q) Chan-in,
// All rights reserved.
//
// Author: Thanawat (Q) Chan-in
// Year: 2022
//
//

package models

///////////////////////////////////////////
//
//  Libraries
//

///////////////////////////////////////////
//
//  Helper functions
//

///////////////////////////////////////////
//
//  Struct
//

type Content struct {
	Id           int32  `json="id"`
	Title        string `json="title"`
	Description  string `json="description"`
	ThumbnailUrl string `json="thumbnailUrl"`
}

///////////////////////////////////////////
//
//  Struct functions
//
