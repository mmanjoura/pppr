package report

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mmanjoura/pppr/configuration/decoder"
)

func writeHeader(pyts []Payment, hdr []string) [][]string {
	data := make([][]string, len(pyts))
	headerRow := make([]string, 0)
	for _, title := range hdr {

		headerRow = append(headerRow, title)
	}
	data[0] = headerRow
	return data
}

// GenerateCsv ...
func GenerateCsv(pyts []Payment, hdr []string, conf decoder.Config, mid string) Report {

	returnVal := GetCsv(pyts, conf, mid)
	data := writeHeader(pyts, hdr)
	for i, pyt := range pyts {

		// Don't overrite the header
		if i == 0 {
			continue
		}
		row := make([]string, 0)
		row = append(row, pyt.TransactionID)
		row = append(row, pyt.MerchantID)
		row = append(row, pyt.TerminalID)
		row = append(row, pyt.CardNumberMasked[11:])
		row = append(row, strconv.FormatFloat(pyt.OriginalTransactionAmount, 'f', 6, 64))
		row = append(row, pyt.LocalCurrency)
		row = append(row, pyt.TransactionDate)
		row = append(row, strconv.FormatFloat(pyt.AcquirerFee, 'f', 6, 64))
		row = append(row, pyt.MarginRate)
		row = append(row, pyt.IsCardPresent)

		data[i] = row

	}

	writeCsvFile(conf.DropFolder+mid+"csvoutput.csv", data)
	return returnVal

}

func GetCsv(pyts []Payment, conf decoder.Config, mid string) Report {
	csvReport := Report{}

	for i, pyt := range pyts {
		csvReport.AcquirerID = pyt.AcquirerID
		csvReport.CreatedDate = time.Now().Format("2006-01-02")
		csvReport.CreatedTime = time.Now().Format("2006-01-02 15:04:05")[11:]
		csvReport.MerchantID = pyt.MerchantID
		csvReport.MerchantName = pyt.MerchantName
		csvReport.PaymentID = pyt.PaymentID
		csvReport.ReportType = strconv.Itoa(DAILY_STATEMENT)
		csvReport.FileName = pyt.MerchantName + ".csv"
		csvReport.FileType = "CSV"
		csvReport.FilePath = strings.Replace(conf.DropFolder, ".", "", -1) + strconv.Itoa(i) + pyt.MerchantName + ".csv"
		break
	}
	return csvReport

}

// `intToFloatString` takes an integer `n` and calculates
// the floating point value representing `n/100` as a string.
func intToFloatString(n int) string {
	intgr := n / 100
	frac := n - intgr*100
	return fmt.Sprintf("%d.%d", intgr, frac)
}

func writeCsvFile(name string, rows [][]string) {

	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	// We are going to write to a file, so any errors are relevant and
	// need to be logged. Hence the anonymous func instead of a one-liner.
	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
