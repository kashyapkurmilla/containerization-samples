package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		currentTime := time.Now()
		greeting := getGreeting(currentTime)
		date := currentDate(currentTime)
		quote := getRandomQuote()

		htmlResponse := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Welcome</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f4f4f4;
						margin: 50px;
						padding: 20px;
					}
					h1 {
						color: #333;
					}
					p {
						color: #555;
					}
					blockquote {
						font-style: italic;
						color: #777;
					}
				</style>
				<script>
					function updateTime() {
						var clockElement = document.getElementById('clock');
						var currentTime = new Date();
						var hours = currentTime.getHours();
						var minutes = currentTime.getMinutes();
						var seconds = currentTime.getSeconds();

						minutes = (minutes < 10 ? "0" : "") + minutes;
						seconds = (seconds < 10 ? "0" : "") + seconds;

						var timeString = hours + ":" + minutes + ":" + seconds;
						clockElement.textContent = timeString;
					}

					setInterval(updateTime, 1000);
				</script>
			</head>
			<body>
				<h1>Welcome!</h1>
				<p id="clock" style="font-size: 2em;"></p>
				<p>Date: %s</p>
				<p>Day: %s</p>
				<p>%s</p>
				<blockquote>%s</blockquote>
			</body>
			</html>
		`, date, currentTime.Weekday().String(), greeting, quote)

		return c.HTML(http.StatusOK, htmlResponse)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func getGreeting(t time.Time) string {
	hour := t.Hour()
	switch {
	case hour >= 5 && hour < 12:
		return "Good morning!"
	case hour >= 12 && hour < 18:
		return "Good afternoon!"
	default:
		return "Good evening!"
	}
}

func currentDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func getRandomQuote() string {
	quotes := []string{
		"Be yourself; everyone else is already taken. - Oscar Wilde",
		"In the middle of difficulty lies opportunity. - Albert Einstein",
		"The only limit to our realization of tomorrow will be our doubts of today. - Franklin D. Roosevelt",
		"Believe you can and you're halfway there. - Theodore Roosevelt",
		"Strive not to be a success, but rather to be of value. - Albert Einstein",
		"Everything has beauty, but not everyone can see. - Confucius",
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}
