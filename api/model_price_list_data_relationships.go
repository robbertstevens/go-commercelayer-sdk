/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PriceListDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListDataRelationships{}

// PriceListDataRelationships struct for PriceListDataRelationships
type PriceListDataRelationships struct {
	Prices              *PriceFrequencyTierDataRelationshipsPrice   `json:"prices,omitempty"`
	PriceListSchedulers *MarketDataRelationshipsPriceListSchedulers `json:"price_list_schedulers,omitempty"`
	Attachments         *AuthorizationDataRelationshipsAttachments  `json:"attachments,omitempty"`
	Versions            *AddressDataRelationshipsVersions           `json:"versions,omitempty"`
}

// NewPriceListDataRelationships instantiates a new PriceListDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListDataRelationships() *PriceListDataRelationships {
	this := PriceListDataRelationships{}
	return &this
}

// NewPriceListDataRelationshipsWithDefaults instantiates a new PriceListDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListDataRelationshipsWithDefaults() *PriceListDataRelationships {
	this := PriceListDataRelationships{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *PriceListDataRelationships) GetPrices() PriceFrequencyTierDataRelationshipsPrice {
	if o == nil || IsNil(o.Prices) {
		var ret PriceFrequencyTierDataRelationshipsPrice
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListDataRelationships) GetPricesOk() (*PriceFrequencyTierDataRelationshipsPrice, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *PriceListDataRelationships) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given PriceFrequencyTierDataRelationshipsPrice and assigns it to the Prices field.
func (o *PriceListDataRelationships) SetPrices(v PriceFrequencyTierDataRelationshipsPrice) {
	o.Prices = &v
}

// GetPriceListSchedulers returns the PriceListSchedulers field value if set, zero value otherwise.
func (o *PriceListDataRelationships) GetPriceListSchedulers() MarketDataRelationshipsPriceListSchedulers {
	if o == nil || IsNil(o.PriceListSchedulers) {
		var ret MarketDataRelationshipsPriceListSchedulers
		return ret
	}
	return *o.PriceListSchedulers
}

// GetPriceListSchedulersOk returns a tuple with the PriceListSchedulers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListDataRelationships) GetPriceListSchedulersOk() (*MarketDataRelationshipsPriceListSchedulers, bool) {
	if o == nil || IsNil(o.PriceListSchedulers) {
		return nil, false
	}
	return o.PriceListSchedulers, true
}

// HasPriceListSchedulers returns a boolean if a field has been set.
func (o *PriceListDataRelationships) HasPriceListSchedulers() bool {
	if o != nil && !IsNil(o.PriceListSchedulers) {
		return true
	}

	return false
}

// SetPriceListSchedulers gets a reference to the given MarketDataRelationshipsPriceListSchedulers and assigns it to the PriceListSchedulers field.
func (o *PriceListDataRelationships) SetPriceListSchedulers(v MarketDataRelationshipsPriceListSchedulers) {
	o.PriceListSchedulers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PriceListDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PriceListDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *PriceListDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *PriceListDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *PriceListDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *PriceListDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o PriceListDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.PriceListSchedulers) {
		toSerialize["price_list_schedulers"] = o.PriceListSchedulers
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePriceListDataRelationships struct {
	value *PriceListDataRelationships
	isSet bool
}

func (v NullablePriceListDataRelationships) Get() *PriceListDataRelationships {
	return v.value
}

func (v *NullablePriceListDataRelationships) Set(val *PriceListDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListDataRelationships(val *PriceListDataRelationships) *NullablePriceListDataRelationships {
	return &NullablePriceListDataRelationships{value: val, isSet: true}
}

func (v NullablePriceListDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
