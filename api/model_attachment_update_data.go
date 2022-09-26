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

// AttachmentUpdateData struct for AttachmentUpdateData
type AttachmentUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                `json:"id"`
	Attributes    PATCHAttachmentsAttachmentId200ResponseDataAttributes `json:"attributes"`
	Relationships *AttachmentDataRelationships                          `json:"relationships,omitempty"`
}

// NewAttachmentUpdateData instantiates a new AttachmentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentUpdateData(type_ string, id string, attributes PATCHAttachmentsAttachmentId200ResponseDataAttributes) *AttachmentUpdateData {
	this := AttachmentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAttachmentUpdateDataWithDefaults instantiates a new AttachmentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentUpdateDataWithDefaults() *AttachmentUpdateData {
	this := AttachmentUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *AttachmentUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AttachmentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AttachmentUpdateData) GetAttributes() PATCHAttachmentsAttachmentId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHAttachmentsAttachmentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentUpdateData) GetAttributesOk() (*PATCHAttachmentsAttachmentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AttachmentUpdateData) SetAttributes(v PATCHAttachmentsAttachmentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AttachmentUpdateData) GetRelationships() AttachmentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AttachmentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentUpdateData) GetRelationshipsOk() (*AttachmentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AttachmentUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AttachmentDataRelationships and assigns it to the Relationships field.
func (o *AttachmentUpdateData) SetRelationships(v AttachmentDataRelationships) {
	o.Relationships = &v
}

func (o AttachmentUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableAttachmentUpdateData struct {
	value *AttachmentUpdateData
	isSet bool
}

func (v NullableAttachmentUpdateData) Get() *AttachmentUpdateData {
	return v.value
}

func (v *NullableAttachmentUpdateData) Set(val *AttachmentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentUpdateData(val *AttachmentUpdateData) *NullableAttachmentUpdateData {
	return &NullableAttachmentUpdateData{value: val, isSet: true}
}

func (v NullableAttachmentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
