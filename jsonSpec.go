package main

type PACS008AccEnq struct {
	Messageid                 string `json:"messageId,omitempty"`
	Creationdatetime          string `json:"creationDateTime,omitempty"`
	Numbertransaction         string `json:"NumberTransaction,omitempty"`
	Settlementmethod          string `json:"SettlementMethod,omitempty"`
	Endtoendid                string `json:"EndToEndID,omitempty"`
	Transactionid             string `json:"TransactionID,omitempty"`
	Interbanksettlementamount string `json:"InterBankSettlementAmount,omitempty"`
	Currencycode              string `json:"CurrencyCode,omitempty"`
	Chargebearer              string `json:"ChargeBearer,omitempty"`
	Debtorbankid              string `json:"DebtorBankID,omitempty"`
	Creditorbankid            string `json:"CreditorBankID,omitempty"`
	Customeraccountnumbera    string `json:"CustomerAccountNumbera,omitempty"`
}

type PACS008CreditTransfer struct {
	Messageid                         string `json:"messageId,omitempty"`
	Creationdatetime                  string `json:"creationDateTime,omitempty"`
	Numberoftransaction               string `json:"numberOfTransaction,omitempty"`
	Settlementmethod                  string `json:"settlementMethod,omitempty"`
	Endtoendid                        string `json:"endToEndId,omitempty"`
	Transactionid                     string `json:"transactionId,omitempty"`
	Paymentchannelid                  string `json:"paymentChannelId,omitempty"`
	Categorypurpose                   string `json:"categoryPurpose,omitempty"`
	Interbanksettlementamount         string `json:"InterBankSettlementAmount,omitempty"`
	Currencycode                      string `json:"currencyCode,omitempty"`
	Chargebearer                      string `json:"chargeBearer,omitempty"`
	Debtorname                        string `json:"debtorName,omitempty"`
	Debtororganizationid              string `json:"debtorOrganizationId,omitempty"`
	Debtorprivateid                   string `json:"debtorPrivateId,omitempty"`
	Debtoraccountid                   string `json:"debtorAccountId,omitempty"`
	Debtoraccounttype                 string `json:"debtorAccountType,omitempty"`
	Debtorbankid                      string `json:"debtorBankId,omitempty"`
	Creditorbankid                    string `json:"creditorBankId,omitempty"`
	Creditorname                      string `json:"creditorName,omitempty"`
	Creditororganizationid            string `json:"creditorOrganizationId,omitempty"`
	Creditorprivateid                 string `json:"creditorPrivateId,omitempty"`
	Creditoraccountid                 string `json:"creditorAccountId,omitempty"`
	Creditoraccounttype               string `json:"creditorAccountType,omitempty"`
	Remmitanceinformationunstructured string `json:"remmitanceInformationunstructured,omitempty"`
	Debtortype                        string `json:"DebtorType,omitempty"`
	Debtorresidentstatus              string `json:"DebtorResidentStatus,omitempty"`
	Debtortownname                    string `json:"DebtorTownName,omitempty"`
	Creditortype                      string `json:"CreditorType,omitempty"`
	Creditorresidentstatus            string `json:"CreditorResidentStatus,omitempty"`
	Creditortownname                  string `json:"CreditorTownName,omitempty"`
}

type PACS008CTwProxy struct {
	Messageid                         string `json:"messageId,omitempty"`
	Creationdatetime                  string `json:"creationDateTime,omitempty"`
	Numberoftransaction               string `json:"numberOfTransaction,omitempty"`
	Settlementmethod                  string `json:"settlementMethod,omitempty"`
	Endtoendid                        string `json:"endToEndId,omitempty"`
	Transactionid                     string `json:"transactionId,omitempty"`
	Paymentchannelid                  string `json:"paymentChannelId,omitempty"`
	Categorypurpose                   string `json:"categoryPurpose,omitempty"`
	Interbanksettlementamount         string `json:"InterBankSettlementAmount,omitempty"`
	Currencycode                      string `json:"currencyCode,omitempty"`
	Chargebearer                      string `json:"chargeBearer,omitempty"`
	Debtorname                        string `json:"debtorName,omitempty"`
	Debtororganizationid              string `json:"debtorOrganizationId,omitempty"`
	Debtorprivateid                   string `json:"debtorPrivateId,omitempty"`
	Debtoraccountid                   string `json:"debtorAccountId,omitempty"`
	Debtoraccounttype                 string `json:"debtorAccountType,omitempty"`
	Debtorbankid                      string `json:"debtorBankId,omitempty"`
	Creditorbankid                    string `json:"creditorBankId,omitempty"`
	Creditorname                      string `json:"creditorName,omitempty"`
	Creditororganizationid            string `json:"creditorOrganizationId,omitempty"`
	Creditorprivateid                 string `json:"creditorPrivateId,omitempty"`
	Creditoraccountid                 string `json:"creditorAccountId,omitempty"`
	Creditoraccounttype               string `json:"creditorAccountType,omitempty"`
	Proxycreditoraccounttype          string `json:"proxyCreditorAccountType,omitempty"`
	Proxycreditoraccountid            string `json:"proxyCreditorAccountId,omitempty"`
	Remmitanceinformationunstructured string `json:"remmitanceInformationunstructured,omitempty"`
	Debtortype                        string `json:"DebtorType,omitempty"`
	Debtorresidentstatus              string `json:"DebtorResidentStatus,omitempty"`
	Debtortownname                    string `json:"DebtorTownName,omitempty"`
	Creditortype                      string `json:"CreditorType,omitempty"`
	Creditorresidentstatus            string `json:"CreditorResidentStatus,omitempty"`
	Creditortownname                  string `json:"CreditorTownName,omitempty"`
}
