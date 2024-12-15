package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 创建功能：

// Add 录入新的学生：
func Add(client *gin.Context) {
	var student Students
	if err := client.ShouldBind(&student); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := insertStudent(student)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusOK, gin.H{"message": "学生信息添加成功"})
}

// 注册用户：
func register(client *gin.Context) {
	var user Users
	if err := client.ShouldBind(&user); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := insertUsers(user)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusOK, gin.H{"message": "用户注册成功"})
}

// UpdateAll 更新学生全部信息：
func UpdateAll(client *gin.Context) {
	Updata(client, updateStudent)
}

// UpdateName 更新学生姓名：
func UpdateName(client *gin.Context) {
	Updata(client, updateStudentName)
}

// UpdateBrithday 更新学生生辰：
func UpdateBrithday(client *gin.Context) {
	Updata(client, updateStudentBirthday)
}

// UpdataPassword 更新用户密码:
func UpdataPassword(client *gin.Context) {
	var user Users
	if err := client.ShouldBind(&user); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := updataPassword(user)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	client.JSON(http.StatusOK, gin.H{"message": "密码重制成功"})
}

// Search 查询学生信息：
func Search(client *gin.Context) {
	idstr := client.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := getStudentById(id)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"message": "未找到该学生"})
		return
	}

	client.JSON(http.StatusOK, gin.H{"data": student})
}

// SearchSecurityQuestion 查询密保问题：
func SearchSecurityQuestion(client *gin.Context) {
	Account := client.Query("account")

	user, err := getSecurityQuestionByAccount(Account)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"message": "未找到该用户"})
		return
	}

	client.JSON(http.StatusOK, gin.H{"security_question": user.SecurityQuestion})
}

// DeleteStudent 移除学生：
func DeleteStudent(client *gin.Context) {
	Delete(client, deleteStudent)
}

// DeleteStudentName 删除学生姓名：
func DeleteStudentName(client *gin.Context) {
	Delete(client, deleteStudentName)
}

// DeleteStudentBirthday 删除学生生辰:
func DeleteStudentBirthday(client *gin.Context) {
	Delete(client, deleteStudentBirthday)
}

// loginByPassword 通过密码登录账户：
func loginByPassword(client *gin.Context) {
	login(client, getUsersPasswordByAccount)
}

// loginBySecurityAnswer 通过密保答案登录账户：
func loginBySecurityAnswer(client *gin.Context) {
	login(client, getUsersSecurityAnswerByAccount)
}

// IsLogin 检测是否登录；
func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieValue, err := c.Cookie("islogin")
		if err != nil || cookieValue != "true" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未登录，无权访问"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// logout 登出：
func logout(client *gin.Context) {
	client.SetCookie("islogin", "", -1, "/", "", false, true)
	client.JSON(http.StatusOK, gin.H{"message": "退出登录成功"})
}
