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

// FreeShippingPromotionUpdateData struct for FreeShippingPromotionUpdateData
type FreeShippingPromotionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                                      `json:"id"`
	Attributes    PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes `json:"attributes"`
	Relationships *ExternalPromotionCreateDataRelationships                                   `json:"relationships,omitempty"`
}

// NewFreeShippingPromotionUpdateData instantiates a new FreeShippingPromotionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionUpdateData(type_ string, id string, attributes PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes) *FreeShippingPromotionUpdateData {
	this := FreeShippingPromotionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewFreeShippingPromotionUpdateDataWithDefaults instantiates a new FreeShippingPromotionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionUpdateDataWithDefaults() *FreeShippingPromotionUpdateData {
	this := FreeShippingPromotionUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *FreeShippingPromotionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeShippingPromotionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FreeShippingPromotionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FreeShippingPromotionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *FreeShippingPromotionUpdateData) GetAttributes() PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionUpdateData) GetAttributesOk() (*PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FreeShippingPromotionUpdateData) SetAttributes(v PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FreeShippingPromotionUpdateData) GetRelationships() ExternalPromotionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPromotionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionUpdateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeShippingPromotionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPromotionCreateDataRelationships and assigns it to the Relationships field.
func (o *FreeShippingPromotionUpdateData) SetRelationships(v ExternalPromotionCreateDataRelationships) {
	o.Relationships = &v
}

func (o FreeShippingPromotionUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableFreeShippingPromotionUpdateData struct {
	value *FreeShippingPromotionUpdateData
	isSet bool
}

func (v NullableFreeShippingPromotionUpdateData) Get() *FreeShippingPromotionUpdateData {
	return v.value
}

func (v *NullableFreeShippingPromotionUpdateData) Set(val *FreeShippingPromotionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotionUpdateData(val *FreeShippingPromotionUpdateData) *NullableFreeShippingPromotionUpdateData {
	return &NullableFreeShippingPromotionUpdateData{value: val, isSet: true}
}

func (v NullableFreeShippingPromotionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
