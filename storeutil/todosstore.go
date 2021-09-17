package storeutil

import (
	"todos-vecty/model"

	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

var TodosDispacher *store.Store

var TodosReducer *todosStore

func TodosStoreInit() {
	nextTodoId = 0

	TodosReducer = &todosStore{State: TodosState{State: make([]model.Todo, 0)}}
	TodosDispacher = store.New(TodosReducer)
}

type TodosState struct {
	State []model.Todo
}

type todosStore struct {
	rematch.Reducer
	State TodosState

	Add      *rematch.Action `action:"AddTodo"`
	Complete *rematch.Action `action:"CompleteTodo"`
}

type addTodoAction struct {
	payload model.Todo
}

var nextTodoId int

func NewAddTodoAction(t string) addTodoAction {
	nextTodoId++
	return addTodoAction{payload: model.Todo{Id: nextTodoId, Title: t, Completed: false}}
}

func (t *todosStore) AddTodo(s TodosState, a addTodoAction) TodosState {
	return TodosState{State: append(s.State, a.payload)}
}

type completeTodoAction struct {
	payload int
}

func NewCompleteTodoAction(id int) completeTodoAction {
	return completeTodoAction{payload: id}
}

func (t *todosStore) CompleteTodo(s TodosState, a completeTodoAction) TodosState {
	newState := make([]model.Todo, 0)

	for _, todo := range s.State {
		if todo.Id == a.payload {
			todo.Completed = !todo.Completed
		}
		newState = append(newState, todo)
	}

	return TodosState{State: newState}
}
