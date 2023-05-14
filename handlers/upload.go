package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dslipak/pdf"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) UploadPdf(w http.ResponseWriter, r *http.Request) {
	// Parse the file from the request
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded file
	tempFile, err := ioutil.TempFile("", "uploaded.pdf")
	if err != nil {
		http.Error(w, "Failed to create temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	// Save the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Verify if the uploaded file is in PDF format
	pdfReader, err := pdf.Open(tempFile.Name())
	if err != nil {
		http.Error(w, "Invalid PDF file", http.StatusBadRequest)
		return
	}
	var buf bytes.Buffer
	b, err := pdfReader.GetPlainText()
	if err != nil {
		http.Error(w, "Error reading pdf", http.StatusBadRequest)
		return
	}

	buf.ReadFrom(b)
	// Convert the PDF file to DOCX
	docxPath := tempFile.Name() + ".docx"
	err = ioutil.WriteFile(docxPath, []byte(buf.String()), 0644)
	if err != nil {
		http.Error(w, "Failed to convert PDF to DOCX", http.StatusInternalServerError)
		return
	}

	// Read the converted DOCX file
	docxData, err := ioutil.ReadFile(docxPath)
	if err != nil {
		http.Error(w, "Failed to read converted file", http.StatusInternalServerError)
		return
	}

	// Set appropriate headers for the response
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", "attachment; filename=converted.docx")

	// Write the converted file to the response
	w.Write(docxData)
}
