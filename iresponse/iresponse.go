package iresponse

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/i18n"
	"github.com/rhernandez-itemsoft/ihelpers/iresponse/iresponsestt"
)

//Definition esto se inyecta
type Definition struct {
	Ctx iris.Context //el contexto
}

//New Crea una nueva instancia de HTTPResponse
func New(ctx iris.Context) *Definition {
	return &Definition{
		Ctx: ctx,
	}
}

//JSON retorna una respuesta en formato JSON
func (def *Definition) JSON(_response iresponsestt.Response) {
	if def.Ctx == nil {
		//strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT. statusCode: %v, data: %v, iMessages: %v", statusCode, data, iMessages)
		strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT.")
		fmt.Println(strErr)
		return
	}

	def.Ctx.StatusCode(200)
	def.Ctx.JSON(_response)
}

func (def *Definition) JSONResponse(statusCode int, data interface{}, iMessages ...string) {
	var msgs []string

	if def.Ctx == nil {
		//strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT. statusCode: %v, data: %v, iMessages: %v", statusCode, data, iMessages)
		strErr := fmt.Sprintf("iresponse.JSON - NO RECIBIO EL CONTEXT.")
		msgs = append(msgs, strErr)
	} else {
		for _, message := range iMessages {
			msg := i18n.Translate(def.Ctx, message)

			if msg == "" {
				msgs = append(msgs, message)
			} else {
				msgs = append(msgs, msg)
			}
		}
	}

	def.Ctx.StatusCode(statusCode)
	def.Ctx.JSON(map[string]interface{}{
		"Messages": msgs,
		"Data":     data,
	})
}
