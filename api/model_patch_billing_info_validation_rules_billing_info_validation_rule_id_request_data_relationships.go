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

// checks if the PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships{}

// PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships struct for PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships
type PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships struct {
	Market *POSTBillingInfoValidationRulesRequestDataRelationshipsMarket `json:"market,omitempty"`
}

// NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships instantiates a new PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships() *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships {
	this := PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships{}
	return &this
}

// NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationshipsWithDefaults instantiates a new PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationshipsWithDefaults() *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships {
	this := PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRulesRequestDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRulesRequestDataRelationshipsMarket and assigns it to the Market field.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket) {
	o.Market = &v
}

func (o PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	return toSerialize, nil
}

type NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships struct {
	value *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) Get() *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) Set(val *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships(val *PATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships {
	return &NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
