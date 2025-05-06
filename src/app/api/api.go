package api

import (
	"currency/internal/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New(validator.WithRequiredStructEnabled())

func ResponseError(message string) *Response {
	return &Response{
		Status:  "error",
		Message: message,
	}
}

func ResponseSuccess(data any) *Response {
	return &Response{
		Status:  "success",
		Message: "ok",
		Data:    data,
	}
}

func ValidateRequest(c *fiber.Ctx, req interface{}) error {
	body := string(c.BodyRaw())
	err := utils.ValidateRequest(body, req)
	if err != nil {
		return err
	}

	return nil
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type DataEmpty struct {
}

type Identity struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Role     string `json:"role"`
}

type LangQueryParam struct {
	Lang string `json:"lang" query:"lang" default:"ru"`
}

func (l LangQueryParam) Value() string {
	if strings.ToLower(l.Lang) != "en" {
		return "ru"
	}
	return strings.ToLower(l.Lang)
}
