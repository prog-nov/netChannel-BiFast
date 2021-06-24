package main

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type Request struct {
	BusMsg BusMsg `json:"BusMsg"`
}

type BusMsg struct {
	AppHdr   BusinessApplicationHeaderV01 `json:"AppHdr"`
	Document Document                     `json:"Document"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"FinInstnId" json:"FinInstnId"`
	BrnchId    BranchData2                         `xml:"BrnchId,omitempty" json:"BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"Id,omitempty" json:"Id,omitempty"`
	Nm      Max140Text     `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

type BusinessApplicationHeader1 struct {
	CharSet    string                `xml:"CharSet,omitempty" json:"CharSet,omitempty"`
	Fr         Party9Choice          `xml:"Fr" json:"Fr"`
	To         Party9Choice          `xml:"To" json:"To"`
	BizMsgIdr  Max35Text             `xml:"BizMsgIdr" json:"BizMsgIdr"`
	MsgDefIdr  Max35Text             `xml:"MsgDefIdr" json:"MsgDefIdr"`
	BizSvc     Max35Text             `xml:"BizSvc,omitempty" json:"BizSvc,omitempty"`
	CreDt      ISONormalisedDateTime `xml:"CreDt" json:"CreDt"`
	CpyDplct   CopyDuplicate1Code    `xml:"CpyDplct,omitempty" json:"CpyDplct,omitempty"`
	PssblDplct bool                  `xml:"PssblDplct,omitempty" json:"PssblDplct,omitempty"`
	Prty       string                `xml:"Prty,omitempty" json:"Prty,omitempty"`
	Sgntr      SignatureEnvelope     `xml:"Sgntr,omitempty" json:"Sgntr,omitempty"`
}

type BusinessApplicationHeaderV01 struct {
	CharSet    string                     `xml:"CharSet,omitempty" json:"CharSet,omitempty"`
	Fr         Party9Choice               `xml:"Fr" json:"Fr"`
	To         Party9Choice               `xml:"To" json:"To"`
	BizMsgIdr  Max35Text                  `xml:"BizMsgIdr" json:"BizMsgIdr"`
	MsgDefIdr  Max35Text                  `xml:"MsgDefIdr" json:"MsgDefIdr"`
	BizSvc     Max35Text                  `xml:"BizSvc,omitempty" json:"BizSvc,omitempty"`
	CreDt      ISONormalisedDateTime      `xml:"CreDt" json:"CreDt"`
	CpyDplct   CopyDuplicate1Code         `xml:"CpyDplct,omitempty" json:"CpyDplct,omitempty"`
	PssblDplct bool                       `xml:"PssblDplct,omitempty" json:"PssblDplct,omitempty"`
	Prty       string                     `xml:"Prty,omitempty" json:"Prty,omitempty"`
	Sgntr      SignatureEnvelope          `xml:"Sgntr,omitempty" json:"Sgntr,omitempty"`
	Rltd       BusinessApplicationHeader1 `xml:"Rltd,omitempty" json:"Rltd,omitempty"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"NmPrfx,omitempty" json:"NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"PhneNb,omitempty" json:"PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"MobNb,omitempty" json:"MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"FaxNb,omitempty" json:"FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"EmailAdr,omitempty" json:"EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"BirthDt" json:"BirthDt"`
	PrvcOfBirth Max35Text   `xml:"PrvcOfBirth,omitempty" json:"PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"CityOfBirth" json:"CityOfBirth"`
	CtryOfBirth CountryCode `xml:"CtryOfBirth" json:"CtryOfBirth"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"BICFI,omitempty" json:"BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type ISONormalisedDateTime time.Time

func (t *ISONormalisedDateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISONormalisedDateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification7 struct {
	AnyBIC AnyBICIdentifier                     `xml:"AnyBIC,omitempty" json:"AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type Party10Choice struct {
	OrgId  OrganisationIdentification7 `xml:"OrgId,omitempty" json:"OrgId,omitempty"`
	PrvtId PersonIdentification5       `xml:"PrvtId,omitempty" json:"PrvtId,omitempty"`
}

type Party9Choice struct {
	OrgId PartyIdentification42                        `xml:"OrgId,omitempty" json:"OrgId,omitempty"`
	FIId  BranchAndFinancialInstitutionIdentification5 `xml:"FIId,omitempty" json:"FIId,omitempty"`
}

