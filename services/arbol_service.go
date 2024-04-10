package services

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_arbol_mid/helpers"
	"github.com/udistrital/planeacion_arbol_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func ConsultarArbol(id string) ([]map[string]interface{}, error) {
	var hijos []models.Nodo
	var hijosID []map[string]interface{}
	var res map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err != nil {
		return nil, errors.New(err.Error())
	}

	request.LimpiezaRespuestaRefactor(res, &hijos)
	request.LimpiezaRespuestaRefactor(res, &hijosID)

	resultado, err := ConstruirArbol(hijos, hijosID)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if resultado == nil {
		return nil, errors.New("error del servicio ConsultarArbol: El arbol no contiene datos")
	}

	return resultado, nil
}

func ConstruirArbol(hijos []models.Nodo, hijosID []map[string]interface{}) ([]map[string]interface{}, error) {
	var tree []map[string]interface{}

	for index, hijo := range hijos {
		forkData := make(map[string]interface{})
		forkData["id"] = hijosID[index]["_id"]
		forkData["nombre"] = hijo.Nombre
		forkData["descripcion"] = hijo.Descripcion
		if hijo.Activo {
			forkData["activo"] = "activo"
		} else {
			forkData["activo"] = "inactivo"
		}

		if len(hijo.Hijos) > 0 {
			resultado, err := ConsultarHijos(hijo.Hijos)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			forkData["children"] = resultado
		}

		tree = append(tree, forkData)
	}

	return tree, nil
}

func ConsultarHijos(children []string) ([]map[string]interface{}, error) {
	var res map[string]interface{}
	var nodo models.Nodo
	var nodoId map[string]interface{}
	var childrenTree []map[string]interface{}

	for _, child := range children {
		forkData := make(map[string]interface{})

		err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+child, &res)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		request.LimpiezaRespuestaRefactor(res, &nodo)
		request.LimpiezaRespuestaRefactor(res, &nodoId)
		forkData["id"] = nodoId["_id"]
		forkData["nombre"] = nodo.Nombre
		forkData["descripcion"] = nodo.Descripcion

		if nodo.Activo {
			forkData["activo"] = "activo"
		} else {
			forkData["activo"] = "inactivo"
		}

		if len(nodo.Hijos) > 0 {
			resultado, err := ConsultarHijos(nodo.Hijos)
			if err != nil {
				return nil, errors.New(err.Error())
			}
			forkData["children"] = resultado
		}
		childrenTree = append(childrenTree, forkData)
	}
	return childrenTree, nil
}

func ActivarPlan(id string) (map[string]interface{}, error) {
	var plan map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res)
	if err != nil {
		return nil, errors.New("error del servicio ActivarPlan: no se pudo obtener el plan")
	}

	request.LimpiezaRespuestaRefactor(res, &plan)
	if helpers.EsMapVacio(plan) {
		return nil, errors.New("error del servicio ActivarPlan: La respuesta no contiene datos")
	}

	plan["activo"] = true

	err = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+plan["_id"].(string), "PUT", &resPut, plan)
	if err != nil {
		return nil, errors.New("error del servicio ActivarPlan: no se pudo activar el plan")
	}

	err = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+plan["_id"].(string), &resHijos)
	if err != nil {
		return nil, errors.New("error del servicio ActivarPlan: no se pudieron obtener los hijos del plan")
	}

	request.LimpiezaRespuestaRefactor(resHijos, &hijos)
	_, err = ActivarHijos(hijos)
	if err != nil {
		return nil, errors.New("error del servicio ActivarPlan: no se pudieron activar los hijos del plan")
	}

	return plan, nil
}

func DesactivarPlan(id string) (map[string]interface{}, error) {
	var plan map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &plan)
		if helpers.EsMapVacio(plan) {
			return nil, errors.New("error del servicio DesactivarPlan: La respuesta no contiene datos")
		}
		plan["activo"] = false
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+plan["_id"].(string), "PUT", &resPut, plan); err != nil {
			return nil, errors.New(err.Error())
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+plan["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
		}
		DesactivarHijos(hijos)
		return plan, nil
	} else {
		return nil, errors.New(err.Error())
	}
}

