package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 创建功能：
// 录入新的学生：
func insertStudent(student Students) error {
	sqlStmt := "INSERT INTO students (name, birthday) VALUES (?,?)"
	result, err := db.Exec(sqlStmt, student.Name, student.Birthday)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}

// 创建用户：
func insertUsers(users Users) error {
	if users.Account == "" || users.Password == "" {
		return errors.New("account or Password is empty")
	}
	if len(users.SecurityQuestion) < 3 {
		return fmt.Errorf("security question is too short")
	}
	if len(users.SecurityAnswer) < 3 {
		return fmt.Errorf("security answer is too short")
	}

	sqltmt := "INSERT INTO users(account,password,security_question,security_answer) VALUES (?,?,?,?);"
	result, err := db.Exec(sqltmt, users.Account, users.Password, users.SecurityQuestion, users.SecurityAnswer)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// 更新功能：
// 更新学生全部信息：
func updateStudent(student Students) error {
	err := DBupdate("UPDATE students SET name =?, birthday =? WHERE id =?", student)
	return err
}

// 更新学生姓名：
func updateStudentName(student Students) error {
	err := DBupdate("UPDATE students SET name =? WHERE id =?", student)
	return err
}

// 更新学生生辰：
func updateStudentBirthday(student Students) error {
	err := DBupdate("UPDATE students SET birthday =? WHERE id =?", student)
	return err
}

// 更新密码:
func updataPassword(user Users) error {
	sqlStmt := "UPDATE users SET password =? WHERE account =?"
	result, err := db.Exec(sqlStmt, user.Password, user.Account)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	return err
}

// 查询功能：
// 通过id查询学生：
func getStudentById(Id int) (Students, error) {
	var student Students
	sqlStmt := "SELECT name, birthday ,id FROM students WHERE id=?"
	row := db.QueryRow(sqlStmt, Id)
	err := row.Scan(&student.Name, &student.Birthday, &student.Id)
	return student, err
}

// 通过账户查询密码
func getUsersPasswordByAccount(account, password string) error {
	var queriedPassword string

	sqlStmt := "SELECT password FROM users WHERE account=?"
	row := db.QueryRow(sqlStmt, account)
	err := row.Scan(&queriedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("用户名不存在")
		}
		return errors.New("查询密码时出现数据库错误")
	}

	if queriedPassword != password {
		return errors.New("密码错误")
	}

	return nil
}

// 通过账户查询密保问题
func getSecurityQuestionByAccount(Account string) (Users, error) {
	var user Users
	sqlStmt := "SELECT security_question FROM users WHERE account=?"
	row := db.QueryRow(sqlStmt, Account)
	err := row.Scan(&user.SecurityQuestion)
	return user, err
}

// 通过账户查询密保答案
func getUsersSecurityAnswerByAccount(account, SecurityAnswer string) error {
	var queriedSecurityAnswer string

	sqlStmt := "SELECT security_answer FROM users WHERE account=?"
	row := db.QueryRow(sqlStmt, account)
	err := row.Scan(&queriedSecurityAnswer)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("用户名不存在")
		}
		return errors.New("查询答案时出现数据库错误")
	}

	if queriedSecurityAnswer != SecurityAnswer {
		return errors.New("答案错误")
	}

	return nil
}

// 移除功能：
// 移除学生信息：
func deleteStudent(id int) error {
	err := DBdelete("DELETE FROM students WHERE id =?", id)
	return err
}

// 删除功能：
// 删除学生姓名：
func deleteStudentName(id int) error {
	err := DBdelete("UPDATE students SET name = '' WHERE id = ?;", id)
	return err
}

// 删除学生生辰：
func deleteStudentBirthday(id int) error {
	err := DBdelete("UPDATE students SET birthday = '' WHERE id = ?;", id)
	return err
}
