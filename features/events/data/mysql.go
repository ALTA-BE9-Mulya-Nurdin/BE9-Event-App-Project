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

func (repo *mysqlEventRepository) InsertEvent(input events.Core) (int, error) {
	event := fromCore(input)

	res := repo.db.Create(&event)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert input")
	}
	return int(res.RowsAffected), nil
}

func (repo *mysqlEventRepository) GetAllEvent() (data []events.Core, err error) {
	var getAllData []Event
	tx := repo.db.Find(&getAllData)
	if tx.Error != nil {
		return []events.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}

func (repo *mysqlEventRepository) GetEvent(id int) (data events.Core, err error) {
	var getData Event
	tx := repo.db.First(&getData, id)
	if tx.Error != nil {
		return events.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlEventRepository) DeleteEvent(id int) (row int, err error) {
	var getData Event
	tx := repo.db.Unscoped().Delete(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to deleted data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlEventRepository) UpdateEvent(id int, insert events.Core) (row int, err error) {
	tx := repo.db.Model(&Event{}).Where("id = ?", id).Updates(Event{
		Image:       insert.Image,
		Name:        insert.Name,
		Address:     insert.Address,
		Date:        insert.Date,
		Price:       insert.Price,
		Quota:       insert.Quota,
		Longitude:   insert.Longitude,
		Latitude:    insert.Latitude,
		Link:        insert.Link,
		Description: insert.Description,
		Status:      insert.Status,
	})
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to updated data event")
	}
	return int(tx.RowsAffected), nil
}
