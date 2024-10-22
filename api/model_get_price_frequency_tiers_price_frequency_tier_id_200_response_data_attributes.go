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

// checks if the GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes{}

// GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes struct for GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes
type GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes struct {
	// The price tier's name.
	Name interface{} `json:"name,omitempty"`
	// The price tier's type.
	Type interface{} `json:"type,omitempty"`
	// The tier upper limit, expressed as the line item frequency in days (or frequency label, ie 'monthly'). When 'null' it means infinity (useful to have an always matching tier).
	UpTo interface{} `json:"up_to,omitempty"`
	// The price of this price tier, in cents.
	PriceAmountCents interface{} `json:"price_amount_cents,omitempty"`
	// The price of this price tier, float.
	PriceAmountFloat interface{} `json:"price_amount_float,omitempty"`
	// The price of this price tier, formatted.
	FormattedPriceAmount interface{} `json:"formatted_price_amount,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes instantiates a new GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes() *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes {
	this := GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes{}
	return &this
}

// NewGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributesWithDefaults instantiates a new GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributesWithDefaults() *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes {
	this := GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetUpTo returns the UpTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetUpTo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpTo
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetUpToOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpTo) {
		return nil, false
	}
	return &o.UpTo, true
}

// HasUpTo returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasUpTo() bool {
	if o != nil && IsNil(o.UpTo) {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given interface{} and assigns it to the UpTo field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetUpTo(v interface{}) {
	o.UpTo = v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetPriceAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountCents) {
		return nil, false
	}
	return &o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasPriceAmountCents() bool {
	if o != nil && IsNil(o.PriceAmountCents) {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given interface{} and assigns it to the PriceAmountCents field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetPriceAmountCents(v interface{}) {
	o.PriceAmountCents = v
}

// GetPriceAmountFloat returns the PriceAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetPriceAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PriceAmountFloat
}

// GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetPriceAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PriceAmountFloat) {
		return nil, false
	}
	return &o.PriceAmountFloat, true
}

// HasPriceAmountFloat returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasPriceAmountFloat() bool {
	if o != nil && IsNil(o.PriceAmountFloat) {
		return true
	}

	return false
}

// SetPriceAmountFloat gets a reference to the given interface{} and assigns it to the PriceAmountFloat field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetPriceAmountFloat(v interface{}) {
	o.PriceAmountFloat = v
}

// GetFormattedPriceAmount returns the FormattedPriceAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetFormattedPriceAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedPriceAmount
}

// GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetFormattedPriceAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedPriceAmount) {
		return nil, false
	}
	return &o.FormattedPriceAmount, true
}

// HasFormattedPriceAmount returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasFormattedPriceAmount() bool {
	if o != nil && IsNil(o.FormattedPriceAmount) {
		return true
	}

	return false
}

// SetFormattedPriceAmount gets a reference to the given interface{} and assigns it to the FormattedPriceAmount field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetFormattedPriceAmount(v interface{}) {
	o.FormattedPriceAmount = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpTo != nil {
		toSerialize["up_to"] = o.UpTo
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.PriceAmountFloat != nil {
		toSerialize["price_amount_float"] = o.PriceAmountFloat
	}
	if o.FormattedPriceAmount != nil {
		toSerialize["formatted_price_amount"] = o.FormattedPriceAmount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes struct {
	value *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) Get() *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) Set(val *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes(val *GETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) *NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes {
	return &NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
