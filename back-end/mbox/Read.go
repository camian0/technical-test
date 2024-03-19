package mbox

import (
	"bytes"
	"fmt"
	"github.com/tvanriper/mbox"
	"io/ioutil"
	"net/mail"
	"os"
	"strings"
)

var mboxPath string = "D:\\Camilo\\Downloads\\emailExtracted\\enron in_sent.mbox"

func ReadMbox() {
	//lee todo el archibo mbox
	datosComoBytes, err := os.ReadFile(mboxPath)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	file := bytes.NewBuffer([]byte(datosComoBytes))
	mailReader := mbox.NewReader(file)
	mailReader.Type = mbox.MBOXRD

	mailBytes := bytes.NewBuffer([]byte{})
	for err == nil {
		_, err = mailReader.NextMessage(mailBytes)
		msg, err := mail.ReadMessage(bytes.NewBuffer(mailBytes.Bytes()))

		stringBytes, err := ioutil.ReadAll(msg.Body)
		strMessage := string(stringBytes)
		split := strings.Split(strMessage, "\r\n")
		if err != nil {
			panic(err)
		}

		fmt.Println(split)
		mailBytes.Reset()
	}
}
