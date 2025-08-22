package rest

import (
	"aws-sns-local-go/usecase/dto"
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

type MessagesResponse struct {
	Page      int              `json:"page"`
	Limit     int              `json:"limit"`
	Size      int              `json:"size"`
	TotalPage int              `json:"totalPage"`
	TotalSize int              `json:"totalSize"`
	Items     []dto.MessageDto `json:"items"`
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
	phoneNumber := c.QueryParam("phoneNumber")

	messages, err := h.MessageQuerySvc.FindAll(phoneNumber, limit, offset)
	if err != nil {
		return c.JSON(500, err)
	}

	resp := MessagesResponse{
		Page:  orgSize,
		Limit: limit,
		Size:  len(messages),
		Items: messages,
	}

	return c.JSON(200, resp)
}
