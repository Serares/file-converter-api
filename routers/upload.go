package routers

import (
	"net/http"

	"github.com/Serares/file-converter/handlers"
	"github.com/gorilla/mux"
)

func RegisterUploadRoutes(router *mux.Router, uploadHandler *handlers.UploadHandler) {
	router.HandleFunc("/upload_pdf", uploadHandler.UploadPdf).Methods(http.MethodPost)
}
