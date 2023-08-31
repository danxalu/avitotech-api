package cases

import (
	"avitotech-api/inter"
	"github.com/gofiber/fiber/v2"
)

type UC struct {
	rep R
}

func NewUC(r R) *UC {
	return &UC{
		rep: r,
	}
}

func (use *UC) UC_DeleteSegment(ctx *fiber.Ctx, ask *inter.ChangeSegmentAsk, ans *inter.ChangeSegmentAns) (*inter.ChangeSegmentAns, error) {
	ans, err := use.rep.R_DeleteSegment(ctx, ask, ans)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}

func (use *UC) UC_GetSegment(ctx *fiber.Ctx, ask *inter.CheckSegmentAsk, ans *inter.CheckSegmentAns) (*inter.CheckSegmentAns, error) {
	ans, err := use.rep.R_GetSegment(ctx, ask, ans)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}

func (use *UC) UC_CreateSegment(ctx *fiber.Ctx, ask *inter.ChangeSegmentAsk, ans *inter.ChangeSegmentAns) (*inter.ChangeSegmentAns, error) {
	ans, err := use.rep.R_CreateSegment(ctx, ask, ans)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}

func (use *UC) UC_CreateUser(ctx *fiber.Ctx, ask *inter.CreateUserAsk, ans *inter.CreateUserAns) (*inter.CreateUserAns, error) {
	ans, err := use.rep.R_CreateUser(ctx, ask, ans)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}

func (use *UC) UC_SetUserSegment(ctx *fiber.Ctx, ask *inter.ChangeUserSegmentAsk, ans *inter.ChangeUserSegmentAns) (*inter.ChangeUserSegmentAns, error) {
	ans, err := use.rep.R_SetSegment(ctx, ask, ans)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}
