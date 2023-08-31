package cases

import (
	"avitotech-api/inter"
	"github.com/gofiber/fiber/v2"
)

type (
	R interface {
		R_SetSegment(ctx *fiber.Ctx, request *inter.ChangeUserSegmentAsk, ans *inter.ChangeUserSegmentAns) (*inter.ChangeUserSegmentAns, error)
		R_GetSegment(ctx *fiber.Ctx, request *inter.CheckSegmentAsk, ans *inter.CheckSegmentAns) (*inter.CheckSegmentAns, error)

		R_CreateSegment(ctx *fiber.Ctx, ask *inter.ChangeSegmentAsk, ans *inter.ChangeSegmentAns) (*inter.ChangeSegmentAns, error)
		R_DeleteSegment(ctx *fiber.Ctx, request *inter.ChangeSegmentAsk, ans *inter.ChangeSegmentAns) (*inter.ChangeSegmentAns, error)

		R_CreateUser(ctx *fiber.Ctx, request *inter.CreateUserAsk, ans *inter.CreateUserAns) (*inter.CreateUserAns, error)
		}
)
