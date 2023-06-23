### TinyHS
Simple Go HTTP Server for Hosting Angular Applications

### Description
This Go program allows you to start a simple HTTP server to host Angular web applications. It provides support for Cross-Origin Resource Sharing (CORS) and allows you to configure the hosting folder and port from the command line.

### Usage
server [options]

### Options
  --dir    Hosting directory. (Default: current directory)
  --port   Server port. (Default: 8080)

### Examples
  ths --dir=/path/to/your/angular-app --port=8000
    - Starts the server on port 8000, hosting the Angular application in /path/to/your/angular-app.

### Build
go build
