package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MemoriaRam struct {
	MemoriaOcupada int64 `json:"memoria_ocupada"`
	MemoriaLibre   int64 `json:"memoria_libre"`
	MemoriaTotal   int64 `json:"memoria_total"`
}

var totalRAM int64

func CalcularPorcentajeMemoria(jsonString string) (string, error) {
	var stats MemoriaRam
	err := json.Unmarshal([]byte(jsonString), &stats)
	if err != nil {
		return "", err
	}
	porcentaje := (float64(stats.MemoriaOcupada) / float64(stats.MemoriaTotal)) * 100.0
	porcentajeString := fmt.Sprintf("%.2f", porcentaje)
	porcentajeFloat, err := strconv.ParseFloat(porcentajeString, 64)
	if err != nil {
		return "", err
	}
	totalRAM = stats.MemoriaTotal

	porcentajeJson := map[string]float64{"porcentaje": porcentajeFloat}
	porcentajeStringSalida, err := json.Marshal(porcentajeJson)
	if err != nil {
		return "", err
	}
	return string(porcentajeStringSalida), nil
}

// func conectarBaseDeDatos() (conexion *sql.DB) {
// 	nombreContenedor := "mysql-db"
// 	nombreDB := "logs"
// 	driver := "mysql"
// 	usuario := "root"
// 	contrasena := "1234"
// 	puerto := "3306"

// 	conexion, err := sql.Open(driver, usuario+":"+contrasena+"@tcp("+nombreContenedor+":"+puerto+")/"+nombreDB)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return conexion
// }

func main() {
	// DBconexion := conectarBaseDeDatos()
	for {

		cmd2 := exec.Command("sh", "-c", "cat /host/proc/ram_202010223")
		out2, err := cmd2.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		output2 := string(out2[:])
		fmt.Println("DATOS OBTENIDOS DESDE EL MODULO RAM------------------:")
		fmt.Println("")
		porcentajeString, err := CalcularPorcentajeMemoria(output2)
		if err != nil {
			panic(err)
		}
		fmt.Println(porcentajeString)
		fmt.Println(totalRAM)
		fmt.Println("--------------------------------------------------------:")

		cmd := exec.Command("sh", "-c", "cat /host/proc/cpu_202010223")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		output := string(out[:])
		fmt.Println("DATOS OBTENIDOS DESDE EL MODULO CPU------------------:")
		fmt.Println("")
		reemplazo := strings.ReplaceAll(output, "}{", "},{")
		jsonSTR, err := actualizarUsuario(reemplazo)
		if err != nil {
			fmt.Println("Error actualizando JSON:", err)
			return
		}
		fmt.Println(jsonSTR)
		fmt.Println("--------------------------------------------------------:")

		time.Sleep(1 * time.Second)
	}
}

func actualizarUsuario(jsonStr string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("error decodificando JSON: %s", err)
	}

	for _, proceso := range data["procesos"].([]interface{}) {
		usuarioID := proceso.(map[string]interface{})["usuario"].(string)
		// fmt.Println("Valor actual de usuario ID:", usuarioID)

		usuario, err := user.LookupId(usuarioID)
		if err != nil {
			return "", fmt.Errorf("error obteniendo usuario: %s", err)
		}
		// fmt.Println("Nombre de usuario:", usuario.Username)

		proceso.(map[string]interface{})["usuario"] = usuario.Username
		// fmt.Println("Nuevo valor de usuario:", proceso.(map[string]interface{})["usuario"].(string))

		ramStr := proceso.(map[string]interface{})["ram"].(string)
		// fmt.Println(ramStr)
		ram, err := strconv.ParseFloat(ramStr, 64)
		if err != nil {
			return "", fmt.Errorf("error convirtiendo RAM a número: %s", err)
		}
		ram /= float64(totalRAM)
		ram = math.Round(ram*100) / 100
		proceso.(map[string]interface{})["ram"] = fmt.Sprintf("%.2f", ram*100)

	}

	newData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error codificando JSON: %s", err)
	}

	return string(newData), nil
}
