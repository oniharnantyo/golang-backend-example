package middleware

import (
	"fmt"
	"net/http"

	"github.com/oniharnantyo/golang-backend-example/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jwtHeader struct {
	Authorization string `header:"Authorization"`
}

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var h jwtHeader

		err := ctx.ShouldBindHeader(&h)
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				ctx.JSON(http.StatusForbidden, util.Response{
					Data: nil,
					Errors: []string{
						"Something went wrong",
					},
				})

				ctx.Abort()
				return
			}

			var errsMessage []string
			for _, err := range errs {
				errsMessage = append(errsMessage, err.Error())
			}

			ctx.JSON(http.StatusForbidden, util.Response{
				Data:   nil,
				Errors: errsMessage,
			})

			ctx.Abort()
			return
		}

		fmt.Println(h)

		ctx.Next()
	}
}
