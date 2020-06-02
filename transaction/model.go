/* *********************************************************************
			FHR	File Header
			UHR	User Header [Per Acquirer Model]
			MHR	Merchant Header
			TDR	Transaction Detail Record
			TAR	Transaction Additional Record [Contains Acquirer Model]
			EM*	EMV Supplemental Record
			ECR	Ecommerce Record
			TIF	Interchange Fee Record
			TSF	Scheme Fee Record
			TAF	Acquirer Fee Record
			BTR	Batch Trailer
			MTR	Merchant Trailer
			UTR	User Trailer
			UHR	User Header [Per Acquirer Model]
			….
			UTR	User Trailer
			FTR	File Trailer

***********************************************************************/

package transaction

type RunParams struct {
	ID          int      `json:"id, omitempty"`
	AcquirerID  string   `json:"acquirerid ,omitempty"`
	StartDate   string   `json:"startdate ,omitempty"`
	EndDate     string   `json:"enddate, omitempty"`
	CreatedDate string   `json:"createddate, omitempty"`
	CreatedTime string   `json:"createdtime, omitempty"`
	Acquirers   []string `json:"acquirers, omitempty"`
}

// Acquirer Model
type Acquirer struct {
	ID        int        `json:"id, omitempty"`
	Name      string     `json:"name"`
	StartDate string     `json:"startdate ,omitempty"`
	EndDate   string     `json:"enddate"`
	Merchants []Merchant `json: "merchants, omitempty"`
}

// Merchant is a field of a Acquirer.
// ID should be unique within the profile (at a minimum).
type Merchant struct {
	ID         string    `json:"id, omitempty"`
	MerchantId string    `json:"merchantid,omitempty"`
	Payments   []Payment `json:"transactions,omitempty"`
}

// Transaction Model
type Payment struct{}

// type Message struct {
// 	ID            int    `json:"id, omitempty"`
// 	CreatedAt     string `json:"created_at" bson:"creaed_at"`
// 	Level         string `json:"level"   bson:"level"`
// 	ServiceName   string `json:"service_name"   bson:"service_name"`
// 	CallingMethod string `json:"calling_method"   bsn:"calling_method"`
// 	Host          string `json:"host"   bson:"host"`
// 	Body          string `json:"body"   bson:"message"`
// 	Latency       string `json:"latency"   bson:"latency"`
// }

