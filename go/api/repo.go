package main

import "fmt"

var (
	todos     Todos
	currentID int
)

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	return Todo{}
}

func RepoFindTodoNotyet() Todos {
	var todosNotyet Todos

	for _, t := range todos {
		if !t.Completed {
			todosNotyet = append(todosNotyet, t)
		}
	}

	return todosNotyet
}

func RepoCreateTodo(t Todo) Todo {
	currentID += 1
	t.ID = currentID
	todos = append(todos, t)
	return t
}

func RepoTodoCompleted(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = true
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
