package entities

import (
	"encoding/json"
	"io"
)

type Settings struct {
	InAppProducts               []Product `json:"inAppProducts,omitempty"`
	NativeProducts              []Product `json:"nativeProducts,omitempty"`
	WebPaymentUrl               string    `json:"webPaymentUrl,omitempty"`
	PromoCodesEnabled           bool      `json:"promoCodesEnabled,omitempty"`
	WebPaymentMonthProductPrice Price     `json:"webPaymentMonthProductPrice"`
}

type SettingsResult struct {
	ResponseWithoutResult
	Result Settings
}

func (s *SettingsResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
