// package main;

// import(
// 	"log"
// 	"github.com/markbates/goth"	
// 	"github.com/markbates/goth/gothic"	
// 	"github.com/markbates/goth/providers/google"	
// 	"github.com/markbates/goth/providers/twitter"	
// 	"github.com/spf13/viper"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middeware/session"
// )

// func initConfig(){
// 	viper.SetConfigFile(".env")
// 	viper.SetConfigType("env")
// 	viper.AddConfigPath("../../")
// 	viper.AutomaticEnv()
// 	viper.BindEnv("USER_SERVICE_PORT")
// 	viper.BindEnv("CLIENT_ID")
// 	viper.BindEnv("CLIENT_SECRET")
// 	if err:= viper.ReadInConfig() ; err !=nil {
// 		log.Fatal("Error Reading Config Env !")
// 	}
// }
// func main(){
// 	initConfig()

// 	authPort := viper.GetString("USER_SERVICE_PORT")
// 	client := viper.GetString("CLIENT_ID")
// 	secret := viper.GetString("CLIENT_SECRET")
// 	app:= fiber.New()
// 	goth.UseProviders(
// 		google.New(client,secret,"http://localhost")
// 	)
// 	app.Get("/register",  Register)
// 	app.Get("/login", Login)
// //	log.Fatal(app.Listen(fmt.Sprintf(":%s",port)))
// 	err := app.Listen(":"+authPort)
// 	if err != nil {
// 		log.Fatal("Port Listening Error")
// 	}
// }