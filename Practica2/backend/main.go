package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"os/user"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

		cmd := exec.Command("sh", "-c", "cat /proc/cpu_202010223")
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

		cmd2 := exec.Command("sh", "-c", "cat /proc/ram_202010223")
		out2, err := cmd2.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		output2 := string(out2[:])
		fmt.Println("DATOS OBTENIDOS DESDE EL MODULO RAM------------------:")
		fmt.Println("")
		fmt.Println(output2)
		fmt.Println("--------------------------------------------------------:")
		time.Sleep(80 * time.Second)
	}
}

func actualizarUsuario(jsonStr string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", fmt.Errorf("Error decodificando JSON: %s", err)
	}

	for _, proceso := range data["procesos"].([]interface{}) {
		usuarioID := proceso.(map[string]interface{})["usuario"].(string)
		fmt.Println("Valor actual de usuario ID:", usuarioID)

		usuario, err := user.LookupId(usuarioID)
		if err != nil {
			return "", fmt.Errorf("Error obteniendo usuario: %s", err)
		}
		fmt.Println("Nombre de usuario:", usuario.Username)

		proceso.(map[string]interface{})["usuario"] = usuario.Username
		fmt.Println("Nuevo valor de usuario:", proceso.(map[string]interface{})["usuario"].(string))
	}

	newData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Error codificando JSON: %s", err)
	}

	return string(newData), nil
}
