package controller

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/web"
	"ahya37/xyz_multifinance/service"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type UploadControllerImpl struct {
	UploadService   service.UploadService
	KonsumenService service.KonsumenService
}

func NewUploadController(uploadService service.UploadService, konsumenService service.KonsumenService) *UploadControllerImpl {
	return &UploadControllerImpl{
		UploadService:   uploadService,
		KonsumenService: konsumenService,
	}
}

func (controller *UploadControllerImpl) UploadKTP(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenIdStr := params.ByName("konsumenId")
	konsumenId, err := strconv.Atoi(konsumenIdStr)
	helper.PanicIfError(err)

	// cek konsumen if exist
	controller.KonsumenService.FindById(request.Context(), konsumenId)

	// limit 10  MB
	err = request.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		helper.PanicIfError(err)
	}

	file, header, err := request.FormFile("foto_ktp")
	helper.PanicIfError(err)
	defer file.Close()

	if !strings.HasPrefix(header.Header.Get("Content-type"), "image/") {
		helper.PanicIfError(err)
	}

	timestamp := time.Now().UnixNano() // generate a unique filename
	filename := fmt.Sprintf("ktp_%d_%s", konsumenId, strconv.FormatInt(timestamp, 10)+filepath.Ext(header.Filename))

	uploadDir := "./uploads/ktp"
	filePath := filepath.Join(uploadDir, filename)

	// create a directory if doesn't exist
	err = os.MkdirAll(uploadDir, os.ModePerm)
	helper.PanicIfError(err)

	// save the file
	dst, err := os.Create(filePath)
	helper.PanicIfError(err)
	defer dst.Close()

	_, err = io.Copy(dst, file)
	helper.PanicIfError(err)

	err = controller.UploadService.UpdateKTPFilename(request.Context(), konsumenId, filename)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": "KTP uploaded successfully", "filename": filename},
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *UploadControllerImpl) UploadFotoSelfie(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenIdStr := params.ByName("konsumenId")
	konsumenId, err := strconv.Atoi(konsumenIdStr)
	helper.PanicIfError(err)

	// cek konsumen if exist
	controller.KonsumenService.FindById(request.Context(), konsumenId)

	// limit 10  MB
	err = request.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		helper.PanicIfError(err)
	}

	file, header, err := request.FormFile("foto_selfie")
	helper.PanicIfError(err)
	defer file.Close()

	if !strings.HasPrefix(header.Header.Get("Content-type"), "image/") {
		helper.PanicIfError(err)
	}

	timestamp := time.Now().UnixNano() // generate a unique filename
	filename := fmt.Sprintf("foto_selfie_%d_%s", konsumenId, strconv.FormatInt(timestamp, 10)+filepath.Ext(header.Filename))

	uploadDir := "./uploads/foto/selfie"
	filePath := filepath.Join(uploadDir, filename)

	// create a directory if doesn't exist
	err = os.MkdirAll(uploadDir, os.ModePerm)
	helper.PanicIfError(err)

	// save the file
	dst, err := os.Create(filePath)
	helper.PanicIfError(err)
	defer dst.Close()

	_, err = io.Copy(dst, file)
	helper.PanicIfError(err)

	err = controller.UploadService.UpdateKTPFilename(request.Context(), konsumenId, filename)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": "Foto selfie uploaded successfully", "filename": filename},
	}

	helper.WriteToResponseBody(writer, webResponse)

}
