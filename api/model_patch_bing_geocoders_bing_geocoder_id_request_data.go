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

// checks if the PATCHBingGeocodersBingGeocoderIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBingGeocodersBingGeocoderIdRequestData{}

// PATCHBingGeocodersBingGeocoderIdRequestData struct for PATCHBingGeocodersBingGeocoderIdRequestData
type PATCHBingGeocodersBingGeocoderIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                           `json:"id"`
	Attributes    PATCHBingGeocodersBingGeocoderIdRequestDataAttributes `json:"attributes"`
	Relationships interface{}                                           `json:"relationships,omitempty"`
}

// NewPATCHBingGeocodersBingGeocoderIdRequestData instantiates a new PATCHBingGeocodersBingGeocoderIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBingGeocodersBingGeocoderIdRequestData(type_ interface{}, id interface{}, attributes PATCHBingGeocodersBingGeocoderIdRequestDataAttributes) *PATCHBingGeocodersBingGeocoderIdRequestData {
	this := PATCHBingGeocodersBingGeocoderIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHBingGeocodersBingGeocoderIdRequestDataWithDefaults instantiates a new PATCHBingGeocodersBingGeocoderIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBingGeocodersBingGeocoderIdRequestDataWithDefaults() *PATCHBingGeocodersBingGeocoderIdRequestData {
	this := PATCHBingGeocodersBingGeocoderIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetAttributes() PATCHBingGeocodersBingGeocoderIdRequestDataAttributes {
	if o == nil {
		var ret PATCHBingGeocodersBingGeocoderIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetAttributesOk() (*PATCHBingGeocodersBingGeocoderIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) SetAttributes(v PATCHBingGeocodersBingGeocoderIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *PATCHBingGeocodersBingGeocoderIdRequestData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o PATCHBingGeocodersBingGeocoderIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBingGeocodersBingGeocoderIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHBingGeocodersBingGeocoderIdRequestData struct {
	value *PATCHBingGeocodersBingGeocoderIdRequestData
	isSet bool
}

func (v NullablePATCHBingGeocodersBingGeocoderIdRequestData) Get() *PATCHBingGeocodersBingGeocoderIdRequestData {
	return v.value
}

func (v *NullablePATCHBingGeocodersBingGeocoderIdRequestData) Set(val *PATCHBingGeocodersBingGeocoderIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBingGeocodersBingGeocoderIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBingGeocodersBingGeocoderIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBingGeocodersBingGeocoderIdRequestData(val *PATCHBingGeocodersBingGeocoderIdRequestData) *NullablePATCHBingGeocodersBingGeocoderIdRequestData {
	return &NullablePATCHBingGeocodersBingGeocoderIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHBingGeocodersBingGeocoderIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBingGeocodersBingGeocoderIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
