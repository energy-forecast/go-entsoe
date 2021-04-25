//go:generate go run ./tools/extract_types.go

package goentsoe

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type PsrType string

const (
	PsrTypeMixed                      PsrType = "A03"
	PsrTypeGeneration                 PsrType = "A04"
	PsrTypeLoad                       PsrType = "A05"
	PsrTypeBiomass                    PsrType = "B01"
	PsrTypeFossilBrownCoalLignite     PsrType = "B02"
	PsrTypeFossilCoalDerivedGas       PsrType = "B03"
	PsrTypeFossilGas                  PsrType = "B04"
	PsrTypeFossilHardCoal             PsrType = "B05"
	PsrTypeFossilOil                  PsrType = "B06"
	PsrTypeFossilOilShale             PsrType = "B07"
	PsrTypeFossilPeat                 PsrType = "B08"
	PsrTypeGeothermal                 PsrType = "B09"
	PsrTypeHydroPumpedStorage         PsrType = "B10"
	PsrTypeHydroRunOfRiverAndPoundage PsrType = "B11"
	PsrTypeHydroWaterReservoir        PsrType = "B12"
	PsrTypeMarine                     PsrType = "B13"
	PsrTypeNuclear                    PsrType = "B14"
	PsrTypeOtherRenewable             PsrType = "B15"
	PsrTypeSolar                      PsrType = "B16"
	PsrTypeWaste                      PsrType = "B17"
	PsrTypeWindOffshore               PsrType = "B18"
	PsrTypeWindOnshore                PsrType = "B19"
	PsrTypeOther                      PsrType = "B20"
	PsrTypeACLink                     PsrType = "B21"
	PsrTypeDCLink                     PsrType = "B22"
	PsrTypeSubstation                 PsrType = "B23"
	PsrTypeTransformer                PsrType = "B24"
)

type BusinessType string

const (
	BusinessTypeGeneralCapacityInformation           BusinessType = "A25"
	BusinessTypeAlreadyAllocatedCapacity             BusinessType = "A29"
	BusinessTypeRequestedCapacity                    BusinessType = "A43"
	BusinessTypeSystemOperatorRedispatching          BusinessType = "A46"
	BusinessTypePlannedMaintenance                   BusinessType = "A53"
	BusinessTypeUnplannedOutage                      BusinessType = "A54"
	BusinessTypeInternalRedispatch                   BusinessType = "A85"
	BusinessTypeFrequencyContainmentReserve          BusinessType = "A95"
	BusinessTypeAutomaticFrequencyRestorationReserve BusinessType = "A96"
	BusinessTypeManualFrequencyRestorationReserve    BusinessType = "A97"
	BusinessTypeReplacementReserve                   BusinessType = "A98"
	BusinessTypeInterconnectorNetworkEvolution       BusinessType = "B01"
	BusinessTypeInterconnectorNetworkDismantling     BusinessType = "B02"
	BusinessTypeCounterTrade                         BusinessType = "B03"
	BusinessTypeCongestionCosts                      BusinessType = "B04"
	BusinessTypeCapacityAllocated                    BusinessType = "B05"
	BusinessTypeAuctionRevenue                       BusinessType = "B07"
	BusinessTypeTotalNominatedCapacity               BusinessType = "B08"
	BusinessTypeNetPosition                          BusinessType = "B09"
	BusinessTypeCongestionIncome                     BusinessType = "B10"
	BusinessTypeProductionUnit                       BusinessType = "B11"
	BusinessTypeAreaControlError                     BusinessType = "B33"
	BusinessTypeProcuredCapacity                     BusinessType = "B95"
	BusinessTypeSharedBalancingReserveCapacity       BusinessType = "C22"
	BusinessTypeShareOfReserveCapacity               BusinessType = "C23"
	BusinessTypeActualReserveCapacity                BusinessType = "C24"
)

type ProcessType string

