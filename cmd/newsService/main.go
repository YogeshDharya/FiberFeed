package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
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
func main(){
	//Load Environment Variable 
	if err := godotenv.Load(); err != nil {
		log.Fatal("Changed Loading .env File !")
	}
	engine := html.New(public, ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	apiKey := os.Getenv("API_TOKEN")
	//port := os.Getenv("FIBER_PORT")
	if apiKey == "" {
		log.Fatal("API key not Set !")
	}
	log.Print("News Blog Works Fine Now !")
	app.Get("/news",func(c *fiber.Ctx) error {
		news, err := fetchNews(apiKey)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
        }
	return c.Render("index",news)
	})
	app.Static("/","./public")
	log.Fatal(app.Listen(":3000"))
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