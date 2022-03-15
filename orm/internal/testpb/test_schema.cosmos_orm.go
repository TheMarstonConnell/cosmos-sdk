// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package testpb

import (
	context "context"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type ExampleTableTable interface {
	Insert(ctx context.Context, exampleTable *ExampleTable) error
	Update(ctx context.Context, exampleTable *ExampleTable) error
	Save(ctx context.Context, exampleTable *ExampleTable) error
	Delete(ctx context.Context, exampleTable *ExampleTable) error
	Has(ctx context.Context, u32 uint32, i64 int64, str string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, u32 uint32, i64 int64, str string) (*ExampleTable, error)
	HasByU64Str(ctx context.Context, u64 uint64, str string) (found bool, err error)
	// GetByU64Str returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByU64Str(ctx context.Context, u64 uint64, str string) (*ExampleTable, error)
	List(ctx context.Context, prefixKey ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error)
	ListRange(ctx context.Context, from, to ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error)
	DeleteBy(ctx context.Context, prefixKey ExampleTableIndexKey) error
	DeleteRange(ctx context.Context, from, to ExampleTableIndexKey) error

	doNotImplement()
}

type ExampleTableIterator struct {
	ormtable.Iterator
}

func (i ExampleTableIterator) Value() (*ExampleTable, error) {
	var exampleTable ExampleTable
	err := i.UnmarshalMessage(&exampleTable)
	return &exampleTable, err
}

type ExampleTableIndexKey interface {
	id() uint32
	values() []interface{}
	exampleTableIndexKey()
}

// primary key starting index..
type ExampleTablePrimaryKey = ExampleTableU32I64StrIndexKey

type ExampleTableU32I64StrIndexKey struct {
	vs []interface{}
}

func (x ExampleTableU32I64StrIndexKey) id() uint32            { return 0 }
func (x ExampleTableU32I64StrIndexKey) values() []interface{} { return x.vs }
func (x ExampleTableU32I64StrIndexKey) exampleTableIndexKey() {}

func (this ExampleTableU32I64StrIndexKey) WithU32(u32 uint32) ExampleTableU32I64StrIndexKey {
	this.vs = []interface{}{u32}
	return this
}

func (this ExampleTableU32I64StrIndexKey) WithU32I64(u32 uint32, i64 int64) ExampleTableU32I64StrIndexKey {
	this.vs = []interface{}{u32, i64}
	return this
}

func (this ExampleTableU32I64StrIndexKey) WithU32I64Str(u32 uint32, i64 int64, str string) ExampleTableU32I64StrIndexKey {
	this.vs = []interface{}{u32, i64, str}
	return this
}

type ExampleTableU64StrIndexKey struct {
	vs []interface{}
}

func (x ExampleTableU64StrIndexKey) id() uint32            { return 1 }
func (x ExampleTableU64StrIndexKey) values() []interface{} { return x.vs }
func (x ExampleTableU64StrIndexKey) exampleTableIndexKey() {}

func (this ExampleTableU64StrIndexKey) WithU64(u64 uint64) ExampleTableU64StrIndexKey {
	this.vs = []interface{}{u64}
	return this
}

func (this ExampleTableU64StrIndexKey) WithU64Str(u64 uint64, str string) ExampleTableU64StrIndexKey {
	this.vs = []interface{}{u64, str}
	return this
}

type ExampleTableStrU32IndexKey struct {
	vs []interface{}
}

func (x ExampleTableStrU32IndexKey) id() uint32            { return 2 }
func (x ExampleTableStrU32IndexKey) values() []interface{} { return x.vs }
func (x ExampleTableStrU32IndexKey) exampleTableIndexKey() {}

func (this ExampleTableStrU32IndexKey) WithStr(str string) ExampleTableStrU32IndexKey {
	this.vs = []interface{}{str}
	return this
}

func (this ExampleTableStrU32IndexKey) WithStrU32(str string, u32 uint32) ExampleTableStrU32IndexKey {
	this.vs = []interface{}{str, u32}
	return this
}

type ExampleTableBzStrIndexKey struct {
	vs []interface{}
}

func (x ExampleTableBzStrIndexKey) id() uint32            { return 3 }
func (x ExampleTableBzStrIndexKey) values() []interface{} { return x.vs }
func (x ExampleTableBzStrIndexKey) exampleTableIndexKey() {}

func (this ExampleTableBzStrIndexKey) WithBz(bz []byte) ExampleTableBzStrIndexKey {
	this.vs = []interface{}{bz}
	return this
}

func (this ExampleTableBzStrIndexKey) WithBzStr(bz []byte, str string) ExampleTableBzStrIndexKey {
	this.vs = []interface{}{bz, str}
	return this
}

type exampleTableTable struct {
	table ormtable.Table
}

func (this exampleTableTable) Insert(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Insert(ctx, exampleTable)
}

func (this exampleTableTable) Update(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Update(ctx, exampleTable)
}

func (this exampleTableTable) Save(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Save(ctx, exampleTable)
}

func (this exampleTableTable) Delete(ctx context.Context, exampleTable *ExampleTable) error {
	return this.table.Delete(ctx, exampleTable)
}

func (this exampleTableTable) Has(ctx context.Context, u32 uint32, i64 int64, str string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, u32, i64, str)
}

