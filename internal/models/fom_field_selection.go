package models

import "gorm.io/gorm"

// FormFieldSelection represents the boolean field selection for the form
type FormFieldSelection struct {
	gorm.Model
	Title             bool `json:"title"`
	PersonalInfo      bool `json:"personalInfo"`
	EducationInfo     bool `json:"educationInfo"`
	WrokExperience    bool `json:"workExperience"`
	AddtionalIdentity bool `json:"additionalIdentity"`
	MarrigeInfo       bool `json:"marriageInfo"`
	Descrition        bool `json:"description"`
}
