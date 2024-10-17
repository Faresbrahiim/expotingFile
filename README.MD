# ASCII Art Web

## Description
ASCII Art Web is a web-based application that renders text as ASCII art. It provides a simple and interactive way for users to generate text-based artwork using pre-defined ASCII banners.

## Authors
- **[Youssef Hayyani]** 
- **[Ibrahim Fares]** 

m
## Usage: How to Run
1. **Install Dependencies**: Make sure you have Go installed on your system.
2. **Run the Web Server**:
   - Open your terminal in the project directory.
   - Run the following command:
     ```bash
     go run .
     ```
   - The web server will start, typically on port `8080` (you can adjust this in the code if needed).
3. **Access the Application**:
   - Open your browser and go to: `http://localhost:8080`
   - Input text in the web interface to generate ASCII art.
   
## Implementation Details: Algorithm
The core algorithm behind the ASCII Art Web application is based on the idea of transforming input text into ASCII art using pre-defined "banners" stored in the `banners/` directory.

- **Banners**: The banners directory contains different ASCII styles, where each character is represented by an art equivalent. When a user inputs a string, the application reads the corresponding ASCII banner for each character and combines them.
  
- **Template Rendering**: The application uses Go's `html/template` package to dynamically render the HTML page. The templates are stored in the `templates/` directory.

- **HTTP Server**: The Go `net/http` package is used to handle the web server functionality. It listens for user requests, processes them by generating the ASCII art, and serves the final result as a rendered HTML page.

