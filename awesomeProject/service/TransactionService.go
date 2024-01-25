package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

type TransactionService struct {
	TransactionDAO *dao.TransactionDAO
	CustomerDAO    *dao.CustomerDAO
	ItemDAO        *dao.ItemDAO
}

func (service *TransactionService) RecordTransaction(customerNIK int, tenorMonths int, loan int, codeItem string) (errS error) {
	approve, trxModel, err := dao.CustomerDAO.ApplyLoan(customerNIK, tenorMonths, loan)
	if err != nil {
		errS = err
		return
	}

	modelItem, err := dao.ItemDAO.GetItemByCode(codeItem)
	if err != nil {
		errS = err
		return
	}

	otr := calculateOTR(modelItem.Type, modelItem.Price)
	assetName := modelItem.Name
	trxModel.OTR = otr
	trxModel.AssetName = assetName

	if approve != 0 {
		errS = service.TransactionDAO.RecordTransaction(trxModel)
	}

	return
}

func (service *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	return service.TransactionDAO.GetAllTransactions()
}

const (
	DefaultBiayaAdministrasiPercentMotor      = 1.5
	DefaultBiayaAdministrasiPercentMobil      = 2.5
	DefaultBiayaAdministrasiPercentWhiteGoods = 1.0

	DefaultPajakPercentMotor      = 4.0
	DefaultPajakPercentMobil      = 6.0
	DefaultPajakPercentWhiteGoods = 2.0

	DefaultAsuransiPercentMotor      = 1.0
	DefaultAsuransiPercentMobil      = 2.0
	DefaultAsuransiPercentWhiteGoods = 0.5
)

func calculateOTR(assetType string, hargaBarang float64) float64 {
	var otr float64

	switch assetType {
	case "WhiteGoods":
		otr = ((DefaultBiayaAdministrasiPercentWhiteGoods / 100) * hargaBarang) + ((DefaultPajakPercentWhiteGoods / 100) * hargaBarang) + ((DefaultAsuransiPercentWhiteGoods / 100) * hargaBarang)
	case "Motor":
		otr = ((DefaultBiayaAdministrasiPercentMotor / 100) * hargaBarang) + ((DefaultPajakPercentMotor / 100) * hargaBarang) + ((DefaultAsuransiPercentMotor / 100) * hargaBarang)
	case "Mobil":
		otr = ((DefaultBiayaAdministrasiPercentMobil / 100) * hargaBarang) + ((DefaultPajakPercentMobil / 100) * hargaBarang) + ((DefaultAsuransiPercentMobil / 100) * hargaBarang)
	default:
		return 0.0
	}

	return otr
}
