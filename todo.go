package main

import "github.com/google/uuid"

// TodoItem basic structure
type TodoItem struct {
	ID        string
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
	URL       string `json:"url"`
}

// Todo holding all the state
type Todo map[string]*TodoItem

func (t Todo) nextID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

// Create a new TodoItem and add it to the list
func (t Todo) Create(item TodoItem, baseURL string) *TodoItem {
	item.ID = t.nextID()
	item.URL = "http://" + baseURL + "/" + item.ID
	t[item.ID] = &item
	return &item
}

// GetAll will return all TodoItems
func (t Todo) GetAll() []*TodoItem {
	items := []*TodoItem{}
	for _, item := range t {
		items = append(items, item)
	}
	return items
}

// Get will return a single item based on the Id
func (t Todo) Get(ID string) *TodoItem {
	return t[ID]
}

// Update single item
func (t Todo) Update(ID string, update TodoItem) {
	item := t.Get(ID)
	if len(update.Title) > 0 {
		item.Title = update.Title
	}

	if update.Order != item.Order {
		item.Order = update.Order
	}

	if update.Completed != item.Completed {
		item.Completed = update.Completed
	}

	t[ID] = item
}

// DeleteAll removes all Todos
func (t Todo) DeleteAll() {
	for k := range t {
		delete(t, k)
	}
}

// Delete a single Todo
func (t Todo) Delete(ID string) *TodoItem {
	result := t.Get(ID)
	if result == nil {
		return nil
	}

	delete(t, ID)
	return result
}
