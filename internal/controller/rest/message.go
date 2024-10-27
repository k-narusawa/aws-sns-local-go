package rest

import (
	"aws-sns-local-go/usecase/query"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	MessageQuerySvc query.IMessageQuery
}

func NewMessageHandler(e *echo.Echo, messageQuerySvc query.IMessageQuery) {
	handler := &MessageHandler{
		MessageQuerySvc: messageQuerySvc,
	}

	e.GET("/messages", handler.FindAll)
}

func (h *MessageHandler) FindAll(c echo.Context) error {
	qPage := c.QueryParam("page")
	qLimit := c.QueryParam("limit")

	if qPage == "" {
		qPage = "1"
	}

	if qLimit == "" {
		qLimit = "10"
	}

	orgSize, _ := strconv.Atoi(qPage)
	isize := orgSize - 1
	limit, _ := strconv.Atoi(qLimit)
	offset := isize * limit

	messages, err := h.MessageQuerySvc.FindAll(limit, offset)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, messages)
}