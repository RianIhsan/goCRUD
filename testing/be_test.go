package testing_test

import (
	"os"
	"testing"

	"github.com/RianIhsan/ex-go-crud-icc/database"
	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Simulate loading environment variables
	os.Setenv("DSN", "root:jMDtaKFBirHCuoN9pzNY@tcp(containers-us-west-76.railway.app:7357)/railway?charset=utf8mb4&parseTime=True&loc=Local")

	// Call the ConnectDB function
	database.ConnectDB()

	// Ensure DB variable is not nil
	assert.NotNil(t, database.DB)

	// Close the database connection
	sqlDB, err := database.DB.DB()
	assert.NoError(t, err)
	sqlDB.Close()
}
