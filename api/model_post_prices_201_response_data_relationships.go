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

// checks if the POSTPrices201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPrices201ResponseDataRelationships{}

// POSTPrices201ResponseDataRelationships struct for POSTPrices201ResponseDataRelationships
type POSTPrices201ResponseDataRelationships struct {
	PriceList           *POSTMarkets201ResponseDataRelationshipsPriceList                        `json:"price_list,omitempty"`
	Sku                 *POSTInStockSubscriptions201ResponseDataRelationshipsSku                 `json:"sku,omitempty"`
	PriceTiers          *POSTPrices201ResponseDataRelationshipsPriceTiers                        `json:"price_tiers,omitempty"`
	PriceVolumeTiers    *POSTPrices201ResponseDataRelationshipsPriceVolumeTiers                  `json:"price_volume_tiers,omitempty"`
	PriceFrequencyTiers *POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers               `json:"price_frequency_tiers,omitempty"`
	Attachments         *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions            *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
	JwtCustomer         *POSTPrices201ResponseDataRelationshipsJwtCustomer                       `json:"jwt_customer,omitempty"`
	JwtMarkets          *POSTPrices201ResponseDataRelationshipsJwtMarkets                        `json:"jwt_markets,omitempty"`
	JwtStockLocations   *POSTPrices201ResponseDataRelationshipsJwtStockLocations                 `json:"jwt_stock_locations,omitempty"`
}

// NewPOSTPrices201ResponseDataRelationships instantiates a new POSTPrices201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPrices201ResponseDataRelationships() *POSTPrices201ResponseDataRelationships {
	this := POSTPrices201ResponseDataRelationships{}
	return &this
}

// NewPOSTPrices201ResponseDataRelationshipsWithDefaults instantiates a new POSTPrices201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPrices201ResponseDataRelationshipsWithDefaults() *POSTPrices201ResponseDataRelationships {
	this := POSTPrices201ResponseDataRelationships{}
	return &this
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetPriceList() POSTMarkets201ResponseDataRelationshipsPriceList {
	if o == nil || IsNil(o.PriceList) {
		var ret POSTMarkets201ResponseDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetPriceListOk() (*POSTMarkets201ResponseDataRelationshipsPriceList, bool) {
	if o == nil || IsNil(o.PriceList) {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasPriceList() bool {
	if o != nil && !IsNil(o.PriceList) {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given POSTMarkets201ResponseDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *POSTPrices201ResponseDataRelationships) SetPriceList(v POSTMarkets201ResponseDataRelationshipsPriceList) {
	o.PriceList = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTPrices201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetPriceTiers returns the PriceTiers field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetPriceTiers() POSTPrices201ResponseDataRelationshipsPriceTiers {
	if o == nil || IsNil(o.PriceTiers) {
		var ret POSTPrices201ResponseDataRelationshipsPriceTiers
		return ret
	}
	return *o.PriceTiers
}

// GetPriceTiersOk returns a tuple with the PriceTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetPriceTiersOk() (*POSTPrices201ResponseDataRelationshipsPriceTiers, bool) {
	if o == nil || IsNil(o.PriceTiers) {
		return nil, false
	}
	return o.PriceTiers, true
}

// HasPriceTiers returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasPriceTiers() bool {
	if o != nil && !IsNil(o.PriceTiers) {
		return true
	}

	return false
}

// SetPriceTiers gets a reference to the given POSTPrices201ResponseDataRelationshipsPriceTiers and assigns it to the PriceTiers field.
func (o *POSTPrices201ResponseDataRelationships) SetPriceTiers(v POSTPrices201ResponseDataRelationshipsPriceTiers) {
	o.PriceTiers = &v
}

// GetPriceVolumeTiers returns the PriceVolumeTiers field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetPriceVolumeTiers() POSTPrices201ResponseDataRelationshipsPriceVolumeTiers {
	if o == nil || IsNil(o.PriceVolumeTiers) {
		var ret POSTPrices201ResponseDataRelationshipsPriceVolumeTiers
		return ret
	}
	return *o.PriceVolumeTiers
}

// GetPriceVolumeTiersOk returns a tuple with the PriceVolumeTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetPriceVolumeTiersOk() (*POSTPrices201ResponseDataRelationshipsPriceVolumeTiers, bool) {
	if o == nil || IsNil(o.PriceVolumeTiers) {
		return nil, false
	}
	return o.PriceVolumeTiers, true
}

// HasPriceVolumeTiers returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasPriceVolumeTiers() bool {
	if o != nil && !IsNil(o.PriceVolumeTiers) {
		return true
	}

	return false
}

// SetPriceVolumeTiers gets a reference to the given POSTPrices201ResponseDataRelationshipsPriceVolumeTiers and assigns it to the PriceVolumeTiers field.
func (o *POSTPrices201ResponseDataRelationships) SetPriceVolumeTiers(v POSTPrices201ResponseDataRelationshipsPriceVolumeTiers) {
	o.PriceVolumeTiers = &v
}

// GetPriceFrequencyTiers returns the PriceFrequencyTiers field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetPriceFrequencyTiers() POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers {
	if o == nil || IsNil(o.PriceFrequencyTiers) {
		var ret POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers
		return ret
	}
	return *o.PriceFrequencyTiers
}

// GetPriceFrequencyTiersOk returns a tuple with the PriceFrequencyTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetPriceFrequencyTiersOk() (*POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers, bool) {
	if o == nil || IsNil(o.PriceFrequencyTiers) {
		return nil, false
	}
	return o.PriceFrequencyTiers, true
}

