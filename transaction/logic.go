package transaction

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/mmanjoura/pppr/configuration-svc/decoder"
	"github.com/mmanjoura/pppr/logging"

	loggingpb "github.com/mmanjoura/pppr/logging/proto"
	paymentpb "github.com/mmanjoura/pppr/payment/proto"

	errs "github.com/pkg/errors"
	"google.golang.org/grpc"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrTransactionInvalid = errors.New("Transaction Invalid")

	loggingClient loggingpb.LogServiceClient
	paymentClient paymentpb.PaymentServiceClient
	// reportClient  reportpb.ReportServiceClient

	requestCtx  context.Context
	requestOpts grpc.DialOption
)

type service struct {
	batchFileRepo Repository
}

func NewService(batchFileRepo Repository) Service {
	return &service{
		batchFileRepo,
	}
}

func (s *service) Save(meta *Meta, collection string) (map[string]string, error) {

	var m = Meta{}

	if err := validate.Validate(meta); err != nil {
		return nil, errs.Wrap(ErrTransactionInvalid, "service.Transaction.Save")
	}

	trxConfig, err := decoder.Decode(decoder.TransactionService)
	if err != nil {
		return nil, err
	}

	dir, err := os.Open(trxConfig.SourcePath)
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	dateTime := make(map[string]string)

	for _, file := range files {

		f, err := os.Open(trxConfig.SourcePath + file.Name())
		if err != nil {
			return nil, err
		}
		trxs, err := getBatchTransactions(m, trxConfig.SourcePath, f)
		if err != nil {
			return nil, err
		}

		// Iterate through the tfx and store in DB

		for _, trx := range trxs {

			m = trx
			_, err := s.batchFileRepo.Save(&m, collection)
			if err != nil {
				return nil, err
			}
			dateTime[trx.CreatedTime] = trx.CreatedDate
		}
		f.Close()
		err = moveFile(f.Name(), trxConfig.ArchivePath+file.Name())
		if err != nil {
			return nil, err
		}
	}

	return dateTime, nil
}

