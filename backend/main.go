package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type saludo struct {
	Saludo string
}

type resultado struct {
	Resultado string
}

type formato struct {
	Simbolo string
	Numero1 string
	Numero2 string
}

type registros struct {
	Numero1   string
	Numero2   string
	Operacion string
	Resultado string
	Fecha     string
}

var operaciones []registros

func getVerificacion(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	enviar := saludo{Saludo: "Bienvenido el servidor esta ejecutandose correctamente!!"}
	json.NewEncoder(res).Encode(enviar)
}

func endPointRegistros(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(operaciones)
}

func endPointSuma(res http.ResponseWriter, req *http.Request) {
	entrada := &formato{}
	_ = json.NewDecoder(req.Body).Decode(entrada)
	valor1, error1 := strconv.ParseFloat(entrada.Numero1, 64) //convertir de string a float64
	valor2, error2 := strconv.ParseFloat(entrada.Numero2, 64) //convertir de string a float64
	var respuesta resultado
	if error1 != nil || error2 != nil {
		respuesta := resultado{Resultado: "Error"}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(respuesta)
		return
	}

	suma := strconv.FormatFloat(valor1+valor2, 'f', -1, 64) //convertir de float a string
	//fmt.Println(suma)

	respuesta = resultado{Resultado: suma}

	actual := time.Now()
	actualstring := actual.Format("02/01/2006 15:04:05")
	registro := registros{Numero1: entrada.Numero1, Numero2: entrada.Numero2, Operacion: entrada.Simbolo, Resultado: suma, Fecha: actualstring}
	operaciones = append(operaciones, registro)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respuesta)

}

func endPointResta(res http.ResponseWriter, req *http.Request) {
	entrada := &formato{}
	_ = json.NewDecoder(req.Body).Decode(entrada)
	valor1, error1 := strconv.ParseFloat(entrada.Numero1, 64) //convertir de string a float64
	valor2, error2 := strconv.ParseFloat(entrada.Numero2, 64) //convertir de string a float64
	var respuesta resultado
	if error1 != nil || error2 != nil {
		respuesta := resultado{Resultado: "Error"}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(respuesta)
		return
	}

	resta := strconv.FormatFloat(valor1-valor2, 'f', -1, 64) //convertir de float a string
	//fmt.Println(resta)

	respuesta = resultado{Resultado: resta}
	actual := time.Now()
	actualstring := actual.Format("02/01/2006 15:04:05")
	registro := registros{Numero1: entrada.Numero1, Numero2: entrada.Numero2, Operacion: entrada.Simbolo, Resultado: resta, Fecha: actualstring}
	operaciones = append(operaciones, registro)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respuesta)

}

func endPointMultiplicacion(res http.ResponseWriter, req *http.Request) {
	entrada := &formato{}
	_ = json.NewDecoder(req.Body).Decode(entrada)
	valor1, error1 := strconv.ParseFloat(entrada.Numero1, 64) //convertir de string a float64
	valor2, error2 := strconv.ParseFloat(entrada.Numero2, 64) //convertir de string a float64
	var respuesta resultado
	if error1 != nil || error2 != nil {
		respuesta := resultado{Resultado: "Error"}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(respuesta)
		return
	}

	multiplicacion := strconv.FormatFloat(valor1*valor2, 'f', -1, 64) //convertir de float a string
	//fmt.Println(multiplicacion)

	respuesta = resultado{Resultado: multiplicacion}
	actual := time.Now()
	actualstring := actual.Format("02/01/2006 15:04:05")
	registro := registros{Numero1: entrada.Numero1, Numero2: entrada.Numero2, Operacion: entrada.Simbolo, Resultado: multiplicacion, Fecha: actualstring}
	operaciones = append(operaciones, registro)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respuesta)

}

func endPointDivision(res http.ResponseWriter, req *http.Request) {
	entrada := &formato{}
	_ = json.NewDecoder(req.Body).Decode(entrada)
	valor1, error1 := strconv.ParseFloat(entrada.Numero1, 64) //convertir de string a float64
	valor2, error2 := strconv.ParseFloat(entrada.Numero2, 64) //convertir de string a float64
	var respuesta resultado
	if error1 != nil || error2 != nil || valor2 == 0 {
		respuesta := resultado{Resultado: "Error"}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(respuesta)
		return
	}

	division := strconv.FormatFloat(valor1/valor2, 'f', -1, 64) //convertir de float a string
	//fmt.Println(division)

	respuesta = resultado{Resultado: division}
	actual := time.Now()
	actualstring := actual.Format("02/01/2006 15:04:05")
	registro := registros{Numero1: entrada.Numero1, Numero2: entrada.Numero2, Operacion: entrada.Simbolo, Resultado: division, Fecha: actualstring}
	operaciones = append(operaciones, registro)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respuesta)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", getVerificacion).Methods("GET")
	router.HandleFunc("/registros", endPointRegistros).Methods("GET")
	router.HandleFunc("/suma", endPointSuma).Methods("POST")
	router.HandleFunc("/resta", endPointResta).Methods("POST")
	router.HandleFunc("/multiplicacion", endPointMultiplicacion).Methods("POST")
	router.HandleFunc("/division", endPointDivision).Methods("POST")

	parametro1 := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}) // aceptar todos los heders del fetch
	parametro2 := handlers.AllowedMethods([]string{"GET", "POST", "HEAD", "OPTIONS"})                                                   // indico que metodos aceptare
	parametro3 := handlers.AllowedOrigins([]string{"*"})                                                                                //que acepte todos los orignes
	http.ListenAndServe(":5000", handlers.CORS(parametro1, parametro2, parametro3)(router))

}
