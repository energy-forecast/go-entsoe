package goentsoe

import "encoding/xml"

type TransmissionNetworkMarketDocument struct {
	XMLName                     xml.Name `xml:"TransmissionNetwork_MarketDocument"`
	Text                        string   `xml:",chardata"`
	Xmlns                       string   `xml:"xmlns,attr"`
	MRID                        string   `xml:"mRID"`                // 792709571696419ab4cbb2dfa...
	RevisionNumber              string   `xml:"revisionNumber"`      // 1
	Type                        string   `xml:"type"`                // A92
	ProcessProcessType          string   `xml:"process.processType"` // A16
	CreatedDateTime             string   `xml:"createdDateTime"`     // 2020-08-08T19:54:43Z
	SenderMarketParticipantMRID struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"sender_MarketParticipant.mRID"`
	SenderMarketParticipantMarketRoleType string `xml:"sender_MarketParticipant.marketRole.type"` // A32
	ReceiverMarketParticipantMRID         struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"receiver_MarketParticipant.mRID"`
	ReceiverMarketParticipantMarketRoleType string `xml:"receiver_MarketParticipant.marketRole.type"` // A33
	PeriodTimeInterval                      struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start"` // 2016-01-01T00:00Z
		End   string `xml:"end"`   // 2017-01-01T00:00Z
	} `xml:"period.timeInterval"`
	TimeSeries []struct {
		Text         string `xml:",chardata"`
		MRID         string `xml:"mRID"`         // 1, 2, 3, 4, 5, 6, 7, 8, 9...
		BusinessType string `xml:"businessType"` // B03, B03, B03, B03, B03, ...
		InDomainMRID struct {
			Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"in_Domain.mRID"`
		OutDomainMRID struct {
			Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"out_Domain.mRID"`
		CurveType string `xml:"curveType"` // A01, A01, A01, A01, A01, ...
		Period    struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start"` // 2016-01-01T00:00Z, 2016-0...
				End   string `xml:"end"`   // 2016-02-01T00:00Z, 2016-0...
			} `xml:"timeInterval"`
			Resolution string `xml:"resolution"` // P1M, P1M, P1M, P1M, P1M, ...
			Point      struct {
				Text     string `xml:",chardata"`
				Position string `xml:"position"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			} `xml:"Point"`
		} `xml:"Period"`
	} `xml:"TimeSeries"`
}

type BalancingMarketDocument struct {
	XMLName                     xml.Name `xml:"Balancing_MarketDocument"`
	Text                        string   `xml:",chardata"`
	Xmlns                       string   `xml:"xmlns,attr"`
	MRID                        string   `xml:"mRID"`                // 267de175a6e8476882c72c380...
	RevisionNumber              string   `xml:"revisionNumber"`      // 1, 1, 1, 1, 1, 1, 1, 1, 1...
	Type                        string   `xml:"type"`                // A86, A85, A86, A86, A87, ...
	ProcessProcessType          string   `xml:"process.processType"` // A16, A16, A16, A16, A16, ...
	SenderMarketParticipantMRID struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"sender_MarketParticipant.mRID"`
	SenderMarketParticipantMarketRoleType string `xml:"sender_MarketParticipant.marketRole.type"` // A32, A32, A32, A32, A32, ...
	ReceiverMarketParticipantMRID         struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"receiver_MarketParticipant.mRID"`
	ReceiverMarketParticipantMarketRoleType string `xml:"receiver_MarketParticipant.marketRole.type"` // A33, A33, A33, A33, A33, ...
	CreatedDateTime                         string `xml:"createdDateTime"`                            // 2020-08-08T19:54:50Z, 202...
	AreaDomainMRID                          struct {
		Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YAT-A...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"area_Domain.mRID"`
	PeriodTimeInterval struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start"` // 2019-12-19T00:00Z, 2015-1...
		End   string `xml:"end"`   // 2019-12-19T00:10Z, 2016-1...
	} `xml:"period.timeInterval"`
	TimeSeries []struct {
		Text                    string `xml:",chardata"`
		MRID                    string `xml:"mRID"`                       // 1, 1, 2, 3, 4, 5, 6, 7, 8...
		BusinessType            string `xml:"businessType"`               // B33, A19, A19, A19, A19, ...
		FlowDirectionDirection  string `xml:"flowDirection.direction"`    // A02, A01, A02, A01, A01, ...
		QuantityMeasureUnitName string `xml:"quantity_Measure_Unit.name"` // MAW, MWH, MWH, MWH, MWH, ...
		CurveType               string `xml:"curveType"`                  // A01, A01, A01, A01, A01, ...
		Period                  struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start"` // 2019-12-19T00:00Z, 2015-1...
				End   string `xml:"end"`   // 2019-12-19T00:10Z, 2016-0...
			} `xml:"timeInterval"`
			Resolution string `xml:"resolution"` // PT1M, PT60M, PT60M, PT60M...
			Point      []struct {
				Text                   string `xml:",chardata"`
				Position               string `xml:"position"`                 // 1, 2, 3, 4, 5, 6, 7, 8, 9...
				Quantity               string `xml:"quantity"`                 // 78.39, 75.53, 59.41, 61.2...
				ImbalancePriceAmount   string `xml:"imbalance_Price.amount"`   // -514, -562, -548, -314, -...
				ImbalancePriceCategory string `xml:"imbalance_Price.category"` // A04, A04, A04, A04, A04, ...
				FinancialPrice         []struct {
					Text      string `xml:",chardata"`
					Amount    string `xml:"amount"`    // 464071702, 39706489, 4416...
					Direction string `xml:"direction"` // A01, A02, A01, A02, A01, ...
				} `xml:"Financial_Price"`
				SecondaryQuantity      string `xml:"secondaryQuantity"`        // 0, 0, 0, 18, 0, 0, 0, 27,...
				ProcurementPriceAmount string `xml:"procurement_Price.amount"` // 605, 672, 755, 781, 686, ...
				ActivationPriceAmount  string `xml:"activation_Price.amount"`  // 0, 0, 0, 0, 0, 0, 0, 0, 0...
			} `xml:"Point"`
		} `xml:"Period"`
		CurrencyUnitName                       string `xml:"currency_Unit.name"`                       // CZK, CZK, CZK, CZK, CZK, ...
		PriceMeasureUnitName                   string `xml:"price_Measure_Unit.name"`                  // MWH, MWH, MWH, MWH, MWH, ...
		StandardMarketProductMarketProductType string `xml:"standard_MarketProduct.marketProductType"` // A01, A01
		TypeMarketAgreementType                string `xml:"type_MarketAgreement.type"`                // A01, A01, A01, A01
		MktPSRTypePsrType                      string `xml:"mktPSRType.psrType"`                       // A04, A04, A04, A04, A04, ...
	} `xml:"TimeSeries"`
	ControlAreaDomainMRID struct {
		Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"controlArea_Domain.mRID"`
	DocStatus struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value"` // A01, A02, A01
	} `xml:"docStatus"`
}