// SCAN TXF file and build Transaction Meta Model
func getBatchTransactions(meta Meta, sourcePath string, file io.Reader) ([]Meta, error) {

	scanner := bufio.NewScanner(file)

	batchFileTxrList := []Meta{}
	meta.CreatedDate = time.Now().Format("2006-01-02")
	meta.CreatedTime = time.Now().Format("2006.01.02 15:04:05")[10:]
	for scanner.Scan() {
		line := scanner.Text()
		lineIdentifier := strings.TrimSpace(line[0:3])

		if lineIdentifier == "FHR" {
			meta.FHRFileRecordIdentifier = strings.TrimSpace(line[0:3])
			meta.FHRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.FHRFileCreateDate = strings.TrimSpace(line[10:18])
			meta.FHRFileCreateTime = strings.TrimSpace(line[18:26])
			meta.FHRFileCreateZone = strings.TrimSpace(line[26:30])
			meta.FHRFileCreateZoneOffset = strings.TrimSpace(line[30:36])
			meta.FHRClientCode = strings.TrimSpace(line[36:46])
			meta.FHRFileVolumeIdentifier = strings.TrimSpace(line[46:52])
			meta.FHRFileIssueIdentifier = strings.TrimSpace(line[52:54])
			meta.FHRClientContact = strings.TrimSpace(line[54:79])
			meta.FHRClientPhone = strings.TrimSpace(line[79:99])
			meta.FHRClientFax = strings.TrimSpace(line[99:119])
			meta.FHRConfirmationTime = strings.TrimSpace(line[119:122])
			meta.FHRRecordCount = strings.TrimSpace(line[122:129])
			meta.FHRTXFVersion = strings.TrimSpace(line[129:135])
			meta.FHRTXFType = strings.TrimSpace(line[135:138])
			meta.FHRRecordFiller = strings.TrimSpace(line[138:300])

			continue
		}
		if lineIdentifier == "UHR" {
			meta.UHRRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.UHRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.UHRUserCode = strings.TrimSpace(line[10:15])
			meta.UHRUserBaseCurrencyCode = strings.TrimSpace(line[15:18])
			meta.UHRUserVolumeNumber = strings.TrimSpace(line[18:24])
			meta.UHRUserIssuerIdentifer = strings.TrimSpace(line[24:26])
			meta.UHRClientCode = strings.TrimSpace(line[26:36])
			meta.UHRAcquirerID = strings.TrimSpace(line[36:72])
			meta.UHRRecordFiller = strings.TrimSpace(line[72:300])
			continue
		}
		if lineIdentifier == "MHR" {
			meta.MHRRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.MHRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.MHRMerchantCreateDate = strings.TrimSpace(line[10:16])
			meta.MHRMerchantIdentifier = strings.TrimSpace(line[16:31])
			meta.MHRMerchantBaseCurrCode = strings.TrimSpace(line[31:34])
			meta.MHRMerchantSubmissionId = strings.TrimSpace(line[34:40])
			meta.MHRMerchantClientRef = strings.TrimSpace(line[40:65])
			meta.MHRMerchantCategoryCode = strings.TrimSpace(line[65:69])
			meta.MHRMerchantOutletName = strings.TrimSpace(line[69:99])
			meta.MHRMerchantOutletStreet = strings.TrimSpace(line[99:149])
			meta.MHRMerchantTrace = strings.TrimSpace(line[149:157])
			meta.MHRMerchantDBAOverride = strings.TrimSpace(line[157:158])
			meta.MHRRecordFiller = strings.TrimSpace(line[158:300])

			continue
		}
		if lineIdentifier == "TDR" {
			meta.TDRRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.TDRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.TDRTranCardPAN = strings.TrimSpace(line[10:29])
			meta.TDRTranTypeCode = strings.TrimSpace(line[29:31])
			meta.TDRTranSubType = strings.TrimSpace(line[31:34])
			meta.TDRTranTerminalID = strings.TrimSpace(line[34:42])
			meta.TDRTranTerminalBatchSeqNo = strings.TrimSpace(line[42:48])
			meta.TDRTranTerminalBatchTransSeqNo = strings.TrimSpace(line[48:55])
			meta.TDRTranCardExpiryDate = strings.TrimSpace(line[55:59])
			meta.TDRTranCardServiceCode = strings.TrimSpace(line[59:65])
			meta.TDRTranMerchantAmount = strings.TrimSpace(line[65:77])
			meta.TDRTranMerchantCurrencyCode = strings.TrimSpace(line[77:79])
			meta.TDRTranCardholderAmount = strings.TrimSpace(line[80:92])
			meta.TDRTranCardholderCurrency = strings.TrimSpace(line[92:95])
			meta.TDRTranConversionRate = strings.TrimSpace(line[95:107])
			meta.TDRTranDate = strings.TrimSpace(line[107:113])
			meta.TDRTranTime = strings.TrimSpace(line[113:119])
			meta.TDRTranAuthMethod = strings.TrimSpace(line[119:121])
			meta.TDRTranAuthCode = strings.TrimSpace(line[121:127])
			meta.TDRTranStatementMessage1 = strings.TrimSpace(line[127:152])
			meta.TDRTranStatementMessage2 = strings.TrimSpace(line[152:165])
			meta.TDRTranCardIssueDigit1 = strings.TrimSpace(line[165:166])
			meta.TDRTerminalClass = strings.TrimSpace(line[166:167])
			meta.TDRTranCardIssueDigit2 = strings.TrimSpace(line[167:168])
			meta.TDRTranEntryMode = strings.TrimSpace(line[168:170])
			meta.TDRTranCashCurr = strings.TrimSpace(line[170:173])
			meta.TDRTranCashAmount = strings.TrimSpace(line[173:185])
			meta.TDRTranCountryCode = strings.TrimSpace(line[185:188])
			meta.TDRTranSupplement = strings.TrimSpace(line[188:189])
			meta.TDRTerminalCRTEntry = strings.TrimSpace(line[189:195])
			meta.TDRTranEMV = strings.TrimSpace(line[195:196])
			meta.TDRTranCVM = strings.TrimSpace(line[196:197])
			meta.TDRTranSecurity = strings.TrimSpace(line[197:198])
			meta.TDRTranAuthMerchantID = strings.TrimSpace(line[198:213])
			meta.TDRTranReasonCode = strings.TrimSpace(line[213:219])
			meta.TDRTranDisputeCaseID = strings.TrimSpace(line[219:239])
			meta.TDRTranChargebackRefNum = strings.TrimSpace(line[239:251])
			meta.TDRTranPresentment = strings.TrimSpace(line[251:254])
			meta.TDRTranReversalIndicator = strings.TrimSpace(line[254:255])
			meta.TDRTranPartialIndicator = strings.TrimSpace(line[255:256])
			meta.TDRTranTaxRecredit = strings.TrimSpace(line[256:257])
			meta.TDRTranIndMarginvalue = strings.TrimSpace(line[257:269])
			meta.TDRTranIndMarginpercent = strings.TrimSpace(line[269:277])
			meta.TDRRecordFiller = strings.TrimSpace(line[277:300])
			continue
		}
		if lineIdentifier == "TAR" {
			meta.TARRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.TARRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.TARTranID = strings.TrimSpace(line[10:25])
			meta.TARValidationCode = strings.TrimSpace(line[25:29])
			meta.TARUniqueID = strings.TrimSpace(line[29:39])
			meta.TARAcqRefNum = strings.TrimSpace(line[39:62])
			meta.TARRetRefNum = strings.TrimSpace(line[62:74])
			meta.TARMASPostingCurrentStatus = strings.TrimSpace(line[74:75])
			meta.TARMASResponseTextStatus = strings.TrimSpace(line[75:125])
			meta.TARFinancialInstitution = strings.TrimSpace(line[125:128])
			meta.TARGatewayTransactionID = strings.TrimSpace(line[128:164])
			meta.TARIntAMPSTranID = strings.TrimSpace(line[164:200])
			meta.TARAcqBIN = strings.TrimSpace(line[200:210])
			meta.TARCardIssCntry = strings.TrimSpace(line[210:213])
			meta.AcquirerID = strings.TrimSpace(line[213:249])
			meta.TARAcqPlanetCode = strings.TrimSpace(line[249:252])
			//meta.TARFiller = strings.TrimSpace(line[252:255])
			meta.TARCardtype = strings.TrimSpace(line[255:258])
			meta.TARCardIssCurrency = strings.TrimSpace(line[258:261])
			meta.TAREligibleCurrency = strings.TrimSpace(line[261:262])
			meta.TAROriginalUniqueID = strings.TrimSpace(line[262:272])
			meta.TARFiller = strings.TrimSpace(line[272:300])
			continue
		}
		if lineIdentifier == "EM1" {
			meta.EM1RecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.EM1RecordSeqNo = strings.TrimSpace(line[3:10])
			meta.EM1EMVPANSequenceNumber = strings.TrimSpace(line[10:12])
			meta.EM1EMVAuthRespCode = strings.TrimSpace(line[12:14])
			meta.EM1EMVTransAmt = strings.TrimSpace(line[14:26])
			meta.EM1EMVTransType = strings.TrimSpace(line[26:28])
			meta.EM1EMVTransDate = strings.TrimSpace(line[28:34])
			meta.EM1EMVTransCurr = strings.TrimSpace(line[34:37])
			meta.EM1EMVTermCountry = strings.TrimSpace(line[37:40])
			meta.EM1EMVTransCryptogram = strings.TrimSpace(line[40:56])
			meta.EM1EMVAIP = strings.TrimSpace(line[56:60])
			meta.EM1EMVATC = strings.TrimSpace(line[60:64])
			meta.EM1EMVUnpredictNumber = strings.TrimSpace(line[64:72])
			meta.EM1EMVTVR = strings.TrimSpace(line[72:82])
			meta.EM1EMVIAD = strings.TrimSpace(line[82:146])
			meta.EM1EMVAppUsageControl = strings.TrimSpace(line[146:150])
			meta.EM1EMVCryptoInfoData = strings.TrimSpace(line[150:152])
			meta.EM1EMVCVMR = strings.TrimSpace(line[152:158])
			meta.EM1EMVFileName = strings.TrimSpace(line[158:190])
			meta.EM1Filler = strings.TrimSpace(line[190:300])
			continue
		}
		if lineIdentifier == "EM2" {
			meta.EM2RecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.EM2RecordSeqNo = strings.TrimSpace(line[3:10])
			meta.EM2EMVAID = strings.TrimSpace(line[10:42])
			meta.EM2EMVAppVersion = strings.TrimSpace(line[42:46])
			meta.EM2EMVTransStatusIinfo = strings.TrimSpace(line[46:50])
			meta.EM2EMVTermType = strings.TrimSpace(line[50:52])
			meta.EM2EMVTermCapabilities = strings.TrimSpace(line[52:58])
			meta.EM2EMVPOSEntryMode = strings.TrimSpace(line[58:60])
			meta.EM2EMVFillerOCD = strings.TrimSpace(line[60:86])
			meta.EM2EMVIACdefault = strings.TrimSpace(line[86:96])
			meta.EM2EMVIACdenial = strings.TrimSpace(line[96:106])
			//meta.EM2EMVIACon = strings.TrimSpace(line[106:116])
			meta.EM2EMVCardIssuerCntry = strings.TrimSpace(line[116:119])
			meta.EM2EMVIFDSerialNo = strings.TrimSpace(line[119:135])
			meta.EM2EMVTransCatCode = strings.TrimSpace(line[135:137])
			meta.EM2EMVTermAppVersion = strings.TrimSpace(line[137:141])
			meta.EM2EMVTransSeqCounter = strings.TrimSpace(line[141:149])
			meta.EM2EMVIssScriptRespartII = strings.TrimSpace(line[149:169])
			//meta.EM2EMVIssScriptRespartII = strings.TrimSpace(line[169:187])
			meta.EM2EMV3Follows = strings.TrimSpace(line[187:188])
			meta.EM2Filler = strings.TrimSpace(line[190:300])
			continue
		}
		if lineIdentifier == "EM3" {
			meta.EM3RecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.EM3RecordSeqNo = strings.TrimSpace(line[3:10])
			meta.EM3EMVCtlsFormFactor = strings.TrimSpace(line[10:74])
			meta.EM3EMVCtlsDiscrData = strings.TrimSpace(line[74:138])
			meta.EM3Filler = strings.TrimSpace(line[138:300])
			continue
		}
		if lineIdentifier == "ECR" {
			meta.ECRRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.ECRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.ECREcommerceTranSecurity = strings.TrimSpace(line[10:11])
			meta.ECRAVSResponseCode = strings.TrimSpace(line[11:12])
			meta.ECRCVV2ResultCode = strings.TrimSpace(line[12:13])
			meta.ECRECSIndicator = strings.TrimSpace(line[13:16])
			meta.ECRAdditionalUCAFData = strings.TrimSpace(line[16:48])
			meta.ECRMerchantURL = strings.TrimSpace(line[48:98])
			meta.ECRMerchantTelephoneNumber = strings.TrimSpace(line[98:118])
			meta.ECREcommerceGoodsIndicator = strings.TrimSpace(line[118:119])
			meta.ECRDateTransactionAuthorised = strings.TrimSpace(line[119:125])
			meta.ECRDateTransactionCaptured = strings.TrimSpace(line[125:131])
			meta.ECRRecordFiller = strings.TrimSpace(line[131:300])
			continue
		}
		if lineIdentifier == "TIF" {
			meta.TIFRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.TIFRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.TIFIFAmount = strings.TrimSpace(line[10:22])
			meta.TIFIFCurrencyCode = strings.TrimSpace(line[22:25])
			meta.TIFIFRate = strings.TrimSpace(line[25:33])
			meta.TIFIFPPT = strings.TrimSpace(line[33:42])
			meta.TIFIFParameterCode = strings.TrimSpace(line[42:82])
			meta.TIFIFIndicator1 = strings.TrimSpace(line[82:92])
			meta.TIFFiller = strings.TrimSpace(line[92:300])
			continue
		}
		if lineIdentifier == "TSF" {
			meta.TSFRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.TSFRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.TSFSFAmount = strings.TrimSpace(line[10:22])
			meta.TSFSFCurrencyCode = strings.TrimSpace(line[22:25])
			meta.TSFSFRate = strings.TrimSpace(line[25:33])
			meta.TSFSFPPT = strings.TrimSpace(line[33:42])
			meta.TSFSFParameterCode = strings.TrimSpace(line[42:82])
			meta.TSFFiller = strings.TrimSpace(line[82:300])
			continue
		}
		if lineIdentifier == "TAF" {
			meta.TAFRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.TAFRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.TAFAFAmount = strings.TrimSpace(line[10:22])
			meta.TAFAFCurrencyCode = strings.TrimSpace(line[22:25])
			meta.TAFAFRate = strings.TrimSpace(line[25:33])
			meta.TAFAFPPT = strings.TrimSpace(line[33:42])
			meta.TAFAFParameterCode = strings.TrimSpace(line[42:82])
			meta.TAFAFCardScheme = strings.TrimSpace(line[82:85])
			meta.TAFAFRegion = strings.TrimSpace(line[85:95])
			meta.TAFAFProduct = strings.TrimSpace(line[95:98])
			meta.TAFAFCardPresence = strings.TrimSpace(line[98:101])
			meta.TAFAFType = strings.TrimSpace(line[101:104])
			meta.TAFAFStatus = strings.TrimSpace(line[104:107])
			meta.TAFAFAdditional1 = strings.TrimSpace(line[107:110])
			meta.TAFAFAdditional2 = strings.TrimSpace(line[110:113])
			meta.TAFAFAdditional3 = strings.TrimSpace(line[113:116])
			meta.TAFFiller = strings.TrimSpace(line[116:300])

			batchFileTxrList = append(batchFileTxrList, meta)

			continue
		}
		if lineIdentifier == "BTR" {
			meta.BTRRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.BTRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.BTRBatchSettleBankCode = strings.TrimSpace(line[10:12])
			meta.BTRBatchSettleBankData = strings.TrimSpace(line[12:37])
			meta.BTRBatchSettleBankType = strings.TrimSpace(line[37:39])
			meta.BTRBatchSettlementCurrChoice = strings.TrimSpace(line[39:40])
			meta.BTRBatchSettleCode = strings.TrimSpace(line[40:42])
			//meta.BTRBatchSettleBanklineRef = strings.TrimSpace(line[42:57])
			meta.BTRBatchSettlementCurrencyCode = strings.TrimSpace(line[57:60])
			meta.BTRBatchSettlementTotalDebit = strings.TrimSpace(line[60:75])
			meta.BTRBatchSettlementTotalCredit = strings.TrimSpace(line[75:90])
			meta.BTRBatchSettlementCountDebit = strings.TrimSpace(line[90:97])
			meta.BTRBatchSettlementCountCredit = strings.TrimSpace(line[97:104])
			meta.BTRRecordFiller = strings.TrimSpace(line[104:300])

			continue
		}
		if lineIdentifier == "MTR" {
			meta.MTRRecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.MTRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.MTRMerchantBaseDebitTotal = strings.TrimSpace(line[10:23])
			meta.MTRMerchantBaseCreditTotal = strings.TrimSpace(line[23:36])
			meta.MTRMerchantTxnCount = strings.TrimSpace(line[36:43])
			meta.MTRMerchantCurrencyTotalHash = strings.TrimSpace(line[43:58])
			meta.MTRMerchantBatchCount = strings.TrimSpace(line[58:61])
			meta.MTRRecordFiller = strings.TrimSpace(line[61:300])
			continue
		}
		if lineIdentifier == "UTR" {
			meta.UTRRecordLabelIdentifer = strings.TrimSpace(line[0:3])
			meta.UTRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.UTRUserHeaderRecordsMerchant = strings.TrimSpace(line[10:17])
			meta.UTRUserCardTxns = strings.TrimSpace(line[17:24])
			meta.UTRValueofDebitsBaseCurr = strings.TrimSpace(line[24:37])
			meta.UTRValueofCreditsBaseCurr = strings.TrimSpace(line[37:50])
			meta.UTRCountofDebits = strings.TrimSpace(line[50:57])
			meta.UTRCountofCredits = strings.TrimSpace(line[57:64])
			meta.UTRRecordFiller = strings.TrimSpace(line[64:300])
			continue
		}
		//Not used for now.
		if lineIdentifier == "AL1" {
			meta.AL1RecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.AL1RecordSeqNo = strings.TrimSpace(line[3:10])
			meta.AL1TicketNumber = strings.TrimSpace(line[10:24])
			meta.AL1CarrierName = strings.TrimSpace(line[24:43])
			meta.AL1TravelAgencyCode = strings.TrimSpace(line[43:51])
			meta.AL1TicketAgencyName = strings.TrimSpace(line[51:76])
			meta.AL1AirlinePlanNumber = strings.TrimSpace(line[76:78])
			meta.AL1AirlineInvoiceNumber = strings.TrimSpace(line[78:84])
			meta.AL1OriginalCurrency = strings.TrimSpace(line[84:87])
			meta.AL1PassengerName = strings.TrimSpace(line[87:113])
			meta.AL1CustomerReferenceNumber = strings.TrimSpace(line[113:125])
			meta.AL1OriginalTransactionAmount = strings.TrimSpace(line[125:138])
			meta.AL1TicketIssuerAddress = strings.TrimSpace(line[138:154])
			meta.AL1TripLeg1DepartAirport = strings.TrimSpace(line[154:157])
			meta.AL1TripLeg1CarrierCode = strings.TrimSpace(line[157:159])
			meta.AL1TripLeg1FareBasisCode = strings.TrimSpace(line[159:165])
			meta.AL1TripLeg1ClassOfTravel = strings.TrimSpace(line[165:166])
			meta.AL1TripLeg1StopOverCode = strings.TrimSpace(line[166:167])
			meta.AL1TripLeg1DestinationCode = strings.TrimSpace(line[167:170])
			meta.AL1TripLeg1DateOfTravel = strings.TrimSpace(line[170:176])
			meta.AL1TripLeg1DepartTax = strings.TrimSpace(line[176:188])
			meta.AL1Filler = strings.TrimSpace(line[188:300])
			continue
		}
		if lineIdentifier == "AL2" {
			meta.AL2RecordLabelIdentifier = strings.TrimSpace(line[0:3])
			meta.AL2RecordSeqNo = strings.TrimSpace(line[3:10])
			//meta.AL2TripLeg2DepartAirport = strings.TrimSpace(line[10:13])
			meta.AL2TripLeg2CarrierCode = strings.TrimSpace(line[13:15])
			meta.AL2TripLeg2FareBasisCode = strings.TrimSpace(line[15:21])
			meta.AL2TripLeg2ClassOfTravel = strings.TrimSpace(line[21:22])
			meta.AL2TripLeg2StopOverCode = strings.TrimSpace(line[22:23])
			meta.AL2TripLeg2DestinationCode = strings.TrimSpace(line[23:26])
			meta.AL2TripLeg2DateOfTravel = strings.TrimSpace(line[26:32])
			meta.AL2TripLeg2DepartTax = strings.TrimSpace(line[32:44])
			meta.AL2TripLeg2DepartAirport = strings.TrimSpace(line[44:47])
			meta.AL2TripLeg3CarrierCode = strings.TrimSpace(line[47:49])
			meta.AL2TripLeg3FareBasisCode = strings.TrimSpace(line[49:55])
			meta.AL2TripLeg3ClassOfTravel = strings.TrimSpace(line[55:56])
			meta.AL2TripLeg3StopOverCode = strings.TrimSpace(line[56:57])
			meta.AL2TripLeg3DestinationCode = strings.TrimSpace(line[57:60])
			meta.AL2TripLeg3DateOfTravel = strings.TrimSpace(line[60:66])
			meta.AL2TripLeg3DepartTax = strings.TrimSpace(line[66:78])
			meta.AL2TripLeg4DepartAirport = strings.TrimSpace(line[78:81])
			meta.AL2TripLeg4CarrierCode = strings.TrimSpace(line[81:83])
			meta.AL2TRIPLEG4FAREBASISCODE = strings.TrimSpace(line[83:89])
			meta.AL2TRIPLEG4CLASSOFTRAVEL = strings.TrimSpace(line[89:90])
			meta.AL2TRIPLEG4STOPOVERCODE = strings.TrimSpace(line[90:91])
			meta.AL2TRIPLEG4DESTINATIONCODE = strings.TrimSpace(line[91:94])
			meta.AL2TRIPLEG4DATEOFTRAVEL = strings.TrimSpace(line[94:100])
			meta.AL2TRIPLEG4DEPARTTAX = strings.TrimSpace(line[100:112])
			meta.AL2TRIPLEG5DEPARTAIRPORT = strings.TrimSpace(line[112:115])
			meta.AL2TRIPLEG5CARRIERCODE = strings.TrimSpace(line[115:117])
			meta.AL2TRIPLEG5FAREBASISCODE = strings.TrimSpace(line[117:123])
			meta.AL2TRIPLEG5CLASSOFTRAVEL = strings.TrimSpace(line[123:124])
			meta.AL2TRIPLEG5STOPOVERCODE = strings.TrimSpace(line[124:125])
			meta.AL2TRIPLEG5DESTINATIONCODE = strings.TrimSpace(line[125:128])
			meta.AL2TRIPLEG5DATEOFTRAVEL = strings.TrimSpace(line[128:134])
			meta.AL2TRIPLEG5DEPARTTAX = strings.TrimSpace(line[134:146])
			meta.AL2TRIPLEG6DEPARTAIRPORT = strings.TrimSpace(line[146:149])
			meta.AL2TRIPLEG6CARRIERCODE = strings.TrimSpace(line[149:151])
			meta.AL2TRIPLEG6FAREBASISCODE = strings.TrimSpace(line[151:157])
			meta.AL2TRIPLEG65CLASSOFTRAVEL = strings.TrimSpace(line[157:158])
			meta.AL2TRIPLEG6STOPOVERCODE = strings.TrimSpace(line[158:159])
			meta.AL2TRIPLEG6DESTINATIONCODE = strings.TrimSpace(line[159:162])
			meta.AL2TRIPLEG6DATEOFTRAVEL = strings.TrimSpace(line[162:168])
			meta.AL2TRIPLEG6DEPARTTAX = strings.TrimSpace(line[168:180])
			meta.AL2Filler = strings.TrimSpace(line[180:300])
			continue
		}
		if lineIdentifier == "FTR" {
			meta.FTRFileRecordIdentifier = strings.TrimSpace(line[0:3])
			meta.FTRRecordSeqNo = strings.TrimSpace(line[3:10])
			meta.FTRFileUserRecords = strings.TrimSpace(line[10:17])
			meta.FTRFileMerchantRecords = strings.TrimSpace(line[17:24])
			meta.FTRFileCardTxns = strings.TrimSpace(line[24:31])
			meta.FTRFileHashNetValue = strings.TrimSpace(line[31:45])
			meta.FTRRecordFiller = strings.TrimSpace(line[45:300])
			continue
		}

	}

	return batchFileTxrList, nil
}

