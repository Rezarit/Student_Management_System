package main

type Students struct {
	Name     string `json:"name"     form:"name"`
	Id       int    `json:"id"       form:"id"        require:"id"`
	Birthday string `json:"birthday" form:"birthday"`
}

type Users struct {
	Account          string `json:"account"  form:"account"  require:"account"`
	Password         string `json:"password" form:"password" require:"password"`
	SecurityQuestion string `json:"security_question" form:"security_question" require:"security_question"`
	SecurityAnswer   string `json:"security_answer"   form:"security_answer"   require:"security_answer"`
}
