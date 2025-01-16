package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
)

type ApiResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	News         []struct {
		ID          string   `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		URL         string   `json:"url"`
		Author      string   `json:"author"`
		Image       string   `json:"image"`
		Language    string   `json:"language"`
		Category    []string `json:"category"`
		Published   string   `json:"published"`
	} `json:"news"`
}

type SearchServiceI interface{
	Search(query String) ([]string,error)
}
type searchService struct{}
func main() {
	// Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Used the html/template engine with absolute path
	publicDir, err := filepath.Abs("./public")
	if err != nil {
		log.Fatal("Failed to get absolute path for public directory:", err)
	}
	engine := html.New(publicDir, ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	apiKey := os.Getenv("API_TOKEN")
	port := os.Getenv("FIBER_PORT")
	if apiKey == "" {
		log.Fatal("API key not set!")
	}
	log.Print("News Blog works fine now!")
	loggger := log.NewLogfmtLogger(log.NewsSyncWriter(os.Stdout))
	svc := searchService{} 
 
	app.Get("/news", func(c *fiber.Ctx) error {
		news, err := fetchNews(apiKey)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Render("index", news)
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		query := c.Query("q")
		if query == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Query parameter 'q' is required")
		}
		searchResults, err := searchNews(apiKey, query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Render("index", searchResults)
	})

	app.Static("/", "./public")

	log.Fatal(app.Listen(":" + port))
}
type searchRequest struct{
	Query string `json:"query"`
}
func createSearchEndpoint(svc searchService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(searchRequest)
		results, err := svc.Search(req.Query)
		return ApiResponse{Results: results, Err: err}, nil
	}
}
func decodeSearchRequest(_ context.Context, r *http.Request) (interface{} , error){
	//TODO 
	//we don't have an interface in the return statement 
	var req searchRequest
	err := json.NewDecoder(r.body).Decode(&req)
	return req
}
func encodeSearchRequest(_ context.Context, r *http.Request)
func fetchNews(APIkey string) (*ApiResponse, error) {
	url := "https://api.currentsapi.services/v1/latest-news?apiKey=" + APIkey
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var newsResponse ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&newsResponse); err != nil {
		return nil, err
	}
	return &newsResponse, nil
}

func (service searchService) searchNews(APIkey, query string) (*ApiResponse, error) {
	url := "https://api.currentsapi.services/v1/search?apiKey=" + APIkey + "&keywords=" + query
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var searchResponse ApiResponse
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}
	//TODO 
	//below value should be incorporated with returning value 4 go kit 
	//[]string{"Result1 for " + query, "Result2 for " + query}, nil
	return &searchResponse, nil
}
