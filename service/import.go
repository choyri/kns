package service

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/choyri/kns/model"
	"github.com/choyri/kns/store"
	"github.com/choyri/kns/support"
	"github.com/jinzhu/gorm"
	"github.com/t-tiger/gorm-bulk-insert"
	"io"
	"regexp"
	"strings"
)

func ImportFromReader(r io.Reader) (uint, error) {
	file, err := excelize.OpenReader(r)
	if err != nil {
		return 0, fmt.Errorf("解析 XLSX 时出现了错误：%w", err)
	}

	_, err = isFileValid(file)
	if err != nil {
		return 0, err
	}

	var (
		tx             = store.GetMySQL().Begin()
		importRecordID uint
		originalAmount uint
		currentAmount  uint
	)

	defer tx.Rollback()

	originalAmount = getOrderAmount(tx)

	importRecordID, err = createImportRecord(tx, file)
	if err != nil {
		return 0, err
	}

	err = createOrders(tx, file, importRecordID)
	if err != nil {
		return 0, err
	}

	currentAmount = getOrderAmount(tx)

	return currentAmount - originalAmount, tx.Commit().Error
}

func getOrderAmount(tx *gorm.DB) uint {
	var ret uint
	tx.Model(&model.Order{}).Count(&ret)
	return ret
}

func createImportRecord(tx *gorm.DB, file *excelize.File) (uint, error) {
	var (
		record model.ImportRecord
		err    error
		a2     string
	)

	a2, err = file.GetCellValue(SheetName, "A2")
	if err != nil {
		return 0, fmt.Errorf("获取单元格 A2 失败：%w", err)
	}

	dates := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).FindAllString(a2, -1)
	if len(dates) != 2 {
		return 0, errors.New("单元格 A2 找不到开始和结束日期，请重新确认")
	}

	for _, v := range dates {
		_, err := support.ParseTime(v, support.DateDefaultLayout)
		if err != nil {
			return 0, fmt.Errorf("单元格 A2 中的日期解析失败：%s", err)
		}
	}

	record = model.ImportRecord{
		StartDate: dates[0],
		EndDate:   dates[1],
	}

	err = tx.Create(&record).Error
	if err != nil {
		err = fmt.Errorf("创建 import_records 失败：%w", err)
	}

	return record.ID, err
}

func createOrders(tx *gorm.DB, file *excelize.File, importRecordID uint) error {
	rows, err := file.GetRows(SheetName)
	if err != nil {
		return fmt.Errorf("获取 Rows 失败：%w", err)
	}

	fieldMap, err := renderFieldMap(rows)
	if err != nil {
		return fmt.Errorf("渲染 FieldMap 失败：%w", err)
	}

	var (
		records []interface{}
		rowLen  = len(rows)
	)

	for k, row := range rows {
		if k < HeaderLines {
			continue
		}

		if k == rowLen-1 {
			var shouldBreak bool

			for _, col := range row {
				if strings.Contains(col, FooterSubtotal) {
					shouldBreak = true
					break
				}
			}

			if shouldBreak {
				break
			}
		}

		record := model.Order{
			ImportRecordID:        importRecordID,
			CustomerName:          row[fieldMap[FieldNameCustomerName]],
			Salesman:              row[fieldMap[FieldNameSalesman]],
			CustomerOrderNumber:   row[fieldMap[FieldNameCustomerOrderNumber]],
			Brand:                 row[fieldMap[FieldNameBrand]],
			OrderNumber:           row[fieldMap[FieldNameOrderNumber]],
			SerialNumber:          support.Str2Uint(row[fieldMap[FieldNameSerialNumber]]),
			ProductNameCode:       row[fieldMap[FieldNameProductNameCode]],
			ProductNameChinese:    row[fieldMap[FieldNameProductNameChinese]],
			ProductNameEnglish:    row[fieldMap[FieldNameProductNameEnglish]],
			Ingredient:            row[fieldMap[FieldNameIngredient]],
			Specification:         row[fieldMap[FieldNameSpecification]],
			Color:                 row[fieldMap[FieldNameColor]],
			ColorNumber:           row[fieldMap[FieldNameColorNumber]],
			CustomerVersionNumber: row[fieldMap[FieldNameCustomerVersionNumber]],
		}

		record.TrimSpace()

		records = append(records, record)
	}

	err = gormbulk.BulkInsert(tx, records, 500)
	if err != nil {
		return fmt.Errorf("创建 orders 失败：%w", err)
	}

	return nil
}

func isFileValid(file *excelize.File) (bool, error) {
	a1, err := file.GetCellValue(SheetName, "A1")
	if err != nil {
		return false, fmt.Errorf("获取单元格 A1 失败：%w", err)
	}

	if strings.Contains(a1, A1Title) {
		return true, nil
	}

	return false, fmt.Errorf("单元格 A1 不包含「%s」，请重新确认", A1Title)
}

func renderFieldMap(rows [][]string) (map[string]int, error) {
	if len(rows) < HeaderLines+1 {
		return nil, fmt.Errorf("表格至少需要 %d 行，请重新确认", HeaderLines+1)
	}

	ret := make(map[string]int)

	for k, v := range rows[2] {
		ret[v] = k
	}

	return ret, nil
}
