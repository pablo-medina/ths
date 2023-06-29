### TinyHS
Simple Go HTTP/HTTPS Server designed to Host Angular Applications

### Description
This Go program allows you to start a simple HTTP server to host Angular web applications. It provides support for Cross-Origin Resource Sharing (CORS) and allows you to configure the hosting folder and port from the command line.

### Usage
ths [options]

### Options
  --dir    Hosting directory. (Default: current directory)
  --port   Server port. (Default: 8080)
  --cors   Enable CORS. (Default: disabled)
  --https  Enable HTTPS. Requires certificates path definition.
  --cert   Certificate File Path (for HTTPS).
  --key    Certificate Key File Path (for HTTPS).  

### Examples
  ths --dir=/path/to/your/angular-app --port=8000
    - Starts the server on port 8000, hosting the Angular application in /path/to/your/angular-app.

### Build
go build
