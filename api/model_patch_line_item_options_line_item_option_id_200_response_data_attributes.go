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

// checks if the PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{}

// PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes struct for PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes
type PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes struct {
	// The name of the line item option. When blank, it gets populated with the name of the associated SKU option.
	Name interface{} `json:"name,omitempty"`
	// The line item option's quantity.
	Quantity interface{} `json:"quantity,omitempty"`
	// The unit amount of the line item option, in cents. When you add a line item option to an order, this is automatically populated from associated SKU option's price. Cannot be passed by sales channels.
	UnitAmountCents interface{} `json:"unit_amount_cents,omitempty"`
	// Set of key-value pairs that represent the selected options.
	Options interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes instantiates a new PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes() *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	this := PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributesWithDefaults instantiates a new PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemOptionsLineItemOptionId200ResponseDataAttributesWithDefaults() *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	this := PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetUnitAmountCents returns the UnitAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetUnitAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitAmountCents
}

// GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitAmountCents) {
		return nil, false
	}
	return &o.UnitAmountCents, true
}

// HasUnitAmountCents returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasUnitAmountCents() bool {
	if o != nil && IsNil(o.UnitAmountCents) {
		return true
	}

	return false
}

// SetUnitAmountCents gets a reference to the given interface{} and assigns it to the UnitAmountCents field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetUnitAmountCents(v interface{}) {
	o.UnitAmountCents = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.UnitAmountCents != nil {
		toSerialize["unit_amount_cents"] = o.UnitAmountCents
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes struct {
	value *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) Get() *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) Set(val *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes(val *PATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes {
	return &NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
