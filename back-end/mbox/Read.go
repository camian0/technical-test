package mbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tvanriper/mbox"
	"io/ioutil"
	"net/mail"
	"os"
	"strings"
)

var path string = "D:\\Camilo\\Downloads\\emailExtracted\\"

// funcion para leer una archibo mbox, como ruta tiene una constante con la ruta predeterminada
func ReadMbox() {
	//lee todo el archibo mbox
	datosComoBytes, err := os.ReadFile(path + "enron in_sent.mbox")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	file := bytes.NewBuffer([]byte(datosComoBytes))
	mailReader := mbox.NewReader(file)
	mailReader.Type = mbox.MBOXRD
	mailBytes := bytes.NewBuffer([]byte{})
	jsonsData := []map[string]string{}
	for err == nil {
		_, err = mailReader.NextMessage(mailBytes)
		msg, err := mail.ReadMessage(bytes.NewBuffer(mailBytes.Bytes()))

		stringBytes, err := ioutil.ReadAll(msg.Body)
		if err != nil {
			panic(err)
		}
		strMessage := string(stringBytes)
		split := strings.Split(strMessage, "\r\n")
		aJson := processString(split)
		jsonsData = append(jsonsData, aJson)

		mailBytes.Reset()
	}

	createJsonFile(jsonsData)
}

func processString(splited []string) map[string]string {
	newJson := map[string]string{}
	for i := 0; i < len(splited); i++ {
		if splited[i] != "" && len(splited)-1 != i {
			before, after, _ := strings.Cut(splited[i], ":")
			newJson[before] = after
		}
		if i == len(splited)-1 {
			newJson["Content"] = splited[i]
		}
	}
	return newJson
}

// funcion para crear n archivo json, con un arreglo de maps
func createJsonFile(data []map[string]string) {
	// Crear o abrir el archivo para escritura
	file, err := os.Create(path + "data.json")
	if err != nil {
		fmt.Println("Error al crear el archivo json:", err)
		return
	}
	defer file.Close()

	// Codificar el arreglo de mapas a formato JSON
	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(data)
	if err != nil {
		fmt.Println("Error al codificar el JSON:", err)
		return
	}

	fmt.Println("Archivo JSON creado con éxito.")
}
