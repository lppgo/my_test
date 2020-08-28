package ctxlog

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ctx参数转换
func HandleCtxParams() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			reqID := ""

			xreq := c.Get(echo.HeaderXRequestID)
			if xreq == nil {
				reqID = generateUuid()
			} else if xreq == "" {
				reqID = generateUuid()
			}
			ctx = context.WithValue(ctx, "ClientIP", c.RealIP())
			ctx = context.WithValue(ctx, "RequestID", reqID)
			ctx = context.WithValue(ctx, "DentifyID", c.Get("userId"))
			ctx = context.WithValue(ctx, "ServersID", "mountain_go")
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)
			err := next(c)
			c.Response().Header().Set(echo.HeaderXRequestID, reqID)
			return err
		}
	}
}

func generateUuid() (id string) {
	rawId := uuid.New().String()
	Ids := strings.Split(rawId, "-")
	id = Ids[0] + Ids[4]
	return id
}
