package report

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mmanjoura/pppr/configuration/decoder"
)

func writeXlsHeader(hdr []string) *excelize.File {
	f := excelize.NewFile()
	// Based on the header number of Columns, in here we have 10
	cell := "A1,B1,C1,D1,E1,F1,G1,H1,I1,J1"
	cells := strings.Split(cell, ",")
	for i, title := range hdr {
		f.SetCellValue("Sheet1", cells[i], title)
	}
	return f
}

// GenerateXls ...
func GenerateXls(pyts []Payment, hdr []string, conf decoder.Config, mid string) Report {

	returnVal := GetXls(pyts, conf, mid)
	f := writeXlsHeader(hdr)
	cell := "A,B,C,D,E,F,G,H,I,J"
	cells := strings.Split(cell, ",")
	i := 2
	for _, pyt := range pyts {

		f.SetCellValue("Sheet1", cells[0]+strconv.Itoa(i), pyt.TransactionID)
		f.SetCellValue("Sheet1", cells[1]+strconv.Itoa(i), pyt.MerchantID)
		f.SetCellValue("Sheet1", cells[2]+strconv.Itoa(i), pyt.TerminalID)
		f.SetCellValue("Sheet1", cells[3]+strconv.Itoa(i), pyt.CardNumberMasked[10:])
		f.SetCellValue("Sheet1", cells[4]+strconv.Itoa(i), strconv.FormatFloat(pyt.OriginalTransactionAmount, 'f', 6, 64))
		f.SetCellValue("Sheet1", cells[5]+strconv.Itoa(i), pyt.LocalCurrency)
		f.SetCellValue("Sheet1", cells[6]+strconv.Itoa(i), pyt.TransactionDate)
		f.SetCellValue("Sheet1", cells[7]+strconv.Itoa(i), strconv.FormatFloat(pyt.AcquirerFee, 'f', 6, 64))
		f.SetCellValue("Sheet1", cells[8]+strconv.Itoa(i), pyt.MarginRate)
		f.SetCellValue("Sheet1", cells[9]+strconv.Itoa(i), pyt.IsCardPresent)
		i++

	}

	if err := f.SaveAs(conf.DropFolder + mid + "xlxoutput.xlsx"); err != nil {
		fmt.Println(err)

	}
	return returnVal
}

func GetXls(pyts []Payment, conf decoder.Config, mid string) Report {
	xlsReport := Report{}

	for i, pyt := range pyts {
		xlsReport.AcquirerID = pyt.AcquirerID
		xlsReport.CreatedDate = time.Now().Format("2006-01-02")
		xlsReport.CreatedTime = time.Now().Format("2006-01-02 15:04:05")[11:]
		xlsReport.MerchantID = pyt.MerchantID
		xlsReport.MerchantName = pyt.MerchantName
		xlsReport.PaymentID = pyt.PaymentID
		xlsReport.ReportType = strconv.Itoa(DAILY_STATEMENT)
		xlsReport.FileName = pyt.MerchantName + ".xls"
		xlsReport.FileType = "xls"
		xlsReport.FilePath = strings.Replace(conf.DropFolder, ".", "", -1) + strconv.Itoa(i) + pyt.MerchantName + ".xls"
		break
	}
	return xlsReport

}
