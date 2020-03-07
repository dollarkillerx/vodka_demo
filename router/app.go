/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package router

import (
	"vodka/controller"
	"vodka/core/router"
)

func Registry(app *router.Router) {

	app.Run1(controller.Run1)

	app.Run2(controller.Run2)

	app.Run3(controller.Run3)

	app.Run4(controller.Run4)

}
