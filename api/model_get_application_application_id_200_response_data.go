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

// GETApplicationApplicationId200ResponseData struct for GETApplicationApplicationId200ResponseData
type GETApplicationApplicationId200ResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                               `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                `json:"links,omitempty"`
	Attributes    *GETApplicationApplicationId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships map[string]interface{}                                `json:"relationships,omitempty"`
}

// NewGETApplicationApplicationId200ResponseData instantiates a new GETApplicationApplicationId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETApplicationApplicationId200ResponseData() *GETApplicationApplicationId200ResponseData {
	this := GETApplicationApplicationId200ResponseData{}
	return &this
}

// NewGETApplicationApplicationId200ResponseDataWithDefaults instantiates a new GETApplicationApplicationId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETApplicationApplicationId200ResponseDataWithDefaults() *GETApplicationApplicationId200ResponseData {
	this := GETApplicationApplicationId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETApplicationApplicationId200ResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETApplicationApplicationId200ResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETApplicationApplicationId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseData) GetAttributes() GETApplicationApplicationId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETApplicationApplicationId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseData) GetAttributesOk() (*GETApplicationApplicationId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETApplicationApplicationId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *GETApplicationApplicationId200ResponseData) SetAttributes(v GETApplicationApplicationId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *GETApplicationApplicationId200ResponseData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o GETApplicationApplicationId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableGETApplicationApplicationId200ResponseData struct {
	value *GETApplicationApplicationId200ResponseData
	isSet bool
}

func (v NullableGETApplicationApplicationId200ResponseData) Get() *GETApplicationApplicationId200ResponseData {
	return v.value
}

func (v *NullableGETApplicationApplicationId200ResponseData) Set(val *GETApplicationApplicationId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETApplicationApplicationId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETApplicationApplicationId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETApplicationApplicationId200ResponseData(val *GETApplicationApplicationId200ResponseData) *NullableGETApplicationApplicationId200ResponseData {
	return &NullableGETApplicationApplicationId200ResponseData{value: val, isSet: true}
}

func (v NullableGETApplicationApplicationId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETApplicationApplicationId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