const (
	ProcessTypeDayAhead                             ProcessType = "A01"
	ProcessTypeIntraDayIncremental                  ProcessType = "A02"
	ProcessTypeRealised                             ProcessType = "A16"
	ProcessTypeIntradayTotal                        ProcessType = "A18"
	ProcessTypeWeekAhead                            ProcessType = "A31"
	ProcessTypeMonthAhead                           ProcessType = "A32"
	ProcessTypeYearAhead                            ProcessType = "A33"
	ProcessTypeSynchronisationProcess               ProcessType = "A39"
	ProcessTypeIntradayProcess                      ProcessType = "A40"
	ProcessTypeReplacementReserve                   ProcessType = "A46"
	ProcessTypeManualFrequencyRestorationReserve    ProcessType = "A47"
	ProcessTypeAutomaticFrequencyRestorationReserve ProcessType = "A51"
	ProcessTypeFrequencyContainmentReserve          ProcessType = "A52"
	ProcessTypeFrequencyRestorationReserve          ProcessType = "A56"
)

type DocStatus string

const (
	DocStatusIntermediate DocStatus = "A01"
	DocStatusFinal        DocStatus = "A02"
	DocStatusActive       DocStatus = "A05"
	DocStatusCancelled    DocStatus = "A09"
	DocStatusEstimated    DocStatus = "X01"
)

type DocumentType string

const (
	DocumentTypeFinalisedSchedule                        DocumentType = "A09"
	DocumentTypeAggregatedEnergyDataReport               DocumentType = "A11"
	DocumentTypeAcquiringSystemOperatorReserveSchedule   DocumentType = "A15"
	DocumentTypeBidDocument                              DocumentType = "A24"
	DocumentTypeAllocationResultDocument                 DocumentType = "A25"
	DocumentTypeCapacityDocument                         DocumentType = "A26"
	DocumentTypeAgreedCapacity                           DocumentType = "A31"
	DocumentTypeReserveAllocationResultDocument          DocumentType = "A38"
	DocumentTypePriceDocument                            DocumentType = "A44"
	DocumentTypeEstimatedNetTransferCapacity             DocumentType = "A61"
	DocumentTypeRedispatchNotice                         DocumentType = "A63"
	DocumentTypeSystemTotalLoad                          DocumentType = "A65"
	DocumentTypeInstalledGenerationPerType               DocumentType = "A68"
	DocumentTypeWindAndSolarForecast                     DocumentType = "A69"
	DocumentTypeLoadForecastMargin                       DocumentType = "A70"
	DocumentTypeGenerationForecast                       DocumentType = "A71"
	DocumentTypeReservoirFillingInformation              DocumentType = "A72"
	DocumentTypeActualGeneration                         DocumentType = "A73"
	DocumentTypeWindAndSolarGeneration                   DocumentType = "A74"
	DocumentTypeActualGenerationPerType                  DocumentType = "A75"
	DocumentTypeLoadUnavailability                       DocumentType = "A76"
	DocumentTypeProductionUnavailability                 DocumentType = "A77"
	DocumentTypeTransmissionUnavailability               DocumentType = "A78"
	DocumentTypeOffshoreGridInfrastructureUnavailability DocumentType = "A79"
	DocumentTypeGenerationUnavailability                 DocumentType = "A80"
	DocumentTypeContractedReserves                       DocumentType = "A81"
	DocumentTypeAcceptedOffers                           DocumentType = "A82"
	DocumentTypeActivatedBalancingQuantities             DocumentType = "A83"
	DocumentTypeActivatedBalancingPrices                 DocumentType = "A84"
	DocumentTypeImbalancePrices                          DocumentType = "A85"
	DocumentTypeImbalanceVolume                          DocumentType = "A86"
	DocumentTypeFinancialSituation                       DocumentType = "A87"
	DocumentTypeCrossBorderBalancing                     DocumentType = "A88"
	DocumentTypeContractedReservePrices                  DocumentType = "A89"
	DocumentTypeInterconnectionNetworkExpansion          DocumentType = "A90"
	DocumentTypeCounterTradeNotice                       DocumentType = "A91"
	DocumentTypeCongestionCosts                          DocumentType = "A92"
	DocumentTypeDcLinkCapacity                           DocumentType = "A93"
	DocumentTypeNonEuAllocations                         DocumentType = "A94"
	DocumentTypeConfigurationDocument                    DocumentType = "A95"
	DocumentTypeFlowBasedAllocations                     DocumentType = "B11"
)

type DomainType = string

