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

	// Get the blog directory and read blogs
	blogDir := getBlogDir()
	blogs := readBlogs(blogDir)

	// Route to render the main page with blog titles
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Writings": blogs,
		})
	})

	// Route to render individual blog posts
	router.GET("/blogs/:title", func(c *gin.Context) {
		title := c.Param("title")
		content, exists := blogs[title]

		if !exists {
			c.HTML(http.StatusNotFound, "blog.html", gin.H{
				"Title":   "Writing Not Found",
				"Content": "Sorry, the stuff you're looking for does not exist.",
			})
			return
		}

		c.HTML(http.StatusOK, "blog.html", gin.H{
			"Title":   title,
			"Content": content,
		})
	})

	// Start the server
	router.Run(":8080")
}

// Get the directory path for blogs
func getBlogDir() string {
	dir, err := os.Getwd()
	CheckNilError(err)
	return filepath.Join(dir, "..", "blogs")
}

// Reads blog files and stores them as a map[title]content
func readBlogs(blogDir string) map[string]string {
	files, err := os.ReadDir(blogDir)
	CheckNilError(err)

	blogs := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			title := strings.TrimSuffix(file.Name(), ".txt")
			path := filepath.Join(blogDir, file.Name())

			content, err := os.ReadFile(path)
			CheckNilError(err)

			blogs[title] = string(content)
		}
	}

	return blogs
}

// Handle errors, it pretty handy
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
