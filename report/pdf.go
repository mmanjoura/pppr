package report

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/boombuler/barcode/qr"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
	"github.com/mmanjoura/pppr/configuration-svc/decoder"
)

//GeneratePdf ...
func GeneratePdf(pyts []Payment, hdr []string, conf decoder.Config, mid string) Report {

	returnVal := GetPdf(pyts, conf, mid)
	data := make([][]string, len(pyts))
	var pdf = &gofpdf.Fpdf{}
	for i, pyt := range pyts {
		row := make([]string, 0)
		row = append(row, pyt.TransactionID)
		row = append(row, pyt.MerchantID)
		row = append(row, pyt.TerminalID)
		row = append(row, pyt.CardNumberMasked[10:])
		row = append(row, strconv.FormatFloat(pyt.OriginalTransactionAmount, 'f', 6, 64))
		row = append(row, pyt.LocalCurrency)
		row = append(row, pyt.TransactionDate)
		row = append(row, strconv.FormatFloat(pyt.AcquirerFee, 'f', 6, 64))
		row = append(row, pyt.MarginRate)
		row = append(row, pyt.IsCardPresent)

		data[i] = row

		//template, pdf := CreateTemplate(hdr)

		// Invoice Title
		pdf.SetDrawColor(62, 181, 95)
		pdf.SetY(20)
		pdf.Cell(50, 0, "")
		pdf.CellFormat(80, 20, "Invoice Summary", "1", 0, "TL", false, 0, "")

		// print 1000 transaction at time
		pdf = table(data, hdr)

		if pdf.Err() {
			log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
		}

	}
	err := savePDF(pdf, conf.DropFolder+mid+"-"+time.Now().Format("20060102150405")+"-Pdfoutput")
	if err != nil {
		log.Fatalf("Cannot save PDF: %s|n", err)

	}
	return returnVal
}

func GetPdf(pyts []Payment, conf decoder.Config, mid string) Report {
	PdfReport := Report{}

	for i, pyt := range pyts {
		PdfReport.AcquirerID = pyt.AcquirerID
		PdfReport.CreatedDate = time.Now().Format("2006-01-02")
		PdfReport.CreatedTime = time.Now().Format("2006-01-02 15:04:05")[11:]
		PdfReport.MerchantID = pyt.MerchantID
		PdfReport.MerchantName = pyt.MerchantName
		PdfReport.PaymentID = pyt.PaymentID
		PdfReport.ReportType = strconv.Itoa(DAILY_STATEMENT)
		PdfReport.FileName = pyt.MerchantName + ".pdf"
		PdfReport.FileType = "pdf"
		PdfReport.FilePath = strings.Replace(conf.DropFolder, ".", "", -1) + strconv.Itoa(i) + pyt.MerchantName + ".pdf"
		break
	}
	return PdfReport

}

/* ### How Cell() and Ln() advance the output position

As mentioned in the comments, the `Cell()` method takes no coordinates.
Instead, the PDF document maintains the current output position internally,
and advances it to the right by the length of the cell being written.

Method `Ln()` moves the output position back to the left border and down
by the provided value. (Passing `-1` uses the height of the recently written cell.)

*/

// Having created the initial document, we can now create the table header.
// This time, we generate a formatted cell with a light grey as the
// background color.
func header(pdf *gofpdf.Fpdf, hdr []string) (*gofpdf.Fpdf, map[int]float64) {
	//pdf.SetFont("Arial", "B", 5)
	//pdf.SetFillColor(240, 240, 240)
	m := make(map[int]float64)
	for i, str := range hdr {
		// The `CellFormat()` method takes a couple of parameters to format
		// the cell. We make use of this to create a visible border around
		// the cell, and to enable the background fill.
		cellWidth := float64(len(str))
		//headerWidth = append(headerWidth, cellWidth)
		m[i] = cellWidth

		pdf.CellFormat(float64(240/12), 7, str, "0", 0, "", false, 0, "")
	}

	// Passing `-1` to `Ln()` uses the height of the last printed cell as
	// the line height.
	pdf.Ln(-1)
	return pdf, m
}

