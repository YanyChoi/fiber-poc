package routers

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/YanyChoi/fiber-poc/internal"
	"github.com/gofiber/fiber/v2"
)


func addOAuthRouters(app *fiber.App) {

	internal.GithubConfig()
	
	router := app.Group("")
	router.Get("/github_login", GithubLogin)
	router.Get("/github_callback", GithubCallback)
}

func GithubLogin(c *fiber.Ctx) error {
	url := internal.AppConfig.GithubLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GithubCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	githubcon := internal.GithubConfig()

	token, err := githubcon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Error exchanging code for token")
	}
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return c.SendString("Error creating request")
	}
	request.Header.Set("Authorization", "token "+token.AccessToken)
	request.Host = "api.github.com"
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return c.SendString("Error getting user from Github")
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
	return c.SendString("Welcome to Github, " + token.AccessToken)
}
