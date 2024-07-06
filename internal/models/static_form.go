package models

import "gorm.io/gorm"

type StaticForm struct {
	gorm.Model
	Title string

	//fields
	PersonalInformation           *PersonalInformation           `json:"personalInformation"`
	EducationInformation          []EducationInformation         `json:"educationInformation"`
	WrokInformation               []WrokExperienceInformation    `json:"workExperience"`
	AdditionalIdentityInformation *AdditionalIdentityInformation `json:"additionalIdentityInformation"`
	MarrigeInformation            *MarrigeInformation            `json:"marriageInformation"`
	AttachedFiles                 []File                         `json:"attachedFiles"`
	Descrition                    string                         `json:"description"`
}

type File struct {
	gorm.Model
	RelationID  string `json:"relationID"`
	Title       string `json:"title"`
	FileUrl     string `json:"fileUrl"`
	Description string `json:"description"`
}

type PersonalInformation struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	FatherName    string `json:"fatherName"`
	BirthDate     int64  `json:"birthDate"`
	NationalCode  string `json:"nationalCode"`
	PhoneNumber   string `json:"phoneNumber"`
	Address       string `json:"address"`
	IsMale        bool   `json:"isMale"`
	AttachedFiles []File `json:"attachedFiles"`
}

type EducationInformation struct {
	LastDegeree        EducationDegree `json:"lastDegeree"`
	LastEducationPlace string          `json:"lastEducationPlace"`
	EducationField     string          `json:"educationField"`
	Average            int64           `json:"average"`
}

type EducationDegree int

const (
	HighSchool EducationDegree = iota
	BachelorDegree
	MasterDegree
	DoctorateDegree
	PostDoctorateDegree
)

type WrokExperienceInformation struct {
	WorkPosition     string `json:"workPosition"`
	ExperienceByYear int64  `json:"experienceByYear"`
	CompanyName      string `json:"companyName"`
	PhoneNumber      string `json:"phoneNumber"`
	Address          string `json:"address"`
	AttachedFiles    []File `json:"attachedFiles"`
}

type AdditionalIdentityInformation struct {
	PassportNumber string    `json:"passportNumber"`
	BirthLand      *Location `json:"birthLand"`
}

type Location struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type MarrigeInformation struct {
	IsMarried                     bool                          `json:"isMarried"`
	PartnersFullName              string                        `json:"partnersFullName"`
	NationalCode                  string                        ` json:"nationalCode"`
	PhoneNumber                   string                        `json:"phoneNumber"`
	EducationInformation          EducationInformation          `json:"educationInformation"`
	WrokExperienceInformation     WrokExperienceInformation     `json:"workExperience"`
	AdditionalIdentityInformation AdditionalIdentityInformation `json:"additionalIdentity"`
}
