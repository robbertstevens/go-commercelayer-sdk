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

// checks if the POSTCoupons201ResponseDataRelationshipsPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCoupons201ResponseDataRelationshipsPromotionRule{}

// POSTCoupons201ResponseDataRelationshipsPromotionRule struct for POSTCoupons201ResponseDataRelationshipsPromotionRule
type POSTCoupons201ResponseDataRelationshipsPromotionRule struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *POSTCoupons201ResponseDataRelationshipsPromotionRuleData `json:"data,omitempty"`
}

// NewPOSTCoupons201ResponseDataRelationshipsPromotionRule instantiates a new POSTCoupons201ResponseDataRelationshipsPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCoupons201ResponseDataRelationshipsPromotionRule() *POSTCoupons201ResponseDataRelationshipsPromotionRule {
	this := POSTCoupons201ResponseDataRelationshipsPromotionRule{}
	return &this
}

// NewPOSTCoupons201ResponseDataRelationshipsPromotionRuleWithDefaults instantiates a new POSTCoupons201ResponseDataRelationshipsPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCoupons201ResponseDataRelationshipsPromotionRuleWithDefaults() *POSTCoupons201ResponseDataRelationshipsPromotionRule {
	this := POSTCoupons201ResponseDataRelationshipsPromotionRule{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) GetData() POSTCoupons201ResponseDataRelationshipsPromotionRuleData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCoupons201ResponseDataRelationshipsPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) GetDataOk() (*POSTCoupons201ResponseDataRelationshipsPromotionRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCoupons201ResponseDataRelationshipsPromotionRuleData and assigns it to the Data field.
func (o *POSTCoupons201ResponseDataRelationshipsPromotionRule) SetData(v POSTCoupons201ResponseDataRelationshipsPromotionRuleData) {
	o.Data = &v
}

func (o POSTCoupons201ResponseDataRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCoupons201ResponseDataRelationshipsPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule struct {
	value *POSTCoupons201ResponseDataRelationshipsPromotionRule
	isSet bool
}

func (v NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule) Get() *POSTCoupons201ResponseDataRelationshipsPromotionRule {
	return v.value
}

func (v *NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule) Set(val *POSTCoupons201ResponseDataRelationshipsPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCoupons201ResponseDataRelationshipsPromotionRule(val *POSTCoupons201ResponseDataRelationshipsPromotionRule) *NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule {
	return &NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule{value: val, isSet: true}
}

func (v NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCoupons201ResponseDataRelationshipsPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
