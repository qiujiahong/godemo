package routers

import (
    "github.com/astaxie/beego"
    "mbook/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})

    //首页&分类&详情
    beego.Router("/",&controllers.HomeController{},"get:Index")         //首页
    beego.Router("/explore",&controllers.ExploreController{},"get:Index")   //分类
    //beego.Router("/books/:key",&controllers.DocumentController{},"get:Index")   //图书目录&详情

    //读书
    //beego.Router("/read/:key/:id",&controllers.DocumentController{},"*:Read")
    //beego.Router("/read/:key/search",&controllers.DocumentController{},"post:Search")
    //
    ////图书编辑
    //beego.Router("/api/:key/content/?:id",&controllers.DocumentController{},"*:Content")
    //beego.Router("/api/:key/edit/?:id",&controllers.DocumentController{},"*:Edit")
    //beego.Router("/api/upload",&controllers.DocumentController{},"post:Upload")
    //beego.Router("/api/:key/create",&controllers.DocumentController{},"post:Create")
    //beego.Router("/api/:key/delete",&controllers.DocumentController{},"post:Delete")
    //
    ////搜索
    //beego.Router("/search",&controllers.SearchController{},"get:Search")
    //beego.Router("/search/result",&controllers.SearchController{},"get:Result")
    //
    ////login
    //beego.Router("/login",&controllers.AccountController{},"*:Login")
    //beego.Router("/regist",&controllers.AccountController{},"*:Regist")
    //beego.Router("/logout",&controllers.AccountController{},"*:Logout")
    //
    //
    ////用户图书管理
    //beego.Router("/book",&controllers.BookController{},"*:Index")
    //beego.Router("/book/create",&controllers.BookController{},"post:Create")
    //beego.Router("/book/:key/setting",&controllers.BookController{},"*:Setting")
    //beego.Router("/book/setting/upload",&controllers.BookController{},"post:Upload")
    //beego.Router("/book/star/:id",&controllers.BookController{},"*:Collection")
    //beego.Router("/book/setting/save",&controllers.BookController{},"post:SaveBook")
    //beego.Router("/book/:key/release",&controllers.BookController{},"post:Release")
    //beego.Router("/book/setting/token",&controllers.BookController{},"post:CreateToken")
    //
    ////个人中心
    //beego.Router("/user/:username",&controllers.UserController{},"get:Index")   //分享
    //beego.Router("/user/:username/collection",&controllers.UserController{},"get:Collection")  //收藏
    //beego.Router("/user/:username/follow",&controllers.UserController{},"get:Follow")      //关注
    //beego.Router("/user/:username/fans",&controllers.UserController{},"get:Fans")      //粉丝
    //beego.Router("/follow/:uid",&controllers.UserController{},"get:SetFollow")      //关注或者取消关注
    //beego.Router("/book/score/:id",&controllers.UserController{},"*:Score")      //评分
    //beego.Router("/book/comment/:id",&controllers.UserController{},"post:Comment")      //评论
    //
    ////个人设置模块
    //beego.Router("/setting",&controllers.SettingController{},"*:Index")
    //beego.Router("/setting/upload",&controllers.SettingController{},"*:Upload")
    //
    ////管理员后台
    //beego.Router("/manager/category",&controllers.ManagerController{},"post,get:Category")
    //beego.Router("/manager/update-cate",&controllers.ManagerController{},"get:UpdateCate")
    //beego.Router("/manager/del-cate",&controllers.ManagerController{},"get:DelCate")
    //beego.Router("/manager/icon-cate",&controllers.ManagerController{},"get:UpdateCateIcon")

}