func (this exampleTableTable) Get(ctx context.Context, u32 uint32, i64 int64, str string) (*ExampleTable, error) {
	var exampleTable ExampleTable
	err := this.table.PrimaryKey().Get(ctx, &exampleTable, u32, i64, str)
	if err != nil {
		return nil, err
	}
	return &exampleTable, nil
}

func (this exampleTableTable) HasByU64Str(ctx context.Context, u64 uint64, str string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		u64,
		str,
	)
}

func (this exampleTableTable) GetByU64Str(ctx context.Context, u64 uint64, str string) (*ExampleTable, error) {
	var exampleTable ExampleTable
	err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &exampleTable,
		u64,
		str,
	)
	if err != nil {
		return nil, err
	}
	return &exampleTable, nil
}

func (this exampleTableTable) List(ctx context.Context, prefixKey ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ExampleTableIterator{it}, err
}

func (this exampleTableTable) ListRange(ctx context.Context, from, to ExampleTableIndexKey, opts ...ormlist.Option) (ExampleTableIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ExampleTableIterator{it}, err
}

func (this exampleTableTable) DeleteBy(ctx context.Context, prefixKey ExampleTableIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this exampleTableTable) DeleteRange(ctx context.Context, from, to ExampleTableIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this exampleTableTable) doNotImplement() {}

var _ ExampleTableTable = exampleTableTable{}

func NewExampleTableTable(db ormtable.Schema) (ExampleTableTable, error) {
	table := db.GetTable(&ExampleTable{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleTable{}).ProtoReflect().Descriptor().FullName()))
	}
	return exampleTableTable{table}, nil
}

