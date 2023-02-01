/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BingGeocoderCreateData struct for BingGeocoderCreateData
type BingGeocoderCreateData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    POSTBingGeocoders201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                     `json:"relationships,omitempty"`
}

// NewBingGeocoderCreateData instantiates a new BingGeocoderCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderCreateData(type_ string, attributes POSTBingGeocoders201ResponseDataAttributes) *BingGeocoderCreateData {
	this := BingGeocoderCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBingGeocoderCreateDataWithDefaults instantiates a new BingGeocoderCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderCreateDataWithDefaults() *BingGeocoderCreateData {
	this := BingGeocoderCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *BingGeocoderCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BingGeocoderCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BingGeocoderCreateData) GetAttributes() POSTBingGeocoders201ResponseDataAttributes {
	if o == nil {
		var ret POSTBingGeocoders201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderCreateData) GetAttributesOk() (*POSTBingGeocoders201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BingGeocoderCreateData) SetAttributes(v POSTBingGeocoders201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BingGeocoderCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BingGeocoderCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *BingGeocoderCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o BingGeocoderCreateData) MarshalJSON() ([]byte, error) {
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

type NullableBingGeocoderCreateData struct {
	value *BingGeocoderCreateData
	isSet bool
}

func (v NullableBingGeocoderCreateData) Get() *BingGeocoderCreateData {
	return v.value
}

func (v *NullableBingGeocoderCreateData) Set(val *BingGeocoderCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderCreateData(val *BingGeocoderCreateData) *NullableBingGeocoderCreateData {
	return &NullableBingGeocoderCreateData{value: val, isSet: true}
}

func (v NullableBingGeocoderCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
