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

// KlarnaGatewayUpdateData struct for KlarnaGatewayUpdateData
type KlarnaGatewayUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                      `json:"id"`
	Attributes    PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *KlarnaGatewayCreateDataRelationships                       `json:"relationships,omitempty"`
}

// NewKlarnaGatewayUpdateData instantiates a new KlarnaGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayUpdateData(type_ string, id string, attributes PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) *KlarnaGatewayUpdateData {
	this := KlarnaGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewKlarnaGatewayUpdateDataWithDefaults instantiates a new KlarnaGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayUpdateDataWithDefaults() *KlarnaGatewayUpdateData {
	this := KlarnaGatewayUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *KlarnaGatewayUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KlarnaGatewayUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *KlarnaGatewayUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KlarnaGatewayUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *KlarnaGatewayUpdateData) GetAttributes() PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayUpdateData) GetAttributesOk() (*PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *KlarnaGatewayUpdateData) SetAttributes(v PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *KlarnaGatewayUpdateData) GetRelationships() KlarnaGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret KlarnaGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayUpdateData) GetRelationshipsOk() (*KlarnaGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *KlarnaGatewayUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given KlarnaGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *KlarnaGatewayUpdateData) SetRelationships(v KlarnaGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o KlarnaGatewayUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableKlarnaGatewayUpdateData struct {
	value *KlarnaGatewayUpdateData
	isSet bool
}

func (v NullableKlarnaGatewayUpdateData) Get() *KlarnaGatewayUpdateData {
	return v.value
}

func (v *NullableKlarnaGatewayUpdateData) Set(val *KlarnaGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayUpdateData(val *KlarnaGatewayUpdateData) *NullableKlarnaGatewayUpdateData {
	return &NullableKlarnaGatewayUpdateData{value: val, isSet: true}
}

func (v NullableKlarnaGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
