package ormtable

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/internal/fieldnames"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"

	"github.com/cosmos/cosmos-sdk/orm/encoding/encodeutil"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/orm/encoding/ormkv"
)

// primaryKeyIndex defines an UniqueIndex for the primary key.
type primaryKeyIndex struct {
	*ormkv.PrimaryKeyCodec
	fields         fieldnames.FieldNames
	indexers       []indexer
	getBackend     func(context.Context) (Backend, error)
	getReadBackend func(context.Context) (ReadBackend, error)
}

func (p primaryKeyIndex) List(ctx context.Context, prefixKey []interface{}, options ...ormlist.Option) (Iterator, error) {
	backend, err := p.getReadBackend(ctx)
	if err != nil {
		return nil, err
	}

	return prefixIterator(backend.CommitmentStoreReader(), backend, p, p.KeyCodec, prefixKey, options)
}

func (p primaryKeyIndex) ListRange(ctx context.Context, from, to []interface{}, options ...ormlist.Option) (Iterator, error) {
	backend, err := p.getReadBackend(ctx)
	if err != nil {
		return nil, err
	}

	return rangeIterator(backend.CommitmentStoreReader(), backend, p, p.KeyCodec, from, to, options)
}

func (p primaryKeyIndex) doNotImplement() {}

func (p primaryKeyIndex) Has(ctx context.Context, key ...interface{}) (found bool, err error) {
	backend, err := p.getReadBackend(ctx)
	if err != nil {
		return false, err
	}

	return p.has(backend, encodeutil.ValuesOf(key...))
}

func (p primaryKeyIndex) has(backend ReadBackend, values []protoreflect.Value) (found bool, err error) {
	keyBz, err := p.EncodeKey(values)
	if err != nil {
		return false, err
	}

	return backend.CommitmentStoreReader().Has(keyBz)
}

func (p primaryKeyIndex) Get(ctx context.Context, message proto.Message, values ...interface{}) (found bool, err error) {
	backend, err := p.getReadBackend(ctx)
	if err != nil {
		return false, err
	}

	return p.get(backend, message, encodeutil.ValuesOf(values...))
}

func (p primaryKeyIndex) get(backend ReadBackend, message proto.Message, values []protoreflect.Value) (found bool, err error) {
	key, err := p.EncodeKey(values)
	if err != nil {
		return false, err
	}

	return p.getByKeyBytes(backend, key, values, message)
}

func (p primaryKeyIndex) DeleteBy(ctx context.Context, primaryKeyValues ...interface{}) error {
	if len(primaryKeyValues) == len(p.GetFieldNames()) {
		return p.doDelete(ctx, encodeutil.ValuesOf(primaryKeyValues...))
	}

	it, err := p.List(ctx, primaryKeyValues)
	if err != nil {
		return err
	}

	return p.deleteByIterator(ctx, it)
}

func (p primaryKeyIndex) DeleteRange(ctx context.Context, from, to []interface{}) error {
	it, err := p.ListRange(ctx, from, to)
	if err != nil {
		return err
	}

	return p.deleteByIterator(ctx, it)
}

func (p primaryKeyIndex) doDelete(ctx context.Context, primaryKeyValues []protoreflect.Value) error {
	backend, err := p.getBackend(ctx)
	if err != nil {
		return err
	}

	// delete object
	writer := newBatchIndexCommitmentWriter(backend)
	defer writer.Close()

	pk, err := p.EncodeKey(primaryKeyValues)
	if err != nil {
		return err
	}

	msg := p.MessageType().New().Interface()
	found, err := p.getByKeyBytes(backend, pk, primaryKeyValues, msg)
	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	err = p.doDeleteWithWriteBatch(backend, writer, pk, msg)
	if err != nil {
		return err
	}

	return writer.WriteUpdates()
}

func (p primaryKeyIndex) doDeleteWithWriteBatch(backend Backend, writer *batchIndexCommitmentWriter, primaryKeyBz []byte, message proto.Message) error {
	if hooks := backend.Hooks(); hooks != nil {
		err := hooks.OnDelete(message)
		if err != nil {
			return err
		}
	}

	// delete object
	err := writer.CommitmentStore().Delete(primaryKeyBz)
	if err != nil {
		return err
	}

	// clear indexes
	mref := message.ProtoReflect()
	indexStoreWriter := writer.IndexStore()
	for _, idx := range p.indexers {
		err := idx.onDelete(indexStoreWriter, mref)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p primaryKeyIndex) getByKeyBytes(store ReadBackend, key []byte, keyValues []protoreflect.Value, message proto.Message) (found bool, err error) {
	bz, err := store.CommitmentStoreReader().Get(key)
	if err != nil {
		return false, err
	}

	if bz == nil {
		return false, nil
	}

	return true, p.Unmarshal(keyValues, bz, message)
}

func (p primaryKeyIndex) readValueFromIndexKey(_ ReadBackend, primaryKey []protoreflect.Value, value []byte, message proto.Message) error {
	return p.Unmarshal(primaryKey, value, message)
}

func (p primaryKeyIndex) Fields() string {
	return p.fields.String()
}

func (p primaryKeyIndex) deleteByIterator(ctx context.Context, it Iterator) error {
	backend, err := p.getBackend(ctx)
	if err != nil {
		return err
	}

	// we batch writes while the iterator is still open
	writer := newBatchIndexCommitmentWriter(backend)
	defer writer.Close()

	for it.Next() {
		err = p.deleteOneByIterator(backend, writer, it)
		if err != nil {
			return err
		}
	}

	// close iterator
	it.Close()
	// then write batch
	return writer.WriteUpdates()
}

func (p primaryKeyIndex) deleteOneByIterator(backend Backend, writer *batchIndexCommitmentWriter, it Iterator) error {
	_, pk, err := it.Keys()
	if err != nil {
		return err
	}

	msg, err := it.GetMessage()
	if err != nil {
		return err
	}

	pkBz, err := p.EncodeKey(pk)
	if err != nil {
		return err
	}

	err = p.doDeleteWithWriteBatch(backend, writer, pkBz, msg)
	if err != nil {
		return err
	}

	return nil
}

var _ UniqueIndex = &primaryKeyIndex{}
