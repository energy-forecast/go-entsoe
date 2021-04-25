package goentsoe

import (
	"log"
	"testing"
	"time"

	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

// 4.1. Load domain

// 4.1.1. Actual Total Load [6.1.A]
func TestGetActualTotalLoad(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetActualTotalLoad(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.1.2. Day-Ahead Total Load Forecast [6.1.B]
func TestGetDayAheadTotalLoadForecast(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetDayAheadTotalLoadForecast(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.1.3. Week-Ahead Total Load Forecast [6.1.C]
func TestGetWeekAheadTotalLoadForecast(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetWeekAheadTotalLoadForecast(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.1.4. Month-Ahead Total Load Forecast [6.1.D]
func TestGetMonthAheadTotalLoadForecast(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetMonthAheadTotalLoadForecast(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.1.5. Year-Ahead Total Load Forecast [6.1.E]
func TestGetYearAheadTotalLoadForecast(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetYearAheadTotalLoadForecast(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.1.6. Year-Ahead Forecast Margin [8.1]
func TestGetYearAheadForecastMargin(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetYearAheadForecastMargin(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2. Transmission domain

// 4.2.1. Expansion and Dismantling Projects [9.1]

func TestGetExpansionAndDismantlingProjects(t *testing.T) {
	t.Skip("TODO: always returns \"No matching data found for Data item Expansion and Dismantling Projects [9.1]\"")
	c := NewEntsoeClientFromEnv()
	businessType := BusinessTypeInterconnectorNetworkEvolution
	doc, err := c.GetExpansionAndDismantlingProjects(
		DomainCZ,
		DomainSK,
		genTime("201512312300"),
		genTime("201612312300"),
		&businessType,
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.2. Forecasted Capacity [11.1.A]
func TestGetForecastedCapacity(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetForecastedCapacity(
		ContractMarketAgreementTypeDaily,
		DomainCZ,
		DomainSK,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.3. Offered Capacity [11.1.A]
func TestGetOfferedCapacity(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetOfferedCapacity(
		AuctionTypeImplicit,
		ContractMarketAgreementTypeDaily,
		DomainSK,
		DomainCZ,
		genTime("201601012300"),
		genTime("201601022300"),
		nil,
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.4. Flow-based Parameters [11.1.B]
func TestGetFlowBasedParameters(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetFlowBasedParameters(
		ProcessTypeDayAhead,
		"10YDOM-REGION-1V",
		genTime("201512312300"),
		genTime("201601012300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.5. Intraday Transfer Limits [11.3]
func TestGetIntradayTransferLimits(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetIntradayTransferLimits(
		DomainFR,
		DomainGB,
		genTime("201512312300"),
		genTime("201601012300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.6. Explicit Allocation Information (Capacity) [12.1.A]
func TestExplicitAllocationInformationCapacity(t *testing.T) {
	t.Skip("TODO: always returns \"No matching data found for Data item Explicit Allocations\"")
	c := NewEntsoeClientFromEnv()
	auctionCategory := AuctionCategoryBase
	doc, err := c.GetExplicitAllocationInformation(
		BusinessTypeCapacityAllocated,
		ContractMarketAgreementTypeDaily,
		DomainSK,
		DomainCZ,
		genTime("201601012300"),
		genTime("201601022300"),
		&auctionCategory,
		pointy.Int(1),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.7. Explicit Allocation Information (Revenue only) [12.1.A]
func TestExplicitAllocationInformationRevenueOnly(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetExplicitAllocationInformation(
		BusinessTypeAuctionRevenue,
		ContractMarketAgreementTypeDaily,
		DomainAT,
		DomainCZ,
		genTime("201601012300"),
		genTime("201601022300"),
		nil,
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.8. Total Capacity Nominated [12.1.B]
func TestGetTotalCapacityNominated(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetTotalCapacityNominated(
		BusinessTypeTotalNominatedCapacity,
		DomainCZ,
		DomainSK,
		genTime("201601012300"),
		genTime("201601022300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.9. Total Capacity Already Allocated [12.1.C]
func TestGetTotalCapacityAlreadyAllocated(t *testing.T) {
	t.Skip("TODO: always returns \"The combination of [DocumentType=A26,BusinessType=A29] is not valid, or the requested data is not allowed to be fetched via this service.\"")
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetTotalCapacityAlreadyAllocated(
		BusinessTypeAlreadyAllocatedCapacity,
		ContractMarketAgreementTypeIntraday,
		DomainSK,
		DomainCZ,
		genTime("201601012300"),
		genTime("201601022300"),
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.10. Day Ahead Prices [12.1.D]
func TestGetDayAheadPrices(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetDayAheadPrices(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.11. Implicit Auction — Net Positions [12.1.E]
func TestGetImplicitAuctionNetPositions(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetImplicitAuction(
		BusinessTypeNetPosition,
		ContractMarketAgreementTypeDaily,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.12. Implicit Auction — Congestion Income [12.1.E]
func TestGetImplicitAuctionCongestionIncome(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetImplicitAuction(
		BusinessTypeCongestionIncome,
		ContractMarketAgreementTypeDaily,
		"10YDOM-1001A083J",
		genTime("201601012300"),
		genTime("201601022300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.13. Total Commercial Schedules [12.1.F]
func TestGetTotalCommercialSchedules(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetTotalCommercialSchedules(
		DomainCZ,
		DomainSK,
		genTime("201512312300"),
		genTime("201612312300"),
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.14. Day-ahead Commercial Schedules [12.1.F]
func TestGetDayAheadCommercialSchedules(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetDayAheadCommercialSchedules(
		DomainCZ,
		DomainSK,
		genTime("201601012300"),
		genTime("201601022300"),
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.15. Physical Flows [12.1.G]
func TestGetPhysicalFlows(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetPhysicalFlows(
		DomainCZ,
		DomainSK,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.2.16. Capacity Allocated Outside EU [12.1.H]
func TestGetCapacityAllocatedOutsideEu(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	auctionCategory := AuctionCategoryHourly
	doc, err := c.GetCapacityAllocatedOutsideEu(
		AuctionTypeExplicit,
		ContractMarketAgreementTypeDaily,
		DomainSK,
		DomainUA,
		genTime("201601012300"),
		genTime("201601022300"),
		&auctionCategory,
		pointy.Int(1),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.3. Congestion domain

// 4.3.1. Redispatching [13.1.A]
func TestGetRedispatching(t *testing.T) {
	t.Skip("TODO: always returns \"The combination of [DocumentType=A63] is not valid, or the requested data is not allowed to be fetched via this service.\"")
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetRedispatching(
		DomainCZ,
		DomainSK,
		genTime("201512312300"),
		genTime("201612312300"),
		nil,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.3.2. Countertrading [13.1.B]
func TestGetCountertrading(t *testing.T) {
	t.Skip("TODO: always returns \"No matching data found for Data item Countertrading [13.1.B]\"")
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetCountertrading(
		DomainCZ,
		DomainSK,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.3.3. Costs of Congestion Management [13.1.C]
func TestGetCostsOfCongestionManagement(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	businessType := BusinessTypeCounterTrade
	doc, err := c.GetCostsOfCongestionManagement(
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
		&businessType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4. Generation domain

// 4.4.1. Installed Generation Capacity Aggregated [14.1.A]
func TestGetInstalledGenerationCapacityAggregated(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	psrType := PsrTypeSolar
	doc, err := c.GetInstalledGenerationCapacityAggregated(
		ProcessTypeYearAhead,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
		&psrType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.2. Installed Generation Capacity per Unit [14.1.B]
func TestGetInstalledGenerationCapacityPerUnit(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	psrType := PsrTypeFossilBrownCoalLignite
	doc, err := c.GetInstalledGenerationCapacityPerUnit(
		ProcessTypeYearAhead,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
		&psrType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.3. Day-ahead Aggregated Generation [14.1.C]
func TestGetDayAheadAggregatedGeneration(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetDayAheadAggregatedGeneration(
		ProcessTypeDayAhead,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.4. Day-ahead Generation Forecasts for Wind and Solar [14.1.D]
func TestDayAheadGenerationForecastsForWindAndSolar(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	psrType := PsrTypeSolar
	doc, err := c.GetGenerationForecastsForWindAndSolar(
		ProcessTypeDayAhead,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
		&psrType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.5. Current Generation Forecasts for Wind and Solar [14.1.D]
func TestCurrentGenerationForecastsForWindAndSolar(t *testing.T) {
	t.Skip("TODO: always returns \"No matching data found for Data item Current Generation Forecasts for Wind and Solar [14.1.D]\"")
	c := NewEntsoeClientFromEnv()
	psrType := PsrTypeSolar
	doc, err := c.GetGenerationForecastsForWindAndSolar(
		ProcessTypeIntradayTotal,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
		&psrType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.6. Intraday Generation Forecasts for Wind and Solar [14.1.D]
func TestIntradayGenerationForecastsForWindAndSolar(t *testing.T) {
	t.Skip("TODO: always returns \"No matching data found for Data item Current Generation Forecasts for Wind and Solar [14.1.D]\"")
	c := NewEntsoeClientFromEnv()
	psrType := PsrTypeSolar
	doc, err := c.GetGenerationForecastsForWindAndSolar(
		ProcessTypeIntradayProcess,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
		&psrType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.7. Actual Generation Output per Generation Unit [16.1.A]
func TestActualGenerationOutputPerGenerationUnit(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	psrType := PsrTypeFossilBrownCoalLignite
	doc, err := c.GetActualGenerationOutputPerGenerationUnit(
		ProcessTypeRealised,
		DomainCZ,
		genTime("201512302300"),
		genTime("201512312300"),
		&psrType,
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.8. Aggregated Generation per Type [16.1.B&C]
func TestAggregatedGenerationPerType(t *testing.T) {
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetAggregatedGenerationPerType(
		ProcessTypeRealised,
		PsrTypeFossilBrownCoalLignite,
		DomainCZ,
		genTime("201512302300"),
		genTime("201512312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

// 4.4.9. Aggregated Filling Rate of Water Reservoirs and Hydro Storage Plants [16.1.D]
func TestAggregatedFillingRateOfWaterReservoirsAndHydroStoragePlants(t *testing.T) {
	t.Skip("TODO: always returns \"No matching data found for Data item Aggregate Filling Rate of Water Reservoirs and Hydro Storage Plants [16.1.D]\"")
	c := NewEntsoeClientFromEnv()
	doc, err := c.GetAggregatedFillingRateOfWaterReservoirsAndHydroStoragePlants(
		ProcessTypeRealised,
		DomainCZ,
		genTime("201512312300"),
		genTime("201612312300"),
	)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
}

func genTime(timeString string) time.Time {
	t, err := time.Parse("200601021504", timeString)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
