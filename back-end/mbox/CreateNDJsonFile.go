package mbox

import (
	"bufio"
	"encoding/json"
	"fmt"
	"myapp.com/enron/helpers"
	"os"
	"strconv"
)

const PATHNDFILE = "ndGo\\data"
const EXTENSIONFILE = ".ndjson"

func CreateNJFile() {
	jsonFile, err := os.ReadFile(helpers.PATH + "datos.json")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	//crear un arreglo de mapas, para guardar lo leido del archivo json
	jsonMaps := []map[string]string{}
	err = json.Unmarshal(jsonFile, &jsonMaps)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	//crear el primer archivos nd
	fileIndex := 0
	ndFile, err := os.Create(helpers.PATH + PATHNDFILE + strconv.Itoa(fileIndex) + EXTENSIONFILE)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer ndFile.Close()
	// Crear un escritor de buffer para el archivo
	writer := bufio.NewWriter(ndFile)
	indexSplit := 0
	indexAdded := 0
	quantity := 45
	startSplit := 0
	finishSplit := quantity
	length := len(jsonMaps) - 1
	var i int = 0
	for _, line := range jsonMaps {
		i += 1
		j, err := json.Marshal(line)
		if err != nil {
			fmt.Println("Error al codificar el archivo json:", err)
			return
		}
		/*comprobamos primero siempre el limite del archivo para escribir en el actual
		  o en un nuevo archivo
		*/
		if indexSplit > finishSplit {
			writer.Flush()
			ndFile.Close()
			indexSplit = 0
			fileIndex += 1
			startSplit = finishSplit
			/*
			  comprobar siempre la cantidad a dividir no supere el tamaño del arreglo,
			  si la supera, el indice final debe cambiar a la cantidad restante
			*/
			finishSplit += quantity
			if (length - startSplit) < quantity {
				finishSplit = startSplit + (length - startSplit)
			}
			ndFile, err = os.Create(helpers.PATH + PATHNDFILE + strconv.Itoa(fileIndex) + EXTENSIONFILE)
			writer = bufio.NewWriter(ndFile)
			if err != nil {
				fmt.Println("Error al crear el archivo:", err)
				return
			}
		}
		fmt.Fprintln(writer, helpers.INDEX_TABLE)
		fmt.Fprintln(writer, string(j))
		indexAdded += 1
		indexSplit += 1
	}

	// Asegurarse de que todas las líneas se han escrito en el archivo
	writer.Flush()
	//Escribir un reporte cuanto termine de procesar
	fmt.Printf("Total de archivos: %d \n", i)
	fmt.Printf("Total de archivos agregados divididos : %d \n", indexAdded)
}
