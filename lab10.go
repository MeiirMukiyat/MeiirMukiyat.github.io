package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Todo struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Status bool    `json:"status"`
}

var todoList []Todo

func main() {
    router := gin.Default()

    // создание задачи
    router.POST("/todos", func(c *gin.Context) {
        var todo Todo
        if err := c.ShouldBindJSON(&todo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        todo.ID = strconv.Itoa(len(todoList) + 1)
        todoList = append(todoList, todo)
        c.JSON(http.StatusOK, todo)
    })

    // получение списка задач
    router.GET("/todos", func(c *gin.Context) {
        c.JSON(http.StatusOK, todoList)
    })

    // получение задачи по ID
    router.GET("/todos/:id", func(c *gin.Context) {
        id := c.Param("id")
        for _, todo := range todoList {
            if todo.ID == id {
                c.JSON(http.StatusOK, todo)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
    })

    // обновление задачи по ID
    router.PUT("/todos/:id", func(c *gin.Context) {
        id := c.Param("id")
        for i, todo := range todoList {
            if todo.ID == id {
                var updatedTodo Todo
                if err := c.ShouldBindJSON(&updatedTodo); err != nil {
                    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                    return
                }
                updatedTodo.ID = todo.ID
                todoList[i] = updatedTodo
                c.JSON(http.StatusOK, updatedTodo)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
    })

    // удаление задачи по ID
    router.DELETE("/todos/:id", func(c *gin.Context) {
        id := c.Param("id")
        for i, todo := range todoList {
            if todo.ID == id {
                todoList = append(todoList[:i], todoList[i+1:]...)
                c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
    })

    router.Run(":8080")
}
