package controllers

import (
	pb "GoWebApi/rpcdatacontract"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/astaxie/beego"
)

const (
	address = "localhost:60000"
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
// func (this *MainController) Get() {
// 	logs.GetLogger().Println("entering main controller, method: GET")
// 	id := this.Ctx.Input.Params()[":id"]

// 	beego.Info(reflect.TypeOf(id))

// 	p := models.Person{
// 		ID:    id,headers
// 		Name:  "godking",
// 		Age:   10,
// 		Title: "SE",
// 	}
// 	name := this.Ctx.Input.Query("name")
// 	beego.Info(this.Ctx.Request)
// 	if name == "" {
// 		// this.Ctx.WriteString("name is empty")
// 		//http.Error(output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
// 		http.Error(this.Ctx.ResponseWriter, "name parameter should not be empty", http.StatusBadRequest)
// 		return
// 	}
// 	beego.Info(name)
// 	bm := cacheSingleton.Get()
// 	godking := bm.Get("godking")
// 	fmt.Println(godking)
// 	this.Ctx.Output.JSON(&p, false, false)
// 	logs.GetLogger("main").Println("leaving main controller, method: GET")
// }

// @router /rpc
func (this *MainController) Get() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Process(ctx, &pb.Request{Question: "Ask a question"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// log.Printf("Greeting: %s", r.GetTables())

	// data, _ := json.MarshalIndent(r.GetTables(), "", "    ")
	// log.Printf("%s\n", data)

	data := r.GetTables()

	tables := make([]TableData, 0)

	for _, table := range data {
		columns := table.GetColumns()
		rows := table.GetRows()
		headers := getHeaders(columns)
		body := getBody(columns, rows)

		table := TableData{Header: headers, Body: body}
		tables = append(tables, table)
	}

	this.Data["json"] = &tables
	this.ServeJSON()
}

func getHeaders(columns []*pb.Column) []map[string]string {
	var headers []map[string]string
	for _, col := range columns {
		header := make(map[string]string)
		header["key"] = col.EnglishName
		header["title"] = col.ChineseName
		headers = append(headers, header)
	}
	return headers
}

func getBody(columns []*pb.Column, rows []*pb.Row) []map[string]interface{} {
	var resultRows []map[string]interface{}
	for _, row := range rows {
		resultRow := make(map[string]interface{})
		for index, col := range columns {
			cell := row.GetCells()[index]

			_, ok := cell.GetValue().(*pb.Cell_ValInteger)
			if ok {
				fmt.Println("is integer")
				resultRow[col.EnglishName] = cell.GetValInteger()
			} else {
				fmt.Println("not integer")
				resultRow[col.EnglishName] = cell.GetValString()
			}

			// resultRow[col.EnglishName] = cell.GetValue().(pb.Cell_ValInteger) ? cell.GetValInteger() : cell.GetValString()
		}
		resultRows = append(resultRows, resultRow)
	}
	return resultRows
}

type TableData struct {
	Body   []map[string]interface{} `json:"body"`
	Header []map[string]string      `json:"header"`
}
