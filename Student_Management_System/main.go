package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	Router := gin.Default()

	//登入前功能
	Router.POST("/register", register)                                  //注册
	Router.POST("/login", loginByPassword)                              //密码登入
	Router.POST("/logout", logout)                                      //登出
	Router.POST("/loginBySecurityAnswer", loginBySecurityAnswer)        //密保登录
	Router.GET("/getSecurityQuestionByAccount", SearchSecurityQuestion) //查询密保问题
	//{"account":"123","password":"123","security_question":"123","security_answer":"123"}

	protectedRouter := Router.Group("/")
	protectedRouter.Use(IsLogin())
	{
		//登入后功能
		protectedRouter.POST("/add", Add)                                       //录入学生
		protectedRouter.POST("/updataAll", UpdateAll)                           //更新学生信息
		protectedRouter.POST("/updataName", UpdateName)                         //更新学生姓名
		protectedRouter.POST("/updataBrithday", UpdateBrithday)                 //更新学生生辰
		protectedRouter.GET("/search", Search)                                  //查找学生信息
		protectedRouter.DELETE("/delete", DeleteStudent)                        //移除学生
		protectedRouter.DELETE("/deleteStudentName", DeleteStudentName)         //删除学生姓名
		protectedRouter.DELETE("/deleteStudentBirthday", DeleteStudentBirthday) //删除学生生辰
		protectedRouter.POST("/updataPassword", UpdataPassword)                 //登陆后修改密码
	}

	err = Router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
