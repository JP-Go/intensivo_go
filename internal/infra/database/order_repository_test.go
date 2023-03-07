package database

import (
	"database/sql"
	"testing"

	"github.com/JP-Go/go-intensivo/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL,tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 12.0)
	suite.NoError(err)
	suite.NoError(order.TotalPrice())

	println("Got Here")
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

}