type PartyIdentification42 struct {
	Nm        Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Id        Party10Choice   `xml:"Id,omitempty" json:"Id,omitempty"`
	CtryOfRes CountryCode     `xml:"CtryOfRes,omitempty" json:"CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"CtctDtls,omitempty" json:"CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"DtAndPlcOfBirth,omitempty" json:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"AdrTp,omitempty" json:"AdrTp,omitempty"`
	Dept        Max70Text        `xml:"Dept,omitempty" json:"Dept,omitempty"`
	SubDept     Max70Text        `xml:"SubDept,omitempty" json:"SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"StrtNm,omitempty" json:"StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"BldgNb,omitempty" json:"BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"PstCd,omitempty" json:"PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"CtrySubDvsn,omitempty" json:"CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"Ctry,omitempty" json:"Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"AdrLine,omitempty" json:"AdrLine,omitempty"`
}

type SignatureEnvelope struct {
	Item string `xml:",any" json:",any"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"IBAN,omitempty" json:"IBAN,omitempty"`
	Othr GenericAccountIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                          `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata" json:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr" json:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata" json:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr" json:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry GenericIdentification30 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BI_AddtlCstmrInf struct {
	Tp       Max35Text `xml:"Tp,omitempty" json:"Tp,omitempty"`
	RsdntSts Max35Text `xml:"RsdntSts,omitempty" json:"RsdntSts,omitempty"`
	TwnNm    Max35Text `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`
}

type BI_SupplementaryData1 struct {
	PlcAndNm Max350Text                    `xml:"PlcAndNm,omitempty" json:"PlcAndNm,omitempty"`
	Envlp    BI_SupplementaryDataEnvelope1 `xml:"Envlp" json:"Envlp"`
}

type BI_SupplementaryDataEnvelope1 struct {
	Dbtr            BI_AddtlCstmrInf `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	Cdtr            BI_AddtlCstmrInf `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	OrgnlEndtoEndId Max34Text        `xml:"OrgnlEndtoEndId,omitempty" json:"OrgnlEndtoEndId,omitempty"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId" json:"FinInstnId"`
	BrnchId    BranchData3                          `xml:"BrnchId,omitempty" json:"BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"Id,omitempty" json:"Id,omitempty"`
	LEI     LEIIdentifier   `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Nm      Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"Id" json:"Id"`
	Tp   CashAccountType2Choice       `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:"Ccy,omitempty"`
	Nm   Max70Text                    `xml:"Nm,omitempty" json:"Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"Prxy,omitempty" json:"Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt" json:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt" json:"Agt"`
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                                 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:"ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"MmbId" json:"MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"NmPrfx,omitempty" json:"NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"PhneNb,omitempty" json:"PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"MobNb,omitempty" json:"MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"FaxNb,omitempty" json:"FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"EmailAdr,omitempty" json:"EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"EmailPurp,omitempty" json:"EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"JobTitl,omitempty" json:"JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"Rspnsblty,omitempty" json:"Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"Dept,omitempty" json:"Dept,omitempty"`
	Othr      []OtherContact1             `xml:"Othr,omitempty" json:"Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:"PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of CRDT, DBIT
type CreditDebitCode string

type CreditTransferTransaction39 struct {
	PmtId             PaymentIdentification7                       `xml:"PmtId" json:"PmtId"`
	PmtTpInf          PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt    ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt" json:"IntrBkSttlmAmt"`
	IntrBkSttlmDt     ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmPrty         Priority3Code                                `xml:"SttlmPrty,omitempty" json:"SttlmPrty,omitempty"`
	SttlmTmIndctn     SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty" json:"SttlmTmIndctn,omitempty"`
	SttlmTmReq        SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty" json:"SttlmTmReq,omitempty"`
	AccptncDtTm       ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	PoolgAdjstmntDt   ISODate                                      `xml:"PoolgAdjstmntDt,omitempty" json:"PoolgAdjstmntDt,omitempty"`
	InstdAmt          ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty" json:"InstdAmt,omitempty"`
	XchgRate          float64                                      `xml:"XchgRate,omitempty" json:"XchgRate,omitempty"`
	ChrgBr            ChargeBearerType1Code                        `xml:"ChrgBr" json:"ChrgBr"`
	ChrgsInf          []Charges7                                   `xml:"ChrgsInf,omitempty" json:"ChrgsInf,omitempty"`
	PrvsInstgAgt1     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	IntrmyAgt1        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr         PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	InitgPty          PartyIdentification135                       `xml:"InitgPty,omitempty" json:"InitgPty,omitempty"`
	Dbtr              PartyIdentification135                       `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct          CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"DbtrAgt"`
	DbtrAgtAcct       CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"CdtrAgt"`
	CdtrAgtAcct       CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr              PartyIdentification135                       `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct          CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr         PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []InstructionForCreditorAgent1               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt    []InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Purp              Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
	RgltryRptg        []RegulatoryReporting3                       `xml:"RgltryRptg,omitempty" json:"RgltryRptg,omitempty"`
	Tax               TaxInformation8                              `xml:"Tax,omitempty" json:"Tax,omitempty"`
	RltdRmtInf        []RemittanceLocation7                        `xml:"RltdRmtInf,omitempty" json:"RltdRmtInf,omitempty"`
	RmtInf            RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	SplmtryData       []BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ref Max35Text              `xml:"Ref,omitempty" json:"Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text         `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry" json:"CdOrPrtry"`
	Issr      Max35Text                    `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"BirthDt" json:"BirthDt"`
	PrvcOfBirth Max35Text   `xml:"PrvcOfBirth,omitempty" json:"PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"CityOfBirth" json:"CityOfBirth"`
	CtryOfBirth CountryCode `xml:"CtryOfBirth" json:"CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"FrDt" json:"FrDt"`
	ToDt ISODate `xml:"ToDt" json:"ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                       `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type Document struct {
	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV08       `xml:"FIToFICstmrCdtTrf" json:"FIToFICstmrCdtTrf"`
	FIToFIPmtStsRpt   FIToFIPaymentStatusReportV10          `xml:"FIToFIPmtStsRpt" json:"FIToFIPmtStsRpt"`
	FICdtTrf          FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf" json:"FICdtTrf"`
	FIToFIPmtStsReq   FIToFIPaymentStatusRequestV04         `xml:"FIToFIPmtStsReq" json:"FIToFIPmtStsReq"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"CdtDbtInd,omitempty" json:"CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Nb     Max35Text         `xml:"Nb,omitempty" json:"Nb,omitempty"`
	RltdDt ISODate           `xml:"RltdDt,omitempty" json:"RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id" json:"Id"`
	Desc Max2048Text                   `xml:"Desc,omitempty" json:"Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"Amt,omitempty" json:"Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry" json:"CdOrPrtry"`
	Issr      Max35Text               `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                     `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// May be no more than 4 items long
type ExternalAccountIdentification1Code string

// May be no more than 4 items long
type ExternalCashAccountType1Code string

// May be no more than 3 items long
type ExternalCashClearingSystem1Code string

// May be no more than 4 items long
type ExternalCategoryPurpose1Code string

// May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// May be no more than 4 items long
type ExternalDiscountAmountType1Code string

// May be no more than 4 items long
type ExternalDocumentLineType1Code string

// May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be no more than 4 items long
type ExternalGarnishmentType1Code string

// May be no more than 35 items long
type ExternalLocalInstrument1Code string

// May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// May be no more than 4 items long
type ExternalPersonIdentification1Code string

// May be no more than 4 items long
type ExternalProxyAccountType1Code string

// May be no more than 4 items long
type ExternalPurpose1Code string

// May be no more than 4 items long
type ExternalServiceLevel1Code string

// May be no more than 4 items long
type ExternalTaxAmountType1Code string

type FIToFICustomerCreditTransferV08 struct {
	GrpHdr      GroupHeader93                 `xml:"GrpHdr" json:"GrpHdr"`
	CdtTrfTxInf []CreditTransferTransaction39 `xml:"CdtTrfTxInf" json:"CdtTrfTxInf"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                                       `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"BICFI,omitempty" json:"BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Nm          Max140Text                          `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"Tp" json:"Tp"`
	Grnshee           PartyIdentification135            `xml:"Grnshee,omitempty" json:"Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:"GrnshmtAdmstr,omitempty"`
	RefNb             Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Dt                ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"FmlyMdclInsrncInd,omitempty" json:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"MplyeeTermntnInd,omitempty" json:"MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry" json:"CdOrPrtry"`
	Issr      Max35Text              `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"Id" json:"Id"`
	SchmeNm AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"Id" json:"Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"Id" json:"Id"`
	Issr    Max35Text              `xml:"Issr" json:"Issr"`
	SchmeNm Max35Text              `xml:"SchmeNm,omitempty" json:"SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"Id" json:"Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"Id" json:"Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type GroupHeader93 struct {
	MsgId             Max35Text                                    `xml:"MsgId" json:"MsgId"`
	CreDtTm           ISODateTime                                  `xml:"CreDtTm" json:"CreDtTm"`
	BtchBookg         bool                                         `xml:"BtchBookg,omitempty" json:"BtchBookg,omitempty"`
	NbOfTxs           Max15NumericText                             `xml:"NbOfTxs" json:"NbOfTxs"`
	CtrlSum           float64                                      `xml:"CtrlSum,omitempty" json:"CtrlSum,omitempty"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt,omitempty" json:"TtlIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt     ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmInf          SettlementInstruction7                       `xml:"SttlmInf" json:"SttlmInf"`
	PmtTpInf          PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// May be one of CHQB, HOLD, PHOB, TELB
type Instruction3Code string

// May be one of PHOA, TELA
type Instruction4Code string

type InstructionForCreditorAgent1 struct {
	Cd       Instruction3Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	InstrInf Max140Text       `xml:"InstrInf,omitempty" json:"InstrInf,omitempty"`
}

type InstructionForNextAgent1 struct {
	Cd       Instruction4Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	InstrInf Max140Text       `xml:"InstrInf,omitempty" json:"InstrInf,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                    `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be no more than 10 items long
type Max10Text string

// May be no more than 128 items long
type Max128Text string

// May be no more than 140 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// May be no more than 16 items long
type Max16Text string

// May be no more than 2048 items long
type Max2048Text string

// May be no more than 34 items long
type Max34Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 4 items long
type Max4Text string

// May be no more than 70 items long
type Max70Text string

type NameAndAddress16 struct {
	Nm  Max140Text      `xml:"Nm" json:"Nm"`
	Adr PostalAddress24 `xml:"Adr" json:"Adr"`
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"AnyBIC,omitempty" json:"AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"LEI,omitempty" json:"LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                               `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"ChanlTp" json:"ChanlTp"`
	Id      Max128Text `xml:"Id,omitempty" json:"Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId,omitempty" json:"OrgId,omitempty"`
	PrvtId PersonIdentification13       `xml:"PrvtId,omitempty" json:"PrvtId,omitempty"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"Id,omitempty" json:"Id,omitempty"`
	CtryOfRes CountryCode     `xml:"CtryOfRes,omitempty" json:"CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"CtctDtls,omitempty" json:"CtctDtls,omitempty"`
}

type PaymentIdentification7 struct {
	InstrId    Max35Text        `xml:"InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"EndToEndId" json:"EndToEndId"`
	TxId       Max35Text        `xml:"TxId,omitempty" json:"TxId,omitempty"`
	UETR       UUIDv4Identifier `xml:"UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  Max35Text        `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

type PaymentTypeInformation28 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty" json:"InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:"ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth,omitempty" json:"DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:"Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                         `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"AdrTp,omitempty" json:"AdrTp,omitempty"`
	Dept        Max70Text          `xml:"Dept,omitempty" json:"Dept,omitempty"`
	SubDept     Max70Text          `xml:"SubDept,omitempty" json:"SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"StrtNm,omitempty" json:"StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"BldgNb,omitempty" json:"BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"BldgNm,omitempty" json:"BldgNm,omitempty"`
	Flr         Max70Text          `xml:"Flr,omitempty" json:"Flr,omitempty"`
	PstBx       Max16Text          `xml:"PstBx,omitempty" json:"PstBx,omitempty"`
	Room        Max70Text          `xml:"Room,omitempty" json:"Room,omitempty"`
	PstCd       Max16Text          `xml:"PstCd,omitempty" json:"PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"TwnNm,omitempty" json:"TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"TwnLctnNm,omitempty" json:"TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"DstrctNm,omitempty" json:"DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"CtrySubDvsn,omitempty" json:"CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"Ctry,omitempty" json:"Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"AdrLine,omitempty" json:"AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of HIGH, NORM
type Priority2Code string

// May be one of URGT, HIGH, NORM
type Priority3Code string

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Id Max2048Text             `xml:"Id" json:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                     `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text            `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Nb       Max35Text                  `xml:"Nb,omitempty" json:"Nb,omitempty"`
	RltdDt   ISODate                    `xml:"RltdDt,omitempty" json:"RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:"LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text         `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry" json:"CdOrPrtry"`
	Issr      Max35Text                   `xml:"Issr,omitempty" json:"Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   Max140Text  `xml:"Nm,omitempty" json:"Nm,omitempty"`
	Ctry CountryCode `xml:"Ctry,omitempty" json:"Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"DbtCdtRptgInd,omitempty" json:"DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"Authrty,omitempty" json:"Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"Dtls,omitempty" json:"Dtls,omitempty"`
}

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:"DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:"DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:"CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:"DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:"DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:"CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []Max140Text                        `xml:"Ustrd,omitempty" json:"Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:"Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       Max35Text                 `xml:"RmtId,omitempty" json:"RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:"RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd" json:"Mtd"`
	ElctrncAdr Max2048Text                   `xml:"ElctrncAdr,omitempty" json:"ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type SettlementDateTimeIndication1 struct {
	DbtDtTm ISODateTime `xml:"DbtDtTm,omitempty" json:"DbtDtTm,omitempty"`
	CdtDtTm ISODateTime `xml:"CdtDtTm,omitempty" json:"CdtDtTm,omitempty"`
}

type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"SttlmMtd" json:"SttlmMtd"`
	SttlmAcct            CashAccount38                                `xml:"SttlmAcct,omitempty" json:"SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty" json:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type SettlementTimeRequest2 struct {
	CLSTm  ISOTime `xml:"CLSTm,omitempty" json:"CLSTm,omitempty"`
	TillTm ISOTime `xml:"TillTm,omitempty" json:"TillTm,omitempty"`
	FrTm   ISOTime `xml:"FrTm,omitempty" json:"FrTm,omitempty"`
	RjctTm ISOTime `xml:"RjctTm,omitempty" json:"RjctTm,omitempty"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   Max35Text                         `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Dt   ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	Ctry CountryCode                       `xml:"Ctry,omitempty" json:"Ctry,omitempty"`
	Cd   Max10Text                         `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:"Amt,omitempty"`
	Inf  []Max35Text                       `xml:"Inf,omitempty" json:"Inf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:"RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"RfrdDocAmt,omitempty" json:"RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty" json:"CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"Invcr,omitempty" json:"Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"Invcee,omitempty" json:"Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"TaxRmt,omitempty" json:"TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"GrnshmtRmt,omitempty" json:"GrnshmtRmt,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"AddtlRmtInf,omitempty" json:"AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"PlcAndNm,omitempty" json:"PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp" json:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any" json:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"Rate,omitempty" json:"Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:"TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:"TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"Dtls,omitempty" json:"Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                  `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"Titl,omitempty" json:"Titl,omitempty"`
	Nm   Max140Text `xml:"Nm,omitempty" json:"Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:"Rcrd,omitempty"`
}

type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:"Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  Max35Text `xml:"TaxId,omitempty" json:"TaxId,omitempty"`
	RegnId Max35Text `xml:"RegnId,omitempty" json:"RegnId,omitempty"`
	TaxTp  Max35Text `xml:"TaxTp,omitempty" json:"TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   Max35Text         `xml:"TaxId,omitempty" json:"TaxId,omitempty"`
	RegnId  Max35Text         `xml:"RegnId,omitempty" json:"RegnId,omitempty"`
	TaxTp   Max35Text         `xml:"TaxTp,omitempty" json:"TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"Authstn,omitempty" json:"Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     ISODate              `xml:"Yr,omitempty" json:"Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty" json:"Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"FrToDt,omitempty" json:"FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       Max35Text  `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Ctgy     Max35Text  `xml:"Ctgy,omitempty" json:"Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"CtgyDtls,omitempty" json:"CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"DbtrSts,omitempty" json:"DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"CertId,omitempty" json:"CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"FrmsCd,omitempty" json:"FrmsCd,omitempty"`
	Prd      TaxPeriod2 `xml:"Prd,omitempty" json:"Prd,omitempty"`
	TaxAmt   TaxAmount2 `xml:"TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"Prd,omitempty" json:"Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
}

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      Max35Text                                    `xml:"OrgnlMndtId,omitempty" json:"OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:"OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:"OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:"OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:"OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:"OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:"OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:"OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt ISODate                                      `xml:"OrgnlFnlColltnDt,omitempty" json:"OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:"OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:"OrgnlRsn,omitempty"`
	OrgnlTrckgDays   Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty" json:"OrgnlTrckgDays,omitempty"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty" json:"InstdAmt,omitempty"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt,omitempty" json:"EqvtAmt,omitempty"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"CcyOfTrf" json:"CcyOfTrf"`
}

// Must match the pattern [0-9]{2}
type Exact2NumericText string

// May be no more than 4 items long
type ExternalMandateSetupReason1Code string

// May be no more than 4 items long
type ExternalPaymentGroupStatus1Code string

// May be no more than 4 items long
type ExternalPaymentTransactionStatus1Code string

// May be no more than 4 items long
type ExternalStatusReason1Code string

type FIToFIPaymentStatusReportV10 struct {
	GrpHdr            GroupHeader91           `xml:"GrpHdr" json:"GrpHdr"`
	OrgnlGrpInfAndSts []OriginalGroupHeader17 `xml:"OrgnlGrpInfAndSts,omitempty" json:"OrgnlGrpInfAndSts,omitempty"`
	TxInfAndSts       []PaymentTransaction110 `xml:"TxInfAndSts,omitempty" json:"TxInfAndSts,omitempty"`
	SplmtryData       []SupplementaryData1    `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp,omitempty" json:"Tp,omitempty"`
	Prd    FrequencyPeriod1    `xml:"Prd,omitempty" json:"Prd,omitempty"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm,omitempty" json:"PtInTm,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"Tp" json:"Tp"`
	PtInTm Exact2NumericText `xml:"PtInTm" json:"PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp" json:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd" json:"CntPerPrd"`
}

type GroupHeader91 struct {
	MsgId    Max35Text                                    `xml:"MsgId" json:"MsgId"`
	CreDtTm  ISODateTime                                  `xml:"CreDtTm" json:"CreDtTm"`
	InstgAgt BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        Max35Text                     `xml:"MndtId,omitempty" json:"MndtId,omitempty"`
	DtOfSgntr     ISODate                       `xml:"DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	AmdmntInd     bool                          `xml:"AmdmntInd,omitempty" json:"AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:"AmdmntInfDtls,omitempty"`
	ElctrncSgntr  Max1025Text                   `xml:"ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstColltnDt  ISODate                       `xml:"FrstColltnDt,omitempty" json:"FrstColltnDt,omitempty"`
	FnlColltnDt   ISODate                       `xml:"FnlColltnDt,omitempty" json:"FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	TrckgDays     Exact2NumericText             `xml:"TrckgDays,omitempty" json:"TrckgDays,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max70Text                       `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be no more than 1025 items long
type Max1025Text string

// May be no more than 105 items long
type Max105Text string

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs Max15NumericText                      `xml:"DtldNbOfTxs" json:"DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `xml:"DtldSts" json:"DtldSts"`
	DtldCtrlSum float64                               `xml:"DtldCtrlSum,omitempty" json:"DtldCtrlSum,omitempty"`
}

type OriginalGroupHeader17 struct {
	OrgnlMsgId    Max35Text                        `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId  Max35Text                        `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm  ISODateTime                      `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  Max15NumericText                 `xml:"OrgnlNbOfTxs,omitempty" json:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                          `xml:"OrgnlCtrlSum,omitempty" json:"OrgnlCtrlSum,omitempty"`
	GrpSts        ExternalPaymentGroupStatus1Code  `xml:"GrpSts,omitempty" json:"GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty" json:"NbOfTxsPerSts,omitempty"`
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   Max35Text   `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId Max35Text   `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"Dt,omitempty" json:"Dt,omitempty"`
	DtTm ISODateTime `xml:"DtTm,omitempty" json:"DtTm,omitempty"`
}

type OriginalTransactionReference28 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:"IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"Amt,omitempty" json:"Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:"CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:"SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedInformation14                  `xml:"MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty,omitempty" json:"Pty,omitempty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:"Agt,omitempty"`
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

type PaymentTransaction110 struct {
	StsId             Max35Text                                    `xml:"StsId,omitempty" json:"StsId,omitempty"`
	OrgnlGrpInf       OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId      Max35Text                                    `xml:"OrgnlInstrId,omitempty" json:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId   Max35Text                                    `xml:"OrgnlEndToEndId,omitempty" json:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId         Max35Text                                    `xml:"OrgnlTxId,omitempty" json:"OrgnlTxId,omitempty"`
	OrgnlUETR         UUIDv4Identifier                             `xml:"OrgnlUETR,omitempty" json:"OrgnlUETR,omitempty"`
	TxSts             ExternalPaymentTransactionStatus1Code        `xml:"TxSts,omitempty" json:"TxSts,omitempty"`
	StsRsnInf         []StatusReasonInformation12                  `xml:"StsRsnInf,omitempty" json:"StsRsnInf,omitempty"`
	ChrgsInf          []Charges7                                   `xml:"ChrgsInf,omitempty" json:"ChrgsInf,omitempty"`
	AccptncDtTm       ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	FctvIntrBkSttlmDt DateAndDateTime2Choice                       `xml:"FctvIntrBkSttlmDt,omitempty" json:"FctvIntrBkSttlmDt,omitempty"`
	AcctSvcrRef       Max35Text                                    `xml:"AcctSvcrRef,omitempty" json:"AcctSvcrRef,omitempty"`
	ClrSysRef         Max35Text                                    `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	OrgnlTxRef        OriginalTransactionReference28               `xml:"OrgnlTxRef,omitempty" json:"OrgnlTxRef,omitempty"`
	SplmtryData       []BI_SupplementaryData1                      `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty" json:"InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:"ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"SeqTp,omitempty" json:"SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
}

// May be one of FRST, RCUR, FNAL, OOFF, RPRE
type SequenceType3Code string

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                 `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

type StatusReasonInformation12 struct {
	Orgtr    PartyIdentification135 `xml:"Orgtr,omitempty" json:"Orgtr,omitempty"`
	Rsn      StatusReason6Choice    `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
	AddtlInf []Max105Text           `xml:"AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type CreditTransferTransaction44 struct {
	PmtId              PaymentIdentification13                      `xml:"PmtId" json:"PmtId"`
	PmtTpInf           PaymentTypeInformation28                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt     ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt" json:"IntrBkSttlmAmt"`
	IntrBkSttlmDt      ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmPrty          Priority3Code                                `xml:"SttlmPrty,omitempty" json:"SttlmPrty,omitempty"`
	SttlmTmIndctn      SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty" json:"SttlmTmIndctn,omitempty"`
	SttlmTmReq         SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty" json:"SttlmTmReq,omitempty"`
	PrvsInstgAgt1      BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct  CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2      BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct  CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3      BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct  CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	InstgAgt           BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt           BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	IntrmyAgt1         BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct     CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2         BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct     CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3         BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct     CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr          BranchAndFinancialInstitutionIdentification6 `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr               BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct           CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt            BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct        CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt            BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct        CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr               BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct           CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr          BranchAndFinancialInstitutionIdentification6 `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt    []InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt     []InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Purp               Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
	RmtInf             RemittanceInformation2                       `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UndrlygCstmrCdtTrf CreditTransferTransaction45                  `xml:"UndrlygCstmrCdtTrf,omitempty" json:"UndrlygCstmrCdtTrf,omitempty"`
	SplmtryData        []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type InstructionForCreditorAgent3 struct {
	Cd       ExternalCreditorAgentInstruction1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	InstrInf Max140Text                            `xml:"InstrInf,omitempty" json:"InstrInf,omitempty"`
}

// May be no more than 4 items long
type ExternalCreditorAgentInstruction1Code string

type CreditTransferTransaction45 struct {
	UltmtDbtr         PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	InitgPty          PartyIdentification135                       `xml:"InitgPty,omitempty" json:"InitgPty,omitempty"`
	Dbtr              PartyIdentification135                       `xml:"Dbtr" json:"Dbtr"`
	DbtrAcct          CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"DbtrAgt"`
	DbtrAgtAcct       CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	PrvsInstgAgt1     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct CashAccount38                                `xml:"PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct CashAccount38                                `xml:"PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct CashAccount38                                `xml:"PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	IntrmyAgt1        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	CdtrAgt           BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"CdtrAgt"`
	CdtrAgtAcct       CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr              PartyIdentification135                       `xml:"Cdtr" json:"Cdtr"`
	CdtrAcct          CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr         PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []InstructionForCreditorAgent3               `xml:"InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt    []InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Tax               TaxInformation8                              `xml:"Tax,omitempty" json:"Tax,omitempty"`
	RmtInf            RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	InstdAmt          ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty" json:"InstdAmt,omitempty"`
}

type FinancialInstitutionCreditTransferV09 struct {
	GrpHdr      GroupHeader93                 `xml:"GrpHdr" json:"GrpHdr"`
	CdtTrfTxInf []CreditTransferTransaction44 `xml:"CdtTrfTxInf" json:"CdtTrfTxInf"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type PaymentIdentification13 struct {
	InstrId    Max35Text        `xml:"InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"EndToEndId" json:"EndToEndId"`
	TxId       Max35Text        `xml:"TxId,omitempty" json:"TxId,omitempty"`
	UETR       UUIDv4Identifier `xml:"UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  Max35Text        `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

type RemittanceInformation2 struct {
	Ustrd []Max140Text `xml:"Ustrd,omitempty" json:"Ustrd,omitempty"`
}

type CreditTransferMandateData1 struct {
	MndtId       Max35Text                 `xml:"MndtId,omitempty" json:"MndtId,omitempty"`
	Tp           MandateTypeInformation2   `xml:"Tp,omitempty" json:"Tp,omitempty"`
	DtOfSgntr    ISODate                   `xml:"DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	DtOfVrfctn   ISODateTime               `xml:"DtOfVrfctn,omitempty" json:"DtOfVrfctn,omitempty"`
	ElctrncSgntr Max10KBinary              `xml:"ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstPmtDt    ISODate                   `xml:"FrstPmtDt,omitempty" json:"FrstPmtDt,omitempty"`
	FnlPmtDt     ISODate                   `xml:"FnlPmtDt,omitempty" json:"FnlPmtDt,omitempty"`
	Frqcy        Frequency36Choice         `xml:"Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn          MandateSetupReason1Choice `xml:"Rsn,omitempty" json:"Rsn,omitempty"`
}

type FIToFIPaymentStatusRequestV04 struct {
	GrpHdr      GroupHeader91                `xml:"GrpHdr" json:"GrpHdr"`
	OrgnlGrpInf []OriginalGroupInformation27 `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	TxInf       []PaymentTransaction121      `xml:"TxInf,omitempty" json:"TxInf,omitempty"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    MandateClassification1Code `xml:"Cd,omitempty" json:"Cd,omitempty"`
	Prtry Max35Text                  `xml:"Prtry,omitempty" json:"Prtry,omitempty"`
}

// May be one of FIXE, USGB, VARI
type MandateClassification1Code string

type MandateRelatedData1Choice struct {
	DrctDbtMndt MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:"DrctDbtMndt,omitempty"`
	CdtTrfMndt  CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:"CdtTrfMndt,omitempty"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"Clssfctn,omitempty" json:"Clssfctn,omitempty"`
}

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type OriginalGroupInformation27 struct {
	OrgnlMsgId   Max35Text        `xml:"OrgnlMsgId" json:"OrgnlMsgId"`
	OrgnlMsgNmId Max35Text        `xml:"OrgnlMsgNmId" json:"OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime      `xml:"OrgnlCreDtTm,omitempty" json:"OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs Max15NumericText `xml:"OrgnlNbOfTxs,omitempty" json:"OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum float64          `xml:"OrgnlCtrlSum,omitempty" json:"OrgnlCtrlSum,omitempty"`
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:"IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"Amt,omitempty" json:"Amt,omitempty"`
	IntrBkSttlmDt  ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:"ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:"CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:"SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty" json:"RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty" json:"Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"Cdtr,omitempty" json:"Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty" json:"Purp,omitempty"`
}

type PaymentTransaction121 struct {
	StsReqId        Max35Text                                    `xml:"StsReqId,omitempty" json:"StsReqId,omitempty"`
	OrgnlGrpInf     OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:"OrgnlGrpInf,omitempty"`
	OrgnlInstrId    Max35Text                                    `xml:"OrgnlInstrId,omitempty" json:"OrgnlInstrId,omitempty"`
	OrgnlEndToEndId Max35Text                                    `xml:"OrgnlEndToEndId,omitempty" json:"OrgnlEndToEndId,omitempty"`
	OrgnlTxId       Max35Text                                    `xml:"OrgnlTxId,omitempty" json:"OrgnlTxId,omitempty"`
	OrgnlUETR       UUIDv4Identifier                             `xml:"OrgnlUETR,omitempty" json:"OrgnlUETR,omitempty"`
	AccptncDtTm     ISODateTime                                  `xml:"AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	ClrSysRef       Max35Text                                    `xml:"ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
	InstgAgt        BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt        BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	OrgnlTxRef      OriginalTransactionReference31               `xml:"OrgnlTxRef,omitempty" json:"OrgnlTxRef,omitempty"`
	SplmtryData     []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}
