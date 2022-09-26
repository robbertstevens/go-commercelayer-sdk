/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShippingZoneData struct for ShippingZoneData
type ShippingZoneData struct {
	// The resource's type
	Type          string                                         `json:"type"`
	Attributes    GETShippingZones200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ShippingZoneDataRelationships                 `json:"relationships,omitempty"`
}

// NewShippingZoneData instantiates a new ShippingZoneData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingZoneData(type_ string, attributes GETShippingZones200ResponseDataInnerAttributes) *ShippingZoneData {
	this := ShippingZoneData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingZoneDataWithDefaults instantiates a new ShippingZoneData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingZoneDataWithDefaults() *ShippingZoneData {
	this := ShippingZoneData{}
	return &this
}

// GetType returns the Type field value
func (o *ShippingZoneData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingZoneData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingZoneData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingZoneData) GetAttributes() GETShippingZones200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETShippingZones200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingZoneData) GetAttributesOk() (*GETShippingZones200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingZoneData) SetAttributes(v GETShippingZones200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingZoneData) GetRelationships() ShippingZoneDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ShippingZoneDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingZoneData) GetRelationshipsOk() (*ShippingZoneDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingZoneData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingZoneDataRelationships and assigns it to the Relationships field.
func (o *ShippingZoneData) SetRelationships(v ShippingZoneDataRelationships) {
	o.Relationships = &v
}

func (o ShippingZoneData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableShippingZoneData struct {
	value *ShippingZoneData
	isSet bool
}

func (v NullableShippingZoneData) Get() *ShippingZoneData {
	return v.value
}

func (v *NullableShippingZoneData) Set(val *ShippingZoneData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingZoneData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingZoneData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingZoneData(val *ShippingZoneData) *NullableShippingZoneData {
	return &NullableShippingZoneData{value: val, isSet: true}
}

func (v NullableShippingZoneData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingZoneData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
