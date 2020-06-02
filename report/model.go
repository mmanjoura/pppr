/* *********************************************************************

			TDR	Transaction Detail Record
			TAR	Transaction Additional Record [Contains Acquirer Model]
			EM*	EMV Supplemental Record
			ECR	Ecommerce Record
			TIF	Interchange Fee Record
			TSF	Scheme Fee Record
			TAF	Acquirer Fee Record

***********************************************************************/

package report

const (
	DAILY_STATEMENT = iota
	MONTHLY_INVOINCE
	MONTHLY_REBATE
)
const y = iota

// Acquirer represents a single Acquirer profile.
// ID should be globally unique.
type Acquirer struct {
	ID        string     `json:"id"`
	Name      string     `json:"name,omitempty"`
	Merchants []Merchant `json:"merchants,omitempty"`
}

// Merchant is a field of a Acquirer.
// ID should be unique within the profile (at a minimum).
type Merchant struct {
	ID         string    `json:"id"`
	MerchantId string    `json:"merchantid,omitempty"`
	Payments   []Payment `json:"transactions,omitempty"`
}

type ProcessState struct {
	ID            string `json:"id" bson:"id"`
	ProcessType   int    `json:"processtype" bson:"processtype"`
	CreatedDate   string `json:"createddate" bson:"createddate"`
	CreatedTime   string `json:"createdtime" bson:"createdtime"`
	Approved      bool   `json:"approved" bson:"approved"`
	ProcessTypeID string `json:"processtypeid" bson:"processtypeid"` // This can paymentID, ReportID, StatementID ...

}

// Report
type Report struct {
	ReportID     string `json:"reportid,omitempty" bson:"reportid"`
	PaymentID    string `json:"paymentid,omitempty" bson:"paymentid"`
	CreatedDate  string `json:"createddate" bson:"createddate"`
	CreatedTime  string `json:"createdtime" bson:"createdtime"`
	AcquirerID   string `json:"acquirerid,omitempty" bson:"acquirerid"`
	MerchantID   string `json:"merchantid,omitempty" bson:"merchantid"`
	MerchantName string `json:"merchantname,omitempty" bson:"merchantname"`

	ReportType string `json:"reporttype,omitempty" bson:"reporttype"`
	FileType   string `json:"filetype,omitempty" bson:"filetype"`
	FileName   string `json:"filename,omitempty" bson:"filename"`
	FilePath   string `json:"filepath,omitempty" bson:"filepath"`
}

// type Report struct {
// 	ID         int    `json:"id,omitempty"`
// 	AcquirerID string `json:"acquirerid,omitempty"`
// 	StartDate  string `json:"startdate,omitempty"`
// 	EndDate    string `json:"enddate,omitempty"`
// 	ReportType string `json:"reporttype,omitempty"`
// }

type Message struct {
	CreatedAt     string `json:"created_at" bson:"creaed_at"`
	Level         string `json:"level" bson:"level"`
	ServiceName   string `json:"service_name"   bson:"service_name"`
	CallingMethod string `json:"calling_method"   bsn:"calling_method"`
	Host          string `json:"host"   bson:"host"`
	Body          string `json:"body"   bson:"message"`
	Latency       string `json:"latency"   bson:"latency"`
}

