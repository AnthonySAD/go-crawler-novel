package shubao69

type Chapter struct {
	ID          uint 			`gorm:"primary_key"`
	NovelID		uint
	Order       uint
	Title		string
	Content     string			`gorm:"type:text"`
	WordsTotal	int
	NextUrl     string			`sql:index`
}

func (Chapter) TableName() string{
	return "chapters"
}