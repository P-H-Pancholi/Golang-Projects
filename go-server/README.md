# Go server project

In this project, I have created a simple golang HTTP server that listens on port 8000.
The server has 3 routes:
1. \ : It will return content of index.html page 
2. \hello : It will execute the `helloHandler()`
3. \form.html : It will output form according to static/form.html & execute `formHandler()`
 - - - -
## How to Run the project ?
1. Clone this Repo
2. `go mod init <module_name>` or place it inside the Go folder.
3. `go build` , it will create an executable binary as per your OS
4. Execute the binary or `go run main.go`
 - - - -
In this project, the modules used are,

1. [net/http](https://pkg.go.dev/net/http@go1.21.4)
2. [fmt](https://pkg.go.dev/fmt@go1.21.4)
3. [log](https://pkg.go.dev/log@go1.21.4)
 - - - -

Feel free to add any feedbacks