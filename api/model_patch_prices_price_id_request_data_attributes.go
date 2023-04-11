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

// checks if the PATCHPricesPriceIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPricesPriceIdRequestDataAttributes{}

// PATCHPricesPriceIdRequestDataAttributes struct for PATCHPricesPriceIdRequestDataAttributes
type PATCHPricesPriceIdRequestDataAttributes struct {
	// The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The SKU price amount for the associated price list, in cents.
	AmountCents interface{} `json:"amount_cents,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents interface{} `json:"compare_at_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHPricesPriceIdRequestDataAttributes instantiates a new PATCHPricesPriceIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPricesPriceIdRequestDataAttributes() *PATCHPricesPriceIdRequestDataAttributes {
	this := PATCHPricesPriceIdRequestDataAttributes{}
	return &this
}

// NewPATCHPricesPriceIdRequestDataAttributesWithDefaults instantiates a new PATCHPricesPriceIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPricesPriceIdRequestDataAttributesWithDefaults() *PATCHPricesPriceIdRequestDataAttributes {
	this := PATCHPricesPriceIdRequestDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPricesPriceIdRequestDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPricesPriceIdRequestDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *PATCHPricesPriceIdRequestDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *PATCHPricesPriceIdRequestDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPricesPriceIdRequestDataAttributes) GetAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPricesPriceIdRequestDataAttributes) GetAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return &o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *PATCHPricesPriceIdRequestDataAttributes) HasAmountCents() bool {
	if o != nil && IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given interface{} and assigns it to the AmountCents field.
func (o *PATCHPricesPriceIdRequestDataAttributes) SetAmountCents(v interface{}) {
	o.AmountCents = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPricesPriceIdRequestDataAttributes) GetCompareAtAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPricesPriceIdRequestDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CompareAtAmountCents) {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *PATCHPricesPriceIdRequestDataAttributes) HasCompareAtAmountCents() bool {
	if o != nil && IsNil(o.CompareAtAmountCents) {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given interface{} and assigns it to the CompareAtAmountCents field.
func (o *PATCHPricesPriceIdRequestDataAttributes) SetCompareAtAmountCents(v interface{}) {
	o.CompareAtAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPricesPriceIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPricesPriceIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPricesPriceIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPricesPriceIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPricesPriceIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPricesPriceIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPricesPriceIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPricesPriceIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPricesPriceIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPricesPriceIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPricesPriceIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPricesPriceIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHPricesPriceIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPricesPriceIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
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
	return toSerialize, nil
}

type NullablePATCHPricesPriceIdRequestDataAttributes struct {
	value *PATCHPricesPriceIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHPricesPriceIdRequestDataAttributes) Get() *PATCHPricesPriceIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHPricesPriceIdRequestDataAttributes) Set(val *PATCHPricesPriceIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPricesPriceIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPricesPriceIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPricesPriceIdRequestDataAttributes(val *PATCHPricesPriceIdRequestDataAttributes) *NullablePATCHPricesPriceIdRequestDataAttributes {
	return &NullablePATCHPricesPriceIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPricesPriceIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPricesPriceIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