func moveFile(source string, destination string) error {
	err := os.Rename(source, destination)
	if err != nil {
		return err
	}
	return nil
}

func LogMessage(srvlog logging.LogMessage, conf decoder.Config) error {

	Log := &loggingpb.Log{}

	// Prepare connection to Log gRP Server
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()
	gRPCServer := fmt.Sprintf("%s:%d", conf.GRPCServer, conf.GRPCPort)
	conn, err := grpc.Dial(gRPCServer, requestOpts)
	if err != nil {
		conn.Close()
		return err
	}

	loggingClient = loggingpb.NewLogServiceClient(conn)

	Log.Level = srvlog.Level
	Log.CallingMethod = srvlog.CallingMethod
	Log.ServiceName = srvlog.ServiceName
	Log.Host = srvlog.Host
	Log.Body = srvlog.Body
	Log.Latency = srvlog.Latency

	_, err = loggingClient.SaveLog(
		context.TODO(),
		&loggingpb.SaveLogReq{
			Log: Log,
		},
	)

	if err != nil {
		conn.Close()
		return err
	}
	conn.Close()
	return nil
}

func RunPayment(runParams RunParams, conf decoder.Config) error {

	payment := &paymentpb.Payment{}

	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()
	gRPCServer := fmt.Sprintf("%s:%d", conf.GRPCServer, conf.GRPCPort)
	conn, err := grpc.Dial(gRPCServer, requestOpts)
	if err != nil {
		conn.Close()
		return err
	}

	paymentClient = paymentpb.NewPaymentServiceClient(conn)

	payment.AcquirerId = runParams.AcquirerID
	payment.StartDate = runParams.StartDate
	payment.EndDate = runParams.StartDate

	_, err = paymentClient.RunPayment(
		context.TODO(),
		&paymentpb.RunPaymentReq{
			Payment: payment,
		},
	)

	if err != nil {
		conn.Close()
		return err
	}
	conn.Close()
	return err
}

func (s *service) Get(date, time string) ([]string, error) {
	if err := validate.Validate(date); err != nil {
		return nil, errs.Wrap(ErrTransactionInvalid, "service.Transaction.Get")
	}
	acquirers, err := s.batchFileRepo.Get(date, time)
	if err != nil {
		return nil, errs.Wrap(ErrTransactionInvalid, "service.Transaction.Get")
	}
	return acquirers, nil
}
