package data

import "errors"

// Datastore is an in-memory database.
type Datastore struct {
	Todos []*Todo `json:"todos"`
}

// GetTodoByName returns the todo object with a specific name if it exists.
func (d *Datastore) GetTodoByName(name string) (*Todo, error) {
	for _, todo := range d.Todos {
		if todo.Name == name {
			return todo, nil
		}
	}

	return nil, errors.New("todo not found")
}

// GetTodos returns the todos list eventually filter by passed tags.
func (d *Datastore) GetTodos(tags ...string) []*Todo {
	selectedTodos := []*Todo{}

	for _, todo := range d.Todos {
		if todo.checkTags(tags...) {
			selectedTodos = append(selectedTodos, todo)
		}
	}

	return selectedTodos
}

// AddTodo creates a new todo object in the datastore.
func (d *Datastore) AddTodo(name string, tags ...string) error {
	_, err := d.GetTodoByName(name)
	if err != nil {
		if err.Error() == "todo not found" {
			d.Todos = append(d.Todos, &Todo{
				Name: name,
				Tags: tags,
			})

			return nil
		}
	}

	return errors.New("todo already exists")
}