type ExampleAutoIncrementTableTable interface {
	Insert(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	InsertReturningID(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) (uint64, error)
	Update(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	Save(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	Delete(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*ExampleAutoIncrementTable, error)
	HasByX(ctx context.Context, x string) (found bool, err error)
	// GetByX returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByX(ctx context.Context, x string) (*ExampleAutoIncrementTable, error)
	List(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error)
	ListRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error)
	DeleteBy(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey) error
	DeleteRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey) error

	doNotImplement()
}

type ExampleAutoIncrementTableIterator struct {
	ormtable.Iterator
}

func (i ExampleAutoIncrementTableIterator) Value() (*ExampleAutoIncrementTable, error) {
	var exampleAutoIncrementTable ExampleAutoIncrementTable
	err := i.UnmarshalMessage(&exampleAutoIncrementTable)
	return &exampleAutoIncrementTable, err
}

type ExampleAutoIncrementTableIndexKey interface {
	id() uint32
	values() []interface{}
	exampleAutoIncrementTableIndexKey()
}

// primary key starting index..
type ExampleAutoIncrementTablePrimaryKey = ExampleAutoIncrementTableIdIndexKey

type ExampleAutoIncrementTableIdIndexKey struct {
	vs []interface{}
}

func (x ExampleAutoIncrementTableIdIndexKey) id() uint32                         { return 0 }
func (x ExampleAutoIncrementTableIdIndexKey) values() []interface{}              { return x.vs }
func (x ExampleAutoIncrementTableIdIndexKey) exampleAutoIncrementTableIndexKey() {}

func (this ExampleAutoIncrementTableIdIndexKey) WithId(id uint64) ExampleAutoIncrementTableIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type ExampleAutoIncrementTableXIndexKey struct {
	vs []interface{}
}

func (x ExampleAutoIncrementTableXIndexKey) id() uint32                         { return 1 }
func (x ExampleAutoIncrementTableXIndexKey) values() []interface{}              { return x.vs }
func (x ExampleAutoIncrementTableXIndexKey) exampleAutoIncrementTableIndexKey() {}

func (this ExampleAutoIncrementTableXIndexKey) WithX(x string) ExampleAutoIncrementTableXIndexKey {
	this.vs = []interface{}{x}
	return this
}

type exampleAutoIncrementTableTable struct {
	table ormtable.AutoIncrementTable
}

func (this exampleAutoIncrementTableTable) Insert(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Insert(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableTable) Update(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Update(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableTable) Save(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Save(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableTable) Delete(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) error {
	return this.table.Delete(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableTable) InsertReturningID(ctx context.Context, exampleAutoIncrementTable *ExampleAutoIncrementTable) (uint64, error) {
	return this.table.InsertReturningID(ctx, exampleAutoIncrementTable)
}

func (this exampleAutoIncrementTableTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this exampleAutoIncrementTableTable) Get(ctx context.Context, id uint64) (*ExampleAutoIncrementTable, error) {
	var exampleAutoIncrementTable ExampleAutoIncrementTable
	err := this.table.PrimaryKey().Get(ctx, &exampleAutoIncrementTable, id)
	if err != nil {
		return nil, err
	}
	return &exampleAutoIncrementTable, nil
}

func (this exampleAutoIncrementTableTable) HasByX(ctx context.Context, x string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		x,
	)
}

func (this exampleAutoIncrementTableTable) GetByX(ctx context.Context, x string) (*ExampleAutoIncrementTable, error) {
	var exampleAutoIncrementTable ExampleAutoIncrementTable
	err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &exampleAutoIncrementTable,
		x,
	)
	if err != nil {
		return nil, err
	}
	return &exampleAutoIncrementTable, nil
}

func (this exampleAutoIncrementTableTable) List(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ExampleAutoIncrementTableIterator{it}, err
}

func (this exampleAutoIncrementTableTable) ListRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey, opts ...ormlist.Option) (ExampleAutoIncrementTableIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ExampleAutoIncrementTableIterator{it}, err
}

func (this exampleAutoIncrementTableTable) DeleteBy(ctx context.Context, prefixKey ExampleAutoIncrementTableIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this exampleAutoIncrementTableTable) DeleteRange(ctx context.Context, from, to ExampleAutoIncrementTableIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this exampleAutoIncrementTableTable) doNotImplement() {}

var _ ExampleAutoIncrementTableTable = exampleAutoIncrementTableTable{}

func NewExampleAutoIncrementTableTable(db ormtable.Schema) (ExampleAutoIncrementTableTable, error) {
	table := db.GetTable(&ExampleAutoIncrementTable{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleAutoIncrementTable{}).ProtoReflect().Descriptor().FullName()))
	}
	return exampleAutoIncrementTableTable{table.(ormtable.AutoIncrementTable)}, nil
}

// singleton store
type ExampleSingletonTable interface {
	Get(ctx context.Context) (*ExampleSingleton, error)
	Save(ctx context.Context, exampleSingleton *ExampleSingleton) error
}

type exampleSingletonTable struct {
	table ormtable.Table
}

var _ ExampleSingletonTable = exampleSingletonTable{}

func (x exampleSingletonTable) Get(ctx context.Context) (*ExampleSingleton, error) {
	exampleSingleton := &ExampleSingleton{}
	err := x.table.Get(ctx, exampleSingleton)
	return exampleSingleton, err
}

func (x exampleSingletonTable) Save(ctx context.Context, exampleSingleton *ExampleSingleton) error {
	return x.table.Save(ctx, exampleSingleton)
}

