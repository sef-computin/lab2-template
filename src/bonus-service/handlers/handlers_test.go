package handlers

import (
	"bytes"
	"encoding/json"
	"lab2/src/bonus-service/models"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreatePrivilege(t *testing.T) {
	privilege := &models.Privilege{
		Username: "tuser1",
		Balance:  0,
	}

	jsonValue, _ := json.Marshal(privilege)
	req, _ := http.NewRequest("POST", "localhost:8050/api/v1/bonus/privilege", bytes.NewBuffer(jsonValue))

	assert.Equal(t, http.StatusOK, req.Response.Status)
}

func TestUpdatePrivilege(t *testing.T) {
	// dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	// 	"postgres", 5432, "postgres", "privileges", "postgres")

	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// bh := &BonusHandler{
	// 	DBHandler: *dbhandler.InitDBHandler(db),
	// }

	// privilege := &models.Privilege{
	// 	Username: "tuser1",
	// 	Balance:  0,
	// }

	// r := SetUpRouter()
	// r.POST("bonus/:username", bh.UpdatePrivilegeHandler)

	// jsonValue, _ := json.Marshal(privilege)
	// req, _ := http.NewRequest("POST", "/bonus/tuser1", bytes.NewBuffer(jsonValue))
	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPrivilegeByUsername(t *testing.T) {

	// dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	// 	"postgres", 5432, "postgres", "privileges", "postgres")

	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// bh := &BonusHandler{
	// 	DBHandler: *dbhandler.InitDBHandler(db),
	// }

	// r := SetUpRouter()
	// r.GET("bonus/:username", bh.GetPrivilegeByUsernameHandler)

	// req, _ := http.NewRequest("GET", "/bonus/tuser1", nil)
	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// var privilege models.Privilege
	// json.Unmarshal(w.Body.Bytes(), &privilege)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.NotEmpty(t, privilege)

	// req, _ = http.NewRequest("GET", "/bonus/tuser99", nil)
	// w = httptest.NewRecorder()
	// r.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreatePrivilegeHistoryHandler(t *testing.T) {
	// dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	// 	"postgres", 5432, "postgres", "privileges", "postgres")

	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// bh := &BonusHandler{
	// 	DBHandler: *dbhandler.InitDBHandler(db),
	// }

	// record := &models.PrivilegeHistory{
	// 	PrivilegeID:   555,
	// 	TicketUID:     "tuid",
	// 	Date:          "12.12.2012",
	// 	BalanceDiff:   -150,
	// 	OperationType: "FILL",
	// }

	// r := SetUpRouter()
	// r.POST("/bonus", bh.CreatePrivilegeHistoryHandler)

	// jsonValue, _ := json.Marshal(record)
	// req, _ := http.NewRequest("POST", "/bonus", bytes.NewBuffer(jsonValue))
	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetHistoryById(t *testing.T) {
	// dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	// 	"postgres", 5432, "postgres", "privileges", "postgres")

	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// bh := &BonusHandler{
	// 	DBHandler: *dbhandler.InitDBHandler(db),
	// }

	// r := SetUpRouter()
	// r.GET("bonus/history/:privilegeId", bh.GetPrivilegeByUsernameHandler)

	// req, _ := http.NewRequest("GET", "bonus/history/555", nil)
	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// var privilege models.Privilege
	// json.Unmarshal(w.Body.Bytes(), &privilege)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.NotEmpty(t, privilege)

	// req, _ = http.NewRequest("GET", "bonus/history/999", nil)
	// w = httptest.NewRecorder()
	// r.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
