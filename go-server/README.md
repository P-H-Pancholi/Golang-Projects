# Go server project

In this project, I have created a simple golang HTTP server that listens on port 8000.
The server has 3 routes:
1. \ : It will return content of index.html page 
2. \hello : It will execute the `helloHandler()`
3. \form.html : It will output form according to static/form.html & execute `formHandler()`

 - - - -

In this project, the modules used are,

1. [net/http](https://pkg.go.dev/net/http@go1.21.4)