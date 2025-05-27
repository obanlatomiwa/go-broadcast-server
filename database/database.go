package database

import (
	"errors"
	"fmt"
	"github.com/obanlatomiwa/go-broadcast-server/models"
	"github.com/obanlatomiwa/go-broadcast-server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitialiseDatabase() {
	var (
		databaseUser     string = utils.GetValueFromConfigFile("DB_USER")
		databasePassword string = utils.GetValueFromConfigFile("DB_PASSWORD")
		databaseHost     string = utils.GetValueFromConfigFile("DB_HOST")
		databasePort     string = utils.GetValueFromConfigFile("DB_PORT")
		databaseName     string = utils.GetValueFromConfigFile("DB_NAME")
	)

	// data source for MySQL
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	// variable to store the error
	var err error

	// create a connection to the database
	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	migrateDB()
}

func migrateDB() {
	var err error
	err = DB.AutoMigrate(&models.Message{}, &models.Client{})
	if err != nil {
		return
	}
}

func CreateItem(clientId string, text string) {
	newMessage := models.Message{
		ClientId: clientId,
		Text:     text,
		Date:     time.Now(),
	}
	DB.Create(&newMessage)
}

func CreateClient(clientId string) {
	newClient := models.Client{
		ClientId: clientId,
		Status:   "ONLINE",
	}
	DB.Create(&newClient)
}

func GetAllMessages() []models.Message {
	// connect to db
	InitialiseDatabase()

	var messages []models.Message
	DB.Order("date desc").Find(&messages)
	return messages
}

func GetAllClients() []models.Client {
	// connect to db
	InitialiseDatabase()

	var clients []models.Client
	DB.Order("").Find(&clients)
	return clients
}

func GetClientById(id string) (models.Client, error) {
	var client models.Client
	result := DB.First(&client, "client_id = ?", id)

	if result.RowsAffected == 0 {
		return models.Client{}, errors.New("Item not found")
	}
	return client, nil
}

func UpdateClient(id string) {
	DB.Model(&models.Client{}).Where("client_id = ?", id).Updates(models.Client{Status: "OFFLINE"})
}

func CleanDatabaseData() {
	// connect to db
	InitialiseDatabase()

	// remove all data inside the messages table
	messages := DB.Exec("TRUNCATE messages")
	clients := DB.Exec("TRUNCATE clients")

	// check if the operation failed
	operationFailed := messages.Error != nil || clients.Error != nil

	if operationFailed {
		panic(errors.New("operation Failed. Cleaning DB Data Failed"))
	}

	fmt.Println("DB Cleaned Successfully")
}
