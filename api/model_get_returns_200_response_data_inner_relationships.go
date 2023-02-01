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

// GETReturns200ResponseDataInnerRelationships struct for GETReturns200ResponseDataInnerRelationships
type GETReturns200ResponseDataInnerRelationships struct {
	Order              *GETAdyenPayments200ResponseDataInnerRelationshipsOrder             `json:"order,omitempty"`
	Customer           *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer       `json:"customer,omitempty"`
	StockLocation      *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation `json:"stock_location,omitempty"`
	OriginAddress      *GETReturns200ResponseDataInnerRelationshipsOriginAddress           `json:"origin_address,omitempty"`
	DestinationAddress *GETReturns200ResponseDataInnerRelationshipsDestinationAddress      `json:"destination_address,omitempty"`
	ReturnLineItems    *GETReturns200ResponseDataInnerRelationshipsReturnLineItems         `json:"return_line_items,omitempty"`
	Attachments        *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments     `json:"attachments,omitempty"`
	Events             *GETAuthorizations200ResponseDataInnerRelationshipsEvents           `json:"events,omitempty"`
}

// NewGETReturns200ResponseDataInnerRelationships instantiates a new GETReturns200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturns200ResponseDataInnerRelationships() *GETReturns200ResponseDataInnerRelationships {
	this := GETReturns200ResponseDataInnerRelationships{}
	return &this
}

// NewGETReturns200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETReturns200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturns200ResponseDataInnerRelationshipsWithDefaults() *GETReturns200ResponseDataInnerRelationships {
	this := GETReturns200ResponseDataInnerRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsOrder and assigns it to the Order field.
func (o *GETReturns200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	o.Order = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETReturns200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *GETReturns200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetOriginAddress returns the OriginAddress field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetOriginAddress() GETReturns200ResponseDataInnerRelationshipsOriginAddress {
	if o == nil || o.OriginAddress == nil {
		var ret GETReturns200ResponseDataInnerRelationshipsOriginAddress
		return ret
	}
	return *o.OriginAddress
}

// GetOriginAddressOk returns a tuple with the OriginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetOriginAddressOk() (*GETReturns200ResponseDataInnerRelationshipsOriginAddress, bool) {
	if o == nil || o.OriginAddress == nil {
		return nil, false
	}
	return o.OriginAddress, true
}

// HasOriginAddress returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasOriginAddress() bool {
	if o != nil && o.OriginAddress != nil {
		return true
	}

	return false
}

// SetOriginAddress gets a reference to the given GETReturns200ResponseDataInnerRelationshipsOriginAddress and assigns it to the OriginAddress field.
func (o *GETReturns200ResponseDataInnerRelationships) SetOriginAddress(v GETReturns200ResponseDataInnerRelationshipsOriginAddress) {
	o.OriginAddress = &v
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetDestinationAddress() GETReturns200ResponseDataInnerRelationshipsDestinationAddress {
	if o == nil || o.DestinationAddress == nil {
		var ret GETReturns200ResponseDataInnerRelationshipsDestinationAddress
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetDestinationAddressOk() (*GETReturns200ResponseDataInnerRelationshipsDestinationAddress, bool) {
	if o == nil || o.DestinationAddress == nil {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasDestinationAddress() bool {
	if o != nil && o.DestinationAddress != nil {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given GETReturns200ResponseDataInnerRelationshipsDestinationAddress and assigns it to the DestinationAddress field.
func (o *GETReturns200ResponseDataInnerRelationships) SetDestinationAddress(v GETReturns200ResponseDataInnerRelationshipsDestinationAddress) {
	o.DestinationAddress = &v
}

// GetReturnLineItems returns the ReturnLineItems field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetReturnLineItems() GETReturns200ResponseDataInnerRelationshipsReturnLineItems {
	if o == nil || o.ReturnLineItems == nil {
		var ret GETReturns200ResponseDataInnerRelationshipsReturnLineItems
		return ret
	}
	return *o.ReturnLineItems
}

// GetReturnLineItemsOk returns a tuple with the ReturnLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetReturnLineItemsOk() (*GETReturns200ResponseDataInnerRelationshipsReturnLineItems, bool) {
	if o == nil || o.ReturnLineItems == nil {
		return nil, false
	}
	return o.ReturnLineItems, true
}

// HasReturnLineItems returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasReturnLineItems() bool {
	if o != nil && o.ReturnLineItems != nil {
		return true
	}

	return false
}

// SetReturnLineItems gets a reference to the given GETReturns200ResponseDataInnerRelationshipsReturnLineItems and assigns it to the ReturnLineItems field.
func (o *GETReturns200ResponseDataInnerRelationships) SetReturnLineItems(v GETReturns200ResponseDataInnerRelationshipsReturnLineItems) {
	o.ReturnLineItems = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETReturns200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETReturns200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETAuthorizations200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturns200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETReturns200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETAuthorizations200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETReturns200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETReturns200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.OriginAddress != nil {
		toSerialize["origin_address"] = o.OriginAddress
	}
	if o.DestinationAddress != nil {
		toSerialize["destination_address"] = o.DestinationAddress
	}
	if o.ReturnLineItems != nil {
		toSerialize["return_line_items"] = o.ReturnLineItems
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturns200ResponseDataInnerRelationships struct {
	value *GETReturns200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETReturns200ResponseDataInnerRelationships) Get() *GETReturns200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETReturns200ResponseDataInnerRelationships) Set(val *GETReturns200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturns200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturns200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturns200ResponseDataInnerRelationships(val *GETReturns200ResponseDataInnerRelationships) *NullableGETReturns200ResponseDataInnerRelationships {
	return &NullableGETReturns200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETReturns200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturns200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
