package controllers

import (
	"GoWebApi/cacheSingleton"
	"GoWebApi/models"
	"fmt"
	"net/http"
	"reflect"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

// Person API for testing
type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {

}

func (this *MainController) Finish() {

}

// @Description Get person information by id
// @Success 200 {object} models.Person
// @Param id	path	int	true	"the id of person"
// @Param name	query	string false	"the name"
// @Failure 404 not found
// @Failure 400 bad request
// @router /:id [get]
func (this *MainController) Get() {
	logs.GetLogger().Println("entering main controller, method: GET")
	id := this.Ctx.Input.Params()[":id"]

	beego.Info(reflect.TypeOf(id))

	p := models.Person{
		ID:    id,
		Name:  "godking",
		Age:   10,
		Title: "SE",
	}
	name := this.Ctx.Input.Query("name")
	beego.Info(this.Ctx.Request)
	if name == "" {
		// this.Ctx.WriteString("name is empty")
		//http.Error(output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		http.Error(this.Ctx.ResponseWriter, "name parameter should not be empty", http.StatusBadRequest)
		return
	}
	beego.Info(name)
	bm := cacheSingleton.Get()
	godking := bm.Get("godking")
	fmt.Println(godking)
	this.Ctx.Output.JSON(&p, false, false)
	logs.GetLogger("main").Println("leaving main controller, method: GET")
}