func ActivarHijos(children []map[string]interface{}) ([]map[string]interface{}, error) {
	var res map[string]interface{}
	var resHijos map[string]interface{}

	for _, child := range children {
		child["activo"] = true
		if err := request.SendJson("http://" + beego.AppConfig.String("PlanesService") + "/subgrupo/" + child["_id"].(string), "PUT", &res, child); err != nil {
			return nil, errors.New("error del servicio ActivarHijos: Error al actualizar el campo 'activo'")
		}

		if child["hijos"] != nil {
			hijos, ok := child["hijos"].([]interface{})
			if !ok {
				return nil, errors.New("error del servicio ActivarHijos: Error al obtener los hijos del map")
			}

			if len(hijos) > 0 {
				if err := request.GetJson("http://" + beego.AppConfig.String("PlanesService") + "/subgrupo?query=padre:" + child["_id"].(string), &resHijos); err != nil {
					return nil, errors.New(err.Error())
				}
				var cleanedHijos []map[string]interface{}
				request.LimpiezaRespuestaRefactor(resHijos, &cleanedHijos)
				if _, err := ActivarHijos(cleanedHijos); err != nil {
					return nil, errors.New(err.Error())
				}
			}
		}
	}
	return children, nil
}

func DesactivarHijos(children []map[string]interface{}) ([]map[string]interface{}, error) {
	var res map[string]interface{}
	var resHijos map[string]interface{}

	for _, child := range children {
		child["activo"] = false

		if err := request.SendJson("http://" + beego.AppConfig.String("PlanesService") + "/subgrupo/" + child["_id"].(string), "PUT", &res, child); err != nil {
			return nil, errors.New("error del servicio DesactivarHijos: Error al actualizar el campo 'activo'")
		}

		if child["hijos"] != nil {
			hijos, ok := child["hijos"].([]interface{})
			if !ok {
				return nil, errors.New("error del servicio DesactivarHijos: Error al obtener los hijos del map")
			}

			if len(hijos) > 0 {
				if err := request.GetJson("http://" + beego.AppConfig.String("PlanesService") + "/subgrupo?query=padre:" + child["_id"].(string), &resHijos); err == nil {
					var cleanedHijos []map[string]interface{}
					request.LimpiezaRespuestaRefactor(resHijos, &cleanedHijos)
					if _, err := DesactivarHijos(cleanedHijos); err != nil {
						return nil, errors.New(err.Error())
					}
				} else {
					return nil, errors.New(err.Error())
				}
			}
		}
	}
	return children, nil
}

func DesactivarNodo(id string) (map[string]interface{}, error) {
	var subgrupo map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		if helpers.EsMapVacio(subgrupo) {
			return nil, errors.New("error del servicio DesactivarNodo: La respuesta no contiene datos")
		}
		subgrupo["activo"] = false
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["_id"].(string), "PUT", &resPut, subgrupo); err != nil {
			return nil, err
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+subgrupo["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
			if _, err := DesactivarHijos(hijos); err != nil {
				return nil, errors.New(err.Error())
			}
		}
		return subgrupo, nil
	} else {
		return nil, errors.New(err.Error())
	}
}

func ActivarNodo(id string) (map[string]interface{}, error) {
	var subgrupo map[string]interface{}
	var res map[string]interface{}
	var resPut map[string]interface{}
	var resHijos map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		if helpers.EsMapVacio(subgrupo) {
			return nil, errors.New("error del servicio ActivarNodo: La respuesta no contiene datos")
		}
		subgrupo["activo"] = true
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["_id"].(string), "PUT", &resPut, subgrupo); err != nil {
			return nil, errors.New(err.Error())
		}
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+subgrupo["_id"].(string), &resHijos); err == nil {
			request.LimpiezaRespuestaRefactor(resHijos, &hijos)
			ActivarHijos(hijos)
		}
		return subgrupo, nil
	} else {
		return nil, errors.New(err.Error())
	}
}