const (
	DomainAL     DomainType = "10YAL-KESH-----5"
	DomainAT     DomainType = "10YAT-APG------L"
	DomainBA     DomainType = "10YBA-JPCC-----D"
	DomainBE     DomainType = "10YBE----------2"
	DomainBG     DomainType = "10YCA-BULGARIA-R"
	DomainBY     DomainType = "10Y1001A1001A51S"
	DomainCH     DomainType = "10YCH-SWISSGRIDZ"
	DomainCZ     DomainType = "10YCZ-CEPS-----N"
	DomainDE     DomainType = "10Y1001A1001A83F"
	DomainDK     DomainType = "10Y1001A1001A65H"
	DomainEE     DomainType = "10Y1001A1001A39I"
	DomainES     DomainType = "10YES-REE------0"
	DomainFI     DomainType = "10YFI-1--------U"
	DomainFR     DomainType = "10YFR-RTE------C"
	DomainGB     DomainType = "10YGB----------A"
	DomainGBNIR  DomainType = "10Y1001A1001A016"
	DomainGR     DomainType = "10YGR-HTSO-----Y"
	DomainHR     DomainType = "10YHR-HEP------M"
	DomainHU     DomainType = "10YHU-MAVIR----U"
	DomainIE     DomainType = "10YIE-1001A00010"
	DomainIT     DomainType = "10YIT-GRTN-----B"
	DomainLT     DomainType = "10YLT-1001A0008Q"
	DomainLU     DomainType = "10YLU-CEGEDEL-NQ"
	DomainLV     DomainType = "10YLV-1001A00074"
	DomainME     DomainType = "10YCS-CG-TSO---S"
	DomainMK     DomainType = "10YMK-MEPSO----8"
	DomainMT     DomainType = "10Y1001A1001A93C"
	DomainNL     DomainType = "10YNL----------L"
	DomainNO     DomainType = "10YNO-0--------C"
	DomainPL     DomainType = "10YPL-AREA-----S"
	DomainPT     DomainType = "10YPT-REN------W"
	DomainRO     DomainType = "10YRO-TEL------P"
	DomainRS     DomainType = "10YCS-SERBIATSOV"
	DomainRU     DomainType = "10Y1001A1001A49F"
	DomainRUKGD  DomainType = "10Y1001A1001A50U"
	DomainSE     DomainType = "10YSE-1--------K"
	DomainSI     DomainType = "10YSI-ELES-----O"
	DomainSK     DomainType = "10YSK-SEPS-----K"
	DomainTR     DomainType = "10YTR-TEIAS----W"
	DomainUA     DomainType = "10YUA-WEPS-----0"
	DomainDEATLU DomainType = "10Y1001A1001A63L"
)

type EntsoeClient struct {
	apiKey string
}

func NewEntsoeClient(apiKey string) *EntsoeClient {
	c := EntsoeClient{
		apiKey: apiKey,
	}
	return &c
}

func NewEntsoeClientFromEnv() *EntsoeClient {
	apiKey := os.Getenv("ENTSOE_API_KEY")
	if apiKey == "" {
		log.Fatal("Environment variable ENTSOE_API_KEY with api key not set")
	}

	c := EntsoeClient{
		apiKey: apiKey,
	}
	return &c
}

type Parameter string

const (
	ParameterDocumentType                                             = "documentType"
	ParameterDocStatus                                                = "docStatus"
	ParameterProcessType                                              = "processType"
	ParameterBusinessType                                             = "businessType"
	ParameterPsrType                                                  = "psrType"
	ParameterTypeMarketAgreementType                                  = "type_MarketAgreement.type"
	ParameterContractMarketAgreementType                              = "contract_MarketAgreement.Type"
	ParameterAuctionType                                              = "auction.Type"
	ParameterAuctionCategory                                          = "auction.Category"
	ParameterClassificationSequenceAttributeInstanceComponentPosition = "classificationSequence_AttributeInstanceComponent.Position"
	ParameterOutBiddingZoneDomain                                     = "outBiddingZone_Domain"
	ParameterBiddingZoneDomain                                        = "biddingZone_Domain"
	ParameterControlAreaDomain                                        = "controlArea_Domain"
	ParameterInDomain                                                 = "in_Domain"
	ParameterOutDomain                                                = "out_Domain"
	ParameterAcquiringDomain                                          = "acquiring_Domain"
	ParameterConnectingDomain                                         = "connecting_Domain"
	ParameterRegisteredResource                                       = "RegisteredResource"
	ParameterTimeInterval                                             = "TimeInterval"
	ParameterPeriodStart                                              = "periodStart"
	ParameterPeriodEnd                                                = "periodEnd"
	ParameterTimeIntervalUpdate                                       = "TimeIntervalUpdate"
	ParameterPeriodStartUpdate                                        = "PeriodStartUpdate"
	ParameterPeriodEndUpdate                                          = "PeriodEndUpdate"
)