func NewExampleSingletonTable(db ormtable.Schema) (ExampleSingletonTable, error) {
	table := db.GetTable(&ExampleSingleton{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleSingleton{}).ProtoReflect().Descriptor().FullName()))
	}
	return &exampleSingletonTable{table}, nil
}

type ExampleTimestampTable interface {
	Insert(ctx context.Context, exampleTimestamp *ExampleTimestamp) error
	InsertReturningID(ctx context.Context, exampleTimestamp *ExampleTimestamp) (uint64, error)
	Update(ctx context.Context, exampleTimestamp *ExampleTimestamp) error
	Save(ctx context.Context, exampleTimestamp *ExampleTimestamp) error
	Delete(ctx context.Context, exampleTimestamp *ExampleTimestamp) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*ExampleTimestamp, error)
	List(ctx context.Context, prefixKey ExampleTimestampIndexKey, opts ...ormlist.Option) (ExampleTimestampIterator, error)
	ListRange(ctx context.Context, from, to ExampleTimestampIndexKey, opts ...ormlist.Option) (ExampleTimestampIterator, error)
	DeleteBy(ctx context.Context, prefixKey ExampleTimestampIndexKey) error
	DeleteRange(ctx context.Context, from, to ExampleTimestampIndexKey) error

	doNotImplement()
}

type ExampleTimestampIterator struct {
	ormtable.Iterator
}

func (i ExampleTimestampIterator) Value() (*ExampleTimestamp, error) {
	var exampleTimestamp ExampleTimestamp
	err := i.UnmarshalMessage(&exampleTimestamp)
	return &exampleTimestamp, err
}

type ExampleTimestampIndexKey interface {
	id() uint32
	values() []interface{}
	exampleTimestampIndexKey()
}

// primary key starting index..
type ExampleTimestampPrimaryKey = ExampleTimestampIdIndexKey

type ExampleTimestampIdIndexKey struct {
	vs []interface{}
}

func (x ExampleTimestampIdIndexKey) id() uint32                { return 0 }
func (x ExampleTimestampIdIndexKey) values() []interface{}     { return x.vs }
func (x ExampleTimestampIdIndexKey) exampleTimestampIndexKey() {}

func (this ExampleTimestampIdIndexKey) WithId(id uint64) ExampleTimestampIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type ExampleTimestampTsIndexKey struct {
	vs []interface{}
}

func (x ExampleTimestampTsIndexKey) id() uint32                { return 1 }
func (x ExampleTimestampTsIndexKey) values() []interface{}     { return x.vs }
func (x ExampleTimestampTsIndexKey) exampleTimestampIndexKey() {}

func (this ExampleTimestampTsIndexKey) WithTs(ts *timestamppb.Timestamp) ExampleTimestampTsIndexKey {
	this.vs = []interface{}{ts}
	return this
}

type exampleTimestampTable struct {
	table ormtable.AutoIncrementTable
}

func (this exampleTimestampTable) Insert(ctx context.Context, exampleTimestamp *ExampleTimestamp) error {
	return this.table.Insert(ctx, exampleTimestamp)
}

func (this exampleTimestampTable) Update(ctx context.Context, exampleTimestamp *ExampleTimestamp) error {
	return this.table.Update(ctx, exampleTimestamp)
}

func (this exampleTimestampTable) Save(ctx context.Context, exampleTimestamp *ExampleTimestamp) error {
	return this.table.Save(ctx, exampleTimestamp)
}

func (this exampleTimestampTable) Delete(ctx context.Context, exampleTimestamp *ExampleTimestamp) error {
	return this.table.Delete(ctx, exampleTimestamp)
}

func (this exampleTimestampTable) InsertReturningID(ctx context.Context, exampleTimestamp *ExampleTimestamp) (uint64, error) {
	return this.table.InsertReturningID(ctx, exampleTimestamp)
}

func (this exampleTimestampTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this exampleTimestampTable) Get(ctx context.Context, id uint64) (*ExampleTimestamp, error) {
	var exampleTimestamp ExampleTimestamp
	err := this.table.PrimaryKey().Get(ctx, &exampleTimestamp, id)
	if err != nil {
		return nil, err
	}
	return &exampleTimestamp, nil
}

