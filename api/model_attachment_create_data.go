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

// checks if the AttachmentCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentCreateData{}

// AttachmentCreateData struct for AttachmentCreateData
type AttachmentCreateData struct {
	// The resource's type
	Type          interface{}                          `json:"type"`
	Attributes    POSTAttachmentsRequestDataAttributes `json:"attributes"`
	Relationships *AttachmentCreateDataRelationships   `json:"relationships,omitempty"`
}

// NewAttachmentCreateData instantiates a new AttachmentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentCreateData(type_ interface{}, attributes POSTAttachmentsRequestDataAttributes) *AttachmentCreateData {
	this := AttachmentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAttachmentCreateDataWithDefaults instantiates a new AttachmentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentCreateDataWithDefaults() *AttachmentCreateData {
	this := AttachmentCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AttachmentCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AttachmentCreateData) GetAttributes() POSTAttachmentsRequestDataAttributes {
	if o == nil {
		var ret POSTAttachmentsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentCreateData) GetAttributesOk() (*POSTAttachmentsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AttachmentCreateData) SetAttributes(v POSTAttachmentsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AttachmentCreateData) GetRelationships() AttachmentCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AttachmentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentCreateData) GetRelationshipsOk() (*AttachmentCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AttachmentCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AttachmentCreateDataRelationships and assigns it to the Relationships field.
func (o *AttachmentCreateData) SetRelationships(v AttachmentCreateDataRelationships) {
	o.Relationships = &v
}

func (o AttachmentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAttachmentCreateData struct {
	value *AttachmentCreateData
	isSet bool
}

func (v NullableAttachmentCreateData) Get() *AttachmentCreateData {
	return v.value
}

func (v *NullableAttachmentCreateData) Set(val *AttachmentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentCreateData(val *AttachmentCreateData) *NullableAttachmentCreateData {
	return &NullableAttachmentCreateData{value: val, isSet: true}
}

func (v NullableAttachmentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
