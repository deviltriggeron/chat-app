package handler

import "chat-app/internal/usecase"

type Handler struct {
	us usecase.Usecase
}

func NewHandler(us usecase.Usecase) Handler {
	return Handler{
		us: us,
	}
}
