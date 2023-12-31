package main

import (
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/silverhand7/money-tracking-app/app"
	"github.com/silverhand7/money-tracking-app/controllers"
	"github.com/silverhand7/money-tracking-app/database/seeders"
	"github.com/silverhand7/money-tracking-app/exceptions"
	"github.com/silverhand7/money-tracking-app/middleware"
	"github.com/silverhand7/money-tracking-app/models/web/requests/validators"
	"github.com/silverhand7/money-tracking-app/repositories"
	"github.com/silverhand7/money-tracking-app/services"
	"github.com/silverhand7/money-tracking-app/ui"
)

func frontendHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

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

	isSeeder := flag.String("seeder", "", "")
	flag.Parse()

	if string(*isSeeder) == "true" {
		seed := seeders.NewSeed(db)
		seed.ExecuteSeeder()
	}

	validate := validator.New()
	initValidator(*validate)

	router := httprouter.New()

	router.GET("/", frontendHandler)
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		frontendHandler(w, r, nil)
	})

	// static files
	staticFS, _ := fs.Sub(ui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	router.Handler(http.MethodGet, "/assets/*filepath", httpFS)

	userRepository := new(repositories.UserRepository)
	userService := services.UserService{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}

	userController := controllers.UserController{
		UserService: &userService,
	}

	categoryRepository := new(repositories.CategoryRepository)
	categoryService := services.CategoryService{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}

	categoryController := controllers.CategoryController{
		CategoryService: &categoryService,
	}

	transactionRepository := new(repositories.TransactionRepository)

	walletRepository := new(repositories.WalletRepository)
	walletService := services.WalletService{
		WalletRepository:      walletRepository,
		TransactionRepository: transactionRepository,
		DB:                    db,
		Validate:              validate,
	}

	walletController := controllers.WalletController{
		WalletService: &walletService,
		UserService:   &userService,
	}

	transactionService := services.TransactionService{
		TransactionRepository: transactionRepository,
		WalletRepository:      walletRepository,
		CategoryRepository:    categoryRepository,
		DB:                    db,
		Validate:              validate,
	}

	transactionController := controllers.TransactionController{
		TransactionService: &transactionService,
		UserService:        &userService,
		WalletService:      &walletService,
	}

	authService := services.AuthService{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}

	loginController := controllers.UserLoginController{
		AuthService: &authService,
	}

	router.GET("/api/users", userController.GetAll)
	router.POST("/api/users", userController.Create)
	router.GET("/api/users/:userId", userController.FindById)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	router.POST("/api/login", loginController.Login)

	router.GET("/api/categories", categoryController.GetAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/wallets", walletController.GetAll)
	router.POST("/api/wallets", walletController.Create)
	router.GET("/api/wallets/:walletId", walletController.FindById)
	router.PUT("/api/wallets/:walletId", walletController.Update)
	router.DELETE("/api/wallets/:walletId", walletController.Delete)
	router.GET("/api/wallets/:walletId/transactions", walletController.GetWalletTransactions)

	router.GET("/api/transactions", transactionController.GetAll)
	router.POST("/api/transactions", transactionController.Create)
	router.GET("/api/transactions/:transactionId", transactionController.FindById)
	router.PUT("/api/transactions/:transactionId", transactionController.Update)
	router.DELETE("/api/transactions/:transactionId", transactionController.Delete)

	router.PanicHandler = exceptions.ErrorHandler

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.NewCorsMiddleware(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func initValidator(v validator.Validate) {
	v.RegisterValidation("passwordValidator", validators.PasswordValidator)
	v.RegisterValidation("categoryTypeValidator", validators.ValidateType)
}
