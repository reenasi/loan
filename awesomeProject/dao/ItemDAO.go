package dao

import (
	"awesomeProject/model"
	"database/sql"
	"errors"
)

type ItemDAO struct {
	DB *sql.DB
}

func (dao *ItemDAO) GetItemByCode(code string) (*model.Item, error) {
	query := "SELECT id, name, code, type, price FROM item WHERE code = ?"

	row := dao.DB.QueryRow(query, code)
	var item model.Item

	err := row.Scan(&item.ID, &item.Name, &item.Code, &item.Type, &item.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("item not found")
		}
		return nil, err
	}

	return &item, nil
}
