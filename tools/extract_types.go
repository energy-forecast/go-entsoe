package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/sirupsen/logrus"
)

/*
- all Acknowledgement_MarketDocument responses return error code 999
*/

// search: ^(\s+//\s(([0-9]|\.)+)\s.*\n)(\s+)//\s(/api.*)$
// replace: $1$4"$2": "$5",

const GEN_GO_TYPES_FILE = "types.go"

var sampleRequests = map[string]string{
	// 4.1.1. Actual Total Load [6.1.A]
	"4.1.1.": "documentType=A65&processType=A16&outBiddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.1.2. Day-Ahead Total Load Forecast [6.1.B]
	"4.1.2.": "documentType=A65&processType=A01&outBiddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.1.3. Week-Ahead Total Load Forecast [6.1.C]
	"4.1.3.": "documentType=A65&processType=A31&outBiddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.1.4. Month-Ahead Total Load Forecast [6.1.D]
	"4.1.4.": "documentType=A65&processType=A32&outBiddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.1.5. Year-Ahead Total Load Forecast [6.1.E]
	"4.1.5.": "documentType=A65&processType=A33&outBiddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.1.6. Year-Ahead Forecast Margin [8.1]
	"4.1.6.": "documentType=A70&processType=A33&outBiddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.1. Expansion and Dismantling Projects [9.1]
	"4.2.1.": "documentType=A90&businessType=B01&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.2. Forecasted Capacity [11.1.A]
	"4.2.2.": "documentType=A61&contract_MarketAgreement.Type=A01&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.3. Offered Capacity [11.1.A]
	"4.2.3.": "documentType=A31&contract_MarketAgreement.Type=A01&in_Domain=10YSK-SEPS-----K&out_Domain=10YCZ-CEPS-----N&auction.Type=A01&periodStart=201601012300&periodEnd=201601022300",

	// 4.2.4. Flow-based Parameters [11.1.B]
	"4.2.4.": "documentType=B11&processType=A01&in_Domain=10YDOM-REGION-1V&out_Domain=10YDOM-REGION-1V&periodStart=201512312300&periodEnd=201601012300",

	// 4.2.5. Intraday Transfer Limits [11.3]
	"4.2.5.": "documentType=A93&in_Domain=10YFR-RTE------C&out_Domain=10YGB----------A&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.6. Explicit Allocation Information (Capacity) [12.1.A]
	"4.2.6.": "documentType=A25&businessType=B05&contract_MarketAgreement.Type=A01&in_Domain=10YSK-SEPS-----K&out_Domain=10YCZ-CEPS-----N&auction.Category=A01&classificationSequence_AttributeInstanceComponent.Position=1&periodStart=201601012300&periodEnd=201601022300",

	// 4.2.7. Explicit Allocation Information (Revenue only) [12.1.A]
	"4.2.7.": "documentType=A25&businessType=B07&contract_MarketAgreement.Type=A01&in_Domain=10YAT-APG------L&out_Domain=10YCZ-CEPS-----N&periodStart=201601012300&periodEnd=201601022300",

	// 4.2.8. Total Capacity Nominated [12.1.B]
	"4.2.8.": "documentType=A26&businessType=B08&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.9. Total Capacity Already Allocated [12.1.C]
	"4.2.9.": "documentType=A26&businessType=A29&contract_MarketAgreement.Type=A07&in_Domain=10YSK-SEPS-----K&out_Domain=10YCZ-CEPS-----N&periodStart=201601012300&periodEnd=201601022300",

	// 4.2.10. Day Ahead Prices [12.1.D]
	"4.2.10.": "documentType=A44&in_Domain=10YCZ-CEPS-----N&out_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.11. Implicit Auction — Net Positions [12.1.E]
	"4.2.11.": "documentType=A25&businessType=B09&contract_MarketAgreement.Type=A01&in_Domain=10YCZ-CEPS-----N&out_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.12. Implicit Auction — Congestion Income [12.1.E]
	"4.2.12.": "documentType=A25&businessType=B10&contract_MarketAgreement.Type=A01&in_Domain=10YDOM-1001A083J&out_Domain=10YDOM-1001A083J&periodStart=201601012300&periodEnd=201601022300",

	// 4.2.13. Total Commercial Schedules [12.1.F]
	"4.2.13.": "documentType=A09&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&contract_MarketAgreement.Type=A05&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.14. Day-ahead Commercial Schedules [12.1.F]
	"4.2.14.": "documentType=A09&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&contract_MarketAgreement.Type=A01&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.15. Physical Flows [12.1.G]
	"4.2.15.": "documentType=A11&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.2.16. Capacity Allocated Outside EU [12.1.H]
	"4.2.16.": "documentType=A94&contract_MarketAgreement.Type=A01&in_Domain=10YSK-SEPS-----K&out_Domain=10YUA-WEPS-----0&auction.Type=A02&auction.Category=A04&classificationSequence_AttributeInstanceComponent.Position=1&periodStart=201601012300&periodEnd=201601022300",

	// 4.3.1. Redispatching [13.1.A]
	"4.3.1.": "documentType=A63&businessType=A46&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.3.2. Countertrading [13.1.B]
	"4.3.2.": "documentType=A91&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.3.3. Costs of Congestion Management [13.1.C]
	"4.3.3.": "documentType=A92&businessType=B03&in_Domain=10YCZ-CEPS-----N&out_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.1. Installed Generation Capacity Aggregated [14.1.A]
	"4.4.1.": "documentType=A68&processType=A33&psrType=B16&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.2. Installed Generation Capacity per Unit [14.1.B]
	"4.4.2.": "documentType=A71&processType=A33&psrType=B02&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.3. Day-ahead Aggregated Generation [14.1.C]
	"4.4.3.": "documentType=A71&processType=A01&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.4. Day-ahead Generation Forecasts for Wind and Solar [14.1.D]
	"4.4.4.": "documentType=A69&processType=A01&psrType=B16&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.5. Current Generation Forecasts for Wind and Solar [14.1.D]
	"4.4.5.": "documentType=A69&processType=A18&psrType=B16&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.6. Intraday Generation Forecasts for Wind and Solar [14.1.D]
	"4.4.6.": "documentType=A69&processType=A40&psrType=B16&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.7. Actual Generation Output per Generation Unit [16.1.A]
	"4.4.7.": "documentType=A73&processType=A16&psrType=B02&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201601012300",

	// 4.4.8. Aggregated Generation per Type [16.1.B&C]
	"4.4.8.": "documentType=A75&processType=A16&psrType=B02&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.4.9. Aggregated Filling Rate of Water Reservoirs and Hydro Storage Plants [16.1.D]
	"4.4.9.": "documentType=A72&processType=A16&in_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.5.1. Production and Generation Units
	// problem with illegal UTF-8 encoding if examples are created by zek
	//"4.5.1.": "documentType=A95&businessType=B11&biddingZone_Domain=10YCZ-CEPS-----N&implementation_DateAndOrTime=2017-01-01",

	// 4.6.1. Current Balancing State [GL EB 12.3.A]
	"4.6.1.": "documentType=A86&businessType=B33&Area_Domain=10YCZ-CEPS-----N&periodStart=201912190000&periodEnd=201912190010",

	// 4.6.2. Aggregated Balancing Energy Bids [GL EB 12.3.E]
	"4.6.2.": "documentType=A24&processType=A51&area_Domain=10YCZ-CEPS-----N&TimeInterval=2019-12-16T13:00Z/2019-12-16T18:00Z",

	// 4.6.3. Prices of Activated Balancing Energy [GL EB 12.3.F]
	"4.6.3.": "documentType=A15&processType=A51&area_Domain=10YCZ-CEPS-----N&TimeInterval=2019-12-31T23:00Z/2020-01-01T00:00Z",

	// 4.6.4. Use of Allocated Cross-Zonal Balancing Capacity [GL EB 12.3.H&I]
	"4.6.4.": "documentType=A38&processType=A46&Acquiring_Domain=10YAT-APG------L&Connecting_Domain=10YCH-SWISSGRIDZ&TimeInterval=2019-12-16T00:00Z/2019-12-17T00:00Z",

	// 4.6.5. Amount of Balancing Reserves Under Contract [17.1.B]
	"4.6.5.": "documentType=A81&type_MarketAgreement.Type=A13&businessType=A95&psrType=A04&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201601012300",

	// 4.6.6. Prices of Procured Balancing Reserves [17.1.C]
	"4.6.6.": "documentType=A89&type_MarketAgreement.Type=A01&businessType=A96&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201601012300",

	// 4.6.7. Accepted Aggregated Offers [17.1.D]
	"4.6.7.": "documentType=A82&businessType=A95&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.6.8. Activated Balancing Energy [17.1.E]
	"4.6.8.": "documentType=A83&businessType=A96&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.6.9. Prices of Activated Balancing Energy [17.1.F]
	"4.6.9.": "documentType=A84&businessType=A96&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.6.10. Imbalance Prices [17.1.G]
	"4.6.10.": "documentType=A85&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.6.11. Total Imbalance Volumes [17.1.H]
	"4.6.11.": "documentType=A86&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.6.12. Financial Expenses and Income for Balancing [17.1.I]
	"4.6.12.": "documentType=A87&controlArea_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.6.13. Cross-border Balancing [17.1.J]
	"4.6.13.": "documentType=A88&acquiring_Domain=10YCZ-CEPS-----N&connecting_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201601010100",

	// 4.6.14. FCR Total capacity [SO GL 187.2]
	"4.6.14.": "documentType=A26&area_Domain=10YEU-CONT-SYNC0&businessType=A25&TimeInterval=2018-12-31T23:00Z/2019-12-31T23:00Z",

	// 4.6.15. Shares of FCR capacity - share of capacity [SO GL 187.2]
	"4.6.15.": "documentType=A26&area_Domain=10YDE-VE-------2&businessType=C23&TimeInterval=2019-12-31T23:00Z/2020-12-31T23:00Z",

	// 4.6.16. Shares of FCR capacity - contracted reserve capacity [SO GL 187.2]
	"4.6.16.": "documentType=A26&businessType=B95&TimeInterval=2019-12-31T23:00Z/2020-12-31T23:00Z&Area_Domain=10YDE-RWENET---I",

	// 4.6.17. FRR Actual Capacity [SO GL 188.4]
	"4.6.17.": "documentType=A26&processType=A56&businessType=C24&Area_Domain=10YAT-APG------L&TimeInterval=2019-12-31T23:00Z/2020-03-31T22:00Z",

	// 4.6.18. RR Actual Capacity [SO GL 189.3]
	"4.6.18.": "documentType=A26&processType=A46&businessType=C24&Area_Domain=10YAT-APG------L&TimeInterval=2019-12-31T23:00Z/2020-03-31T22:00Z",

	// 4.6.19. Sharing of RR and FRR [SO GL 190.1]
	"4.6.19.": "documentType=A26&businessType=C22&TimeInterval=2019-12-31T23:00Z/2020-12-31T23:00Z&Connecting_Domain=10YAT-APG------L&Acquiring_Domain=10YCB-GERMANY--8&processType=A56",

	// 4.7.1. Unavailability of Consumption Units [7.1A&B]
	"4.7.1.": "documentType=A76&biddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.7.2. Unavailability of Transmission Infrastructure [10.1.A&B]
	"4.7.2.": "documentType=A78&businessType=A53&in_Domain=10YCZ-CEPS-----N&out_Domain=10YSK-SEPS-----K&periodStart=201512312300&periodEnd=201612312300",

	// 4.7.3. Unavailability of Offshore Grid Infrastructure [10.1.C]
	"4.7.3.": "documentType=A79&biddingZone_Domain=10YDE-EON------1&periodStart=201512312300&periodEnd=201612312300",

	// 4.7.4. Unavailability of Generation Units [15.1.A&B]
	"4.7.4.": "documentType=A80&businessType=A53&biddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",

	// 4.7.5. Unavailability of Production Units [15.1.C&D]
	"4.7.5.": "documentType=A77&businessType=A53&biddingZone_Domain=10YCZ-CEPS-----N&periodStart=201512312300&periodEnd=201612312300",
}
var log = logrus.New()

