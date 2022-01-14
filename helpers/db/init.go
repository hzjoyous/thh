package db

func InitConnection() {
	connectDB()
	connectLeveldb()
}
