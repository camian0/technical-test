package mbox

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"myapp.com/enron/helpers"
	"os"
	"runtime/pprof"
	"strconv"
)

const PATHNDFILE = "ndGo\\data"
const EXTENSIONFILE = ".ndjson"

func CreateNJFile() {
	/*
		generar archivo para hacer profiling a cpu
	*/
	cpuFile, err := os.Create("cpuCreateND.prof")
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
	fmt.Println("Comenzando a crear archivos ND")
	jsonFile, err := os.ReadFile(helpers.PATH + "data.json")
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
	indexAdded := 0
	quantity := 10000
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
		if startSplit > finishSplit {
			writer.Flush()
			ndFile.Close()
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
		startSplit += 1
	}

	// Asegurarse de que todas las líneas se han escrito en el archivo
	writer.Flush()
	//Escribir un reporte cuanto termine de procesar
	fmt.Println("Finalizado creacion archivos ND")
	fmt.Printf("Total de archivos: %d \n", i)
	fmt.Printf("Total de archivos agregados divididos : %d \n", indexAdded)
}
