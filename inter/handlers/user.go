package handlers

import (
	"avitotech-api/inter"
	"github.com/gofiber/fiber/v2"
)

func (reg *Reg) changeUserSegment(ctx *fiber.Ctx) error {
	ask := new(inter.ChangeUserSegmentAsk)
	ans := new(inter.ChangeUserSegmentAns)

	err := ctx.BodyParser(ask)
	if err != nil {
		return err
	}

	ans, err = reg.Ucase.UC_SetUserSegment(ctx, ask, ans)
	if err != nil {
		return err
	}

	return ctx.JSON(ans)
}

func (reg *Reg) createUser(ctx *fiber.Ctx) error {
	ans := new(inter.CreateUserAns)

	ask := new(inter.CreateUserAsk)

	err := ctx.BodyParser(ask)
	if err != nil {
		return err
	}

	ans, err = reg.Ucase.UC_CreateUser(ctx, ask, ans)
	if err != nil {
		return err
	}

	return ctx.JSON(ans)
}

func (reg *Reg) checkUserSegment(ctx *fiber.Ctx) error {
	ans := new(inter.CheckSegmentAns)

	ask := new(inter.CheckSegmentAsk)

	err := ctx.BodyParser(ask)
	if err != nil {
		return err
	}

	ans, err = reg.Ucase.UC_GetSegment(ctx, ask, ans)
	if err != nil {
		return err
	}

	return ctx.JSON(ans)
}
