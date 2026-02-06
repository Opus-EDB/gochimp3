package gochimp3

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05-07:00"
)

func (address *Address) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Address
		CountryCode string `json:"country_code"`
	}{
		Address:     *address,
		CountryCode: strings.ToUpper(address.CountryCode),
	}
	return json.Marshal(tmp)
}

func (loc *MemberLocation) MarshalJSON() ([]byte, error) {
	tmp := struct {
		MemberLocation
		CountryCode string `json:"country_code"`
	}{
		MemberLocation: *loc,
		CountryCode:    strings.ToUpper(loc.CountryCode),
	}
	return json.Marshal(tmp)
}

func (store *Store) UnmarshalJSON(data []byte) error {
	type Alias Store
	tmp := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias: (*Alias)(store),
	}

	if err := json.Unmarshal(data, tmp); err != nil {
		return err
	}

	if tmp.CreatedAt != "" {
		if t, err := time.Parse(timeFormat, tmp.CreatedAt); err == nil {
			store.CreatedAt = t
		}
	}
	if tmp.UpdatedAt != "" {
		if t, err := time.Parse(timeFormat, tmp.UpdatedAt); err == nil {
			store.UpdatedAt = t
		}
	}

	return nil
}

func (store *Store) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Store
		CurrencyCode string `json:"currency_code"`
	}{
		Store:        *store,
		CurrencyCode: strings.ToUpper(store.CurrencyCode),
	}
	return json.Marshal(tmp)
}

func (cart *Cart) UnmarshalJSON(data []byte) error {
	type Alias Cart
	tmp := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias: (*Alias)(cart),
	}

	if err := json.Unmarshal(data, tmp); err != nil {
		return err
	}

	if tmp.CreatedAt != "" {
		if t, err := time.Parse(timeFormat, tmp.CreatedAt); err == nil {
			cart.CreatedAt = t
		}
	}
	if tmp.UpdatedAt != "" {
		if t, err := time.Parse(timeFormat, tmp.UpdatedAt); err == nil {
			cart.UpdatedAt = t
		}
	}

	return nil
}

func (cart *Cart) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Cart
		CurrencyCode string `json:"currency_code"`
	}{
		Cart:         *cart,
		CurrencyCode: strings.ToUpper(cart.CurrencyCode),
	}
	return json.Marshal(tmp)
}

func (order *Order) UnmarshalJSON(data []byte) error {
	type Alias Order
	tmp := &struct {
		*Alias
		ProcessedAtForeign string `json:"processed_at_foreign"`
		CancelledAtForeign string `json:"cancelled_at_foreign"`
		UpdatedAtForeign   string `json:"updated_at_foreign"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
	}{
		Alias: (*Alias)(order),
	}

	if err := json.Unmarshal(data, tmp); err != nil {
		return err
	}

	if tmp.ProcessedAtForeign != "" {
		if t, err := time.Parse(timeFormat, tmp.ProcessedAtForeign); err == nil {
			order.ProcessedAtForeign = t
		}
	}
	if tmp.CancelledAtForeign != "" {
		if t, err := time.Parse(timeFormat, tmp.CancelledAtForeign); err == nil {
			order.CancelledAtForeign = t
		}
	}
	if tmp.UpdatedAtForeign != "" {
		if t, err := time.Parse(timeFormat, tmp.UpdatedAtForeign); err == nil {
			order.UpdatedAtForeign = t
		}
	}
	if tmp.CreatedAt != "" {
		if t, err := time.Parse(timeFormat, tmp.CreatedAt); err == nil {
			order.CreatedAt = t
		}
	}
	if tmp.UpdatedAt != "" {
		if t, err := time.Parse(timeFormat, tmp.UpdatedAt); err == nil {
			order.UpdatedAt = t
		}
	}

	return nil
}

func (order *Order) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Order
		CurrencyCode       string `json:"currency_code"`
		ProcessedAtForeign string `json:"processed_at_foreign"`
		CancelledAtForeign string `json:"cancelled_at_foreign"`
		UpdatedAtForeign   string `json:"updated_at_foreign"`
	}{
		Order:              *order,
		CurrencyCode:       strings.ToUpper(order.CurrencyCode),
		ProcessedAtForeign: order.ProcessedAtForeign.Format(timeFormat),
		CancelledAtForeign: order.CancelledAtForeign.Format(timeFormat),
		UpdatedAtForeign:   order.UpdatedAtForeign.Format(timeFormat),
	}
	return json.Marshal(tmp)
}

func (product *Product) UnmarshalJSON(data []byte) error {
	type Alias Product
	tmp := &struct {
		*Alias
		PublishedAtForeign string `json:"published_at_foreign"`
	}{
		Alias: (*Alias)(product),
	}

	if err := json.Unmarshal(data, tmp); err != nil {
		return err
	}

	if tmp.PublishedAtForeign != "" {
		t, err := time.Parse(timeFormat, tmp.PublishedAtForeign)
		if err == nil {
			product.PublishedAtForeign = t
		}
	}

	return nil
}

func (product *Product) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Product
		PublishedAtForeign string `json:"published_at_foreign"`
	}{
		Product:            *product,
		PublishedAtForeign: product.PublishedAtForeign.Format(timeFormat),
	}
	return json.Marshal(tmp)
}
