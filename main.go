package main

import (
	"io/fs"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/app"
	"github.com/silverhand7/money-tracking-app/controllers"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/middleware"
	"github.com/silverhand7/money-tracking-app/models/web/requests/validators"
	"github.com/silverhand7/money-tracking-app/repositories"
	"github.com/silverhand7/money-tracking-app/services"
	"github.com/silverhand7/money-tracking-app/ui"
)

func frontendHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

func main() {
	db := app.NewDB()

	validate := validator.New()
	initValidator(*validate)

	router := httprouter.New()

	router.GET("/", frontendHandler)

	staticFS, _ := fs.Sub(ui.StaticFiles, "dist/static")
	router.ServeFiles("/static/*filepath", http.FS(staticFS))

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

	walletRepository := new(repositories.WalletRepository)
	walletService := services.WalletService{
		WalletRepository: walletRepository,
		DB:               db,
		Validate:         validate,
	}

	walletController := controllers.WalletController{
		WalletService: &walletService,
		UserService:   &userService,
	}
	router.GET("/api/wallets", middleware.AuthMiddleware(walletController.GetAll, db))
	router.POST("/api/wallets", middleware.AuthMiddleware(walletController.Create, db))
	router.GET("/api/wallets/:walletId", middleware.AuthMiddleware(walletController.FindById, db))
	router.PUT("/api/wallets/:walletId", middleware.AuthMiddleware(walletController.Update, db))
	router.DELETE("/api/wallets/:walletId", middleware.AuthMiddleware(walletController.Delete, db))

	transactionRepository := new(repositories.TransactionRepository)
	transactionService := services.TransactionService{
		TransactionRepository: transactionRepository,
		DB:                    db,
		Validate:              validate,
	}

	transactionController := controllers.TransactionController{
		TransactionService: &transactionService,
		UserService:        &userService,
		WalletService:      &walletService,
	}
	router.GET("/api/transactions", transactionController.GetAll)
	router.POST("/api/transactions", transactionController.Create)
	router.GET("/api/transactions/:transactionId", transactionController.FindById)
	router.PUT("/api/transactions/:transactionId", transactionController.Update)
	router.DELETE("/api/transactions/:transactionId", transactionController.Delete)

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
