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

// checks if the PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes{}

// PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes struct for PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes
type PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes struct {
	// The price tier's name
	Name interface{} `json:"name,omitempty"`
	// The tier upper limit, expressed as the line item frequency in days (or frequency label, ie 'monthly'). When 'null' it means infinity (useful to have an always matching tier).
	UpTo interface{} `json:"up_to,omitempty"`
	// The price of this price tier, in cents.
	PriceAmountCents interface{} `json:"price_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes instantiates a new PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes() *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes {
	this := PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes{}
	return &this
}

// NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributesWithDefaults instantiates a new PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributesWithDefaults() *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes {
	this := PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetUpTo returns the UpTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetUpTo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpTo
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetUpToOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpTo) {
		return nil, false
	}
	return &o.UpTo, true
}

// HasUpTo returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) HasUpTo() bool {
	if o != nil && IsNil(o.UpTo) {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given interface{} and assigns it to the UpTo field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) SetUpTo(v interface{}) {
	o.UpTo = v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) HasPriceAmountCents() bool {
	if o != nil && IsNil(o.PriceAmountCents) {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given interface{} and assigns it to the PriceAmountCents field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpTo != nil {
		toSerialize["up_to"] = o.UpTo
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
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

type NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes struct {
	value *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) Get() *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) Set(val *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes(val *PATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes {
	return &NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceFrequencyTiersPriceFrequencyTierIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
