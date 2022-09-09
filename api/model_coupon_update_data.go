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

// CouponUpdateData struct for CouponUpdateData
type CouponUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
	Attributes PATCHCouponsCouponId200ResponseDataAttributes `json:"attributes"`
	Relationships *CouponDataRelationships `json:"relationships,omitempty"`
}

// NewCouponUpdateData instantiates a new CouponUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponUpdateData(type_ string, id string, attributes PATCHCouponsCouponId200ResponseDataAttributes) *CouponUpdateData {
	this := CouponUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCouponUpdateDataWithDefaults instantiates a new CouponUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponUpdateDataWithDefaults() *CouponUpdateData {
	this := CouponUpdateData{}
	var type_ string = "coupons"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CouponUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CouponUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CouponUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CouponUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CouponUpdateData) GetAttributes() PATCHCouponsCouponId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCouponsCouponId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponUpdateData) GetAttributesOk() (*PATCHCouponsCouponId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponUpdateData) SetAttributes(v PATCHCouponsCouponId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponUpdateData) GetRelationships() CouponDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponUpdateData) GetRelationshipsOk() (*CouponDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponDataRelationships and assigns it to the Relationships field.
func (o *CouponUpdateData) SetRelationships(v CouponDataRelationships) {
	o.Relationships = &v
}

func (o CouponUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCouponUpdateData struct {
	value *CouponUpdateData
	isSet bool
}

func (v NullableCouponUpdateData) Get() *CouponUpdateData {
	return v.value
}

func (v *NullableCouponUpdateData) Set(val *CouponUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponUpdateData(val *CouponUpdateData) *NullableCouponUpdateData {
	return &NullableCouponUpdateData{value: val, isSet: true}
}

func (v NullableCouponUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