// HasPriceFrequencyTiers returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasPriceFrequencyTiers() bool {
	if o != nil && !IsNil(o.PriceFrequencyTiers) {
		return true
	}

	return false
}

// SetPriceFrequencyTiers gets a reference to the given POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers and assigns it to the PriceFrequencyTiers field.
func (o *POSTPrices201ResponseDataRelationships) SetPriceFrequencyTiers(v POSTPrices201ResponseDataRelationshipsPriceFrequencyTiers) {
	o.PriceFrequencyTiers = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTPrices201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTPrices201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetJwtCustomer returns the JwtCustomer field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetJwtCustomer() POSTPrices201ResponseDataRelationshipsJwtCustomer {
	if o == nil || IsNil(o.JwtCustomer) {
		var ret POSTPrices201ResponseDataRelationshipsJwtCustomer
		return ret
	}
	return *o.JwtCustomer
}

// GetJwtCustomerOk returns a tuple with the JwtCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetJwtCustomerOk() (*POSTPrices201ResponseDataRelationshipsJwtCustomer, bool) {
	if o == nil || IsNil(o.JwtCustomer) {
		return nil, false
	}
	return o.JwtCustomer, true
}

// HasJwtCustomer returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasJwtCustomer() bool {
	if o != nil && !IsNil(o.JwtCustomer) {
		return true
	}

	return false
}

// SetJwtCustomer gets a reference to the given POSTPrices201ResponseDataRelationshipsJwtCustomer and assigns it to the JwtCustomer field.
func (o *POSTPrices201ResponseDataRelationships) SetJwtCustomer(v POSTPrices201ResponseDataRelationshipsJwtCustomer) {
	o.JwtCustomer = &v
}

// GetJwtMarkets returns the JwtMarkets field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetJwtMarkets() POSTPrices201ResponseDataRelationshipsJwtMarkets {
	if o == nil || IsNil(o.JwtMarkets) {
		var ret POSTPrices201ResponseDataRelationshipsJwtMarkets
		return ret
	}
	return *o.JwtMarkets
}

// GetJwtMarketsOk returns a tuple with the JwtMarkets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetJwtMarketsOk() (*POSTPrices201ResponseDataRelationshipsJwtMarkets, bool) {
	if o == nil || IsNil(o.JwtMarkets) {
		return nil, false
	}
	return o.JwtMarkets, true
}

// HasJwtMarkets returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasJwtMarkets() bool {
	if o != nil && !IsNil(o.JwtMarkets) {
		return true
	}

	return false
}

// SetJwtMarkets gets a reference to the given POSTPrices201ResponseDataRelationshipsJwtMarkets and assigns it to the JwtMarkets field.
func (o *POSTPrices201ResponseDataRelationships) SetJwtMarkets(v POSTPrices201ResponseDataRelationshipsJwtMarkets) {
	o.JwtMarkets = &v
}

// GetJwtStockLocations returns the JwtStockLocations field value if set, zero value otherwise.
func (o *POSTPrices201ResponseDataRelationships) GetJwtStockLocations() POSTPrices201ResponseDataRelationshipsJwtStockLocations {
	if o == nil || IsNil(o.JwtStockLocations) {
		var ret POSTPrices201ResponseDataRelationshipsJwtStockLocations
		return ret
	}
	return *o.JwtStockLocations
}

// GetJwtStockLocationsOk returns a tuple with the JwtStockLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPrices201ResponseDataRelationships) GetJwtStockLocationsOk() (*POSTPrices201ResponseDataRelationshipsJwtStockLocations, bool) {
	if o == nil || IsNil(o.JwtStockLocations) {
		return nil, false
	}
	return o.JwtStockLocations, true
}

// HasJwtStockLocations returns a boolean if a field has been set.
func (o *POSTPrices201ResponseDataRelationships) HasJwtStockLocations() bool {
	if o != nil && !IsNil(o.JwtStockLocations) {
		return true
	}

	return false
}

// SetJwtStockLocations gets a reference to the given POSTPrices201ResponseDataRelationshipsJwtStockLocations and assigns it to the JwtStockLocations field.
func (o *POSTPrices201ResponseDataRelationships) SetJwtStockLocations(v POSTPrices201ResponseDataRelationshipsJwtStockLocations) {
	o.JwtStockLocations = &v
}

func (o POSTPrices201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPrices201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceList) {
		toSerialize["price_list"] = o.PriceList
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.PriceTiers) {
		toSerialize["price_tiers"] = o.PriceTiers
	}
	if !IsNil(o.PriceVolumeTiers) {
		toSerialize["price_volume_tiers"] = o.PriceVolumeTiers
	}
	if !IsNil(o.PriceFrequencyTiers) {
		toSerialize["price_frequency_tiers"] = o.PriceFrequencyTiers
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.JwtCustomer) {
		toSerialize["jwt_customer"] = o.JwtCustomer
	}
	if !IsNil(o.JwtMarkets) {
		toSerialize["jwt_markets"] = o.JwtMarkets
	}
	if !IsNil(o.JwtStockLocations) {
		toSerialize["jwt_stock_locations"] = o.JwtStockLocations
	}
	return toSerialize, nil
}

type NullablePOSTPrices201ResponseDataRelationships struct {
	value *POSTPrices201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTPrices201ResponseDataRelationships) Get() *POSTPrices201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTPrices201ResponseDataRelationships) Set(val *POSTPrices201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPrices201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPrices201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPrices201ResponseDataRelationships(val *POSTPrices201ResponseDataRelationships) *NullablePOSTPrices201ResponseDataRelationships {
	return &NullablePOSTPrices201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTPrices201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPrices201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