type ContractMarketAgreementType string

const (
	ContractMarketAgreementTypeDaily    ContractMarketAgreementType = "A01"
	ContractMarketAgreementTypeWeekly   ContractMarketAgreementType = "A02"
	ContractMarketAgreementTypeMonthly  ContractMarketAgreementType = "A03"
	ContractMarketAgreementTypeYearly   ContractMarketAgreementType = "A04"
	ContractMarketAgreementTypeTotal    ContractMarketAgreementType = "A05"
	ContractMarketAgreementTypeLongTerm ContractMarketAgreementType = "A06"
	ContractMarketAgreementTypeIntraday ContractMarketAgreementType = "A07"
	ContractMarketAgreementTypeHourly   ContractMarketAgreementType = "A13"
)

type AuctionType string

const (
	AuctionTypeImplicit AuctionType = "A01"
	AuctionTypeExplicit AuctionType = "A02"
)

type AuctionCategory string

const (
	AuctionCategoryBase    AuctionCategory = "A01"
	AuctionCategoryPeak    AuctionCategory = "A02"
	AuctionCategoryOffPeak AuctionCategory = "A03"
	AuctionCategoryHourly  AuctionCategory = "A04"
)

// 4.1. Load domain

// 4.1.1. Actual Total Load [6.1.A]
func (c *EntsoeClient) GetActualTotalLoad(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeSystemTotalLoad))
	params.Add(ParameterProcessType, string(ProcessTypeRealised))
	params.Add(ParameterOutBiddingZoneDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.1.2. Day-Ahead Total Load Forecast [6.1.B]
func (c *EntsoeClient) GetDayAheadTotalLoadForecast(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeSystemTotalLoad))
	params.Add(ParameterProcessType, string(ProcessTypeDayAhead))
	params.Add(ParameterOutBiddingZoneDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.1.3. Week-Ahead Total Load Forecast [6.1.C]
func (c *EntsoeClient) GetWeekAheadTotalLoadForecast(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeSystemTotalLoad))
	params.Add(ParameterProcessType, string(ProcessTypeWeekAhead))
	params.Add(ParameterOutBiddingZoneDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.1.4. Month-Ahead Total Load Forecast [6.1.D]
func (c *EntsoeClient) GetMonthAheadTotalLoadForecast(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeSystemTotalLoad))
	params.Add(ParameterProcessType, string(ProcessTypeMonthAhead))
	params.Add(ParameterOutBiddingZoneDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.1.5. Year-Ahead Total Load Forecast [6.1.E]
func (c *EntsoeClient) GetYearAheadTotalLoadForecast(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeSystemTotalLoad))
	params.Add(ParameterProcessType, string(ProcessTypeYearAhead))
	params.Add(ParameterOutBiddingZoneDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.1.6. Year-Ahead Forecast Margin [8.1]
func (c *EntsoeClient) GetYearAheadForecastMargin(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeLoadForecastMargin))
	params.Add(ParameterProcessType, string(ProcessTypeYearAhead))
	params.Add(ParameterOutBiddingZoneDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.2. Transmission domain

// 4.2.1. Expansion and Dismantling Projects [9.1]
func (c *EntsoeClient) GetExpansionAndDismantlingProjects(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	business *BusinessType,
	docStatus *DocStatus,
) (*TransmissionNetworkMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeInterconnectionNetworkExpansion))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if business != nil {
		params.Add(ParameterBusinessType, string(*business))
	}
	if docStatus != nil {
		params.Add(ParameterDocStatus, string(*docStatus))
	}
	return c.requestTransmissionNetworkMarketDocument(params)
}

