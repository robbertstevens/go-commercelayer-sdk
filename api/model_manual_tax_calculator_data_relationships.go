/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ManualTaxCalculatorDataRelationships struct for ManualTaxCalculatorDataRelationships
type ManualTaxCalculatorDataRelationships struct {
	TaxCategories *AvalaraAccountDataRelationshipsTaxCategories `json:"tax_categories,omitempty"`
	Markets *AvalaraAccountDataRelationshipsMarkets `json:"markets,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
	TaxRules *ManualTaxCalculatorDataRelationshipsTaxRules `json:"tax_rules,omitempty"`
}

// NewManualTaxCalculatorDataRelationships instantiates a new ManualTaxCalculatorDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorDataRelationships() *ManualTaxCalculatorDataRelationships {
	this := ManualTaxCalculatorDataRelationships{}
	return &this
}

// NewManualTaxCalculatorDataRelationshipsWithDefaults instantiates a new ManualTaxCalculatorDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorDataRelationshipsWithDefaults() *ManualTaxCalculatorDataRelationships {
	this := ManualTaxCalculatorDataRelationships{}
	return &this
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetTaxCategories() AvalaraAccountDataRelationshipsTaxCategories {
	if o == nil || o.TaxCategories == nil {
		var ret AvalaraAccountDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetTaxCategoriesOk() (*AvalaraAccountDataRelationshipsTaxCategories, bool) {
	if o == nil || o.TaxCategories == nil {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasTaxCategories() bool {
	if o != nil && o.TaxCategories != nil {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given AvalaraAccountDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *ManualTaxCalculatorDataRelationships) SetTaxCategories(v AvalaraAccountDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetMarkets() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Markets == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Markets == nil {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasMarkets() bool {
	if o != nil && o.Markets != nil {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Markets field.
func (o *ManualTaxCalculatorDataRelationships) SetMarkets(v AvalaraAccountDataRelationshipsMarkets) {
	o.Markets = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ManualTaxCalculatorDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetTaxRules returns the TaxRules field value if set, zero value otherwise.
func (o *ManualTaxCalculatorDataRelationships) GetTaxRules() ManualTaxCalculatorDataRelationshipsTaxRules {
	if o == nil || o.TaxRules == nil {
		var ret ManualTaxCalculatorDataRelationshipsTaxRules
		return ret
	}
	return *o.TaxRules
}

// GetTaxRulesOk returns a tuple with the TaxRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorDataRelationships) GetTaxRulesOk() (*ManualTaxCalculatorDataRelationshipsTaxRules, bool) {
	if o == nil || o.TaxRules == nil {
		return nil, false
	}
	return o.TaxRules, true
}

// HasTaxRules returns a boolean if a field has been set.
func (o *ManualTaxCalculatorDataRelationships) HasTaxRules() bool {
	if o != nil && o.TaxRules != nil {
		return true
	}

	return false
}

// SetTaxRules gets a reference to the given ManualTaxCalculatorDataRelationshipsTaxRules and assigns it to the TaxRules field.
func (o *ManualTaxCalculatorDataRelationships) SetTaxRules(v ManualTaxCalculatorDataRelationshipsTaxRules) {
	o.TaxRules = &v
}

func (o ManualTaxCalculatorDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxCategories != nil {
		toSerialize["tax_categories"] = o.TaxCategories
	}
	if o.Markets != nil {
		toSerialize["markets"] = o.Markets
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.TaxRules != nil {
		toSerialize["tax_rules"] = o.TaxRules
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculatorDataRelationships struct {
	value *ManualTaxCalculatorDataRelationships
	isSet bool
}

func (v NullableManualTaxCalculatorDataRelationships) Get() *ManualTaxCalculatorDataRelationships {
	return v.value
}

func (v *NullableManualTaxCalculatorDataRelationships) Set(val *ManualTaxCalculatorDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorDataRelationships(val *ManualTaxCalculatorDataRelationships) *NullableManualTaxCalculatorDataRelationships {
	return &NullableManualTaxCalculatorDataRelationships{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


