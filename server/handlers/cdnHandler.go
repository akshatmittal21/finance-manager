package handlers

import (
	"akshu-21/finance-manager/model"
	"akshu-21/finance-manager/services"

	"crypto/rand"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// MaxUploadSize :
const MaxUploadSize = 20 * 1024 * 1024 // 2 mb
// UploadPath :
const UploadPath = ""

// UploadFileHandler :
func UploadFileHandler() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// validate file size
		w := c.Writer
		c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MaxUploadSize)
		if err := c.Request.ParseMultipartForm(MaxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		// check file type, detectcontenttype only needs the first 512 bytes
		detectedFileType := http.DetectContentType(fileBytes)
		switch detectedFileType {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		// fileName := randToken(12)
		fileName := c.Request.FormValue("name")
		fileEndings, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join(UploadPath, fileName+fileEndings[0])
		fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			fmt.Println(err)
			renderError(w, "Creation Error : CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		transaction := model.Transaction{}
		transaction.TransactionID = fileName
		transaction.ReceiptPath = fileName + fileEndings[0]
		services.UpdateTransaction(&transaction)
		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
