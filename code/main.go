package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the router and load templates
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// Get the writings directory and read writings
	writingDir := getWritingDir()
	writings := readWritings(writingDir)

	// Route to render the main page with writing titles
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Writings": writings,
		})
	})

	// Route to render individual writing pieces
	router.GET("/writings/:title", func(c *gin.Context) {
		title := c.Param("title")
		content, exists := writings[title]

		if !exists {
			c.HTML(http.StatusNotFound, "writing.html", gin.H{
				"Title":   "Writing Not Found",
				"Content": "Sorry, the writing you're looking for does not exist.",
			})
			return
		}

		c.HTML(http.StatusOK, "writing.html", gin.H{
			"Title":   title,
			"Content": content,
		})
	})

	// Start the server
	router.Run(":8080")
}

// Get the directory path for writings
func getWritingDir() string {
	dir, err := os.Getwd()
	CheckNilError(err)
	return filepath.Join(dir, "..", "writings")
}

// Reads writing files and stores them as a map[title]content
func readWritings(writingDir string) map[string]string {
	files, err := os.ReadDir(writingDir)
	CheckNilError(err)

	writings := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			title := strings.TrimSuffix(file.Name(), ".txt")
			path := filepath.Join(writingDir, file.Name())

			content, err := os.ReadFile(path)
			CheckNilError(err)

			writings[title] = string(content)
		}
	}

	return writings
}

// Handle errors gracefully
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
