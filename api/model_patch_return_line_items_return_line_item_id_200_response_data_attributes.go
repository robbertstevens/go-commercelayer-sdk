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

// PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes struct for PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes
type PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes struct {
	// The line item quantity.
	Quantity *int32 `json:"quantity,omitempty"`
	// Send this attribute if you want to restock the line item.
	Restock *bool `json:"_restock,omitempty"`
	// Set of key-value pairs that you can use to add details about return reason.
	ReturnReason map[string]interface{} `json:"return_reason,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes instantiates a new PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes() *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes {
	this := PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes{}
	return &this
}

// NewPATCHReturnLineItemsReturnLineItemId200ResponseDataAttributesWithDefaults instantiates a new PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHReturnLineItemsReturnLineItemId200ResponseDataAttributesWithDefaults() *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes {
	this := PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetRestock returns the Restock field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetRestock() bool {
	if o == nil || o.Restock == nil {
		var ret bool
		return ret
	}
	return *o.Restock
}

// GetRestockOk returns a tuple with the Restock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetRestockOk() (*bool, bool) {
	if o == nil || o.Restock == nil {
		return nil, false
	}
	return o.Restock, true
}

// HasRestock returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasRestock() bool {
	if o != nil && o.Restock != nil {
		return true
	}

	return false
}

// SetRestock gets a reference to the given bool and assigns it to the Restock field.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetRestock(v bool) {
	o.Restock = &v
}

// GetReturnReason returns the ReturnReason field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReturnReason() map[string]interface{} {
	if o == nil || o.ReturnReason == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ReturnReason
}

// GetReturnReasonOk returns a tuple with the ReturnReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReturnReasonOk() (map[string]interface{}, bool) {
	if o == nil || o.ReturnReason == nil {
		return nil, false
	}
	return o.ReturnReason, true
}

// HasReturnReason returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasReturnReason() bool {
	if o != nil && o.ReturnReason != nil {
		return true
	}

	return false
}

// SetReturnReason gets a reference to the given map[string]interface{} and assigns it to the ReturnReason field.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReturnReason(v map[string]interface{}) {
	o.ReturnReason = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Restock != nil {
		toSerialize["_restock"] = o.Restock
	}
	if o.ReturnReason != nil {
		toSerialize["return_reason"] = o.ReturnReason
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

type NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes struct {
	value *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) Get() *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) Set(val *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes(val *PATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) *NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes {
	return &NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
