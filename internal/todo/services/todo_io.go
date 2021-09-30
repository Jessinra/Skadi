package services

type CreateNewTodoInput struct {
	Text        string
	Description *string
	UserID      uint64
}
