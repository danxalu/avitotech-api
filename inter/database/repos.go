package database

import (
	"avitotech-api/inter"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Rep struct {
	db *sql.DB
}

func Rep_New(db *sql.DB) *Rep {
	return &Rep{db}
}


func (r *Rep) R_GetSegment(ctx *fiber.Ctx, ask *inter.CheckSegmentAsk, ans *inter.CheckSegmentAns) (*inter.CheckSegmentAns, error) {
	context := ctx.Context()

	rows, err := r.db.QueryContext(context, string(ask.UserID))
	if err != nil {
		return ans, nil
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var segment string
		if err = rows.Scan(&segment); err != nil {
			return ans, nil
		}
		ans.Segment = append(ans.Segment, segment)
	}

	if err = rows.Err(); err != nil {
		return ans, nil
	}

	return ans, nil
}

func (r *Rep) R_SetSegment(ctx *fiber.Ctx, ask *inter.ChangeUserSegmentAsk, ans *inter.ChangeUserSegmentAns) (*inter.ChangeUserSegmentAns, error) {
	context := ctx.Context()

	if ask.AddSlug != nil {
		for i := 0; i < len(ask.AddSlug); i++ {
			var segment string
			err := r.db.QueryRowContext(context, string(ans.UserID), ask.AddSlug[i]).Scan(&ask.UserID, &segment)
			if err != nil {
				fmt.Println(err)
				return ans, nil
			}

			ans.Segment = append(ans.Segment, segment)
		}
	}

	if ask.DeleteSlug != nil {
		for i := 0; i < len(ask.DeleteSlug); i++ {
			var segment string
			err := r.db.QueryRowContext(context, string(ans.UserID), ask.DeleteSlug[i]).Scan(&ask.UserID, &segment)
			if err != nil {
				fmt.Println(err)
			}
			ans.Segment = append(ans.Segment, segment)
		}
	}

	return ans, nil
}

func (r *Rep) R_CreateSegment(ctx *fiber.Ctx, ask *inter.ChangeSegmentAsk, ans *inter.ChangeSegmentAns) (*inter.ChangeSegmentAns, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, ans.Slug).Scan(&ans.SegmentId, &ask.Slug)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}

func (r *Rep) R_DeleteSegment(ctx *fiber.Ctx, ask *inter.ChangeSegmentAsk, ans *inter.ChangeSegmentAns) (*inter.ChangeSegmentAns, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, ans.Slug).Scan(&ans.SegmentId, &ask.Slug)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}

func (r *Rep) R_CreateUser(ctx *fiber.Ctx, ask *inter.CreateUserAsk, ans *inter.CreateUserAns) (*inter.CreateUserAns, error) {
	context := ctx.Context()

	err := r.db.QueryRowContext(context, ask.UserName).Scan(&ans.UserID)
	if err != nil {
		return ans, nil
	}

	return ans, nil
}