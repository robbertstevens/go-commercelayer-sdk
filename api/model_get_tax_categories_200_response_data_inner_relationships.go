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

// GETTaxCategories200ResponseDataInnerRelationships struct for GETTaxCategories200ResponseDataInnerRelationships
type GETTaxCategories200ResponseDataInnerRelationships struct {
	Sku *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"sku,omitempty"`
	TaxCalculator *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"tax_calculator,omitempty"`
	Attachments *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods `json:"attachments,omitempty"`
}

// NewGETTaxCategories200ResponseDataInnerRelationships instantiates a new GETTaxCategories200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxCategories200ResponseDataInnerRelationships() *GETTaxCategories200ResponseDataInnerRelationships {
	this := GETTaxCategories200ResponseDataInnerRelationships{}
	return &this
}

// NewGETTaxCategories200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETTaxCategories200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxCategories200ResponseDataInnerRelationshipsWithDefaults() *GETTaxCategories200ResponseDataInnerRelationships {
	this := GETTaxCategories200ResponseDataInnerRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GETTaxCategories200ResponseDataInnerRelationships) GetSku() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.Sku == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxCategories200ResponseDataInnerRelationships) GetSkuOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GETTaxCategories200ResponseDataInnerRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the Sku field.
func (o *GETTaxCategories200ResponseDataInnerRelationships) SetSku(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.Sku = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *GETTaxCategories200ResponseDataInnerRelationships) GetTaxCalculator() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.TaxCalculator == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxCategories200ResponseDataInnerRelationships) GetTaxCalculatorOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.TaxCalculator == nil {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *GETTaxCategories200ResponseDataInnerRelationships) HasTaxCalculator() bool {
	if o != nil && o.TaxCalculator != nil {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the TaxCalculator field.
func (o *GETTaxCategories200ResponseDataInnerRelationships) SetTaxCalculator(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.TaxCalculator = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETTaxCategories200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.Attachments == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxCategories200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETTaxCategories200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the Attachments field.
func (o *GETTaxCategories200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.Attachments = &v
}

func (o GETTaxCategories200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.TaxCalculator != nil {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETTaxCategories200ResponseDataInnerRelationships struct {
	value *GETTaxCategories200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETTaxCategories200ResponseDataInnerRelationships) Get() *GETTaxCategories200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETTaxCategories200ResponseDataInnerRelationships) Set(val *GETTaxCategories200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxCategories200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxCategories200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxCategories200ResponseDataInnerRelationships(val *GETTaxCategories200ResponseDataInnerRelationships) *NullableGETTaxCategories200ResponseDataInnerRelationships {
	return &NullableGETTaxCategories200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETTaxCategories200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxCategories200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


