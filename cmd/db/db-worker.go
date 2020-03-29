package db

type DbWorker struct {
	File string
}

func (db *DbWorker) Connect() {
	connection, err := gorm.Open("sqlite3", db.File)
	if err != nil {
		panic("failed to connect database")
	}
	defer connection.Close()
}
