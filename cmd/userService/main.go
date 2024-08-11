package main;

import(
	"log"
	
	"github.com/spf13/viper"
	"github.com/gofiber/fiber/v2"
)

func initConfig(){
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()
	viper.BindEnv("USER_SERVICE_PORT")
	if err:= viper.ReadInConfig() ; err !=nil {
		log.Fatal("Error Reading Config Env !")
	}
}
func main(){
	initConfig()
	authPort := viper.GetString("USER_SERVICE_PORT")
	app:= fiber.New()
	app.Get("/register",  Register)
	app.Get("/login", Login)
//	log.Fatal(app.Listen(fmt.Sprintf(":%s",port)))
	err := app.Listen(":"+authPort)
	if err != nil {
		log.Fatal("Port Listening Error")
	}
}