type Payment struct {

	// ID                        string `json:"id"   bson:"id"`
	TransactionID string `json:"transactionid"   bson:"transactionid"`
	PaymentID     string `json:"paymentid"   bson:"paymentid"`
	CreatedDate   string `json:"createddate" bson:"createddate"`
	CreatedTime   string `json:"createdtime" bson:"createdtime"`

	AcquirerID string `json:"acquirerid"   bson:"acquirerid"`
	MerchantID string `json:"merchantid"   bson:"merchantid"`
	TerminalID string `json:"terminalid"   bson:"terminalid"`

	OriginalTransactionAmount float64 `json:"originaltransactionamount"   bson:"originaltransactionamount"` //  00000001225(.01225)
	SettledAmount             float64 `json:"settledamount"   bson:"settledamount"`
	Fee                       float64 `json:"fee"   bson:"fee"`
	//IsDcc                     string `json:"isdcc"   bson:"isdcc"`
	IsCardPresent string `json:"iscardpresent"   bson:"iscardpresent"`
	//PercentageFee             string `json:"percentagefee"   bson:"percentagefee"`
	//PptFee                    string `json:"pptfee"   bson:"pptfee"`
	InterchangeFee float64 `json:"interchangefee"   bson:"interchangefee"`
	SchemeFee      float64 `json:"schemefee"   bson:"schemefee"`
	AcquirerFee    float64 `json:"acquirerfee"   bson:"acquirerfee"`
	//BatchID                   string `json:"batchid"   bson:"batchid"`
	CardNumberMasked string `json:"cardnumbermasked"   bson:"cardnumbermasked"`
	TransactionDate  string `json:"transactiondate"   bson:"transactiondate"`
	LocalCurrency    string `json:"localcurrency"   bson:"localcurrency"`
	ForeignCurrency  string `json:"foreigncurrency"   bson:"foreigncurrency"`
	//ForeignAmount             string `json:"foreignamount"   bson:"foreignamount"`
	TransactionType string `json:"transactiontype"   bson:"transactiontype"` // E1 = Sale, E2 = Refund, E3 = Cash Advance, E4 Cash Withrawal
	AuthCode        string `json:"authcode"   bson:"authcode"`
	//RetailerReference         string `json:"retailerreference"   bson:"retailerreference"`
	//PaymentID                 string `json:"paymentid"   bson:"paymentid"`
	//IsNetSettled              string `json:"isnetsettled"   bson:"isnetsettled"`
	MerchantName      string `json:"merchantname"   bson:"merchantname"`
	CardSchemeCode    string `json:"cardschemecode"   bson:"cardschemecode"`
	CardIssuedCountry string `json:"cardissuedcountry"   bson:"cardissuedcountry"`
	MarginAmount      string `json:"marginamount"   bson:"marginamount"`
	MarginRate        string `json:"marginrate"   bson:"marginrate"`
	//DataSourceCode            string `json:"datasourcecode"   bson:"datasourcecode"`
	TransactionCode string `json:"transactioncode"   bson:"transactioncode"`
	//AcquirerFeeExtended       string `json:"acquirerfeeextended"   bson:"acquirerfeeextended"`
	//SettlementType            string `json:"settlementtype"   bson:"settlementtype"`
	//CustomerReference         string `json:"customerreference"   bson:"customerreference"`
	IsDccCurrencyOffered string `json:"isdcccurrencyoffered"   bson:"isdcccurrencyoffered"`
	//IsFunded             string `json:"isfunded"   bson:"isfunded"`
	CountryCode string `json:"countrycode"   bson:"countrycode"`
	//DCCCurrencyOffered        string `json:"dcccurrencyoffered"   bson:"dcccurrencyoffered"`
	//CardScheme                string `json:"cardscheme"   bson:"cardscheme"`
	//Region                    string `json:"region"   bson:"region"`
	//Channel                   string `json:"channel"   bson:"channel"`
	//IsSecure                  string `json:"issecure"   bson:"issecure"`
	ChargeType string `json:"chargetype"   bson:"chargetype"`
	//ChargePercentage          string `json:"chargepercentage"   bson:"chargepercentage"`
	//ChargePpt                 string `json:"chargeppt"   bson:"chargeppt"`
	//Service                   string `json:"service"   bson:"service"`
	//MscType                   string `json:"msctype"   bson:"msctype"`
	//TransactionSecurityCode   string `json:"transactionsecuritycode"   bson:"transactionsecuritycode"` // E = Ecommerce M = MOTO Blank = No Value
	//Presentement              string `json:"presentement"  bson:"presentement"`                        // Blank/F First esentment R = Reesentment CBR/CBC2R/C2 = Chargeback RET = Retrieva
	IFAmount       string `json:"ifamount"  bson:"ifamount"`
	IFCurrencyCode string `json:"ifcurrencycode"  bson:"ifcurrencycode"`
	//IFRate                    string `json:"ifrate"  bson:"ifrate"`
	//IFPPT                     string `json:"ifppt"  bson:"ifppt"`
	IFParameterCode string `json:"ifparametercode"  bson:"ifparametercode"`
	//SFAmount                  string `json:"sfamount"  bson:"sfamount"`
	SFCurrencyCode string `json:"sfcurrencycode"  bson:"sfcurrencycode"`
	//SFRate                    string `json:"sfrate"  bson:"sfrate"`
	//SFPPT                     string `json:"sfppt"  bson:"sfppt"`
	//AFAmount                  string `json:"afamount"  bson:"afamount"`
	AFCurrencyCode string `json:"afcurrencycode"  bson:"fcurrencycode"`
	//AFRate                    string `json:"afrate"  bson:"afrate"`
	//AFPPT                     string `json:"afppt"  bson:"afppt"`
	AFParameterCode string `json:"afparametercode"  bson:"afparametercode"`
	AFCardScheme    string `json:"afcardscheme"  bson:"afcardscheme"`
	AFRegion        string `json:"afregion"  bson:"afregion"`
	AFProduct       string `json:"afproduct"  bson:"afproduct"`
	AFCardPresence  string `json:"afcardpresence"  bson:"afcardpresence"`

	// ID                         string `json:"id"   bson:"id"`
	// BatchID                    string `json:"batchid"   bson:"batchid"`
	// TerminalID                 string `json:"terminalid"   bson:"terminalid"`
	// PaymentOriginBankAccountID string `json:"paymentoriginbankaccountid"   bson:"paymentoriginbankaccountid"`
	// MerchantID                 string `json:"merchantid"   bson:"merchantid"`
	// HeadOfficeGroupID          string `json:"headofficegroupid"   bson:"headofficegroupid"`
	// BankChargeCode             string `json:"bankchargecode"   bson:"bankchargecode"`
	// DateOfPayment              string `json:"dateofpayment"   bson:"dateofpayment"`
	// Amount                     string `json:"amount"   bson:"amount"`
	// Narrative                  string `json:"narrative"   bson:"narrative"`
	// Fee                        string `json:"fee"   bson:"fee"`
	// IsNetSettled               string `json:"isnetsettled"   bson:"isnetsettled"`
	// TotalTransactionValue      string `json:"totaltransactionvalue"   bson:"totaltransactionvalue"`
	// Reprocessed                string `json:"reprocessed"   bson:"reprocessed"`
	// MscType                    string `json:"msctype"   bson:"msctype"`
	// NarrativeTypestring        string `json:"narrativetype"   bson:"narrativetype"`
}

type POption struct {
	UUID           string `json:"uuid" bson:"uuid"`
	Name           string `json:"name" bson:"name"`
	Scheme         string `json:"scheme" bson:"scheme"`
	Type           string `json:"type" bson:"type"`
	DomesticMSCPPT string `json:"domestic_m_s_c_p_p_t" bson:"domestic_m_s_c_p_p_t"`
	EEAMSCRate     string `json:"e_e_a_m_s_c_rate" bson:"e_e_a_m_s_c_rate"`
	EEAMSCPPT      string `json:"e_e_a_m_s_c_p_p_t" bson:"e_e_a_m_s_c_p_p_t"`
	MSCRate        string `json:"m_s_c_rate" bson:"m_s_c_rate"`
}

type PMerchant struct {
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	POptions    []POption `json:"poptions"`
	Address     []string  `json:"address"`
	Mid         []string  `json:"mid"`
	Tid         []string  `json:"tid"`
	Description string    `json:"description"`
}

type PAcquirer struct {
	UUID        string      `json:"uuid"`
	AquirerID   string      `json:"aquirerid"`
	Name        string      `json:"name"`
	PMerchants  []PMerchant `json:"pmerchants"`
	Description string      `json:"description"`
}
