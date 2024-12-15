package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func login(client *gin.Context, funcDelete func(account, password string) error) {
	var user Users
	if err := client.ShouldBind(&user); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := funcDelete(user.Account, user.Password)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.SetCookie("islogin", "true", 60, "/", "", false, true)
	client.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}

func Delete(client *gin.Context, funcDelete func(id int) error) {
	idstr := client.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = funcDelete(id)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusOK, gin.H{"message": "学生信息删除成功"})
}

func Updata(client *gin.Context, funcUpdate func(student Students) error) {
	var student Students
	if err := client.ShouldBind(&student); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := funcUpdate(student)
	if err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusOK, gin.H{"message": "学生信息更新成功"})
}

func DBdelete(cmdstr string, id int) error {
	result, err := db.Exec(cmdstr, id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}

func DBupdate(cmdstr string, student Students) error {
	result, err := db.Exec(cmdstr, student.Name, student.Birthday, student.Id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}
