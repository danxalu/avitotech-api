package handlers

import (
	"avitotech-api/inter"
	"github.com/gofiber/fiber/v2"
)

func (reg *Reg) createSegment(ctx *fiber.Ctx) error {
	ans := new(inter.ChangeSegmentAns)

	ask := new(inter.ChangeSegmentAsk)

	err := ctx.BodyParser(ask)
	if err != nil {
		return err
	}

	ans, err = reg.Ucase.UC_CreateSegment(ctx, ask, ans)
	if err != nil {
		return err
	}

	return ctx.JSON(ans)
}

func (reg *Reg) deleteSegment(ctx *fiber.Ctx) error {
	ask := new(inter.ChangeSegmentAsk)
	ans := new(inter.ChangeSegmentAns)

	err := ctx.BodyParser(ask)
	if err != nil {
		return err
	}

	ans, err = reg.Ucase.UC_DeleteSegment(ctx, ask, ans)
	if err != nil {
		return err
	}

	return ctx.JSON(ans)
}
