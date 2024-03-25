package mbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tvanriper/mbox"
	"io/ioutil"
	"log"
	"myapp.com/enron/helpers"
	"net/mail"
	"os"
	"runtime/pprof"
	"strings"
)

// funcion para leer una archibo mbox, como ruta tiene una constante con la ruta predeterminada
func ReadMbox() {
	/*
		generar archivo para hacer profiling a cpu
	*/
	cpuFile, err := os.Create("cpuReadMbox.prof")
	if err != nil {
		log.Fatal("No se pudo crear el archivo de CPU profile: ", err)
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("No se pudo iniciar el CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	/*
		finaiza código para generar cpu profiling
	*/
	fmt.Println("Comenzando a leer Mbox")
	//abrir o crear archivo de error (logs)
	errorFile, errLog := os.OpenFile(helpers.PATH+"errorsReading.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if errLog != nil {
		log.Fatal(errLog)
	}
	// Crear un nuevo logger que escribe en el archivo de log
	logger := log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// leer el archivo mbox
	datosComoBytes, err := os.ReadFile(helpers.PATH + "enron.mbox")
	if err != nil {
		logger.Println("Error al leer el archivo: ", err)
		return
	}

	file := bytes.NewBuffer([]byte(datosComoBytes))
	mailReader := mbox.NewReader(file)
	mailReader.Type = mbox.MBOXRD
	mailBytes := bytes.NewBuffer([]byte{})
	jsonsData := []map[string]string{}
	for err == nil {
		_, err = mailReader.NextMessage(mailBytes)
		msg, errMsg := mail.ReadMessage(bytes.NewBuffer(mailBytes.Bytes()))
		if errMsg != nil {
			logger.Println("Error: ", errMsg)
			mailBytes.Reset()
			continue
		}

		stringBytes, errByte := ioutil.ReadAll(msg.Body)
		if errByte != nil {
			logger.Println("Error: ", errByte)
			mailBytes.Reset()
			continue
			//panic(err)
		}
		strMessage := string(stringBytes)
		split := strings.Split(strMessage, "\r\n")
		aJson := processString(split)
		jsonsData = append(jsonsData, aJson)

		mailBytes.Reset()
	}
	fmt.Println("Terminada lectura Mbox")
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
	file, err := os.Create(helpers.PATH + "data.json")
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
