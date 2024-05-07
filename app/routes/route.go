package routes

import (
	"my-task-app/app/middlewares"
	_projectData "my-task-app/features/project/dataProject"
	_projectHandler "my-task-app/features/project/handler"
	_projectService "my-task-app/features/project/service"
	_taskData "my-task-app/features/task/dataTask"
	_taskHandler "my-task-app/features/task/handler"
	_taskService "my-task-app/features/task/service"
	_userData "my-task-app/features/user/data"
	_userHandler "my-task-app/features/user/handler"
	_userService "my-task-app/features/user/service"
	"my-task-app/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	projectData := _projectData.New(db)
	projectService := _projectService.New(projectData)
	projectHandlerAPI := _projectHandler.New(projectService)

	taskdata := _taskData.New(db)
	taskService := _taskService.New(taskdata)
	taskHandlerAPI := _taskHandler.New(taskService)

	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.PUT("/users/:id", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/users", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.POST("/login", userHandlerAPI.Login)

	//projectHandler
	e.GET("/projects", projectHandlerAPI.GetAllProject, middlewares.JWTMiddleware())
	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.PUT("/projects/:id", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
	e.DELETE("/projects/:id", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())
	e.GET("/projects/:id", projectHandlerAPI.GetByIdProject, middlewares.JWTMiddleware())

	//taskHandler
	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:id", taskHandlerAPI.UpdateTask, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:id", taskHandlerAPI.DeleteTask, middlewares.JWTMiddleware())
}
