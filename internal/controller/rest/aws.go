package rest

import (
	"aws-sns-local-go/usecase/aws"

	"aws-sns-local-go/internal/middleware"

	"github.com/labstack/echo/v4"
)

type AwsService interface {
	CreateTopic(in aws.CreateTopicInput) (aws.CreateTopicOutput, error)
	ListTopics() (aws.ListTopicOutput, error)
	Publish(in aws.PublishInput) (aws.PublishOutput, error)
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
	case "Publish":
		in := aws.PublishInput{
			TopicArn:               c.FormValue("TopicArn"),
			TargetArn:              c.FormValue("TargetArn"),
			PhoneNumber:            c.FormValue("PhoneNumber"),
			Message:                c.FormValue("Message"),
			Subject:                c.FormValue("Subject"),
			MessageStructure:       c.FormValue("MessageStructure"),
			MessageAttributes:      c.FormValue("MessageAttributes"),
			MessageDeduplicationId: c.FormValue("MessageDeduplicationId"),
			MessageGroupId:         c.FormValue("MessageGroupId"),
		}

		out, err := h.AwsService.Publish(in)
		if err != nil {
			return middleware.HandleError(c, err)
		}
		return c.XML(200, out.PublishResponse)
	default:
		return c.JSON(400, "Invalid action")
	}
}
