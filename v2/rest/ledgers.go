package rest

import (
	"github.com/outprog/bitfinex-api-go/v2"
	"path"
	"fmt"
)

// LedgerService manages the Ledgers endpoint.
type LedgerService struct {
	requestFactory
	Synchronous
}

// Retrieves all of the past ledger entreies
// see https://docs.bitfinex.com/reference#ledgers for more info
func (s *LedgerService) Ledgers(currency string, start int64, end int64, max int32) (*bitfinex.LedgerSnapshot, error) {
    if max > 500 {
    	return nil, fmt.Errorf("Max request limit is higher then 500 : %#v", max)
    }

	req, err := s.requestFactory.NewAuthenticatedRequestWithData(bitfinex.PermissionRead, path.Join("ledgers", currency, "hist"), map[string]interface{}{"start": start, "end": end, "limit": max})
	if err != nil {
		return nil, err
	}
	raw, err := s.Request(req)

	if err != nil {
		return nil, err
	}

	os, err := bitfinex.NewLedgerSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return os, nil
}