// 4.2.2. Forecasted Capacity [11.1.A]
func (c *EntsoeClient) GetForecastedCapacity(
	contractMarketAgreement ContractMarketAgreementType,
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeEstimatedNetTransferCapacity))
	params.Add(ParameterContractMarketAgreementType, string(contractMarketAgreement))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestPublicationMarketDocument(params)
}

// 4.2.3. Offered Capacity [11.1.A]
func (c *EntsoeClient) GetOfferedCapacity(
	auctionType AuctionType,
	contractMarketAgreement ContractMarketAgreementType,
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	auctionCategory *AuctionCategory,
	classificationSequenceAttributeInstanceComponentPosition *int,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeAgreedCapacity))
	params.Add(ParameterAuctionType, string(auctionType))
	params.Add(ParameterContractMarketAgreementType, string(contractMarketAgreement))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if auctionCategory != nil {
		params.Add(ParameterAuctionCategory, string(*auctionCategory))
	}
	if auctionCategory != nil {
		params.Add(ParameterClassificationSequenceAttributeInstanceComponentPosition, strconv.Itoa(*classificationSequenceAttributeInstanceComponentPosition))
	}
	return c.requestPublicationMarketDocument(params)
}

// 4.2.4. Flow-based Parameters [11.1.B]
func (c *EntsoeClient) GetFlowBasedParameters(
	processType ProcessType,
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*CriticalNetworkElementMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeFlowBasedAllocations))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(domain))
	params.Add(ParameterOutDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestCriticalNetworkElementMarketDocument(params)
}

// 4.2.5. Intraday Transfer Limits [11.3]
func (c *EntsoeClient) GetIntradayTransferLimits(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeDcLinkCapacity))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestPublicationMarketDocument(params)
}

// 4.2.6. Explicit Allocation Information (Capacity) [12.1.A]
// 4.2.7. Explicit Allocation Information (Revenue only) [12.1.A]
func (c *EntsoeClient) GetExplicitAllocationInformation(
	businessType BusinessType,
	contractMarketAgreementType ContractMarketAgreementType,
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	auctionCategory *AuctionCategory,
	classificationSequenceAttributeInstanceComponentPosition *int,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeAllocationResultDocument))
	params.Add(ParameterBusinessType, string(businessType))
	params.Add(ParameterContractMarketAgreementType, string(contractMarketAgreementType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if auctionCategory != nil {
		params.Add(ParameterAuctionCategory, string(*auctionCategory))
	}
	if classificationSequenceAttributeInstanceComponentPosition != nil {
		params.Add(ParameterClassificationSequenceAttributeInstanceComponentPosition, strconv.Itoa(*classificationSequenceAttributeInstanceComponentPosition))
	}
	return c.requestPublicationMarketDocument(params)
}

// 4.2.8. Total Capacity Nominated [12.1.B]
func (c *EntsoeClient) GetTotalCapacityNominated(
	businessType BusinessType,
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeCapacityDocument))
	params.Add(ParameterBusinessType, string(businessType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestPublicationMarketDocument(params)
}

// 4.2.9. Total Capacity Already Allocated [12.1.C]
func (c *EntsoeClient) GetTotalCapacityAlreadyAllocated(
	businessType BusinessType,
	contractMarketAgreementType ContractMarketAgreementType,
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	auctionCategory *AuctionCategory,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeCapacityDocument))
	params.Add(ParameterBusinessType, string(businessType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if auctionCategory != nil {
		params.Add(ParameterAuctionCategory, string(*auctionCategory))
	}
	return c.requestPublicationMarketDocument(params)
}

// 4.2.10. Day Ahead Prices [12.1.D]
func (c *EntsoeClient) GetDayAheadPrices(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypePriceDocument))
	params.Add(ParameterInDomain, string(domain))
	params.Add(ParameterOutDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestPublicationMarketDocument(params)
}

