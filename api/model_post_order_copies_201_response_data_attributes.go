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

// POSTOrderCopies201ResponseDataAttributes struct for POSTOrderCopies201ResponseDataAttributes
type POSTOrderCopies201ResponseDataAttributes struct {
	// Indicates if the target order must be placed upon copy.
	PlaceTargetOrder *bool `json:"place_target_order,omitempty"`
	// Indicates if the source order must be cancelled upon copy.
	CancelSourceOrder *bool `json:"cancel_source_order,omitempty"`
	// Indicates if the payment source within the source order customer's wallet must be copied.
	ReuseWallet *bool `json:"reuse_wallet,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTOrderCopies201ResponseDataAttributes instantiates a new POSTOrderCopies201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201ResponseDataAttributes() *POSTOrderCopies201ResponseDataAttributes {
	this := POSTOrderCopies201ResponseDataAttributes{}
	return &this
}

// NewPOSTOrderCopies201ResponseDataAttributesWithDefaults instantiates a new POSTOrderCopies201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseDataAttributesWithDefaults() *POSTOrderCopies201ResponseDataAttributes {
	this := POSTOrderCopies201ResponseDataAttributes{}
	return &this
}

// GetPlaceTargetOrder returns the PlaceTargetOrder field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrder() bool {
	if o == nil || o.PlaceTargetOrder == nil {
		var ret bool
		return ret
	}
	return *o.PlaceTargetOrder
}

// GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrderOk() (*bool, bool) {
	if o == nil || o.PlaceTargetOrder == nil {
		return nil, false
	}
	return o.PlaceTargetOrder, true
}

// HasPlaceTargetOrder returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasPlaceTargetOrder() bool {
	if o != nil && o.PlaceTargetOrder != nil {
		return true
	}

	return false
}

// SetPlaceTargetOrder gets a reference to the given bool and assigns it to the PlaceTargetOrder field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetPlaceTargetOrder(v bool) {
	o.PlaceTargetOrder = &v
}

// GetCancelSourceOrder returns the CancelSourceOrder field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrder() bool {
	if o == nil || o.CancelSourceOrder == nil {
		var ret bool
		return ret
	}
	return *o.CancelSourceOrder
}

// GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrderOk() (*bool, bool) {
	if o == nil || o.CancelSourceOrder == nil {
		return nil, false
	}
	return o.CancelSourceOrder, true
}

// HasCancelSourceOrder returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasCancelSourceOrder() bool {
	if o != nil && o.CancelSourceOrder != nil {
		return true
	}

	return false
}

// SetCancelSourceOrder gets a reference to the given bool and assigns it to the CancelSourceOrder field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetCancelSourceOrder(v bool) {
	o.CancelSourceOrder = &v
}

// GetReuseWallet returns the ReuseWallet field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWallet() bool {
	if o == nil || o.ReuseWallet == nil {
		var ret bool
		return ret
	}
	return *o.ReuseWallet
}

// GetReuseWalletOk returns a tuple with the ReuseWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWalletOk() (*bool, bool) {
	if o == nil || o.ReuseWallet == nil {
		return nil, false
	}
	return o.ReuseWallet, true
}

// HasReuseWallet returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasReuseWallet() bool {
	if o != nil && o.ReuseWallet != nil {
		return true
	}

	return false
}

// SetReuseWallet gets a reference to the given bool and assigns it to the ReuseWallet field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetReuseWallet(v bool) {
	o.ReuseWallet = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTOrderCopies201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PlaceTargetOrder != nil {
		toSerialize["place_target_order"] = o.PlaceTargetOrder
	}
	if o.CancelSourceOrder != nil {
		toSerialize["cancel_source_order"] = o.CancelSourceOrder
	}
	if o.ReuseWallet != nil {
		toSerialize["reuse_wallet"] = o.ReuseWallet
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTOrderCopies201ResponseDataAttributes struct {
	value *POSTOrderCopies201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrderCopies201ResponseDataAttributes) Get() *POSTOrderCopies201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderCopies201ResponseDataAttributes) Set(val *POSTOrderCopies201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201ResponseDataAttributes(val *POSTOrderCopies201ResponseDataAttributes) *NullablePOSTOrderCopies201ResponseDataAttributes {
	return &NullablePOSTOrderCopies201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
