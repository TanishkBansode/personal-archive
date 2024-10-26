# Personal-archive

A simple personal blog-like application built with Go and Gin, allowing users to view a list of title of writings  and navigate to individual writing.

## Features

- Displays a list of titles of writings
- Each title links to a dedicated page for detailed content
- Easy to extend with new writings.

## Project Structure

```
/templates
  |-- index.html    // Main page listing blog titles
  |-- blog.html     // Template for individual blog pages
.env
go.mod
go.sum
main.go
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Gin framework

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/blog-application.git
   cd personal-archive
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `blogs` directory in the project root and add `.txt` files for each blog post. The filename (without the extension) will be the title.

4. Create a `.env` file if needed for configuration settings.

## Running the Application

Start the server:
```bash
go run main.go
```

Visit `http://localhost:8080/` in your web browser to view the blog application.

## How to Add a New Blog Post

1. Create a new `.txt` file in the `blogs` directory
2. The filename should be the title of the blog (e.g., `my-new-blog.txt`)
3. Write the content of your blog inside the `.txt` file

## Technologies Used

* Go
* Gin framework
* HTML/CSS

