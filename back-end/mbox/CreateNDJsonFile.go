package mbox

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const PATH string = "D:\\Camilo\\Downloads\\emailExtracted\\"

//var INDEX_TABLE string = "{'index': {'_index': 'enron_go'}}"

func CreateNJFile() {
	jsonFile, err := os.ReadFile(path + "datos.json")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	jsonMaps := []map[string]string{}
	err = json.Unmarshal(jsonFile, &jsonMaps)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	fileIndex := 0
	ndFile, err := os.Create(path + "data" + strconv.Itoa(fileIndex) + ".ndjson")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer ndFile.Close()
	// Crear un escritor de buffer para el archivo
	writer := bufio.NewWriter(ndFile)
	for i, line := range jsonMaps {
		j, err := json.Marshal(line)
		if err != nil {
			return
		}
		fmt.Fprintln(writer, INDEX_TABLE)
		fmt.Fprintln(writer, string(j))

		if i%2 == 0 && i > 0 {
			writer.Flush()
			ndFile.Close()
			fileIndex += 1
			ndFile, err = os.Create(path + "data" + strconv.Itoa(fileIndex) + ".ndjson")
			writer = bufio.NewWriter(ndFile)
			if err != nil {
				fmt.Println("Error al crear el archivo:", err)
				return
			}
		}
		if i == 6 {
			break
		}
	}

	// Asegurarse de que todas las líneas se han escrito en el archivo
	writer.Flush()

}
