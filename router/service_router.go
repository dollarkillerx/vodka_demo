/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package router

import (
	"awesome/controller"
	"awesome/core/router"
)

func Registry(app *router.Router) {
	app.Run1(controller.Run1, controller.Run2)
}
