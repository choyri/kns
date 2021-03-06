package service

const (
	A1Title        = "訂單進度表"
	FooterSubtotal = "小計"
	HeaderLines    = 3
	SheetName      = "Sheet1"
)

const (
	FieldNameCustomerName          = "客戶名稱"
	FieldNameSalesman              = "業務"
	FieldNameCustomerOrderNumber   = "客戶單號"
	FieldNameBrand                 = "品牌"
	FieldNameOrderNumber           = "訂單號碼"
	FieldNameSerialNumber          = "序號"
	FieldNameProductNameCode       = "品名代碼"
	FieldNameProductNameChinese    = "中文品名"
	FieldNameProductNameEnglish    = "英文品名"
	FieldNameIngredient            = "成分"
	FieldNameSpecification         = "規格描述"
	FieldNameColor                 = "顏色"
	FieldNameColorNumber           = "色號"
	FieldNameCustomerVersionNumber = "客戶版號"
)

var SearchFields = []string{
	"customer_order_number",
	"brand",
	"order_number",
	"product_name_code",
	"product_name_chinese",
	"product_name_english",
	"color",
	"color_number",
	"customer_version_number",
}

var ExportTitles = []interface{}{
	"序號",
	"LAB NO",
	"客戶單號",
	"品牌",
	"訂單號碼",
	"序號",
	"品名代碼",
	"成分",
	"規格描述",
	"顏色",
	"客戶版號",
}
