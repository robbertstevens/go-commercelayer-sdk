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

// checks if the CouponRecipientCreateDataRelationshipsCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponRecipientCreateDataRelationshipsCustomer{}

// CouponRecipientCreateDataRelationshipsCustomer struct for CouponRecipientCreateDataRelationshipsCustomer
type CouponRecipientCreateDataRelationshipsCustomer struct {
	Data CouponRecipientDataRelationshipsCustomerData `json:"data"`
}

// NewCouponRecipientCreateDataRelationshipsCustomer instantiates a new CouponRecipientCreateDataRelationshipsCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponRecipientCreateDataRelationshipsCustomer(data CouponRecipientDataRelationshipsCustomerData) *CouponRecipientCreateDataRelationshipsCustomer {
	this := CouponRecipientCreateDataRelationshipsCustomer{}
	this.Data = data
	return &this
}

// NewCouponRecipientCreateDataRelationshipsCustomerWithDefaults instantiates a new CouponRecipientCreateDataRelationshipsCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponRecipientCreateDataRelationshipsCustomerWithDefaults() *CouponRecipientCreateDataRelationshipsCustomer {
	this := CouponRecipientCreateDataRelationshipsCustomer{}
	return &this
}

// GetData returns the Data field value
func (o *CouponRecipientCreateDataRelationshipsCustomer) GetData() CouponRecipientDataRelationshipsCustomerData {
	if o == nil {
		var ret CouponRecipientDataRelationshipsCustomerData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponRecipientCreateDataRelationshipsCustomer) GetDataOk() (*CouponRecipientDataRelationshipsCustomerData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponRecipientCreateDataRelationshipsCustomer) SetData(v CouponRecipientDataRelationshipsCustomerData) {
	o.Data = v
}

func (o CouponRecipientCreateDataRelationshipsCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponRecipientCreateDataRelationshipsCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCouponRecipientCreateDataRelationshipsCustomer struct {
	value *CouponRecipientCreateDataRelationshipsCustomer
	isSet bool
}

func (v NullableCouponRecipientCreateDataRelationshipsCustomer) Get() *CouponRecipientCreateDataRelationshipsCustomer {
	return v.value
}

func (v *NullableCouponRecipientCreateDataRelationshipsCustomer) Set(val *CouponRecipientCreateDataRelationshipsCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponRecipientCreateDataRelationshipsCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponRecipientCreateDataRelationshipsCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponRecipientCreateDataRelationshipsCustomer(val *CouponRecipientCreateDataRelationshipsCustomer) *NullableCouponRecipientCreateDataRelationshipsCustomer {
	return &NullableCouponRecipientCreateDataRelationshipsCustomer{value: val, isSet: true}
}

func (v NullableCouponRecipientCreateDataRelationshipsCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponRecipientCreateDataRelationshipsCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
