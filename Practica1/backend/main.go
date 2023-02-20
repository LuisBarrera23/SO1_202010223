package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

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

func conectarBaseDeDatos() (conexion *sql.DB) {
	nombreContenedor := "mysql-db"
	nombreDB := "logs"
	driver := "mysql"
	usuario := "root"
	contrasena := "1234"
	puerto := "3306"

	conexion, err := sql.Open(driver, usuario+":"+contrasena+"@tcp("+nombreContenedor+":"+puerto+")/"+nombreDB)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func guardarOperacion(datos registros) {
	DBconexion := conectarBaseDeDatos()
	datalogs, err := DBconexion.Prepare("INSERT INTO registros(numero1,numero2,operacion,resultado,fecha) VALUES('" + datos.Numero1 + "','" + datos.Numero2 + "','" + datos.Operacion + "','" + datos.Resultado + "','" + datos.Fecha + "');")
	if err != nil {
		panic(err.Error())
	}
	datalogs.Exec()
	DBconexion.Close()
	actualizarArchivo(datos)

}

func actualizarArchivo(datos registros) {
	archivo, err := os.OpenFile("/registros/archivo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //abrimos el archivo en modo append
	if err != nil {
		log.Fatal(err) //en caso de error
	}
	defer archivo.Close() // comando para que al finalizar el enviroment se cierre el archivo
	actual := time.Now()
	actualstring := actual.Format("02/01/2006")

	var operacion string

	if datos.Operacion == "+" {
		operacion = "suma"
	} else if datos.Operacion == "-" {
		operacion = "resta"
	} else if datos.Operacion == "*" {
		operacion = "multiplicacion"
	} else if datos.Operacion == "/" {
		operacion = "division"
	} else {
		operacion = ""
	}
	linea := "Numero1: " + datos.Numero1 + "," + "Numero2: " + datos.Numero2 + "," + "Operacion: " + operacion + "," + "Resultado: " + datos.Resultado + "," + "Fecha: " + actualstring + "\n"
	_, err = archivo.WriteString(linea) //escribimos la linea nueva al final del archivo
	if err != nil {
		log.Fatal(err)
	}

}

func getVerificacion(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	respuesta := saludo{Saludo: "Bienvenido el servidor esta ejecutandose correctamente!!"}
	json.NewEncoder(res).Encode(respuesta)
}

func endPointRegistros(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	DBconexion := conectarBaseDeDatos()
	datalogs, err := DBconexion.Query("SELECT * FROM registros")
	if err != nil {
		panic(err.Error())
	}

	var operaciones []registros

	for datalogs.Next() {
		nuevo := registros{}
		var n1 string
		var n2 string
		var s string
		var r string
		var d string
		var id string
		err = datalogs.Scan(&id, &n1, &n2, &s, &r, &d)
		if err != nil {
			respuesta := resultado{Resultado: "Error"}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(respuesta)
			return
		}
		nuevo.Numero1 = n1
		nuevo.Numero2 = n2
		nuevo.Operacion = s
		nuevo.Resultado = r
		nuevo.Fecha = d
		operaciones = append(operaciones, nuevo)

	}

	DBconexion.Close()
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
	guardarOperacion(registro)
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
	guardarOperacion(registro)
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
	guardarOperacion(registro)
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
		actual := time.Now()
		actualstring := actual.Format("02/01/2006 15:04:05")
		registro := registros{Numero1: entrada.Numero1, Numero2: entrada.Numero2, Operacion: entrada.Simbolo, Resultado: "error", Fecha: actualstring}
		guardarOperacion(registro)
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
	guardarOperacion(registro)
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
