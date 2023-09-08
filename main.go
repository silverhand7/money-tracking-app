package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/app"
	"github.com/silverhand7/money-tracking-app/controllers"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/models/web/requests/validators"
	"github.com/silverhand7/money-tracking-app/repositories"
	"github.com/silverhand7/money-tracking-app/services"
)

func main() {
	db := app.NewDB()

	validate := validator.New()
	initValidator(*validate)

	router := httprouter.New()

	userRepository := new(repositories.UserRepository)
	userService := services.UserService{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}

	userController := controllers.UserController{
		UserService: &userService,
	}
	router.GET("/api/users", userController.GetAll)
	router.POST("/api/users", userController.Create)
	router.GET("/api/users/:userId", userController.FindById)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	categoryRepository := new(repositories.CategoryRepository)
	categoryService := services.CategoryService{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}

	categoryController := controllers.CategoryController{
		CategoryService: &categoryService,
	}
	router.GET("/api/categories", categoryController.GetAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exceptions.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func initValidator(v validator.Validate) {
	// Register the custom validation function.
	v.RegisterValidation("passwordValidator", validators.PasswordValidator)
	v.RegisterValidation("categoryTypeValidator", validators.ValidateType)
}
