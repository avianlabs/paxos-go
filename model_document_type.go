/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"fmt"
)

// DocumentType the model 'DocumentType'
type DocumentType string

// List of DocumentType
const (
	DocumentType_OTHER_DOCUMENTS DocumentType = "OTHER_DOCUMENTS"
	DocumentType_APPLICATION DocumentType = "APPLICATION"
	DocumentType_ORGANIZATIONAL_DOCUMENTS DocumentType = "ORGANIZATIONAL_DOCUMENTS"
	DocumentType_CERTIFICATE_OF_GOOD_STANDING DocumentType = "CERTIFICATE_OF_GOOD_STANDING"
	DocumentType_TAX_IDENTIFICATION DocumentType = "TAX_IDENTIFICATION"
	DocumentType_IDENTITY_VERIFICATION DocumentType = "IDENTITY_VERIFICATION"
	DocumentType_PROOF_OF_RESIDENCY DocumentType = "PROOF_OF_RESIDENCY"
	DocumentType_PROOF_OF_FUNDS DocumentType = "PROOF_OF_FUNDS"
	DocumentType_SAMPLE_INVOICE DocumentType = "SAMPLE_INVOICE"
	DocumentType_OPERATING_AGREEMENT DocumentType = "OPERATING_AGREEMENT"
	DocumentType_TRUST_DOCUMENTS DocumentType = "TRUST_DOCUMENTS"
	DocumentType_MONEY_SERVICE_DOCUMENTS DocumentType = "MONEY_SERVICE_DOCUMENTS"
	DocumentType_INVESTMENT_DOCUMENTS DocumentType = "INVESTMENT_DOCUMENTS"
	DocumentType_CURP DocumentType = "CURP"
	DocumentType_AML_DOCUMENTS DocumentType = "AML_DOCUMENTS"
	DocumentType_FUND_STRUCTURE_CHART DocumentType = "FUND_STRUCTURE_CHART"
	DocumentType_FUND_MANAGER_REGISTRATION DocumentType = "FUND_MANAGER_REGISTRATION"
	DocumentType_MEMORANDUM_OF_ASSOCIATION DocumentType = "MEMORANDUM_OF_ASSOCIATION"
	DocumentType_ORGANIZATIONAL_CHART DocumentType = "ORGANIZATIONAL_CHART"
	DocumentType_FOUNDATION_BY_LAWS DocumentType = "FOUNDATION_BY_LAWS"
	DocumentType_APPOINTMENT_OF_GUARDIAN_EVIDENCE DocumentType = "APPOINTMENT_OF_GUARDIAN_EVIDENCE"
	DocumentType_LEGAL_DOMICILE_OF_BENEFICIAL_OWNERS DocumentType = "LEGAL_DOMICILE_OF_BENEFICIAL_OWNERS"
	DocumentType_GOVERNING_BODY_MEMBER_NAMES DocumentType = "GOVERNING_BODY_MEMBER_NAMES"
)

// All allowed values of DocumentType enum
var AllowedDocumentTypeEnumValues = []DocumentType{
	"OTHER_DOCUMENTS",
	"APPLICATION",
	"ORGANIZATIONAL_DOCUMENTS",
	"CERTIFICATE_OF_GOOD_STANDING",
	"TAX_IDENTIFICATION",
	"IDENTITY_VERIFICATION",
	"PROOF_OF_RESIDENCY",
	"PROOF_OF_FUNDS",
	"SAMPLE_INVOICE",
	"OPERATING_AGREEMENT",
	"TRUST_DOCUMENTS",
	"MONEY_SERVICE_DOCUMENTS",
	"INVESTMENT_DOCUMENTS",
	"CURP",
	"AML_DOCUMENTS",
	"FUND_STRUCTURE_CHART",
	"FUND_MANAGER_REGISTRATION",
	"MEMORANDUM_OF_ASSOCIATION",
	"ORGANIZATIONAL_CHART",
	"FOUNDATION_BY_LAWS",
	"APPOINTMENT_OF_GUARDIAN_EVIDENCE",
	"LEGAL_DOMICILE_OF_BENEFICIAL_OWNERS",
	"GOVERNING_BODY_MEMBER_NAMES",
}

func (v *DocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentType(value)
	for _, existing := range AllowedDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentType", value)
}

// NewDocumentTypeFromValue returns a pointer to a valid DocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentTypeFromValue(v string) (*DocumentType, error) {
	ev := DocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentType: valid values are %v", v, AllowedDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentType) IsValid() bool {
	for _, existing := range AllowedDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentType value
func (v DocumentType) Ptr() *DocumentType {
	return &v
}

type NullableDocumentType struct {
	value *DocumentType
	isSet bool
}

func (v NullableDocumentType) Get() *DocumentType {
	return v.value
}

func (v *NullableDocumentType) Set(val *DocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentType(val *DocumentType) *NullableDocumentType {
	return &NullableDocumentType{value: val, isSet: true}
}

func (v NullableDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