func (this exampleTimestampTable) List(ctx context.Context, prefixKey ExampleTimestampIndexKey, opts ...ormlist.Option) (ExampleTimestampIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ExampleTimestampIterator{it}, err
}

func (this exampleTimestampTable) ListRange(ctx context.Context, from, to ExampleTimestampIndexKey, opts ...ormlist.Option) (ExampleTimestampIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ExampleTimestampIterator{it}, err
}

func (this exampleTimestampTable) DeleteBy(ctx context.Context, prefixKey ExampleTimestampIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this exampleTimestampTable) DeleteRange(ctx context.Context, from, to ExampleTimestampIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this exampleTimestampTable) doNotImplement() {}

var _ ExampleTimestampTable = exampleTimestampTable{}

func NewExampleTimestampTable(db ormtable.Schema) (ExampleTimestampTable, error) {
	table := db.GetTable(&ExampleTimestamp{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ExampleTimestamp{}).ProtoReflect().Descriptor().FullName()))
	}
	return exampleTimestampTable{table.(ormtable.AutoIncrementTable)}, nil
}

type SimpleExampleTable interface {
	Insert(ctx context.Context, simpleExample *SimpleExample) error
	Update(ctx context.Context, simpleExample *SimpleExample) error
	Save(ctx context.Context, simpleExample *SimpleExample) error
	Delete(ctx context.Context, simpleExample *SimpleExample) error
	Has(ctx context.Context, name string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, name string) (*SimpleExample, error)
	HasByUnique(ctx context.Context, unique string) (found bool, err error)
	// GetByUnique returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByUnique(ctx context.Context, unique string) (*SimpleExample, error)
	List(ctx context.Context, prefixKey SimpleExampleIndexKey, opts ...ormlist.Option) (SimpleExampleIterator, error)
	ListRange(ctx context.Context, from, to SimpleExampleIndexKey, opts ...ormlist.Option) (SimpleExampleIterator, error)
	DeleteBy(ctx context.Context, prefixKey SimpleExampleIndexKey) error
	DeleteRange(ctx context.Context, from, to SimpleExampleIndexKey) error

	doNotImplement()
}

type SimpleExampleIterator struct {
	ormtable.Iterator
}

func (i SimpleExampleIterator) Value() (*SimpleExample, error) {
	var simpleExample SimpleExample
	err := i.UnmarshalMessage(&simpleExample)
	return &simpleExample, err
}

type SimpleExampleIndexKey interface {
	id() uint32
	values() []interface{}
	simpleExampleIndexKey()
}

// primary key starting index..
type SimpleExamplePrimaryKey = SimpleExampleNameIndexKey

type SimpleExampleNameIndexKey struct {
	vs []interface{}
}

func (x SimpleExampleNameIndexKey) id() uint32             { return 0 }
func (x SimpleExampleNameIndexKey) values() []interface{}  { return x.vs }
func (x SimpleExampleNameIndexKey) simpleExampleIndexKey() {}

func (this SimpleExampleNameIndexKey) WithName(name string) SimpleExampleNameIndexKey {
	this.vs = []interface{}{name}
	return this
}

type SimpleExampleUniqueIndexKey struct {
	vs []interface{}
}

func (x SimpleExampleUniqueIndexKey) id() uint32             { return 1 }
func (x SimpleExampleUniqueIndexKey) values() []interface{}  { return x.vs }
func (x SimpleExampleUniqueIndexKey) simpleExampleIndexKey() {}

func (this SimpleExampleUniqueIndexKey) WithUnique(unique string) SimpleExampleUniqueIndexKey {
	this.vs = []interface{}{unique}
	return this
}

type simpleExampleTable struct {
	table ormtable.Table
}

func (this simpleExampleTable) Insert(ctx context.Context, simpleExample *SimpleExample) error {
	return this.table.Insert(ctx, simpleExample)
}

func (this simpleExampleTable) Update(ctx context.Context, simpleExample *SimpleExample) error {
	return this.table.Update(ctx, simpleExample)
}

func (this simpleExampleTable) Save(ctx context.Context, simpleExample *SimpleExample) error {
	return this.table.Save(ctx, simpleExample)
}

