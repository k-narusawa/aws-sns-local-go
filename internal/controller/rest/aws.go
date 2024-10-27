package rest

import (
	"aws-sns-local-go/usecase/aws"

	"aws-sns-local-go/internal/middleware"

	"github.com/labstack/echo/v4"
)

type AwsService interface {
	CreateTopic(in aws.CreateTopicInput) (aws.CreateTopicOutput, error)
	ListTopics() (aws.ListTopicOutput, error)
}

type AwsHandler struct {
	AwsService AwsService
}

func NewAwsHandler(e *echo.Echo, awsService AwsService) {
	handler := &AwsHandler{
		AwsService: awsService,
	}

	e.POST("/", handler.Sns)
}

func (h *AwsHandler) Sns(c echo.Context) error {
	action := c.FormValue("Action")

	switch action {
	case "CreateTopic":
		in := aws.CreateTopicInput{
			Name:       c.FormValue("Name"),
			Tags:       c.FormValue("Tags"),
			Attributes: c.FormValue("Attributes"),
		}
		out, err := h.AwsService.CreateTopic(in)
		if err != nil {
			return middleware.HandleError(c, err)

		}
		return c.XML(200, out.CreateTopicResponse)
	case "ListTopics":
		out, err := h.AwsService.ListTopics()
		if err != nil {
			return middleware.HandleError(c, err)
		}
		return c.XML(200, out.ListTopicResponse)
	default:
		return c.JSON(400, "Invalid action")
	}
}
