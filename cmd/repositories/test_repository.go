package repositories

import "github.com/quark_lt/cmd/models"

type TestRepository struct {
	connection *db.DbWorker
}

func NewTestRepository(connection *db.DbWorker) *TestRepository {
return &TestRepository{
	connection:connection
}
}

func (rep *TestRepository) FindAll() {

}
func (rep *TestRepository) FindById(id int) {

}
func (rep *TestRepository) Add(model *models.TestModel) {

}
func (rep *TestRepository) Remove(id int) {

}

func (rep *TestRepository) Update(model *models.TestModel) {

}
