package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_arbol_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

// ArbolController operations for Arbol
type ArbolController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArbolController) URLMapping() {
	c.Mapping("ConsultarArbol", c.ConsultarArbol)
	c.Mapping("DesactivarPlan", c.DesactivarPlan)
	c.Mapping("DesactivarNodo", c.DesactivarNodo)
	c.Mapping("ActivarPlan", c.ActivarPlan)
	c.Mapping("ActivarNodo", c.ActivarNodo)
}

// ConsultarArbol ...
// @Title ConsultarArbol
// @Description Consulta el árbol por id
// @Param	id		path 	string	true		"Id del árbol que quiere consultar"
// @Success 200 {object} models.Arbol
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArbolController) ConsultarArbol() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.ConsultarArbol(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// DesactivarPlan ...
// @Title DesactivarPlan
// @Description Desactiva el plan árbol
// @Param	id		path 	string	true		"Id del plan que quiere desactivar"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /plan/:id/desactivar [delete]
func (c *ArbolController) DesactivarPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.DesactivarPlan(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 400, nil, err.Error())
	}
	
	c.ServeJSON()
}

// DesactivarNodo ...
// @Title DesactivarNodo
// @Description Desactiva el nodo árbol
// @Param	id		path 	string	true		"Id del nodo que quiere desactivar"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /nodo/:id/desactivar [delete]
func (c *ArbolController) DesactivarNodo() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.DesactivarNodo(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 400, nil, err.Error())
	}
	c.ServeJSON()
}

// ActivarPlan ...
// @Title ActivarPlan
// @Description Activa el plan árbol
// @Param	id		path 	string	true		"Id del plan que quiere activar"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /plan/:id/activar [put]
func (c *ArbolController) ActivarPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	
	if resultado, err := services.ActivarPlan(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 400, nil, err.Error())
	}
	c.ServeJSON()
}

// ActivarNodo ...
// @Title ActivarNodo
// @Description Activa el nodo arbol
// @Param	id		path 	string	true		"Id del nodo que quiere activar"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /nodo/:id/activar [put]
func (c *ArbolController) ActivarNodo() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	
	if resultado, err := services.ActivarNodo(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 400, nil, err.Error())
	}

	c.ServeJSON()
}
