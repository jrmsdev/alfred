// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package controller

import (
	"encoding/json"
	"github.com/jrmsdev/alfred/internal/server/core/model"

	"github.com/astaxie/beego"
)

type ObjectController struct {
	beego.Controller
}

func (o *ObjectController) Post() {
	var ob model.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := model.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := model.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

func (o *ObjectController) GetAll() {
	obs := model.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob model.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := model.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	model.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
