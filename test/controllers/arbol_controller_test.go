package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestConsultarArbol(t *testing.T) {
	endpoint := "http://localhost:8083/v1/arbol/"
	idArbol := "611db9fcd40348bf0438b647"
	if response, err := http.Get(endpoint + idArbol); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error ConsultarArbol Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("ConsultarArbol Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error ConsultarArbol:", err.Error())
		t.Fail()
	}
}

func TestDesactivarPlan(t *testing.T) {
	endpoint := "http://localhost:8083/v1/arbol/plan/"
	idPlan := "611e4a2dd403481fb638b6e9"
	if request, err := http.NewRequest(http.MethodDelete, endpoint+idPlan+"/desactivar", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestDesactivarPlan Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestDesactivarPlan Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}

func TestDesactivarNodo(t *testing.T) {
	endpoint := "http://localhost:8083/v1/arbol/nodo/"
	idNodo := "611db9b4d403482fec38b637"
	if request, err := http.NewRequest(http.MethodDelete, endpoint+idNodo+"/desactivar", nil); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestDesactivarNodo Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestDesactivarNodo Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud DELETE: ", err.Error())
		t.Fail()
	}
}

func TestActivarPlan(t *testing.T) {
	endpoint := "http://localhost:8083/v1/arbol/plan/"
	idPlan := "611e4a2dd403481fb638b6e9"
	body := []byte(`{}`)

	if request, err := http.NewRequest(http.MethodPut, endpoint+idPlan+"/activar", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestActivarPlan Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestActivarPlan Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestActivarNodo(t *testing.T) {
	endpoint := "http://localhost:8083/v1/arbol/nodo/"
	idNodo := "611db9b4d403482fec38b637"
	body := []byte(`{}`)

	if request, err := http.NewRequest(http.MethodPut, endpoint+idNodo+"/activar", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestActivarNodo Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestActivarNodo Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}
