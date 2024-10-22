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

// checks if the PATCHLineItemsLineItemId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHLineItemsLineItemId200ResponseDataAttributes{}

// PATCHLineItemsLineItemId200ResponseDataAttributes struct for PATCHLineItemsLineItemId200ResponseDataAttributes
type PATCHLineItemsLineItemId200ResponseDataAttributes struct {
	// The code of the associated SKU.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode interface{} `json:"bundle_code,omitempty"`
	// The line item quantity.
	Quantity interface{} `json:"quantity,omitempty"`
	// When creating or updating a new line item, set this attribute to '1' if you want to inject the unit_amount_cents price from an external source. Any successive price computation will be done externally, until the attribute is reset to '0'.
	ExternalPrice interface{} `json:"_external_price,omitempty"`
	// Send this attribute if you want to reserve the stock for the line item's SKUs quantity. Stock reservations expiration depends on the inventory model's cutoff. When used on update the existing active stock reservations are renewed. Cannot be passed by sales channels.
	ReserveStock interface{} `json:"_reserve_stock,omitempty"`
	// The unit amount of the line item, in cents. Can be specified only via an integration application, or when the item is missing, otherwise is automatically computed by using one of the available methods.
	UnitAmountCents interface{} `json:"unit_amount_cents,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents interface{} `json:"compare_at_amount_cents,omitempty"`
	// The options amount of the line item, in cents. Cannot be passed by sales channels.
	OptionsAmountCents interface{} `json:"options_amount_cents,omitempty"`
	// The name of the line item. When blank, it gets populated with the name of the associated item (if present).
	Name interface{} `json:"name,omitempty"`
	// The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only).
	ImageUrl interface{} `json:"image_url,omitempty"`
	// The frequency which generates a subscription. Must be supported by existing associated subscription_model.
	Frequency interface{} `json:"frequency,omitempty"`
	// Send this attribute if you want to reset the circuit breaker associated to this resource to 'closed' state and zero failures count. Cannot be passed by sales channels.
	ResetCircuit interface{} `json:"_reset_circuit,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHLineItemsLineItemId200ResponseDataAttributes instantiates a new PATCHLineItemsLineItemId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemsLineItemId200ResponseDataAttributes() *PATCHLineItemsLineItemId200ResponseDataAttributes {
	this := PATCHLineItemsLineItemId200ResponseDataAttributes{}
	return &this
}

// NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults instantiates a new PATCHLineItemsLineItemId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemsLineItemId200ResponseDataAttributesWithDefaults() *PATCHLineItemsLineItemId200ResponseDataAttributes {
	this := PATCHLineItemsLineItemId200ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetBundleCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BundleCode) {
		return nil, false
	}
	return &o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasBundleCode() bool {
	if o != nil && IsNil(o.BundleCode) {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given interface{} and assigns it to the BundleCode field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetBundleCode(v interface{}) {
	o.BundleCode = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetExternalPrice returns the ExternalPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetExternalPrice() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExternalPrice
}

// GetExternalPriceOk returns a tuple with the ExternalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetExternalPriceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExternalPrice) {
		return nil, false
	}
	return &o.ExternalPrice, true
}

// HasExternalPrice returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasExternalPrice() bool {
	if o != nil && IsNil(o.ExternalPrice) {
		return true
	}

	return false
}

// SetExternalPrice gets a reference to the given interface{} and assigns it to the ExternalPrice field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetExternalPrice(v interface{}) {
	o.ExternalPrice = v
}

// GetReserveStock returns the ReserveStock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReserveStock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReserveStock
}

// GetReserveStockOk returns a tuple with the ReserveStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReserveStockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReserveStock) {
		return nil, false
	}
	return &o.ReserveStock, true
}

// HasReserveStock returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReserveStock() bool {
	if o != nil && IsNil(o.ReserveStock) {
		return true
	}

	return false
}

// SetReserveStock gets a reference to the given interface{} and assigns it to the ReserveStock field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReserveStock(v interface{}) {
	o.ReserveStock = v
}

// GetUnitAmountCents returns the UnitAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitAmountCents
}

// GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitAmountCents) {
		return nil, false
	}
	return &o.UnitAmountCents, true
}

// HasUnitAmountCents returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasUnitAmountCents() bool {
	if o != nil && IsNil(o.UnitAmountCents) {
		return true
	}

	return false
}

// SetUnitAmountCents gets a reference to the given interface{} and assigns it to the UnitAmountCents field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountCents(v interface{}) {
	o.UnitAmountCents = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CompareAtAmountCents) {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasCompareAtAmountCents() bool {
	if o != nil && IsNil(o.CompareAtAmountCents) {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given interface{} and assigns it to the CompareAtAmountCents field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{}) {
	o.CompareAtAmountCents = v
}

// GetOptionsAmountCents returns the OptionsAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OptionsAmountCents
}

// GetOptionsAmountCentsOk returns a tuple with the OptionsAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OptionsAmountCents) {
		return nil, false
	}
	return &o.OptionsAmountCents, true
}

// HasOptionsAmountCents returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasOptionsAmountCents() bool {
	if o != nil && IsNil(o.OptionsAmountCents) {
		return true
	}

	return false
}

// SetOptionsAmountCents gets a reference to the given interface{} and assigns it to the OptionsAmountCents field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountCents(v interface{}) {
	o.OptionsAmountCents = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetFrequency() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return &o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasFrequency() bool {
	if o != nil && IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given interface{} and assigns it to the Frequency field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetFrequency(v interface{}) {
	o.Frequency = v
}

// GetResetCircuit returns the ResetCircuit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetResetCircuit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetCircuit
}

// GetResetCircuitOk returns a tuple with the ResetCircuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetResetCircuitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetCircuit) {
		return nil, false
	}
	return &o.ResetCircuit, true
}

// HasResetCircuit returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasResetCircuit() bool {
	if o != nil && IsNil(o.ResetCircuit) {
		return true
	}

	return false
}

// SetResetCircuit gets a reference to the given interface{} and assigns it to the ResetCircuit field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetResetCircuit(v interface{}) {
	o.ResetCircuit = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHLineItemsLineItemId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHLineItemsLineItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHLineItemsLineItemId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.BundleCode != nil {
		toSerialize["bundle_code"] = o.BundleCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ExternalPrice != nil {
		toSerialize["_external_price"] = o.ExternalPrice
	}
	if o.ReserveStock != nil {
		toSerialize["_reserve_stock"] = o.ReserveStock
	}
	if o.UnitAmountCents != nil {
		toSerialize["unit_amount_cents"] = o.UnitAmountCents
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
	}
	if o.OptionsAmountCents != nil {
		toSerialize["options_amount_cents"] = o.OptionsAmountCents
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.ResetCircuit != nil {
		toSerialize["_reset_circuit"] = o.ResetCircuit
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

type NullablePATCHLineItemsLineItemId200ResponseDataAttributes struct {
	value *PATCHLineItemsLineItemId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHLineItemsLineItemId200ResponseDataAttributes) Get() *PATCHLineItemsLineItemId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHLineItemsLineItemId200ResponseDataAttributes) Set(val *PATCHLineItemsLineItemId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemsLineItemId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemsLineItemId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemsLineItemId200ResponseDataAttributes(val *PATCHLineItemsLineItemId200ResponseDataAttributes) *NullablePATCHLineItemsLineItemId200ResponseDataAttributes {
	return &NullablePATCHLineItemsLineItemId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHLineItemsLineItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemsLineItemId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
