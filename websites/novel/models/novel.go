package shubao69

import (
	"time"
)

type Novel struct {
	ID            uint 			`gorm:"primary_key"`
	Title         string		`gorm:"size:100"`
	AuthorName    string		`gorm:"size:50"`
	Url           string
	Tag           string
	Interview     string 		`gorm:"type:text"`
	LastUpdatedAt time.Time     `gorm:"DEFAULT:null"`
	ChapterTotal  uint
	WordsTotal    uint
	Status        int		    `gorm:"type:tinyint;DEFAULT:0"`
	Chapters	  []Chapter
}

func (Novel) TableName() string{
	return "novels"
}
