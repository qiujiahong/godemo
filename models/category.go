package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
)

type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"`
	Intro  string //介绍
	Icon   string
	Cnt    int  //统计分类下图书
	Sort   int  //排序
	Status bool //状态，true 显示
}

func (m *Category) TableName() string {
	return TNCategory()
}

func (m *Category) GetCates(pid int, status int) (cates []Category, err error) {
	qs := orm.NewOrm().QueryTable(TNCategory())
	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}

	if status == 0 || status == 1 {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("-status", "sort", "title").All(&cates)
	return
}

//查询分类
func (m *Category) Find(id int) (cate Category) {
	cate.Id = id
	orm.NewOrm().Read(&cate)
	return cate
}

//批量新增分类
func (m *Category) InsertMulti(pid int, cates string) (err error) {
	slice := strings.Split(cates, "\n")
	if len(slice) == 0 {
		return
	}

	o := orm.NewOrm()
	for _, item := range slice {
		if item = strings.TrimSpace(item); item != "" {
			var cate = Category{
				Pid:    pid,
				Title:  item,
				Status: true,
			}
			if o.Read(&cate, "title"); cate.Id == 0 {
				_, err = o.Insert(&cate)
			}
		}
	}
	return
}

//删除分类
func (m *Category) Delete(id int) (err error) {
	var cate = Category{Id: id}

	o := orm.NewOrm()
	if err = o.Read(&cate); cate.Cnt > 0 { //当前分类下文档图书数量不为0，不允许删除
		return errors.New("删除失败，当前分类下的问下图书不为0，不允许删除")
	}

	if _, err = o.Delete(&cate, "id"); err != nil {
		return
	}
	_, err = o.QueryTable(TNCategory()).Filter("pid", id).Delete()
	if err != nil { //删除分类图标
		return
	}

	return
}

//更新分类字段
func (m *Category) UpdateField(id int, field, val string) (err error) {
	_, err = orm.NewOrm().QueryTable(TNCategory()).Filter("id", id).Update(orm.Params{field: val})
	return
}

// 统计分类书籍
var counting = false

type Count struct {
	Cnt        int
	CategoryId int
}

func CountCategory() {
	if counting {
		return
	}
	counting = true
	defer func() {
		counting = false
	}()

	var count []Count

	o := orm.NewOrm()
	sql := "select count(bc.id) cnt, bc.category_id from " + TNBookCategory() + " bc left join " + TNBook() + " b on b.book_id=bc.book_id where b.privately_owned=0 group by bc.category_id"
	o.Raw(sql).QueryRows(&count)
	if len(count) == 0 {
		return
	}

	var cates []Category
	o.QueryTable(TNCategory()).All(&cates, "id", "pid", "cnt")
	if len(cates) == 0 {
		return
	}

	var err error

	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}()

	o.QueryTable(TNCategory()).Update(orm.Params{"cnt": 0})
	cateChild := make(map[int]int)
	for _, item := range count {
		if item.Cnt > 0 {
			cateChild[item.CategoryId] = item.Cnt
			_, err = o.QueryTable(TNCategory()).Filter("id", item.CategoryId).Update(orm.Params{"cnt": item.Cnt})
			if err != nil {
				return
			}
		}
	}
}