type UnavailabilityMarketDocument struct {
	XMLName                     xml.Name `xml:"Unavailability_MarketDocument"`
	Text                        string   `xml:",chardata"`
	Xmlns                       string   `xml:"xmlns,attr"`
	MRID                        string   `xml:"mRID"`                // kjfQMlmtlZVC32VliNsNQg, -...
	RevisionNumber              string   `xml:"revisionNumber"`      // 2, 1, 3, 3, 3, 3, 3, 2, 2...
	Type                        string   `xml:"type"`                // A79, A79, A79, A79, A79, ...
	ProcessProcessType          string   `xml:"process.processType"` // A26, A26, A26, A26, A26, ...
	CreatedDateTime             string   `xml:"createdDateTime"`     // 2016-05-13T06:19:51Z, 201...
	SenderMarketParticipantMRID struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"sender_MarketParticipant.mRID"`
	SenderMarketParticipantMarketRoleType string `xml:"sender_MarketParticipant.marketRole.type"` // A32, A32, A32, A32, A32, ...
	ReceiverMarketParticipantMRID         struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"receiver_MarketParticipant.mRID"`
	ReceiverMarketParticipantMarketRoleType string `xml:"receiver_MarketParticipant.marketRole.type"` // A33, A33, A33, A33, A33, ...
	UnavailabilityTimePeriodTimeInterval    struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start"` // 2015-11-23T17:50Z, 2015-1...
		End   string `xml:"end"`   // 2016-05-12T19:51Z, 2016-0...
	} `xml:"unavailability_Time_Period.timeInterval"`
	TimeSeries struct {
		Text                  string `xml:",chardata"`
		MRID                  string `xml:"mRID"`         // 1, 1, 1, 1, 1, 1, 1, 1, 1...
		BusinessType          string `xml:"businessType"` // A54, A54, A54, A54, A54, ...
		BiddingZoneDomainMRID struct {
			Text         string `xml:",chardata"` // 10YDE-EON------1, 10YDE-E...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"biddingZone_Domain.mRID"`
		StartDateAndOrTimeDate                                          string `xml:"start_DateAndOrTime.date"`   // 2015-11-23, 2015-12-29, 2...
		StartDateAndOrTimeTime                                          string `xml:"start_DateAndOrTime.time"`   // 17:50:00Z, 19:43:00Z, 07:...
		EndDateAndOrTimeDate                                            string `xml:"end_DateAndOrTime.date"`     // 2016-05-12, 2016-01-05, 2...
		EndDateAndOrTimeTime                                            string `xml:"end_DateAndOrTime.time"`     // 19:51:00Z, 19:43:00Z, 16:...
		QuantityMeasureUnitName                                         string `xml:"quantity_Measure_Unit.name"` // MAW, MAW, MAW, MAW, MAW, ...
		CurveType                                                       string `xml:"curveType"`                  // A03, A03, A03, A03, A03, ...
		ProductionRegisteredResourcePSRTypePowerSystemResourcesNominalP struct {
			Text string `xml:",chardata"` // 113, 156, 864, 144, 144, ...
			Unit string `xml:"unit,attr"`
		} `xml:"production_RegisteredResource.pSRType.powerSystemResources.nominalP"`
		AssetRegisteredResource struct {
			Text string `xml:",chardata"`
			MRID struct {
				Text         string `xml:",chardata"` // 11TD2L000000267O, 11TD2L0...
				CodingScheme string `xml:"codingScheme,attr"`
			} `xml:"mRID"`
			Name                string `xml:"name"`                  // L-155-RIFF-EMDB-AC114, L-...
			AssetPSRTypePsrType string `xml:"asset_PSRType.psrType"` // B21, B21, B23, B21, B21, ...
			LocationName        string `xml:"location.name"`         // Riffgat-Emden/Borssum, Bo...
		} `xml:"Asset_RegisteredResource"`
		WindPowerFeedinPeriod struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start"` // 2015-11-23T17:50Z, 2015-1...
				End   string `xml:"end"`   // 2016-05-12T19:51Z, 2016-0...
			} `xml:"timeInterval"`
			Resolution string `xml:"resolution"` // PT1M, PT1M, PT1M, PT1M, P...
			Point      struct {
				Text     string `xml:",chardata"`
				Position string `xml:"position"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
				Quantity string `xml:"quantity"` // 0, 80, 545, 142, 141, 141...
			} `xml:"Point"`
		} `xml:"WindPowerFeedin_Period"`
	} `xml:"TimeSeries"`
	Reason struct {
		Text string `xml:",chardata"`
		Code string `xml:"code"` // B18, B18, B18, B18, B18, ...
	} `xml:"Reason"`
}

type GLMarketDocument struct {
	XMLName                     xml.Name `xml:"GL_MarketDocument"`
	Text                        string   `xml:",chardata"`
	Xmlns                       string   `xml:"xmlns,attr"`
	MRID                        string   `xml:"mRID"`                // 04930bc57afa4a71b474d7a51...
	RevisionNumber              string   `xml:"revisionNumber"`      // 1, 1, 1, 1, 1, 1, 1, 1, 1...
	Type                        string   `xml:"type"`                // A65, A65, A65, A65, A65, ...
	ProcessProcessType          string   `xml:"process.processType"` // A16, A01, A31, A32, A33, ...
	SenderMarketParticipantMRID struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"sender_MarketParticipant.mRID"`
	SenderMarketParticipantMarketRoleType string `xml:"sender_MarketParticipant.marketRole.type"` // A32, A32, A32, A32, A32, ...
	ReceiverMarketParticipantMRID         struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"receiver_MarketParticipant.mRID"`
	ReceiverMarketParticipantMarketRoleType string `xml:"receiver_MarketParticipant.marketRole.type"` // A33, A33, A33, A33, A33, ...
	CreatedDateTime                         string `xml:"createdDateTime"`                            // 2020-08-08T19:54:20Z, 202...
	TimePeriodTimeInterval                  struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start"` // 2015-12-31T23:00Z, 2015-1...
		End   string `xml:"end"`   // 2016-12-31T23:00Z, 2016-1...
	} `xml:"time_Period.timeInterval"`
	TimeSeries []struct {
		Text                     string `xml:",chardata"`
		MRID                     string `xml:"mRID"`              // 1, 2, 3, 4, 5, 6, 7, 8, 9...
		BusinessType             string `xml:"businessType"`      // A04, A04, A04, A04, A04, ...
		ObjectAggregation        string `xml:"objectAggregation"` // A01, A01, A01, A01, A01, ...
		OutBiddingZoneDomainMRID struct {
			Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"outBiddingZone_Domain.mRID"`
		QuantityMeasureUnitName string `xml:"quantity_Measure_Unit.name"` // MAW, MAW, MAW, MAW, MAW, ...
		CurveType               string `xml:"curveType"`                  // A01, A01, A01, A01, A01, ...
		Period                  struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start"` // 2015-12-31T23:00Z, 2016-0...
				End   string `xml:"end"`   // 2016-01-23T11:00Z, 2016-0...
			} `xml:"timeInterval"`
			Resolution string `xml:"resolution"` // PT60M, PT60M, PT60M, PT60...
			Point      []struct {
				Text     string `xml:",chardata"`
				Position string `xml:"position"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
				Quantity string `xml:"quantity"` // 5872, 5784, 5690, 5604, 5...
			} `xml:"Point"`
		} `xml:"Period"`
		InBiddingZoneDomainMRID struct {
			Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"inBiddingZone_Domain.mRID"`
		MktPSRType struct {
			Text                                        string `xml:",chardata"`
			PsrType                                     string `xml:"psrType"` // B16, B02, B02, B02, B02, ...
			VoltagePowerSystemResourcesHighVoltageLimit struct {
				Text string `xml:",chardata"` // 400, 400, 400, 400, 110, ...
				Unit string `xml:"unit,attr"`
			} `xml:"voltage_PowerSystemResources.highVoltageLimit"`
			PowerSystemResources struct {
				Text string `xml:",chardata"`
				MRID struct {
					Text         string `xml:",chardata"` // 27W-GU-ECHVG1--C, 27W-GU-...
					CodingScheme string `xml:"codingScheme,attr"`
				} `xml:"mRID"`
				Name string `xml:"name"` // ECHV_G1____, ECHV_G2____,...
			} `xml:"PowerSystemResources"`
		} `xml:"MktPSRType"`
		RegisteredResourceMRID struct {
			Text         string `xml:",chardata"` // 27W-PU-EPC1----Y, 27W-PU-...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"registeredResource.mRID"`
		RegisteredResourceName string `xml:"registeredResource.name"` // EPC1_______, EME3_______,...
	} `xml:"TimeSeries"`
}

type AcknowledgementMarketDocument struct {
	XMLName                     xml.Name `xml:"Acknowledgement_MarketDocument"`
	Text                        string   `xml:",chardata"`
	Xmlns                       string   `xml:"xmlns,attr"`
	MRID                        string   `xml:"mRID"`            // 281cd746-add3-4, 39345301...
	CreatedDateTime             string   `xml:"createdDateTime"` // 2020-08-08T19:54:23Z, 202...
	SenderMarketParticipantMRID struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"sender_MarketParticipant.mRID"`
	SenderMarketParticipantMarketRoleType string `xml:"sender_MarketParticipant.marketRole.type"` // A32, A32, A32, A32, A32, ...
	ReceiverMarketParticipantMRID         struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"receiver_MarketParticipant.mRID"`
	ReceiverMarketParticipantMarketRoleType string `xml:"receiver_MarketParticipant.marketRole.type"` // A39, A39, A39, A39, A39, ...
	ReceivedMarketDocumentCreatedDateTime   string `xml:"received_MarketDocument.createdDateTime"`    // 2020-08-08T19:54:23Z, 202...
	Reason                                  struct {
		Chardata string `xml:",chardata"`
		Code     string `xml:"code"` // 999, 999, 999, 999, 999, ...
		Text     string `xml:"text"` // No matching data found fo...
	} `xml:"Reason"`
}

type PublicationMarketDocument struct {
	XMLName                     xml.Name `xml:"Publication_MarketDocument"`
	Text                        string   `xml:",chardata"`
	Xmlns                       string   `xml:"xmlns,attr"`
	MRID                        string   `xml:"mRID"`           // afc06ccfb7be4cca9880679a7...
	RevisionNumber              string   `xml:"revisionNumber"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
	Type                        string   `xml:"type"`           // A44, A25, A25, A09, A11, ...
	SenderMarketParticipantMRID struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"sender_MarketParticipant.mRID"`
	SenderMarketParticipantMarketRoleType string `xml:"sender_MarketParticipant.marketRole.type"` // A32, A32, A32, A32, A32, ...
	ReceiverMarketParticipantMRID         struct {
		Text         string `xml:",chardata"` // 10X1001A1001A450, 10X1001...
		CodingScheme string `xml:"codingScheme,attr"`
	} `xml:"receiver_MarketParticipant.mRID"`
	ReceiverMarketParticipantMarketRoleType string `xml:"receiver_MarketParticipant.marketRole.type"` // A33, A33, A33, A33, A33, ...
	CreatedDateTime                         string `xml:"createdDateTime"`                            // 2020-08-08T19:54:24Z, 202...
	PeriodTimeInterval                      struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start"` // 2015-12-31T23:00Z, 2015-1...
		End   string `xml:"end"`   // 2016-12-31T23:00Z, 2016-1...
	} `xml:"period.timeInterval"`
	TimeSeries []struct {
		Text         string `xml:",chardata"`
		MRID         string `xml:"mRID"`         // 1, 2, 3, 4, 5, 6, 7, 8, 9...
		BusinessType string `xml:"businessType"` // A62, A62, A62, A62, A62, ...
		InDomainMRID struct {
			Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"in_Domain.mRID"`
		OutDomainMRID struct {
			Text         string `xml:",chardata"` // 10YCZ-CEPS-----N, 10YCZ-C...
			CodingScheme string `xml:"codingScheme,attr"`
		} `xml:"out_Domain.mRID"`
		CurrencyUnitName     string `xml:"currency_Unit.name"`      // EUR, EUR, EUR, EUR, EUR, ...
		PriceMeasureUnitName string `xml:"price_Measure_Unit.name"` // MWH, MWH, MWH, MWH, MWH, ...
		CurveType            string `xml:"curveType"`               // A01, A01, A01, A01, A01, ...
		Period               struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start"` // 2015-12-31T23:00Z, 2016-0...
				End   string `xml:"end"`   // 2016-01-01T23:00Z, 2016-0...
			} `xml:"timeInterval"`
			Resolution string `xml:"resolution"` // PT60M, PT60M, PT60M, PT60...
			Point      []struct {
				Text        string `xml:",chardata"`
				Position    string `xml:"position"`     // 1, 2, 3, 4, 5, 6, 7, 8, 9...
				PriceAmount string `xml:"price.amount"` // 16.50, 15.50, 14.00, 10.0...
				Quantity    string `xml:"quantity"`     // 226, 87, 104, 189, 217, 8...
			} `xml:"Point"`
		} `xml:"Period"`
		AuctionType                                              string `xml:"auction.type"`                                               // A01, A01, A01, A01, A01, ...
		ContractMarketAgreementType                              string `xml:"contract_MarketAgreement.type"`                              // A01, A01, A01, A01, A01, ...
		QuantityMeasureUnitName                                  string `xml:"quantity_Measure_Unit.name"`                                 // MAW, MAW, MAW, MAW, MAW, ...
		AuctionMRID                                              string `xml:"auction.mRID"`                                               // CP_A_Hourly_SK-UA, CP_A_D...
		AuctionCategory                                          string `xml:"auction.category"`                                           // A04, A04, A01, A01, A01, ...
		ClassificationSequenceAttributeInstanceComponentPosition string `xml:"classificationSequence_AttributeInstanceComponent.position"` // 1, 1
	} `xml:"TimeSeries"`
}