// regex to extract response document
var re = regexp.MustCompile(`<(.+?)\s+xmlns=.+?>`)

func main() {

	apiKey := os.Getenv("ENTSOE_API_KEY")
	if apiKey == "" {
		log.Fatal("Environment variable ENTSOE_API_KEY with api key not set")
	}

	createdDirs := make(map[string]bool)

	keys := make([]string, 0, len(sampleRequests))
	for k := range sampleRequests {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := sampleRequests[key]
		log.Info("processing request " + key)
		resp, err := http.Get("https://transparency.entsoe.eu/api?securityToken=" + apiKey + "&" + value)
		if err != nil {
			log.Fatal(err)
		}
		body := resp.Body
		defer body.Close()
		bodyBytes, err := ioutil.ReadAll(body)
		if err != nil {
			log.Fatal(err)
		}

		contentType := resp.Header.Get("Content-type")
		if contentType == "application/zip" {
			zipReader, err := zip.NewReader(bytes.NewReader(bodyBytes), int64(len(bodyBytes)))
			if err != nil {
				log.Fatal(err)
			}

			// read all the files from zip archive
			for _, zipFile := range zipReader.File {
				log.Infof("reading file %s from zipped content", zipFile.Name)
				unzippedFileBytes, err := readZipFile(zipFile)
				if err != nil {
					log.Fatal(err)
				}
				documentType := processFileContent(key+zipFile.Name, unzippedFileBytes)
				createdDirs[documentType] = true
			}
		} else {
			// content type is "text/xml", "application/xml" or just missing

			documentType := processFileContent(key+".xml", bodyBytes)
			createdDirs[documentType] = true
		}
	}
	log.Info("Run zek...")

	genTypeFile := filepath.Join("..", GEN_GO_TYPES_FILE)
	if _, err := os.Stat(genTypeFile); err == nil {
		os.Remove(GEN_GO_TYPES_FILE)
	}
	f, err := os.OpenFile(genTypeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("package goentsoe\n\n"))
	f.Write([]byte("import \"encoding/xml\"\n\n"))

	for dir := range createdDirs {
		shell := "bash"

		cmdOutput := &bytes.Buffer{}
		cmd := &exec.Cmd{
			Path:   shell,
			Args:   append([]string{shell}, "-c", "zek -e *.xml"),
			Dir:    dir,
			Stdout: cmdOutput,
		}
		if filepath.Base(shell) == shell {
			if lp, err := exec.LookPath(shell); err != nil {
				log.Fatal(err)
			} else {
				cmd.Path = lp
			}
		}

		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		f.Write(cmdOutput.Bytes())
	}
	log.Info("FINISHED")
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func processFileContent(fileName string, content []byte) string {
	matches := re.FindStringSubmatch(string(content))
	if matches == nil || len(matches) != 2 {
		fmt.Println(string(content))
		log.Fatal("Detect more than one document")
	}
	documentType := matches[1]
	os.MkdirAll(documentType, os.ModePerm)

	err := ioutil.WriteFile(filepath.Join(documentType, fileName), content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return documentType
}
