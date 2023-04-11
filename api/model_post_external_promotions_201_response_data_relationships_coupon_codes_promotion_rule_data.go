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

// checks if the POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData{}

// POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData struct for POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData
type POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData instantiates a new POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData() *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData {
	this := POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData{}
	return &this
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleDataWithDefaults instantiates a new POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleDataWithDefaults() *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData {
	this := POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData struct {
	value *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData
	isSet bool
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) Get() *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData {
	return v.value
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) Set(val *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData(val *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData {
	return &NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData{value: val, isSet: true}
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
