package test

import (
	"ahya37/xyz_multifinance/app"
	"ahya37/xyz_multifinance/controller"
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/middleware"
	"ahya37/xyz_multifinance/model/domain"
	"ahya37/xyz_multifinance/model/repository"
	"ahya37/xyz_multifinance/service"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3307)/xyz-multifinance")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	KonsumenRepository := repository.NewKonsumenRepository()
	KonsumenService := service.NewKonsumenService(KonsumenRepository, db, validate)
	konsumenController := controller.NewKonsumenController(KonsumenService)

	router := app.NewRouter(konsumenController, nil)

	return middleware.NewAuthMiddleware(router)
}

func truncateCustomer(db *sql.DB) {
	db.Exec("TRUNCATE customer")

}

func TestCreateCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
					"nik": "1234500890123422",
					"full_name": "Test Nama",
					"legal_name": "Test Nama Juga.",
					"tempat_lahir": "Jakarta",
					"tanggal_lahir": "2024-01-01",
					"gaji": 5000000,
					"foto_ktp": "https://example.com/budi_ktp.jpg",
					"foto_selfie": "https://example.com/budi_selfie.jpg"
				}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost/api/customers", requestBody)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "1234500890123422", responseBody["data"].(map[string]interface{})["nik"])

}

func TestCreateCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
					"nik": "1234500890123422",
					"full_name": "",
					"legal_name": "Test Nama Juga.",
					"tempat_lahir": "Jakarta",
					"tanggal_lahir": "2024-01-01",
					"gaji": 5000000,
					"foto_ktp": "https://example.com/budi_ktp.jpg",
					"foto_selfie": "https://example.com/budi_selfie.jpg"
				}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost/api/customers", requestBody)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	KonsumenRepository := repository.NewKonsumenRepository()
	tanggalLahir, _ := time.Parse("2006-01-02", "2024-01-01")
	customer := KonsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:          "1234500890123422",
		FullName:     "Test Nama Juga",
		LegalName:    "Test Nama Juga.",
		TempatLahir:  "Jakarta",
		TanggalLahir: tanggalLahir,
		Gaji:         5000000,
		FotoKTP:      "https://example.com/budi_ktp.jpg",
		FotoSelfie:   "https://example.com/budi_ktp.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
					"nik": "1234500890123422",
					"full_name": "Test Nama Juga Update",
					"legal_name": "Test Nama Juga.",
					"tempat_lahir": "Jakarta",
					"tanggal_lahir": "2024-01-01",
					"gaji": 5000000,
					"foto_ktp": "https://example.com/budi_ktp.jpg",
					"foto_selfie": "https://example.com/budi_selfie.jpg"
				}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost/api/customers/"+strconv.Itoa(customer.Id), requestBody)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, customer.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Test Nama Juga Update", responseBody["data"].(map[string]interface{})["full_name"])

}

func TestUpdateCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	KonsumenRepository := repository.NewKonsumenRepository()
	tanggalLahir, _ := time.Parse("2006-01-02", "2024-01-01")
	customer := KonsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:          "1234500890123422",
		FullName:     "Test Nama Juga",
		LegalName:    "Test Nama Juga.",
		TempatLahir:  "Jakarta",
		TanggalLahir: tanggalLahir,
		Gaji:         5000000,
		FotoKTP:      "https://example.com/budi_ktp.jpg",
		FotoSelfie:   "https://example.com/budi_ktp.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
					"nik": "1234500890123422",
					"full_name": "",
					"legal_name": "Test Nama Juga.",
					"tempat_lahir": "Jakarta",
					"tanggal_lahir": "2024-01-01",
					"gaji": 5000000,
					"foto_ktp": "https://example.com/budi_ktp.jpg",
					"foto_selfie": "https://example.com/budi_selfie.jpg"
				}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost/api/customers/"+strconv.Itoa(customer.Id), requestBody)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	KonsumenRepository := repository.NewKonsumenRepository()
	tanggalLahir, _ := time.Parse("2006-01-02", "2024-01-01")
	customer := KonsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:          "1234500890123422",
		FullName:     "Test Nama Juga",
		LegalName:    "Test Nama Juga.",
		TempatLahir:  "Jakarta",
		TanggalLahir: tanggalLahir,
		Gaji:         5000000,
		FotoKTP:      "https://example.com/budi_ktp.jpg",
		FotoSelfie:   "https://example.com/budi_ktp.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost/api/customers/"+strconv.Itoa(customer.Id), nil)
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, customer.FullName, responseBody["data"].(map[string]interface{})["full_name"])
}

func TestGetCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost/api/customers/404", nil)
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	KonsumenRepository := repository.NewKonsumenRepository()
	tanggalLahir, _ := time.Parse("2006-01-02", "2024-01-01")
	customer := KonsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:          "1234500890123422",
		FullName:     "Test Nama Juga",
		LegalName:    "Test Nama Juga.",
		TempatLahir:  "Jakarta",
		TanggalLahir: tanggalLahir,
		Gaji:         5000000,
		FotoKTP:      "https://example.com/budi_ktp.jpg",
		FotoSelfie:   "https://example.com/budi_ktp.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost/api/customers/"+strconv.Itoa(customer.Id), nil)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost/api/customers/404", nil)
	request.Header.Add("Content-type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestGetListCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	KonsumenRepository := repository.NewKonsumenRepository()
	tanggalLahir, _ := time.Parse("2006-01-02", "2024-01-01")
	customer := KonsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:          "1234500890123422",
		FullName:     "Test Nama Juga",
		LegalName:    "Test Nama Juga.",
		TempatLahir:  "Jakarta",
		TanggalLahir: tanggalLahir,
		Gaji:         5000000,
		FotoKTP:      "https://example.com/budi_ktp.jpg",
		FotoSelfie:   "https://example.com/budi_ktp.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost/api/customers", nil)
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var customers = responseBody["data"].([]interface{})

	KonsumenResponse := customers[0].(map[string]interface{})

	assert.Equal(t, customer.Id, int(KonsumenResponse["id"].(float64)))
	assert.Equal(t, customer.FullName, KonsumenResponse["full_name"])
}

func TestGetListCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	KonsumenRepository := repository.NewKonsumenRepository()
	tanggalLahir, _ := time.Parse("2006-01-02", "2024-01-01")
	customer := KonsumenRepository.Save(context.Background(), tx, domain.Konsumen{
		Nik:          "1234500890123422",
		FullName:     "Test Nama Juga",
		LegalName:    "Test Nama Juga.",
		TempatLahir:  "Jakarta",
		TanggalLahir: tanggalLahir,
		Gaji:         5000000,
		FotoKTP:      "https://example.com/budi_ktp.jpg",
		FotoSelfie:   "https://example.com/budi_ktp.jpg",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost/api/customers", nil)
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var customers = responseBody["data"].([]interface{})

	KonsumenResponse := customers[0].(map[string]interface{})

	assert.Equal(t, customer.Id, int(KonsumenResponse["id"].(float64)))
	assert.Equal(t, customer.FullName, KonsumenResponse["full_name"])
}

func TestUnAuthorized(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost/api/customers", nil)
	request.Header.Add("X-API-KEY", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
}
