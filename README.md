

# Todo CLI

The Todo CLI is a command-line application written in Go that allows users to manage their tasks efficiently. With this tool, users can add, mark complete, delete, and list tasks directly from the terminal, making it easy to stay organized and productive.

## Features

- **Task Management**: Add new tasks, mark tasks as completed, and delete tasks as needed.
- **Timestamps**: Keep track of when tasks were created and completed.
- **Persistence**: Tasks are saved to a file, allowing users to persist their task list across sessions.
- **Simple and Intuitive Interface**: The CLI provides a clean and easy-to-use interface for managing tasks.

## Installation

To install the Todo CLI, you can use the `go get` command:

```bash
go get github.com/26thavenue/todo-cli
```

## Build

To build run go build in the root directory ad replace the './path to built file' with it


## Usage

### Adding a Task

To add a new task, use the `add` command followed by the task description:

```bash
./path to built file -add 'Buy groceries'
```

### Completing a Task

To mark a task as completed, use the `complete` command followed by the task index:

```bash
./path to built file -complete 1
```

### Deleting a Task

To delete a task, use the `delete` command followed by the task index:

```bash
./path to built file -del 1
```

### Listing Tasks

To list all tasks, simply run the `list` command:

```bash
./path to built file -list
```

## Note
Instead of using the 'path to built file' path you can use `go run main.go -add` or `make run -add` to perform the add operation, any other flag operation can be made by changing the add flag also


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
