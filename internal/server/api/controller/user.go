// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package controller

import (
	"encoding/json"
	"github.com/jrmsdev/alfred/internal/server/api/model"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {
	var user model.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := model.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

func (u *UserController) GetAll() {
	users := model.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := model.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user model.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := model.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	model.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if model.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
