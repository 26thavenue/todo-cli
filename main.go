package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/26thavenue/todo-cli/todo"
)

const todoFile = ".todo.json"

func main() {

	add:=  flag.Bool("add", false, "add a new task")

	complete := flag.Int("complete", 0, "mark a todo as completed")

	delete := flag.Int("del", 0, "delete a todo")

	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch{
		case *add:
			task,err := getInput(os.Stdin, flag.Args()...)

			if err != nil{
				fmt.Fprintln(os.Stderr,err.Error())
				os.Exit(1)
			}

			todos.Add(task)

			err = todos.Save(todoFile)

			if err != nil{
				fmt.Fprintln(os.Stderr,err.Error())
				os.Exit(1)
			}

		case *complete > 0:
			err := todos.Complete(*complete)

			if err != nil{
				fmt.Fprintln(os.Stderr,err.Error())
				os.Exit(1)
			}
			err = todos.Save(todoFile)

			if err != nil{
				fmt.Fprintln(os.Stderr,err.Error())
				os.Exit(1)
			}

		case *delete > 0:
			err := todos.Delete(*delete)

			if err != nil{
				fmt.Fprintln(os.Stderr,err.Error())
				os.Exit(1)
			}
			err = todos.Save(todoFile)

			if err != nil{
				fmt.Fprintln(os.Stderr,err.Error())
				os.Exit(1)
			}

		case *list:
			todos.Print()
			
		default:
			fmt.Fprintln(os.Stdout,"invalid command")
			os.Exit(0)
	}
}


func getInput(r io.Reader, args ...string) (string, error){
	if len(args) > 0{
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil{
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo")
	}

	return text, nil
}