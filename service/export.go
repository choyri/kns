package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/choyri/kns/model"
	"github.com/choyri/kns/store"
	"io/ioutil"
	"os"
)

func Export(ids []uint) (file *os.File, err error) {
	var records []model.Order

	records, err = getRecords(ids)
	if err != nil {
		return
	}

	err = store.GetMySQL().Create(&model.ExportRecord{Amount: uint(len(records))}).Error
	if err != nil {
		err = fmt.Errorf("创建 export_records 失败：%w", err)
	}

	return newXlSX(records)
}

func getRecords(ids []uint) ([]model.Order, error) {
	var (
		err     error
		data    []model.Order
		dataMap = make(map[uint]model.Order)
	)

	err = store.GetMySQL().Where(ids).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("查找指定记录失败：%w", err)
	}

	for k, v := range data {
		dataMap[v.ID] = data[k]
	}

	ret := make([]model.Order, len(ids))

	for k, v := range ids {
		ret[k] = dataMap[v]
	}

	return ret, err
}

func newXlSX(records []model.Order) (file *os.File, err error) {
	var (
		f    = excelize.NewFile()
		axis string
		data = map[int][]interface{}{
			1: ExportTitles,
		}
	)

	for k, v := range records {
		data[k+2] = []interface{}{
			k + 1,
			"",
			v.CustomerOrderNumber,
			v.Brand,
			v.OrderNumber,
			v.SerialNumber,
			v.ProductNameCode,
			v.Ingredient,
			v.Specification,
			v.Color,
			v.CustomerVersionNumber,
		}
	}

	for k, v := range data {
		v := v
		axis, _ = excelize.JoinCellName("A", k)

		err = f.SetSheetRow(SheetName, axis, &v)
		if err != nil {
			err = fmt.Errorf("行赋值失败（Line %d）：%w", k, err)
			return
		}
	}

	tmpFile, err := ioutil.TempFile("", "kns_export")
	if err != nil {
		err = fmt.Errorf("创建临时文件失败：%w", err)
		return
	}

	err = f.SaveAs(tmpFile.Name())
	if err != nil {
		err = fmt.Errorf("写到临时文件失败：%w", err)
		return
	}

	return tmpFile, nil
}
