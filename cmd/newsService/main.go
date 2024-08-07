package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"github.com/spf13/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)
type ApiResponse struct{
	Status string `json:"status"`
	TotalResults int `json:"totalResults"`
	News    []struct {
        ID          string `json:"id"`
        Title       string `json:"title"`
        Description string `json:"description"`
        URL         string `json:"url"`
        Author      string `json:"author"`
        Image       string `json:"image"`
        Language    string `json:"language"`
        Category    []string `json:"category"`
        Published   string `json:"published"`
    } `json:"news"`
} 
func initConfig(){
	viper.SetConfigType(".env")
	viper.AddConfigPath("../.env")
//	viper.AutomaticEnv()
//	viper.SetEnvPrefix("app");
	viper.BindEnv("NEWS_SERVICE_PORT")
	viper.BindEnv("API_TOKEN")
	err := viper.ReadInConfig()
	if err != nil{
		panic("Error Reading Config env file : " + err.Error())
	}
}
func main(){
	//Load Environment Variable 
  //no need to have absolute file path anymore ?? 
	initConfig()
	
	publicDir, err := filepath.Abs("./public")
	if err != nil {
		log.Fatal("Can't resolve html template directory !")
	}
	engine := html.New(publicDir, ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	apiKey := viper.GetString("API_TOKEN")
	port := viper.GetString("FIBER_PORT")
	if apiKey == "" {
		log.Fatal("API key not Set !")
	}
	log.Println("News Blog Works Fine Now !")
	app.Get("/news",func(c *fiber.Ctx) error {
		news, err := fetchNews(apiKey)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }
	return c.Render("index",news)
	})
	app.Static("/","./public")
	log.Fatal(app.Listen(port))
}
func fetchNews(APIkey string) (*ApiResponse, error){
	url := "https://api.currentsapi.services/v1/latest-news?apiKey=" + APIkey 
	res, err := http.Get(url);
	if err != nil {
		return nil , err 
	} 
	defer res.Body.Close() 
	var newsResponse ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&newsResponse); err != nil {
		return nil , err
	}
	return &newsResponse,nil 
}