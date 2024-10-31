package rest

import (
	"aws-sns-local-go/usecase/dto"
	"aws-sns-local-go/usecase/query"

	"github.com/labstack/echo/v4"
)

type TopicHandler struct {
	TopicQueryService query.TopicQueryService
}

type TopicResponse struct {
	Items []dto.TopicDto `json:"items"`
}

func NewTopicHandler(e *echo.Echo, topicQueryService query.TopicQueryService) {
	handler := &TopicHandler{
		TopicQueryService: topicQueryService,
	}

	e.GET("/topics", handler.GetTopic)
}

func (h *TopicHandler) GetTopic(c echo.Context) error {
	topicArn := c.QueryParam("topicArn")

	topic, err := h.TopicQueryService.FindByTopicArn(topicArn)
	if err != nil {
		return c.JSON(500, err)
	}

	resp := TopicResponse{
		Items: topic,
	}

	return c.JSON(200, resp)
}
