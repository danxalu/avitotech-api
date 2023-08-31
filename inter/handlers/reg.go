package handlers

import (
	"avitotech-api/inter/cases"
	"github.com/gofiber/fiber/v2"
)

type Reg struct {
	Ucase *cases.UC
}

func (reg *Reg) Registeration(app fiber.Router) {
	app.Post("/user/check", reg.checkUserSegment)
	app.Post("/user/change", reg.changeUserSegment)
	app.Post("/user", reg.createUser)

	app.Post("/segment", reg.createSegment)
	app.Delete("/segment", reg.deleteSegment)
}