type CriticalNetworkElementMarketDocument struct {
	XMLName                                 xml.Name `xml:"CriticalNetworkElement_MarketDocument"`
	Text                                    string   `xml:",chardata"`
	Xmlns                                   string   `xml:"xmlns,attr"`
	MRID                                    string   `xml:"mRID"`                                       // ed8553b15e3d41fb880134b05...
	RevisionNumber                          string   `xml:"revisionNumber"`                             // 1
	Type                                    string   `xml:"type"`                                       // B11
	ProcessProcessType                      string   `xml:"process.processType"`                        // A01
	SenderMarketParticipantMRID             string   `xml:"sender_MarketParticipant.mRID"`              // 10X1001A1001A450
	SenderMarketParticipantMarketRoleType   string   `xml:"sender_MarketParticipant.marketRole.type"`   // A32
	ReceiverMarketParticipantMRID           string   `xml:"receiver_MarketParticipant.mRID"`            // 10X1001A1001A450
	ReceiverMarketParticipantMarketRoleType string   `xml:"receiver_MarketParticipant.marketRole.type"` // A33
	CreatedDateTime                         string   `xml:"createdDateTime"`                            // 2020-08-08T19:54:31Z
	TimePeriodTimeInterval                  struct {
		Text  string `xml:",chardata"`
		Start string `xml:"start"` // 2015-12-31T23:00Z
		End   string `xml:"end"`   // 2016-01-02T23:00Z
	} `xml:"time_Period.timeInterval"`
	DomainMRID string `xml:"domain.mRID"` // 10YDOM-REGION-1V
	TimeSeries []struct {
		Text         string `xml:",chardata"`
		MRID         string `xml:"mRID"`         // 1, 2
		BusinessType string `xml:"businessType"` // B39, B39
		CurveType    string `xml:"curveType"`    // A01, A01
		Period       struct {
			Text         string `xml:",chardata"`
			TimeInterval struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start"` // 2015-12-31T23:00Z, 2016-0...
				End   string `xml:"end"`   // 2016-01-01T23:00Z, 2016-0...
			} `xml:"timeInterval"`
			Resolution string `xml:"resolution"` // PT60M, PT60M
			Point      []struct {
				Text                 string `xml:",chardata"`
				Position             string `xml:"position"` // 1, 2, 3, 4, 5, 6, 7, 8, 9...
				ConstraintTimeSeries []struct {
					Text                        string `xml:",chardata"`
					MRID                        string `xml:"mRID"`                           // 14648370000, 12144770000,...
					BusinessType                string `xml:"businessType"`                   // B09, B09, B09, B09, B09, ...
					QuantityMeasurementUnitName string `xml:"quantity_Measurement_Unit.name"` // MAW, MAW, MAW, MAW, MAW, ...
					PTDFMeasurementUnitName     string `xml:"pTDF_Measurement_Unit.name"`     // MAW, MAW, MAW, MAW, MAW, ...
					MonitoredRegisteredResource struct {
						Text                                                string `xml:",chardata"`
						FlowBasedStudyDomainMRID                            string `xml:"flowBasedStudy_Domain.mRID"`                              // 10YDOM-REGION-1V, 10YDOM-...
						FlowBasedStudyDomainFlowBasedMarginQuantityQuantity string `xml:"flowBasedStudy_Domain.flowBasedMargin_Quantity.quantity"` // 756, 760, 417, 537, 1116,...
						PTDFDomain                                          []struct {
							Text                 string `xml:",chardata"`
							MRID                 string `xml:"mRID"`                   // 10YBE----------2, 10Y1001...
							PTDFQuantityQuantity string `xml:"pTDF_Quantity.quantity"` // 0.00767, 0.07408, 0.03358...
						} `xml:"PTDF_Domain"`
					} `xml:"Monitored_RegisteredResource"`
				} `xml:"Constraint_TimeSeries"`
			} `xml:"Point"`
		} `xml:"Period"`
	} `xml:"TimeSeries"`
}
