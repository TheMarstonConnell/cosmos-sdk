package bankexample

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type StateStore interface {
	Balance() BalanceStore
	Supply() SupplyStore
}

type BalanceStore interface {
	Insert(ctx context.Context, balance *Balance) error
	Update(ctx context.Context, balance *Balance) error
	Save(ctx context.Context, balance *Balance) error
	Delete(ctx context.Context, balance *Balance) error
	Has(ctx context.Context, address string, denom string) (found bool, err error)
	Get(ctx context.Context, address string, denom string) (*Balance, error)
	List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
}

type BalanceIterator struct {
	ormtable.Iterator
}

func (i BalanceIterator) Value() (*Balance, error) {
	var balance Balance
	err := i.UnmarshalMessage(&balance)
	return &balance, err
}

type BalanceIndexKey interface {
	id() uint32
	values() []interface{}
	balanceIndexKey()
}

type BalanceAddressDenomIndexKey struct{ vs []interface{} }

func (b BalanceAddressDenomIndexKey) id() uint32            { return 0 /* primary key */ }
func (b BalanceAddressDenomIndexKey) values() []interface{} { return b.vs }
func (b BalanceAddressDenomIndexKey) balanceIndexKey()      {}

var _ BalanceIndexKey = BalanceAddressDenomIndexKey{}

func (b BalanceAddressDenomIndexKey) WithAddress(address string) BalanceAddressDenomIndexKey {
	b.vs = []interface{}{address}
	return b
}

func (b BalanceAddressDenomIndexKey) WithAddressDenom(address string, denom string) BalanceAddressDenomIndexKey {
	b.vs = []interface{}{address, denom}
	return b
}

type BalanceDenomAddressIndexKey struct{ vs []interface{} }

func (b BalanceDenomAddressIndexKey) WithDenom(denom string) BalanceDenomAddressIndexKey {
	b.vs = []interface{}{denom}
	return b
}

func (b BalanceDenomAddressIndexKey) WithDenomAddress(denom string, address string) BalanceDenomAddressIndexKey {
	b.vs = []interface{}{denom, address}
	return b
}

func (b BalanceDenomAddressIndexKey) id() uint32            { return 1 }
func (b BalanceDenomAddressIndexKey) values() []interface{} { return b.vs }
func (b BalanceDenomAddressIndexKey) balanceIndexKey()      {}

var _ BalanceIndexKey = BalanceDenomAddressIndexKey{}

type balanceStore struct {
	table ormtable.Table
}

func (b balanceStore) Insert(ctx context.Context, balance *Balance) error {
	return b.table.Insert(ctx, balance)
}

func (b balanceStore) Update(ctx context.Context, balance *Balance) error {
	return b.table.Update(ctx, balance)
}

func (b balanceStore) Save(ctx context.Context, balance *Balance) error {
	return b.table.Save(ctx, balance)
}

func (b balanceStore) Delete(ctx context.Context, balance *Balance) error {
	return b.table.Delete(ctx, balance)
}

func (b balanceStore) Has(ctx context.Context, address string, denom string) (found bool, err error) {
	return b.table.PrimaryKey().Has(ctx, address, denom)
}

func (b balanceStore) Get(ctx context.Context, address string, denom string) (*Balance, error) {
	var balance Balance
	found, err := b.table.PrimaryKey().Get(ctx, &balance, address, denom)
	if !found {
		return nil, err
	}
	return &balance, err
}

func (b balanceStore) List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	opts = append(opts, ormlist.Prefix(prefixKey.values()))
	it, err := b.table.GetIndexByID(prefixKey.id()).Iterator(ctx, opts...)
	return BalanceIterator{it}, err
}

func (b balanceStore) ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	opts = append(opts, ormlist.Start(from.values()), ormlist.End(to))
	it, err := b.table.GetIndexByID(from.id()).Iterator(ctx, opts...)
	return BalanceIterator{it}, err
}

var _ BalanceStore = balanceStore{}

type SupplyStore interface {
	// TODO
}

type supplyStore struct {
	table ormtable.Table
}

type stateStore struct {
	balance *balanceStore
	supply  *supplyStore
}

func (s stateStore) Balance() BalanceStore {
	return s.balance
}

func (s stateStore) Supply() SupplyStore {
	return s.supply
}

var _ StateStore = stateStore{}
