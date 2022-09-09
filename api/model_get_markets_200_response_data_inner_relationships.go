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

// GETMarkets200ResponseDataInnerRelationships struct for GETMarkets200ResponseDataInnerRelationships
type GETMarkets200ResponseDataInnerRelationships struct {
	Merchant *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"merchant,omitempty"`
	PriceList *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"price_list,omitempty"`
	InventoryModel *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"inventory_model,omitempty"`
	TaxCalculator *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"tax_calculator,omitempty"`
	CustomerGroup *GETAddresses200ResponseDataInnerRelationshipsGeocoder `json:"customer_group,omitempty"`
	Attachments *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods `json:"attachments,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationships instantiates a new GETMarkets200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationships() *GETMarkets200ResponseDataInnerRelationships {
	this := GETMarkets200ResponseDataInnerRelationships{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsWithDefaults() *GETMarkets200ResponseDataInnerRelationships {
	this := GETMarkets200ResponseDataInnerRelationships{}
	return &this
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationships) GetMerchant() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.Merchant == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) GetMerchantOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.Merchant == nil {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) HasMerchant() bool {
	if o != nil && o.Merchant != nil {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the Merchant field.
func (o *GETMarkets200ResponseDataInnerRelationships) SetMerchant(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.Merchant = &v
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationships) GetPriceList() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.PriceList == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) GetPriceListOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.PriceList == nil {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) HasPriceList() bool {
	if o != nil && o.PriceList != nil {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the PriceList field.
func (o *GETMarkets200ResponseDataInnerRelationships) SetPriceList(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.PriceList = &v
}

// GetInventoryModel returns the InventoryModel field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationships) GetInventoryModel() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.InventoryModel == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) GetInventoryModelOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.InventoryModel == nil {
		return nil, false
	}
	return o.InventoryModel, true
}

// HasInventoryModel returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) HasInventoryModel() bool {
	if o != nil && o.InventoryModel != nil {
		return true
	}

	return false
}

// SetInventoryModel gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the InventoryModel field.
func (o *GETMarkets200ResponseDataInnerRelationships) SetInventoryModel(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.InventoryModel = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationships) GetTaxCalculator() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.TaxCalculator == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) GetTaxCalculatorOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.TaxCalculator == nil {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) HasTaxCalculator() bool {
	if o != nil && o.TaxCalculator != nil {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the TaxCalculator field.
func (o *GETMarkets200ResponseDataInnerRelationships) SetTaxCalculator(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.TaxCalculator = &v
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationships) GetCustomerGroup() GETAddresses200ResponseDataInnerRelationshipsGeocoder {
	if o == nil || o.CustomerGroup == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoder
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) GetCustomerGroupOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoder, bool) {
	if o == nil || o.CustomerGroup == nil {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) HasCustomerGroup() bool {
	if o != nil && o.CustomerGroup != nil {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoder and assigns it to the CustomerGroup field.
func (o *GETMarkets200ResponseDataInnerRelationships) SetCustomerGroup(v GETAddresses200ResponseDataInnerRelationshipsGeocoder) {
	o.CustomerGroup = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationships) GetAttachments() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.Attachments == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the Attachments field.
func (o *GETMarkets200ResponseDataInnerRelationships) SetAttachments(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.Attachments = &v
}

func (o GETMarkets200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Merchant != nil {
		toSerialize["merchant"] = o.Merchant
	}
	if o.PriceList != nil {
		toSerialize["price_list"] = o.PriceList
	}
	if o.InventoryModel != nil {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	if o.TaxCalculator != nil {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if o.CustomerGroup != nil {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationships struct {
	value *GETMarkets200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationships) Get() *GETMarkets200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationships) Set(val *GETMarkets200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationships(val *GETMarkets200ResponseDataInnerRelationships) *NullableGETMarkets200ResponseDataInnerRelationships {
	return &NullableGETMarkets200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


