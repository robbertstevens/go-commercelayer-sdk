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

// POSTBundles201ResponseDataAttributes struct for POSTBundles201ResponseDataAttributes
type POSTBundles201ResponseDataAttributes struct {
	// The bundle code, that uniquely identifies the bundle within the market.
	Code string `json:"code"`
	// The internal name of the bundle.
	Name string `json:"name"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// An internal description of the bundle.
	Description *string `json:"description,omitempty"`
	// The URL of an image that represents the bundle.
	ImageUrl *string `json:"image_url,omitempty"`
	// The bundle price amount for the associated market, in cents.
	PriceAmountCents int32 `json:"price_amount_cents"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents int32 `json:"compare_at_amount_cents"`
	// Send this attribute if you want to compute the price_amount_cents as the sum of the prices of the bundle SKUs for the market.
	ComputePriceAmount *bool `json:"_compute_price_amount,omitempty"`
	// Send this attribute if you want to compute the compare_at_amount_cents as the sum of the prices of the bundle SKUs for the market.
	ComputeCompareAtAmount *bool `json:"_compute_compare_at_amount,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTBundles201ResponseDataAttributes instantiates a new POSTBundles201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBundles201ResponseDataAttributes(code string, name string, priceAmountCents int32, compareAtAmountCents int32) *POSTBundles201ResponseDataAttributes {
	this := POSTBundles201ResponseDataAttributes{}
	this.Code = code
	this.Name = name
	this.PriceAmountCents = priceAmountCents
	this.CompareAtAmountCents = compareAtAmountCents
	return &this
}

// NewPOSTBundles201ResponseDataAttributesWithDefaults instantiates a new POSTBundles201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBundles201ResponseDataAttributesWithDefaults() *POSTBundles201ResponseDataAttributes {
	this := POSTBundles201ResponseDataAttributes{}
	return &this
}

// GetCode returns the Code field value
func (o *POSTBundles201ResponseDataAttributes) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *POSTBundles201ResponseDataAttributes) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *POSTBundles201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTBundles201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *POSTBundles201ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *POSTBundles201ResponseDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *POSTBundles201ResponseDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value
func (o *POSTBundles201ResponseDataAttributes) GetPriceAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// SetPriceAmountCents sets field value
func (o *POSTBundles201ResponseDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value
func (o *POSTBundles201ResponseDataAttributes) GetCompareAtAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetCompareAtAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// SetCompareAtAmountCents sets field value
func (o *POSTBundles201ResponseDataAttributes) SetCompareAtAmountCents(v int32) {
	o.CompareAtAmountCents = v
}

// GetComputePriceAmount returns the ComputePriceAmount field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetComputePriceAmount() bool {
	if o == nil || o.ComputePriceAmount == nil {
		var ret bool
		return ret
	}
	return *o.ComputePriceAmount
}

// GetComputePriceAmountOk returns a tuple with the ComputePriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetComputePriceAmountOk() (*bool, bool) {
	if o == nil || o.ComputePriceAmount == nil {
		return nil, false
	}
	return o.ComputePriceAmount, true
}

// HasComputePriceAmount returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasComputePriceAmount() bool {
	if o != nil && o.ComputePriceAmount != nil {
		return true
	}

	return false
}

// SetComputePriceAmount gets a reference to the given bool and assigns it to the ComputePriceAmount field.
func (o *POSTBundles201ResponseDataAttributes) SetComputePriceAmount(v bool) {
	o.ComputePriceAmount = &v
}

// GetComputeCompareAtAmount returns the ComputeCompareAtAmount field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetComputeCompareAtAmount() bool {
	if o == nil || o.ComputeCompareAtAmount == nil {
		var ret bool
		return ret
	}
	return *o.ComputeCompareAtAmount
}

// GetComputeCompareAtAmountOk returns a tuple with the ComputeCompareAtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetComputeCompareAtAmountOk() (*bool, bool) {
	if o == nil || o.ComputeCompareAtAmount == nil {
		return nil, false
	}
	return o.ComputeCompareAtAmount, true
}

// HasComputeCompareAtAmount returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasComputeCompareAtAmount() bool {
	if o != nil && o.ComputeCompareAtAmount != nil {
		return true
	}

	return false
}

// SetComputeCompareAtAmount gets a reference to the given bool and assigns it to the ComputeCompareAtAmount field.
func (o *POSTBundles201ResponseDataAttributes) SetComputeCompareAtAmount(v bool) {
	o.ComputeCompareAtAmount = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTBundles201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTBundles201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTBundles201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBundles201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTBundles201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTBundles201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTBundles201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if true {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if true {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
	}
	if o.ComputePriceAmount != nil {
		toSerialize["_compute_price_amount"] = o.ComputePriceAmount
	}
	if o.ComputeCompareAtAmount != nil {
		toSerialize["_compute_compare_at_amount"] = o.ComputeCompareAtAmount
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

type NullablePOSTBundles201ResponseDataAttributes struct {
	value *POSTBundles201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTBundles201ResponseDataAttributes) Get() *POSTBundles201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTBundles201ResponseDataAttributes) Set(val *POSTBundles201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBundles201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBundles201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBundles201ResponseDataAttributes(val *POSTBundles201ResponseDataAttributes) *NullablePOSTBundles201ResponseDataAttributes {
	return &NullablePOSTBundles201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTBundles201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBundles201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


