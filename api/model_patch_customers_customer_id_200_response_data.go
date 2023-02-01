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

// PATCHCustomersCustomerId200ResponseData struct for PATCHCustomersCustomerId200ResponseData
type PATCHCustomersCustomerId200ResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                            `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks             `json:"links,omitempty"`
	Attributes    *PATCHCustomersCustomerId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *GETCustomers200ResponseDataInnerRelationships     `json:"relationships,omitempty"`
}

// NewPATCHCustomersCustomerId200ResponseData instantiates a new PATCHCustomersCustomerId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomersCustomerId200ResponseData() *PATCHCustomersCustomerId200ResponseData {
	this := PATCHCustomersCustomerId200ResponseData{}
	return &this
}

// NewPATCHCustomersCustomerId200ResponseDataWithDefaults instantiates a new PATCHCustomersCustomerId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomersCustomerId200ResponseDataWithDefaults() *PATCHCustomersCustomerId200ResponseData {
	this := PATCHCustomersCustomerId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PATCHCustomersCustomerId200ResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PATCHCustomersCustomerId200ResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *PATCHCustomersCustomerId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseData) GetAttributes() PATCHCustomersCustomerId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PATCHCustomersCustomerId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseData) GetAttributesOk() (*PATCHCustomersCustomerId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHCustomersCustomerId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHCustomersCustomerId200ResponseData) SetAttributes(v PATCHCustomersCustomerId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHCustomersCustomerId200ResponseData) GetRelationships() GETCustomers200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETCustomers200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomersCustomerId200ResponseData) GetRelationshipsOk() (*GETCustomers200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHCustomersCustomerId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETCustomers200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *PATCHCustomersCustomerId200ResponseData) SetRelationships(v GETCustomers200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o PATCHCustomersCustomerId200ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePATCHCustomersCustomerId200ResponseData struct {
	value *PATCHCustomersCustomerId200ResponseData
	isSet bool
}

func (v NullablePATCHCustomersCustomerId200ResponseData) Get() *PATCHCustomersCustomerId200ResponseData {
	return v.value
}

func (v *NullablePATCHCustomersCustomerId200ResponseData) Set(val *PATCHCustomersCustomerId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomersCustomerId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomersCustomerId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomersCustomerId200ResponseData(val *PATCHCustomersCustomerId200ResponseData) *NullablePATCHCustomersCustomerId200ResponseData {
	return &NullablePATCHCustomersCustomerId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHCustomersCustomerId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomersCustomerId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
