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

// CaptureUpdateData struct for CaptureUpdateData
type CaptureUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                          `json:"id"`
	Attributes    PATCHCapturesCaptureId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                          `json:"relationships,omitempty"`
}

// NewCaptureUpdateData instantiates a new CaptureUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureUpdateData(type_ string, id string, attributes PATCHCapturesCaptureId200ResponseDataAttributes) *CaptureUpdateData {
	this := CaptureUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCaptureUpdateDataWithDefaults instantiates a new CaptureUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureUpdateDataWithDefaults() *CaptureUpdateData {
	this := CaptureUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *CaptureUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptureUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CaptureUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaptureUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaptureUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CaptureUpdateData) GetAttributes() PATCHCapturesCaptureId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCapturesCaptureId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CaptureUpdateData) GetAttributesOk() (*PATCHCapturesCaptureId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CaptureUpdateData) SetAttributes(v PATCHCapturesCaptureId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CaptureUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CaptureUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *CaptureUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o CaptureUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableCaptureUpdateData struct {
	value *CaptureUpdateData
	isSet bool
}

func (v NullableCaptureUpdateData) Get() *CaptureUpdateData {
	return v.value
}

func (v *NullableCaptureUpdateData) Set(val *CaptureUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureUpdateData(val *CaptureUpdateData) *NullableCaptureUpdateData {
	return &NullableCaptureUpdateData{value: val, isSet: true}
}

func (v NullableCaptureUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
