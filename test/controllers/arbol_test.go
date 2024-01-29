package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestConsultarArbol(t *testing.T) {
	if response, err := http.Get("http://localhost:9010/v1/arbol/611db9fcd40348bf0438b647"); err == nil {
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

func TestActivarPlan(t *testing.T) {
	body := []byte(`{}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9010/v1/arbol/activar_plan/611e8f44d40348127238b7fa", bytes.NewBuffer(body)); err == nil {
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
	body := []byte(`{}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:9010/v1/arbol/activar_nodo/611db9fcd40348bf0438b647", bytes.NewBuffer(body)); err == nil {
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

func TestDesactivarPlan(t *testing.T) {
	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9010/v1/arbol/desactivar_plan/611e8f44d40348127238b7fa", nil); err == nil {
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
	if request, err := http.NewRequest(http.MethodDelete, "http://localhost:9010/v1/arbol/desactivar_nodo/611db9fcd40348bf0438b647", nil); err == nil {
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
