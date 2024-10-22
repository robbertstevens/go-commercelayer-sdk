/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTBillingInfoValidationRules201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTBillingInfoValidationRules201ResponseDataRelationships{}

// POSTBillingInfoValidationRules201ResponseDataRelationships struct for POSTBillingInfoValidationRules201ResponseDataRelationships
type POSTBillingInfoValidationRules201ResponseDataRelationships struct {
	Market      *POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket        `json:"market,omitempty"`
	Attachments *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions    *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTBillingInfoValidationRules201ResponseDataRelationships instantiates a new POSTBillingInfoValidationRules201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBillingInfoValidationRules201ResponseDataRelationships() *POSTBillingInfoValidationRules201ResponseDataRelationships {
	this := POSTBillingInfoValidationRules201ResponseDataRelationships{}
	return &this
}

// NewPOSTBillingInfoValidationRules201ResponseDataRelationshipsWithDefaults instantiates a new POSTBillingInfoValidationRules201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBillingInfoValidationRules201ResponseDataRelationshipsWithDefaults() *POSTBillingInfoValidationRules201ResponseDataRelationships {
	this := POSTBillingInfoValidationRules201ResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) GetMarket() POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket and assigns it to the Market field.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) SetMarket(v POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket) {
	o.Market = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTBillingInfoValidationRules201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTBillingInfoValidationRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTBillingInfoValidationRules201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTBillingInfoValidationRules201ResponseDataRelationships struct {
	value *POSTBillingInfoValidationRules201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTBillingInfoValidationRules201ResponseDataRelationships) Get() *POSTBillingInfoValidationRules201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTBillingInfoValidationRules201ResponseDataRelationships) Set(val *POSTBillingInfoValidationRules201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBillingInfoValidationRules201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBillingInfoValidationRules201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBillingInfoValidationRules201ResponseDataRelationships(val *POSTBillingInfoValidationRules201ResponseDataRelationships) *NullablePOSTBillingInfoValidationRules201ResponseDataRelationships {
	return &NullablePOSTBillingInfoValidationRules201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTBillingInfoValidationRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBillingInfoValidationRules201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
