Github Crawler
This is a simple web scraping application written in Go. It uses the gin framework to handle HTTP requests and the soup library to parse HTML and fetch information.

The application starts a web server and listens for POST requests on the "/craw" route. When it receives a request, it begins crawling the Github repository at "https://github.com/golang-jwt/jwt". It gathers information such as the number of stars, forks, and contributors.

Dependencies
Gin - HTTP web framework used for handling HTTP requests.
Soup - Library for web scraping HTML data.
How to Run
First, ensure you have the necessary dependencies installed:

bash
Copy code
go get github.com/gin-gonic/gin
go get github.com/anaskhan96/soup
Then, run the Go program:

bash
Copy code
go run main.go
The server will start and listen for POST requests on the "/craw" route.

Usage
Send a POST request to the "/craw" route to start the scraping process. The application will send multiple concurrent requests to Github to gather the data. Once the data is collected, it will be returned as a JSON response with the following structure:

json
Copy code
{
  "stars": "number of stars",
  "forks": "number of forks",
  "contributor": "number of contributors"
}
Note: The actual values will be integers.

Functions
initRouter: Initializes the router and routes for the application.
crawGithub: Handler function for the "/craw" route. It creates goroutines to scrape the Github repository and waits for them to finish. Returns a JSON response with the data.
craw: Fetches the HTML data from the Github repository, parses it, and extracts the required information.
FindContriNo, FindStarsNo, FindForkNo: Helper functions that parse the HTML data and find the number of contributors, stars, and forks respectively.
Error Handling
The application returns a 500 Internal Server Error response if there's an error during the scraping process.

Note
This is a simple application and may not handle all edge cases or errors that could occur during the scraping process. Use it as a starting point for more complex web scraping projects.
