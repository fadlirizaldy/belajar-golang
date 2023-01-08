package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID			string		`json:"id"`
	Item		string		`json:"item"`
	Completed	bool		`json:"completed"`
}

var todos = []todo{
	{ID:"1", Item:"Study Golang", Completed:false},
	{ID:"2", Item:"Eat Noodles", Completed:false},
	{ID:"3", Item:"Skripsi", Completed:false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

//add new todo list
func addTodo(context *gin.Context){
	var newTodo todo

	//bind json ini akan menangkap whatever request body
	if err:= context.BindJSON(&newTodo); err != nil{
		return
	}

	todos = append(todos, newTodo)
	//lalu kirimkan status created dan data yang udah dikirim
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// search todo by id
func findTodoById(id string) (*todo, error){
	for i, tod := range todos {
		if tod.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

//update data, hanya complete statusnya aja
func toggleTodoStatus(context *gin.Context){
	id:=context.Param("id")
	todo, err := findTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodo(context *gin.Context){
	id:=context.Param("id")
	todo, err := findTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// Delete Todo List
func deleteTodo(context *gin.Context){
	id:=context.Param("id")
	for i, tod := range todos {
		if tod.ID == id {
			todos = append(todos[:i], todos[i + 1:]...)
			break
		}
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Success to delete data"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:9090")
}