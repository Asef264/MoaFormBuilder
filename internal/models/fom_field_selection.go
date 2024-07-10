package models

import (
	"errors"
	"moaformbuilder/pkg/infrastructure/tools"

	"gorm.io/gorm"
)

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

// all fields can not be false
func (ffs FormFieldSelection) Validate() error {
	fieldValues := tools.GetFields(ffs)
	falseCount := 0
	for _, v := range fieldValues {
		if v == false {
			falseCount++
		}
	}
	if falseCount == len(fieldValues) {
		return errors.New("form is invalid")
	}
	return nil
}