func table(tbl [][]string, hdr []string) *gofpdf.Fpdf {

	// Every column gets aligned according to its contents.
	align := []string{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"}

	//template, pdf := CreateTemplate(hdr)
	template, pdf := CreateTemplate(hdr)

	pdf.AddPage()
	pdf.UseTemplate(template)
	pdf.Ln(44)
	for j, line := range tbl {
		//fmt.Println("J" + strconv.Itoa(j))

		if j > 0 && j%25 == 0 {
			pdf.AddPage()
			pdf.UseTemplate(template)
			pdf.Ln(44)
		}

		// if j == 200 {
		// 	break
		// }

		for i, str := range line {
			//fmt.Println("i" + strconv.Itoa(i))
			// Again, we need the `CellFormat()` method to create a visible
			// border around the cell. We also use the `alignStr` parameter
			// here to print the cell content either left-aligned or
			// right-aligned.

			// Noice the widht of each line "240/12"
			pdf.CellFormat(float64(240/12), 7, str, "0", 0, align[i], false, 0, "")

		}

		pdf.Ln(-1)

	}
	return pdf
}

// // Next, let's not forget to impress our boss by adding a fancy image.
// func image(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
// 	// The `ImageOptions` method takes a file path, x, y, width, and height
// 	// parameters, and an `ImageOptions` struct to specify a couple of options.
// 	// Scale the barcode to 200x200 pixels
// 	// Create the barcode

// 	pdf.ImageOptions("qrcode.png", 225, 10, 25, 25, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")

// 	//pdf.ImageOptions("stats.png", 225, 10, 25, 25, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
// 	return pdf
// }

func savePDF(pdf *gofpdf.Fpdf, fileName string) error {
	return pdf.OutputFileAndClose(fileName + ".pdf")
}

func planetAddress(tpl *gofpdf.Tpl) {
	tpl.SetY(10)
	tpl.Cell(10, 0, "")
	tpl.CellFormat(30, 3, "Planet Payment Asia Pacific Pte Ltd", "", 0, "M", false, 0, "")

	tpl.SetY(13)
	tpl.Cell(10, 0, "")
	tpl.CellFormat(30, 3, "12 Eu Tong Sen Street", "", 0, "M", false, 0, "")

	tpl.SetY(16)
	tpl.Cell(10, 0, "")
	tpl.CellFormat(30, 3, "04-171 The Central, 059819", "", 0, "M", false, 0, "")

	tpl.SetY(19)
	tpl.Cell(10, 0, "")
	tpl.CellFormat(30, 3, "Singapore", "", 0, "M", false, 0, "")

	tpl.SetY(22)
	tpl.Cell(10, 0, "")
	tpl.CellFormat(30, 3, "Tel: +353 9 616 5272", "", 0, "M", false, 0, "")

	tpl.SetY(25)
	tpl.Cell(10, 0, "")
	tpl.CellFormat(30, 3, "VAT: 1276498764", "", 0, "M", false, 0, "")

}

func merchantAddress_Static(tpl *gofpdf.Tpl) {
	tpl.SetY(10)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "Client Name:", "", 0, "M", false, 0, "")

	tpl.SetY(13)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "Client Number", "", 0, "M", false, 0, "")

	tpl.SetY(16)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "Period From:", "", 0, "M", false, 0, "")

	tpl.SetY(19)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "Period To:", "", 0, "M", false, 0, "")

	tpl.SetY(22)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "Payment Date:", "", 0, "M", false, 0, "")

	tpl.SetY(25)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "Report Date:", "", 0, "M", false, 0, "")

	tpl.SetY(28)
	tpl.Cell(140, 0, "")
	tpl.CellFormat(15, 3, "", "", 0, "M", false, 0, "")
}

func merchantAddress_Dynamic(tpl *gofpdf.Tpl) {
	tpl.SetY(10)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "MARINA BAY SANDS PTE LTD", "", 0, "M", false, 0, "")

	tpl.SetY(13)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "20050729R", "", 0, "M", false, 0, "")

	tpl.SetY(16)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "27/02/20", "", 0, "M", false, 0, "")

	tpl.SetY(19)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "29/02/20", "", 0, "M", false, 0, "")

	tpl.SetY(22)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "02/03/20", "", 0, "M", false, 0, "")

	tpl.SetY(25)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "02/03/20", "", 0, "M", false, 0, "")

	tpl.SetY(28)
	tpl.Cell(155, 0, "")
	tpl.CellFormat(20, 3, "", "", 0, "M", false, 0, "")
}

func invoiceSummary(tpl *gofpdf.Tpl) {

	tpl.SetDrawColor(62, 181, 95)
	tpl.SetY(20)
	tpl.Cell(50, 0, "")
	tpl.CellFormat(80, 20, "Invoice Summary", "1", 0, "TL", false, 0, "")

}

func qrCode(tpl *gofpdf.Tpl) {
	key := barcode.RegisterQR(tpl, "https://www.planetpayment.com/en/home/", qr.H, qr.Unicode)
	barcode.Barcode(tpl, key, 180, 14, 14, 14, false)
}

// CreateTemplate ...
func CreateTemplate(hd []string) (gofpdf.Template, *gofpdf.Fpdf) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	pdf.SetFont("Arial", "", 6)
	pdf.AliasNbPages("")

	pdf.SetFooterFunc(func() {
		pdf.Image("visamastercard.png", 10, 280, 15, 5, false, "", 0, "")
		pdf.SetY(278)
		pdf.CellFormat(0, 10, "Planet Payments is authorised by the Financial Conduct Authority under the Payment Service Regulations 2009 \n(Registry No. 815854) or the provision of payment services. V0.102",
			"", 0, "R", false, 0, "")
		pdf.SetY(285)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
			"", 0, "L", false, 0, "")

	})

	template := pdf.CreateTemplate(func(tpl *gofpdf.Tpl) {

		tpl.Image("planetlogo.png", 6, 6, 10, 0, false, "", 0, "")

		// Planet Address
		planetAddress(tpl)

		// Invoice Summary
		invoiceSummary(tpl)

		// QR Code
		qrCode(tpl)

		// merchant Address
		merchantAddress_Static(tpl)
		merchantAddress_Dynamic(tpl)

		// header position
		tpl.Ln(20)
		for _, str := range hd {

			// The `CellFormat()` method takes a couple of parameters to format
			// the cell. We make use of this to create a visible border around
			// the cell, and to enable the background fill.

			// Noice the width of each line "240/12" this need to match table
			tpl.CellFormat(float64(240/12), 7, str, "0", 0, "", false, 0, "")

		}
		tpl.SetDrawColor(62, 181, 95)
		tpl.SetLineWidth(0.5)
		tpl.Line(10, 53, 203, 53)

	})

	return template, pdf

}
