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

// CouponRecipientUpdateData struct for CouponRecipientUpdateData
type CouponRecipientUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
	Attributes PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes `json:"attributes"`
	Relationships *CouponRecipientCreateDataRelationships `json:"relationships,omitempty"`
}

// NewCouponRecipientUpdateData instantiates a new CouponRecipientUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponRecipientUpdateData(type_ string, id string, attributes PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) *CouponRecipientUpdateData {
	this := CouponRecipientUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCouponRecipientUpdateDataWithDefaults instantiates a new CouponRecipientUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponRecipientUpdateDataWithDefaults() *CouponRecipientUpdateData {
	this := CouponRecipientUpdateData{}
	var type_ string = "coupon_recipients"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CouponRecipientUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponRecipientUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponRecipientUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CouponRecipientUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CouponRecipientUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CouponRecipientUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CouponRecipientUpdateData) GetAttributes() PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponRecipientUpdateData) GetAttributesOk() (*PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponRecipientUpdateData) SetAttributes(v PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponRecipientUpdateData) GetRelationships() CouponRecipientCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponRecipientCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponRecipientUpdateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponRecipientUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientCreateDataRelationships and assigns it to the Relationships field.
func (o *CouponRecipientUpdateData) SetRelationships(v CouponRecipientCreateDataRelationships) {
	o.Relationships = &v
}

func (o CouponRecipientUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableCouponRecipientUpdateData struct {
	value *CouponRecipientUpdateData
	isSet bool
}

func (v NullableCouponRecipientUpdateData) Get() *CouponRecipientUpdateData {
	return v.value
}

func (v *NullableCouponRecipientUpdateData) Set(val *CouponRecipientUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponRecipientUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponRecipientUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponRecipientUpdateData(val *CouponRecipientUpdateData) *NullableCouponRecipientUpdateData {
	return &NullableCouponRecipientUpdateData{value: val, isSet: true}
}

func (v NullableCouponRecipientUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponRecipientUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


