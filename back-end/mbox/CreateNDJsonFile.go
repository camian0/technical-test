package mbox

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

	ndFile, err := os.Create(path + "01.ndjson")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	defer ndFile.Close()
	// Crear un escritor de buffer para el archivo
	writer := bufio.NewWriter(ndFile)

	//m, err := json.Marshal(INDEX_TABLE)
	j, err := json.Marshal(jsonMaps[0])
	//write, err := ndFile.Write(m)
	//write, err = ndFile.Write(j)
	if err != nil {
		return
	}
	for _, line := range jsonMaps {
		fmt.Fprintln(writer, INDEX_TABLE)
		fmt.Fprintln(writer, string(j))
		fmt.Println(line)
		break
	}

	// Asegurarse de que todas las líneas se han escrito en el archivo
	writer.Flush()

}
