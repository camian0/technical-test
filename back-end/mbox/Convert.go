package mbox

import (
	"fmt"
	"log"
	"myapp.com/enron/helpers"
	"os"
	"path/filepath"
	"regexp"
	"runtime/pprof"
	"strings"
	"time"
)

var foldersDic = map[string]string{"_sent_mail": "_sent_mail", "discussion_threads": "discussion_threads", "inbox": "inbox", "sent_items": "sent_items"}

const MAILDIR = helpers.PATH + "enron_mail_20110402\\maildir"
const MBOX = helpers.PATH + "\\enron.mbox"

func ConvertirMbox() {
	/*
		generar archivo para hacer profiling a cpu
	*/
	cpuFile, err := os.Create("cpuCreateMBOX.prof")
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

	fmt.Println("Comenzando a covertir data a mbox")
	startTime := time.Now()
	mbox, err := os.Create(MBOX)
	if err != nil {
		log.Fatal(err)
	}
	defer mbox.Close()

	//abrir o crear archivo de error (logs)
	errorFile, errLog := os.OpenFile(helpers.PATH+"errorsCreateMbox.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if errLog != nil {
		log.Fatal(errLog)
	}
	// Crear un nuevo logger que escribe en el archivo de log
	logger := log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	folders, err := os.ReadDir(MAILDIR)
	if err != nil {
		log.Fatal(err)
	}
	for _, folder := range folders {
		subFolderPath := filepath.Join(MAILDIR, folder.Name())
		openSubfolder, err := os.ReadDir(subFolderPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, subFolder := range openSubfolder {
			if !verfyFolder(subFolder.Name()) {
				continue
			}
			processFiles(filepath.Join(subFolderPath, subFolder.Name()), mbox, logger)
		}
	}

	defer errorFile.Close()
	finish := time.Since(startTime)
	fmt.Println("Finalizado covertir data a mbox")
	log.Printf("Conversion completed. Mbox file written to %s", MBOX)
	log.Printf("Completed on:  %s", finish)

}

func processFiles(receivedPath string, mbox *os.File, logger *log.Logger) {
	archieves, err := os.ReadDir(receivedPath)
	//Expresion regular para sacar la fecha y el remitente
	reFrom := regexp.MustCompile(`From: ([^\n\r]+)`)
	reDate := regexp.MustCompile(`Date: ([^\n\r]+)`)
	if err != nil {
		logger.Println("Error reading folder: ", receivedPath, err)
	}
	//Recorrer todos los archivos de la ruta recibida y agregar los datos al archivo mbox
	for _, file := range archieves {
		filePath := filepath.Join(receivedPath, file.Name())
		messageText, err := os.ReadFile(filePath)
		if err != nil {
			logger.Println("Error reading file", filePath, err)
			continue
		}
		textData := string(messageText)
		if textData == "" {
			continue
		}
		_from := reFrom.FindStringSubmatch(textData)[1]
		_date := reDate.FindStringSubmatch(textData)[1]
		newDate := ""
		var errParse error
		newDate, errParse = convertDate(_date)
		if errParse != nil {
			logger.Println("Error parsing date of file", filePath, errParse)
			continue
		}
		msg := fmt.Sprintf("From %s %s\n\n%s\n\n", _from, newDate, textData)
		_, err = mbox.WriteString(msg)
		if err != nil {
			logger.Println("Error writting to mbox of file: ", filePath, err)
		}
	}
	fmt.Println("terminados archivos de la carpeta: " + receivedPath)

}

// convertir zona horarioa de pdt a pst cuando se requiera
func convertDate(date string) (string, error) {
	if strings.Contains("(PDT)", date) {
		date = date[:len(date)-6]
		//index := strings.IndexAny(date, "-")
		//date = date[0 : index-1]
		zonePDT, err := time.LoadLocation("America/Los_Angeles")

		if err != nil {
			return "", err
		}

		timePDT, err := time.ParseInLocation(time.RFC1123Z, date, zonePDT)
		if err != nil {
			return "", err
		}

		// Convertir a PST (Pacific Standard Time)
		timePST := timePDT.In(time.UTC).Add(-7 * time.Hour)

		// Formatear la fecha y hora en PST
		formattedDate := timePST.Format(time.ANSIC)
		return formattedDate, nil
	}

	date = date[:len(date)-6]
	newDate, err := time.Parse(time.RFC1123Z, date)
	if err != nil {
		return "", err
	}
	formatted := newDate.Format(time.ANSIC)

	return formatted, nil
}

func verfyFolder(name string) bool {
	_, ok := foldersDic[name]
	if ok {
		return true
	}
	return false
}
