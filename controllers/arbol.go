package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_arbol_mid/helpers"
	"github.com/udistrital/planeacion_arbol_mid/models"
	"github.com/udistrital/utils_oas/request"
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
// @Description Consulta el arbol por id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Arbol
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArbolController) ConsultarArbol() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ArbolController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	var res map[string]interface{}
	var hijos []models.Nodo
	var hijosID []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &hijos)
		request.LimpiezaRespuestaRefactor(res, &hijosID)
		tree := helpers.ConstruirArbol(hijos, hijosID)
		if len(tree) != 0 {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": tree}
		} else {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": nil}
		}
	} else {
		panic(err)
	}

	c.ServeJSON()
}

// DesactivarPlan ...
// @Title DesactivarPlan
// @Description desactiva el plan arbol
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /desactivar_plan/:id [delete]
func (c *ArbolController) DesactivarPlan() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ArbolController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	var plan map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &plan)
		plan["activo"] = false
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+plan["_id"].(string), "PUT", &resPut, plan); err != nil {
			panic(map[string]interface{}{"funcion": "DesactivarPlan", "err": "Error actualizacion inactivo \"id\"", "status": "400", "log": err})
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+plan["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": plan}
		helpers.DesactivarHijos(hijos)
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// DesactivarNodo ...
// @Title DesactivarNodo
// @Description desactiva el nodo arbol
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /desactivar_nodo/:id [delete]
func (c *ArbolController) DesactivarNodo() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ArbolController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	var subgrupo map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		subgrupo["activo"] = false

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["_id"].(string), "PUT", &resPut, subgrupo); err != nil {
			panic(map[string]interface{}{"funcion": "DesactivarNodo", "err": "Error actualizacion inactivo \"id\"", "status": "400", "log": err})
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+subgrupo["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
			helpers.DesactivarHijos(hijos)
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": subgrupo}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// ActivarPlan ...
// @Title ActivarPlan
// @Description activar el plan arbol
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /activar_plan/:id [put]
func (c *ArbolController) ActivarPlan() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ArbolController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	var plan map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &plan)
		plan["activo"] = true
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+plan["_id"].(string), "PUT", &resPut, plan); err != nil {
			panic(map[string]interface{}{"funcion": "ActivarPlan", "err": "Error actualizacion activo \"id\"", "status": "400", "log": err})
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+plan["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
			helpers.ActivarHijos(hijos)
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": plan}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// ActivarNodo ...
// @Title ActivarNodo
// @Description activa el nodo arbol
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /activar_nodo/:id [put]
func (c *ArbolController) ActivarNodo() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ArbolController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	var subgrupo map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		subgrupo["activo"] = true
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["_id"].(string), "PUT", &resPut, subgrupo); err != nil {
			panic(map[string]interface{}{"funcion": "ActivarNodo", "err": "Error actualizacion activo \"id\"", "status": "400", "log": err})
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+subgrupo["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
			helpers.ActivarHijos(hijos)
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": subgrupo}
	} else {
		panic(err)
	}
	c.ServeJSON()
}
