package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/skplayer/internal/constant"
	"github.com/run-bigpig/skplayer/internal/types"
	"net/http"
)

func Success(ctx *fiber.Ctx, data interface{}) error {
	res := types.Response{
		Code: 0,
		Data: data,
		Msg:  constant.Success,
	}
	err := ctx.JSON(res)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusOK)
}

func Fail(ctx *fiber.Ctx, code int, error error) error {
	res := types.Response{
		Code: 1,
		Data: nil,
		Msg:  error.Error(),
	}
	err := ctx.JSON(res)
	if err != nil {
		return err
	}
	return ctx.SendStatus(code)
}
