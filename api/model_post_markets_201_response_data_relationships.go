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

// checks if the POSTMarkets201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarkets201ResponseDataRelationships{}

// POSTMarkets201ResponseDataRelationships struct for POSTMarkets201ResponseDataRelationships
type POSTMarkets201ResponseDataRelationships struct {
	Merchant            *POSTMarkets201ResponseDataRelationshipsMerchant                         `json:"merchant,omitempty"`
	PriceList           *POSTMarkets201ResponseDataRelationshipsPriceList                        `json:"price_list,omitempty"`
	BasePriceList       *POSTMarkets201ResponseDataRelationshipsBasePriceList                    `json:"base_price_list,omitempty"`
	InventoryModel      *POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel  `json:"inventory_model,omitempty"`
	SubscriptionModel   *POSTMarkets201ResponseDataRelationshipsSubscriptionModel                `json:"subscription_model,omitempty"`
	TaxCalculator       *POSTMarkets201ResponseDataRelationshipsTaxCalculator                    `json:"tax_calculator,omitempty"`
	CustomerGroup       *POSTCustomers201ResponseDataRelationshipsCustomerGroup                  `json:"customer_group,omitempty"`
	Geocoder            *POSTAddresses201ResponseDataRelationshipsGeocoder                       `json:"geocoder,omitempty"`
	PriceListSchedulers *POSTMarkets201ResponseDataRelationshipsPriceListSchedulers              `json:"price_list_schedulers,omitempty"`
	Attachments         *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions            *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTMarkets201ResponseDataRelationships instantiates a new POSTMarkets201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarkets201ResponseDataRelationships() *POSTMarkets201ResponseDataRelationships {
	this := POSTMarkets201ResponseDataRelationships{}
	return &this
}

// NewPOSTMarkets201ResponseDataRelationshipsWithDefaults instantiates a new POSTMarkets201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarkets201ResponseDataRelationshipsWithDefaults() *POSTMarkets201ResponseDataRelationships {
	this := POSTMarkets201ResponseDataRelationships{}
	return &this
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetMerchant() POSTMarkets201ResponseDataRelationshipsMerchant {
	if o == nil || IsNil(o.Merchant) {
		var ret POSTMarkets201ResponseDataRelationshipsMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetMerchantOk() (*POSTMarkets201ResponseDataRelationshipsMerchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given POSTMarkets201ResponseDataRelationshipsMerchant and assigns it to the Merchant field.
func (o *POSTMarkets201ResponseDataRelationships) SetMerchant(v POSTMarkets201ResponseDataRelationshipsMerchant) {
	o.Merchant = &v
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetPriceList() POSTMarkets201ResponseDataRelationshipsPriceList {
	if o == nil || IsNil(o.PriceList) {
		var ret POSTMarkets201ResponseDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetPriceListOk() (*POSTMarkets201ResponseDataRelationshipsPriceList, bool) {
	if o == nil || IsNil(o.PriceList) {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasPriceList() bool {
	if o != nil && !IsNil(o.PriceList) {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given POSTMarkets201ResponseDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *POSTMarkets201ResponseDataRelationships) SetPriceList(v POSTMarkets201ResponseDataRelationshipsPriceList) {
	o.PriceList = &v
}

// GetBasePriceList returns the BasePriceList field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetBasePriceList() POSTMarkets201ResponseDataRelationshipsBasePriceList {
	if o == nil || IsNil(o.BasePriceList) {
		var ret POSTMarkets201ResponseDataRelationshipsBasePriceList
		return ret
	}
	return *o.BasePriceList
}

// GetBasePriceListOk returns a tuple with the BasePriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetBasePriceListOk() (*POSTMarkets201ResponseDataRelationshipsBasePriceList, bool) {
	if o == nil || IsNil(o.BasePriceList) {
		return nil, false
	}
	return o.BasePriceList, true
}

// HasBasePriceList returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasBasePriceList() bool {
	if o != nil && !IsNil(o.BasePriceList) {
		return true
	}

	return false
}

// SetBasePriceList gets a reference to the given POSTMarkets201ResponseDataRelationshipsBasePriceList and assigns it to the BasePriceList field.
func (o *POSTMarkets201ResponseDataRelationships) SetBasePriceList(v POSTMarkets201ResponseDataRelationshipsBasePriceList) {
	o.BasePriceList = &v
}

// GetInventoryModel returns the InventoryModel field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetInventoryModel() POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel {
	if o == nil || IsNil(o.InventoryModel) {
		var ret POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel
		return ret
	}
	return *o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetInventoryModelOk() (*POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel, bool) {
	if o == nil || IsNil(o.InventoryModel) {
		return nil, false
	}
	return o.InventoryModel, true
}

// HasInventoryModel returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasInventoryModel() bool {
	if o != nil && !IsNil(o.InventoryModel) {
		return true
	}

	return false
}

// SetInventoryModel gets a reference to the given POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel and assigns it to the InventoryModel field.
func (o *POSTMarkets201ResponseDataRelationships) SetInventoryModel(v POSTInventoryReturnLocations201ResponseDataRelationshipsInventoryModel) {
	o.InventoryModel = &v
}

// GetSubscriptionModel returns the SubscriptionModel field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetSubscriptionModel() POSTMarkets201ResponseDataRelationshipsSubscriptionModel {
	if o == nil || IsNil(o.SubscriptionModel) {
		var ret POSTMarkets201ResponseDataRelationshipsSubscriptionModel
		return ret
	}
	return *o.SubscriptionModel
}

// GetSubscriptionModelOk returns a tuple with the SubscriptionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetSubscriptionModelOk() (*POSTMarkets201ResponseDataRelationshipsSubscriptionModel, bool) {
	if o == nil || IsNil(o.SubscriptionModel) {
		return nil, false
	}
	return o.SubscriptionModel, true
}

// HasSubscriptionModel returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasSubscriptionModel() bool {
	if o != nil && !IsNil(o.SubscriptionModel) {
		return true
	}

	return false
}

// SetSubscriptionModel gets a reference to the given POSTMarkets201ResponseDataRelationshipsSubscriptionModel and assigns it to the SubscriptionModel field.
func (o *POSTMarkets201ResponseDataRelationships) SetSubscriptionModel(v POSTMarkets201ResponseDataRelationshipsSubscriptionModel) {
	o.SubscriptionModel = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetTaxCalculator() POSTMarkets201ResponseDataRelationshipsTaxCalculator {
	if o == nil || IsNil(o.TaxCalculator) {
		var ret POSTMarkets201ResponseDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetTaxCalculatorOk() (*POSTMarkets201ResponseDataRelationshipsTaxCalculator, bool) {
	if o == nil || IsNil(o.TaxCalculator) {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasTaxCalculator() bool {
	if o != nil && !IsNil(o.TaxCalculator) {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given POSTMarkets201ResponseDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *POSTMarkets201ResponseDataRelationships) SetTaxCalculator(v POSTMarkets201ResponseDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetCustomerGroup() POSTCustomers201ResponseDataRelationshipsCustomerGroup {
	if o == nil || IsNil(o.CustomerGroup) {
		var ret POSTCustomers201ResponseDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetCustomerGroupOk() (*POSTCustomers201ResponseDataRelationshipsCustomerGroup, bool) {
	if o == nil || IsNil(o.CustomerGroup) {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasCustomerGroup() bool {
	if o != nil && !IsNil(o.CustomerGroup) {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given POSTCustomers201ResponseDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *POSTMarkets201ResponseDataRelationships) SetCustomerGroup(v POSTCustomers201ResponseDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

// GetGeocoder returns the Geocoder field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetGeocoder() POSTAddresses201ResponseDataRelationshipsGeocoder {
	if o == nil || IsNil(o.Geocoder) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoder
		return ret
	}
	return *o.Geocoder
}

// GetGeocoderOk returns a tuple with the Geocoder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetGeocoderOk() (*POSTAddresses201ResponseDataRelationshipsGeocoder, bool) {
	if o == nil || IsNil(o.Geocoder) {
		return nil, false
	}
	return o.Geocoder, true
}

// HasGeocoder returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasGeocoder() bool {
	if o != nil && !IsNil(o.Geocoder) {
		return true
	}

	return false
}

// SetGeocoder gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoder and assigns it to the Geocoder field.
func (o *POSTMarkets201ResponseDataRelationships) SetGeocoder(v POSTAddresses201ResponseDataRelationshipsGeocoder) {
	o.Geocoder = &v
}

// GetPriceListSchedulers returns the PriceListSchedulers field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetPriceListSchedulers() POSTMarkets201ResponseDataRelationshipsPriceListSchedulers {
	if o == nil || IsNil(o.PriceListSchedulers) {
		var ret POSTMarkets201ResponseDataRelationshipsPriceListSchedulers
		return ret
	}
	return *o.PriceListSchedulers
}

// GetPriceListSchedulersOk returns a tuple with the PriceListSchedulers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetPriceListSchedulersOk() (*POSTMarkets201ResponseDataRelationshipsPriceListSchedulers, bool) {
	if o == nil || IsNil(o.PriceListSchedulers) {
		return nil, false
	}
	return o.PriceListSchedulers, true
}

// HasPriceListSchedulers returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasPriceListSchedulers() bool {
	if o != nil && !IsNil(o.PriceListSchedulers) {
		return true
	}

	return false
}

// SetPriceListSchedulers gets a reference to the given POSTMarkets201ResponseDataRelationshipsPriceListSchedulers and assigns it to the PriceListSchedulers field.
func (o *POSTMarkets201ResponseDataRelationships) SetPriceListSchedulers(v POSTMarkets201ResponseDataRelationshipsPriceListSchedulers) {
	o.PriceListSchedulers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTMarkets201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTMarkets201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarkets201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTMarkets201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTMarkets201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTMarkets201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarkets201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !IsNil(o.PriceList) {
		toSerialize["price_list"] = o.PriceList
	}
	if !IsNil(o.BasePriceList) {
		toSerialize["base_price_list"] = o.BasePriceList
	}
	if !IsNil(o.InventoryModel) {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	if !IsNil(o.SubscriptionModel) {
		toSerialize["subscription_model"] = o.SubscriptionModel
	}
	if !IsNil(o.TaxCalculator) {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if !IsNil(o.CustomerGroup) {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	if !IsNil(o.Geocoder) {
		toSerialize["geocoder"] = o.Geocoder
	}
	if !IsNil(o.PriceListSchedulers) {
		toSerialize["price_list_schedulers"] = o.PriceListSchedulers
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTMarkets201ResponseDataRelationships struct {
	value *POSTMarkets201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTMarkets201ResponseDataRelationships) Get() *POSTMarkets201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTMarkets201ResponseDataRelationships) Set(val *POSTMarkets201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarkets201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarkets201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarkets201ResponseDataRelationships(val *POSTMarkets201ResponseDataRelationships) *NullablePOSTMarkets201ResponseDataRelationships {
	return &NullablePOSTMarkets201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTMarkets201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarkets201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
