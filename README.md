# Cobol

Course in spanish by Pablo Tilotta [here](https://www.udemy.com/course/masters-desarrollo/learn/lecture/17171780)

Why GO?
* Easy to learn, clear and improved syntax
* No need to use ';'
* Takes easy advantage of all PC resources
* Doesn't allow the programmer to do bad practices
* Functions can return more than a value
* Can work with sync and async code
* Only 'FOR' exists for iterations
* Is not Object Oriented
* It has structures, functions, methods, interfaces...
* Variable scope is solved with lower and upper case variables
* Has 3 kinds of basic variables: numbers, strings and booleans

## Setup environment

Install using `apt` with this [link](https://github.com/golang/go/wiki/Ubuntu)

I used the following
```shell
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go -y
```

* Verify your installation
```shell
go version
```

* Execute but don't save the executable
```shell
go run main.go
```

* Build executable
```shell
go build main.go
```
