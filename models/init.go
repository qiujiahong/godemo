package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init()  {
	orm.RegisterModel(
		new(Category),
		new(Book),
		//new(Document),
		//new(Attachment),
		//new(DocumentStore),
		new(BookCategory),
		new(Member),
		new(Collection),
		//new(Relationship),
		new(Fans),
		//new(Comments),
		//new(Score),
	)
}

func TNCategory() string {
	return "md_category"
}

func TNBookCategory() string {
	return "md_book_category"
}

func TNBook() string {
	return "md_books"
}

func TNDocuments() string {
	return "md_documents"
}
func TNDocumentStore() string {
	return "md_document_store"
}

func TNAttachment() string {
	return "md_attachment"
}

func TNRelationship() string {
	return "md_relationship"
}

func TNMembers() string {
	return "md_members"
}

func TNCollection() string {
	return "md_star"
}

func TNFans() string {
	return "md_fans"
}

func TNComments() string {
	return "md_comments"
}

func TNScore() string {
	return "md_score"
}

/*
* Tool Funcs
* */
//设置增减
//@param            table           需要处理的数据表
//@param            field           字段
//@param            condition       条件
//@param            incre           是否是增长值，true则增加，false则减少
//@param            step            增或减的步长
func IncOrDec(table string, field string, condition string, incre bool, step ...int) (err error) {
	mark := "-"
	if incre {
		mark = "+"
	}
	s := 1
	if len(step) > 0 {
		s = step[0]
	}
	sql := fmt.Sprintf("update %v set %v=%v%v%v where %v", table, field, field, mark, s, condition)
	_, err = orm.NewOrm().Raw(sql).Exec()
	return
}