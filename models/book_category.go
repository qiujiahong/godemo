package models

type BookCategory struct {
	Id int //自增ID
	BookId int	//book_id
	CategoryId int //category_id
}

func (m * BookCategory) TableName() string {
	return TNBookCategory()
}