func (this simpleExampleTable) Delete(ctx context.Context, simpleExample *SimpleExample) error {
	return this.table.Delete(ctx, simpleExample)
}

func (this simpleExampleTable) Has(ctx context.Context, name string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, name)
}

func (this simpleExampleTable) Get(ctx context.Context, name string) (*SimpleExample, error) {
	var simpleExample SimpleExample
	err := this.table.PrimaryKey().Get(ctx, &simpleExample, name)
	if err != nil {
		return nil, err
	}
	return &simpleExample, nil
}

func (this simpleExampleTable) HasByUnique(ctx context.Context, unique string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		unique,
	)
}

func (this simpleExampleTable) GetByUnique(ctx context.Context, unique string) (*SimpleExample, error) {
	var simpleExample SimpleExample
	err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &simpleExample,
		unique,
	)
	if err != nil {
		return nil, err
	}
	return &simpleExample, nil
}

func (this simpleExampleTable) List(ctx context.Context, prefixKey SimpleExampleIndexKey, opts ...ormlist.Option) (SimpleExampleIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return SimpleExampleIterator{it}, err
}

func (this simpleExampleTable) ListRange(ctx context.Context, from, to SimpleExampleIndexKey, opts ...ormlist.Option) (SimpleExampleIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return SimpleExampleIterator{it}, err
}

func (this simpleExampleTable) DeleteBy(ctx context.Context, prefixKey SimpleExampleIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this simpleExampleTable) DeleteRange(ctx context.Context, from, to SimpleExampleIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this simpleExampleTable) doNotImplement() {}

var _ SimpleExampleTable = simpleExampleTable{}

func NewSimpleExampleTable(db ormtable.Schema) (SimpleExampleTable, error) {
	table := db.GetTable(&SimpleExample{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&SimpleExample{}).ProtoReflect().Descriptor().FullName()))
	}
	return simpleExampleTable{table}, nil
}

type TestSchemaStore interface {
	ExampleTableTable() ExampleTableTable
	ExampleAutoIncrementTableTable() ExampleAutoIncrementTableTable
	ExampleSingletonTable() ExampleSingletonTable
	ExampleTimestampTable() ExampleTimestampTable
	SimpleExampleTable() SimpleExampleTable

	doNotImplement()
}

type testSchemaStore struct {
	exampleTable              ExampleTableTable
	exampleAutoIncrementTable ExampleAutoIncrementTableTable
	exampleSingleton          ExampleSingletonTable
	exampleTimestamp          ExampleTimestampTable
	simpleExample             SimpleExampleTable
}

func (x testSchemaStore) ExampleTableTable() ExampleTableTable {
	return x.exampleTable
}

func (x testSchemaStore) ExampleAutoIncrementTableTable() ExampleAutoIncrementTableTable {
	return x.exampleAutoIncrementTable
}

func (x testSchemaStore) ExampleSingletonTable() ExampleSingletonTable {
	return x.exampleSingleton
}

func (x testSchemaStore) ExampleTimestampTable() ExampleTimestampTable {
	return x.exampleTimestamp
}

func (x testSchemaStore) SimpleExampleTable() SimpleExampleTable {
	return x.simpleExample
}

func (testSchemaStore) doNotImplement() {}

var _ TestSchemaStore = testSchemaStore{}

func NewTestSchemaStore(db ormtable.Schema) (TestSchemaStore, error) {
	exampleTableTable, err := NewExampleTableTable(db)
	if err != nil {
		return nil, err
	}

	exampleAutoIncrementTableTable, err := NewExampleAutoIncrementTableTable(db)
	if err != nil {
		return nil, err
	}

	exampleSingletonTable, err := NewExampleSingletonTable(db)
	if err != nil {
		return nil, err
	}

	exampleTimestampTable, err := NewExampleTimestampTable(db)
	if err != nil {
		return nil, err
	}

	simpleExampleTable, err := NewSimpleExampleTable(db)
	if err != nil {
		return nil, err
	}

	return testSchemaStore{
		exampleTableTable,
		exampleAutoIncrementTableTable,
		exampleSingletonTable,
		exampleTimestampTable,
		simpleExampleTable,
	}, nil
}
