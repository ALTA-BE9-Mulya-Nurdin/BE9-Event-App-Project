package data

import (
	"be9/event/features/events"
	"fmt"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}
}

func (repo *mysqlEventRepository) InsertData(insert events.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlEventRepository) GetAllData() (data []events.Core, err error) {
	var getAllData []Events
	tx := repo.db.Preload("User").Preload("Categorys").Find(&getAllData)
	if tx.Error != nil {
		return []events.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}

func (repo *mysqlEventRepository) GetData(id int) (data events.Core, err error) {
	var getData Events
	tx := repo.db.Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return events.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlEventRepository) DeleteData(id int) (row int, err error) {
	var deleteData Events
	tx := repo.db.Unscoped().Delete(&deleteData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to deleted data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlEventRepository) UpdatedData(id int, insert events.Core) (row int, err error) {
	tx := repo.db.Model(&Events{}).Where("id = ?", id).Updates(Events{Image: insert.Image, Name: insert.Name, Address: insert.Address, Date: insert.Date, Price: insert.Price, Quote: insert.Quote, Longitude: insert.Longitude, Latitude: insert.Latitude, Link: insert.Link, Description: insert.Description, Status: insert.Status})
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to updated data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlEventRepository) GetToken(id int, idToken int) (data events.Core, err error) {
	var getData Events
	tx := repo.db.Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return events.Core{}, tx.Error
	}
	if getData.toCore().User.ID != idToken {
		return events.Core{}, err
	}
	return getData.toCore(), nil
}

func (repo *mysqlEventRepository) GetDataUserEvent(idToken int) (data []events.Core, err error) {
	var getDataEvents []Events
	tx := repo.db.Where("user_id = ?", idToken).Preload("User").Preload("Categorys").Find(&getDataEvents)
	if tx.Error != nil {
		return []events.Core{}, tx.Error
	}
	return toCoreList(getDataEvents), nil
}

func (repo *mysqlEventRepository) GetDataHistoryEvent(idToken int) (data []events.Core, err error) {
	var getHisEvents []Events
	tx := repo.db.Preload("User").Preload("Events").First(&getHisEvents, idToken)
	if tx.Error != nil {
		return []events.Core{}, tx.Error
	}
	return toCoreList(getHisEvents), nil
}
