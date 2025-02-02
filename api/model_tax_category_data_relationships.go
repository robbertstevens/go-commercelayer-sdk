/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TaxCategoryDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCategoryDataRelationships{}

// TaxCategoryDataRelationships struct for TaxCategoryDataRelationships
type TaxCategoryDataRelationships struct {
	Sku           *BundleDataRelationshipsSkus               `json:"sku,omitempty"`
	TaxCalculator *MarketDataRelationshipsTaxCalculator      `json:"tax_calculator,omitempty"`
	Attachments   *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions      *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
}

// NewTaxCategoryDataRelationships instantiates a new TaxCategoryDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryDataRelationships() *TaxCategoryDataRelationships {
	this := TaxCategoryDataRelationships{}
	return &this
}

// NewTaxCategoryDataRelationshipsWithDefaults instantiates a new TaxCategoryDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryDataRelationshipsWithDefaults() *TaxCategoryDataRelationships {
	this := TaxCategoryDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || IsNil(o.Sku) {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *TaxCategoryDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetTaxCalculator() MarketDataRelationshipsTaxCalculator {
	if o == nil || IsNil(o.TaxCalculator) {
		var ret MarketDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetTaxCalculatorOk() (*MarketDataRelationshipsTaxCalculator, bool) {
	if o == nil || IsNil(o.TaxCalculator) {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasTaxCalculator() bool {
	if o != nil && !IsNil(o.TaxCalculator) {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given MarketDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *TaxCategoryDataRelationships) SetTaxCalculator(v MarketDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *TaxCategoryDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *TaxCategoryDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o TaxCategoryDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCategoryDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.TaxCalculator) {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableTaxCategoryDataRelationships struct {
	value *TaxCategoryDataRelationships
	isSet bool
}

func (v NullableTaxCategoryDataRelationships) Get() *TaxCategoryDataRelationships {
	return v.value
}

func (v *NullableTaxCategoryDataRelationships) Set(val *TaxCategoryDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryDataRelationships(val *TaxCategoryDataRelationships) *NullableTaxCategoryDataRelationships {
	return &NullableTaxCategoryDataRelationships{value: val, isSet: true}
}

func (v NullableTaxCategoryDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