// 4.2.11. Implicit Auction — Net Positions [12.1.E]
// 4.2.12. Implicit Auction — Congestion Income [12.1.E]
func (c *EntsoeClient) GetImplicitAuction(
	businessType BusinessType,
	contractMarketAgreementType ContractMarketAgreementType,
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeAllocationResultDocument))
	params.Add(ParameterBusinessType, string(businessType))
	params.Add(ParameterContractMarketAgreementType, string(contractMarketAgreementType))
	params.Add(ParameterInDomain, string(domain))
	params.Add(ParameterOutDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestPublicationMarketDocument(params)
}

// 4.2.13. Total Commercial Schedules [12.1.F]
func (c *EntsoeClient) GetTotalCommercialSchedules(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	contractType *ContractMarketAgreementType,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeFinalisedSchedule))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if contractType != nil {
		params.Add(ParameterContractMarketAgreementType, string(*contractType))
	}
	return c.requestPublicationMarketDocument(params)
}

// 4.2.14. Day-ahead Commercial Schedules [12.1.F]
func (c *EntsoeClient) GetDayAheadCommercialSchedules(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	contractType *ContractMarketAgreementType,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeFinalisedSchedule))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if contractType != nil {
		params.Add(ParameterContractMarketAgreementType, string(*contractType))
	}
	return c.requestPublicationMarketDocument(params)
}

// 4.2.15. Physical Flows [12.1.G]
func (c *EntsoeClient) GetPhysicalFlows(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeAggregatedEnergyDataReport))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestPublicationMarketDocument(params)
}

// 4.2.16. Capacity Allocated Outside EU [12.1.H]
func (c *EntsoeClient) GetCapacityAllocatedOutsideEu(
	auctionType AuctionType,
	contractMarketAgreementType ContractMarketAgreementType,
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	auctionCategory *AuctionCategory,
	classificationSequenceAttributeInstanceComponentPosition *int,
) (*PublicationMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeNonEuAllocations))
	params.Add(ParameterAuctionType, string(auctionType))
	params.Add(ParameterContractMarketAgreementType, string(contractMarketAgreementType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if auctionCategory != nil {
		params.Add(ParameterAuctionCategory, string(*auctionCategory))
	}
	if classificationSequenceAttributeInstanceComponentPosition != nil {
		params.Add(ParameterClassificationSequenceAttributeInstanceComponentPosition, strconv.Itoa(*classificationSequenceAttributeInstanceComponentPosition))
	}
	return c.requestPublicationMarketDocument(params)
}

// 4.3. Congestion domain

// 4.3.1. Redispatching [13.1.A]
func (c *EntsoeClient) GetRedispatching(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	business *BusinessType,
) (*TransmissionNetworkMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeRedispatchNotice))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if business != nil {
		params.Add(ParameterBusinessType, string(*business))
	}
	return c.requestTransmissionNetworkMarketDocument(params)
}

// 4.3.2. Countertrading [13.1.B]
func (c *EntsoeClient) GetCountertrading(
	inDomain DomainType,
	outDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*TransmissionNetworkMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeCounterTradeNotice))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterOutDomain, string(outDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestTransmissionNetworkMarketDocument(params)
}

// 4.3.3. Costs of Congestion Management [13.1.C]
func (c *EntsoeClient) GetCostsOfCongestionManagement(
	domain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	business *BusinessType,
) (*TransmissionNetworkMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeCongestionCosts))
	params.Add(ParameterInDomain, string(domain))
	params.Add(ParameterOutDomain, string(domain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if business != nil {
		params.Add(ParameterBusinessType, string(*business))
	}
	return c.requestTransmissionNetworkMarketDocument(params)
}

// 4.4. Generation domain

// 4.4.1. Installed Generation Capacity Aggregated [14.1.A]
func (c *EntsoeClient) GetInstalledGenerationCapacityAggregated(
	processType ProcessType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	psrType *PsrType,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeInstalledGenerationPerType))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if psrType != nil {
		params.Add(ParameterPsrType, string(*psrType))
	}
	return c.requestGLMarketDocument(params)
}

// 4.4.2. Installed Generation Capacity per Unit [14.1.B]
// TODO: is document type correct?
func (c *EntsoeClient) GetInstalledGenerationCapacityPerUnit(
	processType ProcessType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	psrType *PsrType,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeGenerationForecast))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if psrType != nil {
		params.Add(ParameterPsrType, string(*psrType))
	}
	return c.requestGLMarketDocument(params)
}

