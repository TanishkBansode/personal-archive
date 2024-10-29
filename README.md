# Personal-archive

A minimalist personal blog-like application built with Go and Gin, allowing users to view a list of title of writings  and navigate to individual writing.

![Preview](https://github.com/user-attachments/assets/26341c18-2c8d-42aa-a316-5b5163bd7035)


## Features

- Displays a list of titles of writings
- Each title links to a dedicated page for detailed content
- Easy to extend with new writings.

## Project Structure

```
/templates
  |-- index.html    // Main page listing  titles
  |-- writing.html     // Template for individual writings
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
   git clone https://github.com/TanishkBansode/personal-archive.git
   cd personal-archive
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `writings` directory in the project root and add `.txt` files for each blog post. The filename (without the extension) will be the title.

## Running the Application

Start the server:
```bash
go run main.go
```

Visit `http://localhost:8080/` in your web browser to view the writings application.

## How to Add a New Writing

1. Create a new `.txt` file in the `writings` directory
2. The filename should be the title of the writing (e.g., `my-new-lit.txt`)
3. Write the content of your writing inside the `.txt` file

## Technologies Used

* Go
* Gin framework
* HTML/CSS