// Meta the transaction structure or Model
type Meta struct {
	CreatedDate                    string `json:"createddate" bson:"createddate"`
	CreatedTime                    string `json:"createdtime" bson:"createdtime"`
	FHRFileRecordIdentifier        string `json:"fhrfilerecordidentifier"   bson:"fhrfilerecordidentifier"`
	FHRRecordSeqNo                 string `json:"fhrrecordseqno"   bson:"fhrrecordseqno"`
	FHRFileCreateDate              string `json:"fhrfilecreatedate"   bson:"fhrfilecreatedate"`
	FHRFileCreateTime              string `json:"fhrfilecreatetime"   bson:"fhrfilecreatetime"`
	FHRFileCreateZone              string `json:"fhrfilecreatezone"   bson:"fhrfilecreatezone"`
	FHRFileCreateZoneOffset        string `json:"fhrfilecreatezoneoffset"   bson:"fhrfilecreatezoneoffset"`
	FHRClientCode                  string `json:"fhrclientcode"   bson:"fhrclientcode"`
	FHRFileVolumeIdentifier        string `json:"fhrfilevolumeidentifier"   bson:"fhrfilevolumeidentifier"`
	FHRFileIssueIdentifier         string `json:"fhrfileissueidentifier"   bson:"fhrfileissueidentifier"`
	FHRClientContact               string `json:"fhrclientcontact"   bson:"fhrclientcontact"`
	FHRClientPhone                 string `json:"fhrclientphone"   bson:"fhrclientphone"`
	FHRClientFax                   string `json:"fhrclientfax"   bson:"fhrclientfax"`
	FHRConfirmationTime            string `json:"fhrconfirmationtime"   bson:"fhrconfirmationtime"`
	FHRRecordCount                 string `json:"fhrrecordcount"   bson:"fhrrecordcount"`
	FHRTXFVersion                  string `json:"fhrtxfversion"   bson:"fhrtxfversion"`
	FHRTXFType                     string `json:"fhrtxftype"   bson:"fhrtxftype"`
	FHRRecordFiller                string `json:"fhrrecordfiller"   bson:"fhrrecordfiller"`
	FTRFileRecordIdentifier        string `json:"ftrfilerecordidentifier"   bson:"ftrfilerecordidentifier"`
	FTRRecordSeqNo                 string `json:"ftrrecordseqno"   bson:"ftrrecordseqno"`
	FTRFileUserRecords             string `json:"ftrfileuserrecords"   bson:"ftrfileuserrecords"`
	FTRFileMerchantRecords         string `json:"ftrfilemerchantrecords"   bson:"ftrfilemerchantrecords"`
	FTRFileCardTxns                string `json:"ftrfilecardtxns"   bson:"ftrfilecardtxns"`
	FTRFileHashNetValue            string `json:"ftrfilehashnetvalue"   bson:"ftrfilehashnetvalue"`
	FTRRecordFiller                string `json:"ftrrecordfiller"   bson:"ftrrecordfiller"`
	UHRRecordLabelIdentifier       string `json:"uhrrecordlabelidentifier"   bson:"uhrrecordlabelidentifier"`
	UHRRecordSeqNo                 string `json:"uhrrecordseqno"   bson:"uhrrecordseqno"`
	UHRUserCode                    string `json:"uhrusercode"   bson:"uhrusercode"`
	UHRUserBaseCurrencyCode        string `json:"uhruserbasecurrencycode"   bson:"uhruserbasecurrencycode"`
	UHRUserVolumeNumber            string `json:"uhruservolumenumber"   bson:"uhruservolumenumber"`
	UHRUserIssuerIdentifer         string `json:"uhruserissueridentifer"   bson:"uhruserissueridentifer"`
	UHRClientCode                  string `json:"uhrclientcode"   bson:"uhrclientcode"`
	UHRAcquirerID                  string `json:"uhracquirerid"   bson:"uhracquirerid"`
	UHRRecordFiller                string `json:"uhrrecordfiller"   bson:"uhrrecordfiller"`
	UTRRecordLabelIdentifer        string `json:"utrrecordlabelidentifer"   bson:"utrrecordlabelidentifer"`
	UTRRecordSeqNo                 string `json:"utrrecordseqno"   bson:"utrrecordseqno"`
	UTRUserHeaderRecordsMerchant   string `json:"utruserheaderrecordsmerchant"   bson:"utruserheaderrecordsmerchant"`
	UTRUserCardTxns                string `json:"utrusercardtxns"   bson:"utrusercardtxns"`
	UTRValueofDebitsBaseCurr       string `json:"utrvalueofdebitsbasecurr"   bson:"utrvalueofdebitsbasecurr"`
	UTRValueofCreditsBaseCurr      string `json:"utrvalueofcreditsbasecurr"   bson:"utrvalueofcreditsbasecurr"`
	UTRCountofDebits               string `json:"utrcountofdebits"   bson:"utrcountofdebits"`
	UTRCountofCredits              string `json:"utrcountofcredits"   bson:"utrcountofcredits"`
	UTRRecordFiller                string `json:"utrrecordfiller"   bson:"utrrecordfiller"`
	MHRRecordLabelIdentifier       string `json:"mhrrecordlabelidentifier"   bson:"mhrrecordlabelidentifier"`
	MHRRecordSeqNo                 string `json:"mhrrecordseqno"   bson:"mhrrecordseqno"`
	MHRMerchantCreateDate          string `json:"mhrmerchantcreatedate"   bson:"mhrmerchantcreatedate"`
	MHRMerchantIdentifier          string `json:"mhrmerchantidentifier"   bson:"mhrmerchantidentifier"`
	MHRMerchantBaseCurrCode        string `json:"mhrmerchantbasecurrcode"   bson:"mhrmerchantbasecurrcode"`
	MHRMerchantSubmissionId        string `json:"mhrmerchantsubmissionid"   bson:"mhrmerchantsubmissionid"`
	MHRMerchantClientRef           string `json:"mhrmerchantclientref"   bson:"mhrmerchantclientref"`
	MHRMerchantCategoryCode        string `json:"mhrmerchantcategorycode"   bson:"mhrmerchantcategorycode"`
	MHRMerchantOutletName          string `json:"mhrmerchantoutletname"   bson:"mhrmerchantoutletname"`
	MHRMerchantOutletStreet        string `json:"mhrmerchantoutletstreet"   bson:"mhrmerchantoutletstreet"`
	MHRMerchantTrace               string `json:"mhrmerchanttrace"   bson:"mhrmerchanttrace"`
	MHRMerchantDBAOverride         string `json:"mhrmerchantdbaoverride"   bson:"mhrmerchantdbaoverride"`
	MHRRecordFiller                string `json:"mhrrecordfiller"   bson:"mhrrecordfiller"`
	MTRRecordLabelIdentifier       string `json:"mtrrecordlabelidentifier"   bson:"mtrrecordlabelidentifier"`
	MTRRecordSeqNo                 string `json:"mtrrecordseqno"   bson:"mtrrecordseqno"`
	MTRMerchantBaseDebitTotal      string `json:"mtrmerchantbasedebittotal"   bson:"mtrmerchantbasedebittotal"`
	MTRMerchantBaseCreditTotal     string `json:"mtrmerchantbasecredittotal"   bson:"mtrmerchantbasecredittotal"`
	MTRMerchantTxnCount            string `json:"mtrmerchanttxncount"   bson:"mtrmerchanttxncount"`
	MTRMerchantCurrencyTotalHash   string `json:"mtrmerchantcurrencytotalhash"   bson:"mtrmerchantcurrencytotalhash"`
	MTRMerchantBatchCount          string `json:"mtrmerchantbatchcount"   bson:"mtrmerchantbatchcount"`
	MTRRecordFiller                string `json:"mtrrecordfiller"   bson:"mtrrecordfiller"`
	BTRRecordLabelIdentifier       string `json:"btrrecordlabelidentifier"   bson:"btrrecordlabelidentifier"`
	BTRRecordSeqNo                 string `json:"btrrecordseqno"   bson:"btrrecordseqno"`
	BTRBatchSettleBankCode         string `json:"btrbatchsettlebankcode"   bson:"btrbatchsettlebankcode"`
	BTRBatchSettleBankData         string `json:"btrbatchsettlebankdata"   bson:"btrbatchsettlebankdata"`
	BTRBatchSettleBankType         string `json:"btrbatchsettlebanktype"   bson:"btrbatchsettlebanktype"`
	BTRBatchSettlementCurrChoice   string `json:"btrbatchsettlementcurrchoice"   bson:"btrbatchsettlementcurrchoice"`
	BTRBatchSettleCode             string `json:"btrbatchsettlecode"   bson:"btrbatchsettlecode"`
	BTRBatchSettleBankLineRef      string `json:"btrbatchsettlebanklineref"   bson:"btrbatchsettlebanklineref"`
	BTRBatchSettlementCurrencyCode string `json:"btrbatchsettlementcurrencycode"   bson:"btrbatchsettlementcurrencycode"`
	BTRBatchSettlementTotalDebit   string `json:"btrbatchsettlementtotaldebit"   bson:"btrbatchsettlementtotaldebit"`
	BTRBatchSettlementTotalCredit  string `json:"btrbatchsettlementtotalcredit"   bson:"btrbatchsettlementtotalcredit"`
	BTRBatchSettlementCountDebit   string `json:"btrbatchsettlementcountdebit"   bson:"btrbatchsettlementcountdebit"`
	BTRBatchSettlementCountCredit  string `json:"btrbatchsettlementcountcredit"   bson:"btrbatchsettlementcountcredit"`
	BTRRecordFiller                string `json:"btrrecordfiller"   bson:"btrrecordfiller"`
	TDRRecordLabelIdentifier       string `json:"tdrrecordlabelidentifier"   bson:"tdrrecordlabelidentifier"`
	TDRRecordSeqNo                 string `json:"tdrrecordseqno"   bson:"tdrrecordseqno"`
	TDRTranCardPAN                 string `json:"tdrtrancardpan"   bson:"tdrtrancardpan"`
	TDRTranTypeCode                string `json:"tdrtrantypecode"   bson:"tdrtrantypecode"`
	TDRTranSubType                 string `json:"tdrtransubtype"   bson:"tdrtransubtype"`
	TDRTranTerminalID              string `json:"tdrtranterminalid"   bson:"tdrtranterminalid"`
	TDRTranTerminalBatchSeqNo      string `json:"tdrtranterminalbatchseqno"   bson:"tdrtranterminalbatchseqno"`
	TDRTranTerminalBatchTransSeqNo string `json:"tdrtranterminalbatchtransseqno"   bson:"tdrtranterminalbatchtransseqno"`
	TDRTranCardExpiryDate          string `json:"tdrtrancardexpirydate"   bson:"tdrtrancardexpirydate"`
	TDRTranCardServiceCode         string `json:"tdrtrancardservicecode"   bson:"tdrtrancardservicecode"`
	TDRTranMerchantAmount          string `json:"tdrtranmerchantamount"   bson:"tdrtranmerchantamount"`
	TDRTranMerchantCurrencyCode    string `json:"tdrtranmerchantcurrencycode"   bson:"tdrtranmerchantcurrencycode"`
	TDRTranCardholderAmount        string `json:"tdrtrancardholderamount"   bson:"tdrtrancardholderamount"`
	TDRTranCardholderCurrency      string `json:"tdrtrancardholdercurrency"   bson:"tdrtrancardholdercurrency"`
	TDRTranConversionRate          string `json:"tdrtranconversionrate"   bson:"tdrtranconversionrate"`
	TDRTranDate                    string `json:"tdrtrandate"   bson:"tdrtrandate"`
	TDRTranTime                    string `json:"tdrtrantime"   bson:"tdrtrantime"`
	TDRTranAuthMethod              string `json:"tdrtranauthmethod"   bson:"tdrtranauthmethod"`
	TDRTranAuthCode                string `json:"tdrtranauthcode"   bson:"tdrtranauthcode"`
	TDRTranStatementMessage1       string `json:"tdrtranstatementmessage1"   bson:"tdrtranstatementmessage1"`
	TDRTranStatementMessage2       string `json:"tdrtranstatementmessage2"   bson:"tdrtranstatementmessage2"`
	TDRTranCardIssueDigit1         string `json:"tdrtrancardissuedigit1"   bson:"tdrtrancardissuedigit1"`
	TDRTerminalClass               string `json:"tdrterminalclass"   bson:"tdrterminalclass"`
	TDRTranCardIssueDigit2         string `json:"tdrtrancardissuedigit2"   bson:"tdrtrancardissuedigit2"`
	TDRTranEntryMode               string `json:"tdrtranentrymode"   bson:"tdrtranentrymode"`
	TDRTranCashCurr                string `json:"tdrtrancashcurr"   bson:"tdrtrancashcurr"`
	TDRTranCashAmount              string `json:"tdrtrancashamount"   bson:"tdrtrancashamount"`
	TDRTranCountryCode             string `json:"tdrtrancountrycode"   bson:"tdrtrancountrycode"`
	TDRTranSupplement              string `json:"tdrtransupplement"   bson:"tdrtransupplement"`
	TDRTerminalCRTEntry            string `json:"tdrterminalcrtentry"   bson:"tdrterminalcrtentry"`
	TDRTranEMV                     string `json:"tdrtranemv"   bson:"tdrtranemv"`
	TDRTranCVM                     string `json:"tdrtrancvm"   bson:"tdrtrancvm"`
	TDRTranSecurity                string `json:"tdrtransecurity"   bson:"tdrtransecurity"`
	TDRTranAuthMerchantID          string `json:"tdrtranauthmerchantid"   bson:"tdrtranauthmerchantid"`
	TDRTranReasonCode              string `json:"tdrtranreasoncode"   bson:"tdrtranreasoncode"`
	TDRTranDisputeCaseID           string `json:"tdrtrandisputecaseid"   bson:"tdrtrandisputecaseid"`
	TDRTranChargebackRefNum        string `json:"tdrtranchargebackrefnum"   bson:"tdrtranchargebackrefnum"`
	TDRTranPresentment             string `json:"tdrtranpresentment"   bson:"tdrtranpresentment"`
	TDRTranReversalIndicator       string `json:"tdrtranreversalindicator"   bson:"tdrtranreversalindicator"`
	TDRTranPartialIndicator        string `json:"tdrtranpartialindicator"   bson:"tdrtranpartialindicator"`
	TDRTranTaxRecredit             string `json:"tdrtrantaxrecredit"   bson:"tdrtrantaxrecredit"`
	TDRTranIndMarginvalue          string `json:"tdrtranindmarginvalue"   bson:"tdrtranindmarginvalue"`
	TDRTranIndMarginpercent        string `json:"tdrtranindmarginpercent"   bson:"tdrtranindmarginpercent"`
	TDRRecordFiller                string `json:"tdrrecordfiller"   bson:"tdrrecordfiller"`
	TARRecordLabelIdentifier       string `json:"tarrecordlabelidentifier"   bson:"tarrecordlabelidentifier"`
	TARRecordSeqNo                 string `json:"tarrecordseqno"   bson:"tarrecordseqno"`
	TARTranID                      string `json:"tartranid"   bson:"tartranid"`
	TARValidationCode              string `json:"tarvalidationcode"   bson:"tarvalidationcode"`
	TARUniqueID                    string `json:"taruniqueid"   bson:"taruniqueid"`
	TARAcqRefNum                   string `json:"taracqrefnum"   bson:"taracqrefnum"`
	TARRetRefNum                   string `json:"tarretrefnum"   bson:"tarretrefnum"`
	TARMASPostingCurrentStatus     string `json:"tarmaspostingcurrentstatus"   bson:"tarmaspostingcurrentstatus"`
	TARMASResponseTextStatus       string `json:"tarmasresponsetextstatus"   bson:"tarmasresponsetextstatus"`
	TARFinancialInstitution        string `json:"tarfinancialinstitution"   bson:"tarfinancialinstitution"`
	TARGatewayTransactionID        string `json:"targatewaytransactionid"   bson:"targatewaytransactionid"`
	TARIntAMPSTranID               string `json:"tarintampstranid"   bson:"tarintampstranid"`
	TARAcqBIN                      string `json:"taracqbin"   bson:"taracqbin"`
	TARCardIssCntry                string `json:"tarcardisscntry"   bson:"tarcardisscntry"`
	AcquirerID                     string `json:"acquirerid"   bson:"acquirerid"`
	TARAcqPlanetCode               string `json:"taracqplanetcode"   bson:"taracqplanetcode"`
	//TARFiller                      string `json:"tarfiller"   bson:"tarfiller"`
	TARCardtype                  string `json:"tarcardtype"   bson:"tarcardtype"`
	TARCardIssCurrency           string `json:"tarcardisscurrency"   bson:"tarcardisscurrency"`
	TAREligibleCurrency          string `json:"tareligiblecurrency"   bson:"tareligiblecurrency"`
	TAROriginalUniqueID          string `json:"taroriginaluniqueid"   bson:"taroriginaluniqueid"`
	TARFiller                    string `json:"tarfiller"   bson:"tarfiller"`
	EM1RecordLabelIdentifier     string `json:"em1recordlabelidentifier"   bson:"em1recordlabelidentifier"`
	EM1RecordSeqNo               string `json:"em1recordseqno"   bson:"em1recordseqno"`
	EM1EMVPANSequenceNumber      string `json:"em1emvpansequencenumber"   bson:"em1emvpansequencenumber"`
	EM1EMVAuthRespCode           string `json:"em1emvauthrespcode"   bson:"em1emvauthrespcode"`
	EM1EMVTransAmt               string `json:"em1emvtransamt"   bson:"em1emvtransamt"`
	EM1EMVTransType              string `json:"em1emvtranstype"   bson:"em1emvtranstype"`
	EM1EMVTransDate              string `json:"em1emvtransdate"   bson:"em1emvtransdate"`
	EM1EMVTransCurr              string `json:"em1emvtranscurr"   bson:"em1emvtranscurr"`
	EM1EMVTermCountry            string `json:"em1emvtermcountry"   bson:"em1emvtermcountry"`
	EM1EMVTransCryptogram        string `json:"em1emvtranscryptogram"   bson:"em1emvtranscryptogram"`
	EM1EMVAIP                    string `json:"em1emvaip"   bson:"em1emvaip"`
	EM1EMVATC                    string `json:"em1emvatc"   bson:"em1emvatc"`
	EM1EMVUnpredictNumber        string `json:"em1emvunpredictnumber"   bson:"em1emvunpredictnumber"`
	EM1EMVTVR                    string `json:"em1emvtvr"   bson:"em1emvtvr"`
	EM1EMVIAD                    string `json:"em1emviad"   bson:"em1emviad"`
	EM1EMVAppUsageControl        string `json:"em1emvappusagecontrol"   bson:"em1emvappusagecontrol"`
	EM1EMVCryptoInfoData         string `json:"em1emvcryptoinfodata"   bson:"em1emvcryptoinfodata"`
	EM1EMVCVMR                   string `json:"em1emvcvmr"   bson:"em1emvcvmr"`
	EM1EMVFileName               string `json:"em1emvfilename"   bson:"em1emvfilename"`
	EM1Filler                    string `json:"em1filler"   bson:"em1filler"`
	EM2RecordLabelIdentifier     string `json:"em2recordlabelidentifier"   bson:"em2recordlabelidentifier"`
	EM2RecordSeqNo               string `json:"em2recordseqno"   bson:"em2recordseqno"`
	EM2EMVAID                    string `json:"em2emvaid"   bson:"em2emvaid"`
	EM2EMVAppVersion             string `json:"em2emvappversion"   bson:"em2emvappversion"`
	EM2EMVTransStatusIinfo       string `json:"em2emvtransstatusiinfo"   bson:"em2emvtransstatusiinfo"`
	EM2EMVTermType               string `json:"em2emvtermtype"   bson:"em2emvtermtype"`
	EM2EMVTermCapabilities       string `json:"em2emvtermcapabilities"   bson:"em2emvtermcapabilities"`
	EM2EMVPOSEntryMode           string `json:"em2emvposentrymode"   bson:"em2emvposentrymode"`
	EM2EMVFillerOCD              string `json:"em2emvfillerocd"   bson:"em2emvfillerocd"`
	EM2EMVIACdefault             string `json:"em2emviacdefault"   bson:"em2emviacdefault"`
	EM2EMVIACdenial              string `json:"em2emviacdenial"   bson:"em2emviacdenial"`
	EM2EMVIAConline              string `json:"em2emviaconline"   bson:"em2emviaconline"`
	EM2EMVCardIssuerCntry        string `json:"em2emvcardissuercntry"   bson:"em2emvcardissuercntry"`
	EM2EMVIFDSerialNo            string `json:"em2emvifdserialno"   bson:"em2emvifdserialno"`
	EM2EMVTransCatCode           string `json:"em2emvtranscatcode"   bson:"em2emvtranscatcode"`
	EM2EMVTermAppVersion         string `json:"em2emvtermappversion"   bson:"em2emvtermappversion"`
	EM2EMVTransSeqCounter        string `json:"em2emvtransseqcounter"   bson:"em2emvtransseqcounter"`
	EM2EMVIssScriptRespartII     string `json:"em2emvissscriptrespartii"   bson:"em2emvissscriptrespartii"`
	EM2EMV3Follows               string `json:"em2emv3follows"   bson:"em2emv3follows"`
	EM2Filler                    string `json:"em2filler"   bson:"em2filler"`
	EM3RecordLabelIdentifier     string `json:"em3recordlabelidentifier"   bson:"em3recordlabelidentifier"`
	EM3RecordSeqNo               string `json:"em3recordseqno"   bson:"em3recordseqno"`
	EM3EMVCtlsFormFactor         string `json:"em3emvctlsformfactor"   bson:"em3emvctlsformfactor"`
	EM3EMVCtlsDiscrData          string `json:"em3emvctlsdiscrdata"   bson:"em3emvctlsdiscrdata"`
	EM3Filler                    string `json:"em3filler"   bson:"em3filler"`
	TIFRecordLabelIdentifier     string `json:"tifrecordlabelidentifier"   bson:"tifrecordlabelidentifier"`
	TIFRecordSeqNo               string `json:"tifrecordseqno"   bson:"tifrecordseqno"`
	TIFIFAmount                  string `json:"tififamount"   bson:"tififamount"`
	TIFIFCurrencyCode            string `json:"tififcurrencycode"   bson:"tififcurrencycode"`
	TIFIFRate                    string `json:"tififrate"   bson:"tififrate"`
	TIFIFPPT                     string `json:"tififppt"   bson:"tififppt"`
	TIFIFParameterCode           string `json:"tififparametercode"   bson:"tififparametercode"`
	TIFIFIndicator1              string `json:"tififindicator1"   bson:"tififindicator1"`
	TIFFiller                    string `json:"tiffiller"   bson:"tiffiller"`
	TSFRecordLabelIdentifier     string `json:"tsfrecordlabelidentifier"   bson:"tsfrecordlabelidentifier"`
	TSFRecordSeqNo               string `json:"tsfrecordseqno"   bson:"tsfrecordseqno"`
	TSFSFAmount                  string `json:"tsfsfamount"   bson:"tsfsfamount"`
	TSFSFCurrencyCode            string `json:"tsfsfcurrencycode"   bson:"tsfsfcurrencycode"`
	TSFSFRate                    string `json:"tsfsfrate"   bson:"tsfsfrate"`
	TSFSFPPT                     string `json:"tsfsfppt"   bson:"tsfsfppt"`
	TSFSFParameterCode           string `json:"tsfsfparametercode"   bson:"tsfsfparametercode"`
	TSFFiller                    string `json:"tsffiller"   bson:"tsffiller"`
	TAFRecordLabelIdentifier     string `json:"tafrecordlabelidentifier"   bson:"tafrecordlabelidentifier"`
	TAFRecordSeqNo               string `json:"tafrecordseqno"   bson:"tafrecordseqno"`
	TAFAFAmount                  string `json:"tafafamount"   bson:"tafafamount"`
	TAFAFCurrencyCode            string `json:"tafafcurrencycode"   bson:"tafafcurrencycode"`
	TAFAFRate                    string `json:"tafafrate"   bson:"tafafrate"`
	TAFAFPPT                     string `json:"tafafppt"   bson:"tafafppt"`
	TAFAFParameterCode           string `json:"tafafparametercode"   bson:"tafafparametercode"`
	TAFAFCardScheme              string `json:"tafafcardscheme"   bson:"tafafcardscheme"`
	TAFAFRegion                  string `json:"tafafregion"   bson:"tafafregion"`
	TAFAFProduct                 string `json:"tafafproduct"   bson:"tafafproduct"`
	TAFAFCardPresence            string `json:"tafafcardpresence"   bson:"tafafcardpresence"`
	TAFAFType                    string `json:"tafaftype"   bson:"tafaftype"`
	TAFAFStatus                  string `json:"tafafstatus"   bson:"tafafstatus"`
	TAFAFAdditional1             string `json:"tafafadditional1"   bson:"tafafadditional1"`
	TAFAFAdditional2             string `json:"tafafadditional2"   bson:"tafafadditional2"`
	TAFAFAdditional3             string `json:"tafafadditional3"   bson:"tafafadditional3"`
	TAFFiller                    string `json:"taffiller"   bson:"taffiller"`
	AL1RecordLabelIdentifier     string `json:"al1recordlabelidentifier"   bson:"al1recordlabelidentifier"`
	AL1RecordSeqNo               string `json:"al1recordseqno"   bson:"al1recordseqno"`
	AL1TicketNumber              string `json:"al1ticketnumber"   bson:"al1ticketnumber"`
	AL1CarrierName               string `json:"al1carriername"   bson:"al1carriername"`
	AL1TravelAgencyCode          string `json:"al1travelagencycode"   bson:"al1travelagencycode"`
	AL1TicketAgencyName          string `json:"al1ticketagencyname"   bson:"al1ticketagencyname"`
	AL1AirlinePlanNumber         string `json:"al1airlineplannumber"   bson:"al1airlineplannumber"`
	AL1AirlineInvoiceNumber      string `json:"al1airlineinvoicenumber"   bson:"al1airlineinvoicenumber"`
	AL1OriginalCurrency          string `json:"al1originalcurrency"   bson:"al1originalcurrency"`
	AL1PassengerName             string `json:"al1passengername"   bson:"al1passengername"`
	AL1CustomerReferenceNumber   string `json:"al1customerreferencenumber"   bson:"al1customerreferencenumber"`
	AL1OriginalTransactionAmount string `json:"al1originaltransactionamount"   bson:"al1originaltransactionamount"`
	AL1TicketIssuerAddress       string `json:"al1ticketissueraddress"   bson:"al1ticketissueraddress"`
	AL1TripLeg1DepartAirport     string `json:"al1tripleg1departairport"   bson:"al1tripleg1departairport"`
	AL1TripLeg1CarrierCode       string `json:"al1tripleg1carriercode"   bson:"al1tripleg1carriercode"`
	AL1TripLeg1FareBasisCode     string `json:"al1tripleg1farebasiscode"   bson:"al1tripleg1farebasiscode"`
	AL1TripLeg1ClassOfTravel     string `json:"al1tripleg1classoftravel"   bson:"al1tripleg1classoftravel"`
	AL1TripLeg1StopOverCode      string `json:"al1tripleg1stopovercode"   bson:"al1tripleg1stopovercode"`
	AL1TripLeg1DestinationCode   string `json:"al1tripleg1destinationcode"   bson:"al1tripleg1destinationcode"`
	AL1TripLeg1DateOfTravel      string `json:"al1tripleg1dateoftravel"   bson:"al1tripleg1dateoftravel"`
	AL1TripLeg1DepartTax         string `json:"al1tripleg1departtax"   bson:"al1tripleg1departtax"`
	AL1Filler                    string `json:"al1filler"   bson:"al1filler"`
	AL2RecordLabelIdentifier     string `json:"al2recordlabelidentifier"   bson:"al2recordlabelidentifier"`
	AL2RecordSeqNo               string `json:"al2recordseqno"   bson:"al2recordseqno"`
	//AL2TripLeg2DepartAirport       string `json:"al2tripleg2departairport"   bson:"al2tripleg2departairport"`
	AL2TripLeg2CarrierCode       string `json:"al2tripleg2carriercode"   bson:"al2tripleg2carriercode"`
	AL2TripLeg2FareBasisCode     string `json:"al2tripleg2farebasiscode"   bson:"al2tripleg2farebasiscode"`
	AL2TripLeg2ClassOfTravel     string `json:"al2tripleg2classoftravel"   bson:"al2tripleg2classoftravel"`
	AL2TripLeg2StopOverCode      string `json:"al2tripleg2stopovercode"   bson:"al2tripleg2stopovercode"`
	AL2TripLeg2DestinationCode   string `json:"al2tripleg2destinationcode"   bson:"al2tripleg2destinationcode"`
	AL2TripLeg2DateOfTravel      string `json:"al2tripleg2dateoftravel"   bson:"al2tripleg2dateoftravel"`
	AL2TripLeg2DepartTax         string `json:"al2tripleg2departtax"   bson:"al2tripleg2departtax"`
	AL2TripLeg2DepartAirport     string `json:"al2tripleg2departairport"   bson:"al2tripleg2departairport"`
	AL2TripLeg3CarrierCode       string `json:"al2tripleg3carriercode"   bson:"al2tripleg3carriercode"`
	AL2TripLeg3FareBasisCode     string `json:"al2tripleg3farebasiscode"   bson:"al2tripleg3farebasiscode"`
	AL2TripLeg3ClassOfTravel     string `json:"al2tripleg3classoftravel"   bson:"al2tripleg3classoftravel"`
	AL2TripLeg3StopOverCode      string `json:"al2tripleg3stopovercode"   bson:"al2tripleg3stopovercode"`
	AL2TripLeg3DestinationCode   string `json:"al2tripleg3destinationcode"   bson:"al2tripleg3destinationcode"`
	AL2TripLeg3DateOfTravel      string `json:"al2tripleg3dateoftravel"   bson:"al2tripleg3dateoftravel"`
	AL2TripLeg3DepartTax         string `json:"al2tripleg3departtax"   bson:"al2tripleg3departtax"`
	AL2TripLeg4DepartAirport     string `json:"al2tripleg4departairport"   bson:"al2tripleg4departairport"`
	AL2TripLeg4CarrierCode       string `json:"al2tripleg4carriercode"   bson:"al2tripleg4carriercode"`
	AL2TRIPLEG4FAREBASISCODE     string `json:"al2tripleg4farebasiscode"   bson:"al2tripleg4farebasiscode"`
	AL2TRIPLEG4CLASSOFTRAVEL     string `json:"al2tripleg4classoftravel"   bson:"al2tripleg4classoftravel"`
	AL2TRIPLEG4STOPOVERCODE      string `json:"al2tripleg4stopovercode"   bson:"al2tripleg4stopovercode"`
	AL2TRIPLEG4DESTINATIONCODE   string `json:"al2tripleg4destinationcode"   bson:"al2tripleg4destinationcode"`
	AL2TRIPLEG4DATEOFTRAVEL      string `json:"al2tripleg4dateoftravel"   bson:"al2tripleg4dateoftravel"`
	AL2TRIPLEG4DEPARTTAX         string `json:"al2tripleg4departtax"   bson:"al2tripleg4departtax"`
	AL2TRIPLEG5DEPARTAIRPORT     string `json:"al2tripleg5departairport"   bson:"al2tripleg5departairport"`
	AL2TRIPLEG5CARRIERCODE       string `json:"al2tripleg5carriercode"   bson:"al2tripleg5carriercode"`
	AL2TRIPLEG5FAREBASISCODE     string `json:"al2tripleg5farebasiscode"   bson:"al2tripleg5farebasiscode"`
	AL2TRIPLEG5CLASSOFTRAVEL     string `json:"al2tripleg5classoftravel"   bson:"al2tripleg5classoftravel"`
	AL2TRIPLEG5STOPOVERCODE      string `json:"al2tripleg5stopovercode"   bson:"al2tripleg5stopovercode"`
	AL2TRIPLEG5DESTINATIONCODE   string `json:"al2tripleg5destinationcode"   bson:"al2tripleg5destinationcode"`
	AL2TRIPLEG5DATEOFTRAVEL      string `json:"al2tripleg5dateoftravel"   bson:"al2tripleg5dateoftravel"`
	AL2TRIPLEG5DEPARTTAX         string `json:"al2tripleg5departtax"   bson:"al2tripleg5departtax"`
	AL2TRIPLEG6DEPARTAIRPORT     string `json:"al2tripleg6departairport"   bson:"al2tripleg6departairport"`
	AL2TRIPLEG6CARRIERCODE       string `json:"al2tripleg6carriercode"   bson:"al2tripleg6carriercode"`
	AL2TRIPLEG6FAREBASISCODE     string `json:"al2tripleg6farebasiscode"   bson:"al2tripleg6farebasiscode"`
	AL2TRIPLEG65CLASSOFTRAVEL    string `json:"al2tripleg65classoftravel"   bson:"al2tripleg65classoftravel"`
	AL2TRIPLEG6STOPOVERCODE      string `json:"al2tripleg6stopovercode"   bson:"al2tripleg6stopovercode"`
	AL2TRIPLEG6DESTINATIONCODE   string `json:"al2tripleg6destinationcode"   bson:"al2tripleg6destinationcode"`
	AL2TRIPLEG6DATEOFTRAVEL      string `json:"al2tripleg6dateoftravel"   bson:"al2tripleg6dateoftravel"`
	AL2TRIPLEG6DEPARTTAX         string `json:"al2tripleg6departtax"   bson:"al2tripleg6departtax"`
	AL2Filler                    string `json:"al2filler"   bson:"al2filler"`
	ECRRecordLabelIdentifier     string `json:"ecrrecordlabelidentifier"   bson:"ecrrecordlabelidentifier"`
	ECRRecordSeqNo               string `json:"ecrrecordseqno"   bson:"ecrrecordseqno"`
	ECREcommerceTranSecurity     string `json:"ecrecommercetransecurity"   bson:"ecrecommercetransecurity"`
	ECRAVSResponseCode           string `json:"ecravsresponsecode"   bson:"ecravsresponsecode"`
	ECRCVV2ResultCode            string `json:"ecrcvv2resultcode"   bson:"ecrcvv2resultcode"`
	ECRECSIndicator              string `json:"ecrecsindicator"   bson:"ecrecsindicator"`
	ECRAdditionalUCAFData        string `json:"ecradditionalucafdata"   bson:"ecradditionalucafdata"`
	ECRMerchantURL               string `json:"ecrmerchanturl"   bson:"ecrmerchanturl"`
	ECRMerchantTelephoneNumber   string `json:"ecrmerchanttelephonenumber"   bson:"ecrmerchanttelephonenumber"`
	ECREcommerceGoodsIndicator   string `json:"ecrecommercegoodsindicator"   bson:"ecrecommercegoodsindicator"`
	ECRDateTransactionAuthorised string `json:"ecrdatetransactionauthorised"   bson:"ecrdatetransactionauthorised"`
	ECRDateTransactionCaptured   string `json:"ecrdatetransactioncaptured"   bson:"ecrdatetransactioncaptured"`
	ECRRecordFiller              string `json:"ecrrecordfiller"   bson:"ecrrecordfiller"`
}
