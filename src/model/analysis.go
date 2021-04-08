package model

import "github.com/jinzhu/gorm"

type AnalysisInfo struct {
	gorm.Model
	GrapeType string  `json:"grape_type"`
	Co        float32 `json:"co"`
	S         float32 `json:"s"`
	Guano     float32 `json:"guano"`
	H20       float32 `json:"h2o"`
	Ph        float32 `json:"ph"`
	Bx        float32 `json:"bx"`
	TA        float32 `json:"ta"`
	Kg        uint    `json:"kg"`
}