// 4.4.3. Day-ahead Aggregated Generation [14.1.C]
func (c *EntsoeClient) GetDayAheadAggregatedGeneration(
	processType ProcessType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeGenerationForecast))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.4.4. Day-ahead Generation Forecasts for Wind and Solar [14.1.D]
// 4.4.5. Current Generation Forecasts for Wind and Solar [14.1.D]
// 4.4.6. Intraday Generation Forecasts for Wind and Solar [14.1.D]
func (c *EntsoeClient) GetGenerationForecastsForWindAndSolar(
	processType ProcessType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	psrType *PsrType,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeWindAndSolarForecast))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if psrType != nil {
		params.Add(ParameterPsrType, string(*psrType))
	}
	return c.requestGLMarketDocument(params)
}

// 4.4.7. Actual Generation Output per Generation Unit [16.1.A]
// TODO: registeredResource missing
func (c *EntsoeClient) GetActualGenerationOutputPerGenerationUnit(
	processType ProcessType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
	psrType *PsrType,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeActualGeneration))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	if psrType != nil {
		params.Add(ParameterPsrType, string(*psrType))
	}
	return c.requestGLMarketDocument(params)
}

// 4.4.8. Aggregated Generation per Type [16.1.B&C]
func (c *EntsoeClient) GetAggregatedGenerationPerType(
	processType ProcessType,
	psrType PsrType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeActualGenerationPerType))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterPsrType, string(psrType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

// 4.4.9. Aggregated Filling Rate of Water Reservoirs and Hydro Storage Plants [16.1.D]
func (c *EntsoeClient) GetAggregatedFillingRateOfWaterReservoirsAndHydroStoragePlants(
	processType ProcessType,
	inDomain DomainType,
	periodStart time.Time,
	periodEnd time.Time,
) (*GLMarketDocument, error) {
	params := url.Values{}
	params.Add(ParameterDocumentType, string(DocumentTypeReservoirFillingInformation))
	params.Add(ParameterProcessType, string(processType))
	params.Add(ParameterInDomain, string(inDomain))
	params.Add(ParameterPeriodStart, periodStart.UTC().Format("200601021504"))
	params.Add(ParameterPeriodEnd, periodEnd.UTC().Format("200601021504"))
	return c.requestGLMarketDocument(params)
}

func (c *EntsoeClient) requestGLMarketDocument(params url.Values) (*GLMarketDocument, error) {
	paramStr := params.Encode()
	data, err := c.sendRequest(paramStr)
	if err != nil {
		return nil, err
	}

	var doc GLMarketDocument
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		fmt.Println(string(data))
		return nil, err
	}
	return &doc, nil
}

func (c *EntsoeClient) requestTransmissionNetworkMarketDocument(params url.Values) (*TransmissionNetworkMarketDocument, error) {
	paramStr := params.Encode()
	data, err := c.sendRequest(paramStr)
	if err != nil {
		return nil, err
	}

	var doc TransmissionNetworkMarketDocument
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		fmt.Println(string(data))
		return nil, err
	}
	return &doc, nil
}

func (c *EntsoeClient) requestPublicationMarketDocument(params url.Values) (*PublicationMarketDocument, error) {
	paramStr := params.Encode()
	data, err := c.sendRequest(paramStr)
	if err != nil {
		return nil, err
	}

	var doc PublicationMarketDocument
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		fmt.Println(string(data))
		return nil, err
	}
	return &doc, nil
}

func (c *EntsoeClient) requestCriticalNetworkElementMarketDocument(params url.Values) (*CriticalNetworkElementMarketDocument, error) {
	paramStr := params.Encode()
	data, err := c.sendRequest(paramStr)
	if err != nil {
		return nil, err
	}

	var doc CriticalNetworkElementMarketDocument
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		fmt.Println(string(data))
		return nil, err
	}
	return &doc, nil
}

func (c *EntsoeClient) sendRequest(paramStr string) ([]byte, error) {
	resp, err := http.Get("https://transparency.entsoe.eu/api?securityToken=" + c.apiKey + "&" + paramStr)
	if err != nil {
		return nil, err
	}
	body := resp.Body
	defer body.Close()
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
