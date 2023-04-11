/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTOrderSubscriptionsRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptionsRequestDataRelationships{}

// POSTOrderSubscriptionsRequestDataRelationships struct for POSTOrderSubscriptionsRequestDataRelationships
type POSTOrderSubscriptionsRequestDataRelationships struct {
	Market      *POSTBillingInfoValidationRulesRequestDataRelationshipsMarket `json:"market,omitempty"`
	SourceOrder POSTAdyenPaymentsRequestDataRelationshipsOrder                `json:"source_order"`
}

// NewPOSTOrderSubscriptionsRequestDataRelationships instantiates a new POSTOrderSubscriptionsRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptionsRequestDataRelationships(sourceOrder POSTAdyenPaymentsRequestDataRelationshipsOrder) *POSTOrderSubscriptionsRequestDataRelationships {
	this := POSTOrderSubscriptionsRequestDataRelationships{}
	this.SourceOrder = sourceOrder
	return &this
}

// NewPOSTOrderSubscriptionsRequestDataRelationshipsWithDefaults instantiates a new POSTOrderSubscriptionsRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptionsRequestDataRelationshipsWithDefaults() *POSTOrderSubscriptionsRequestDataRelationships {
	this := POSTOrderSubscriptionsRequestDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTOrderSubscriptionsRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRulesRequestDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptionsRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionsRequestDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRulesRequestDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTOrderSubscriptionsRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket) {
	o.Market = &v
}

// GetSourceOrder returns the SourceOrder field value
func (o *POSTOrderSubscriptionsRequestDataRelationships) GetSourceOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder {
	if o == nil {
		var ret POSTAdyenPaymentsRequestDataRelationshipsOrder
		return ret
	}

	return o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value
// and a boolean to check if the value has been set.
func (o *POSTOrderSubscriptionsRequestDataRelationships) GetSourceOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceOrder, true
}

// SetSourceOrder sets field value
func (o *POSTOrderSubscriptionsRequestDataRelationships) SetSourceOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder) {
	o.SourceOrder = v
}

func (o POSTOrderSubscriptionsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptionsRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	toSerialize["source_order"] = o.SourceOrder
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptionsRequestDataRelationships struct {
	value *POSTOrderSubscriptionsRequestDataRelationships
	isSet bool
}

func (v NullablePOSTOrderSubscriptionsRequestDataRelationships) Get() *POSTOrderSubscriptionsRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTOrderSubscriptionsRequestDataRelationships) Set(val *POSTOrderSubscriptionsRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptionsRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptionsRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptionsRequestDataRelationships(val *POSTOrderSubscriptionsRequestDataRelationships) *NullablePOSTOrderSubscriptionsRequestDataRelationships {
	return &NullablePOSTOrderSubscriptionsRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptionsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptionsRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
