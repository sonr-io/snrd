// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package svcv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AssertionTable interface {
	Insert(ctx context.Context, assertion *Assertion) error
	Update(ctx context.Context, assertion *Assertion) error
	Save(ctx context.Context, assertion *Assertion) error
	Delete(ctx context.Context, assertion *Assertion) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Assertion, error)
	List(ctx context.Context, prefixKey AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error)
	ListRange(ctx context.Context, from, to AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error)
	DeleteBy(ctx context.Context, prefixKey AssertionIndexKey) error
	DeleteRange(ctx context.Context, from, to AssertionIndexKey) error

	doNotImplement()
}

type AssertionIterator struct {
	ormtable.Iterator
}

func (i AssertionIterator) Value() (*Assertion, error) {
	var assertion Assertion
	err := i.UnmarshalMessage(&assertion)
	return &assertion, err
}

type AssertionIndexKey interface {
	id() uint32
	values() []interface{}
	assertionIndexKey()
}

// primary key starting index..
type AssertionPrimaryKey = AssertionIdIndexKey

type AssertionIdIndexKey struct {
	vs []interface{}
}

func (x AssertionIdIndexKey) id() uint32            { return 0 }
func (x AssertionIdIndexKey) values() []interface{} { return x.vs }
func (x AssertionIdIndexKey) assertionIndexKey()    {}

func (this AssertionIdIndexKey) WithId(id string) AssertionIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type assertionTable struct {
	table ormtable.Table
}

func (this assertionTable) Insert(ctx context.Context, assertion *Assertion) error {
	return this.table.Insert(ctx, assertion)
}

func (this assertionTable) Update(ctx context.Context, assertion *Assertion) error {
	return this.table.Update(ctx, assertion)
}

func (this assertionTable) Save(ctx context.Context, assertion *Assertion) error {
	return this.table.Save(ctx, assertion)
}

func (this assertionTable) Delete(ctx context.Context, assertion *Assertion) error {
	return this.table.Delete(ctx, assertion)
}

func (this assertionTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this assertionTable) Get(ctx context.Context, id string) (*Assertion, error) {
	var assertion Assertion
	found, err := this.table.PrimaryKey().Get(ctx, &assertion, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &assertion, nil
}

func (this assertionTable) List(ctx context.Context, prefixKey AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AssertionIterator{it}, err
}

func (this assertionTable) ListRange(ctx context.Context, from, to AssertionIndexKey, opts ...ormlist.Option) (AssertionIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AssertionIterator{it}, err
}

func (this assertionTable) DeleteBy(ctx context.Context, prefixKey AssertionIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this assertionTable) DeleteRange(ctx context.Context, from, to AssertionIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this assertionTable) doNotImplement() {}

var _ AssertionTable = assertionTable{}

func NewAssertionTable(db ormtable.Schema) (AssertionTable, error) {
	table := db.GetTable(&Assertion{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Assertion{}).ProtoReflect().Descriptor().FullName()))
	}
	return assertionTable{table}, nil
}

type ProfileTable interface {
	Insert(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	Save(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, profile *Profile) error
	Has(ctx context.Context, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id string) (*Profile, error)
	List(ctx context.Context, prefixKey ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error)
	ListRange(ctx context.Context, from, to ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error)
	DeleteBy(ctx context.Context, prefixKey ProfileIndexKey) error
	DeleteRange(ctx context.Context, from, to ProfileIndexKey) error

	doNotImplement()
}

type ProfileIterator struct {
	ormtable.Iterator
}

func (i ProfileIterator) Value() (*Profile, error) {
	var profile Profile
	err := i.UnmarshalMessage(&profile)
	return &profile, err
}

type ProfileIndexKey interface {
	id() uint32
	values() []interface{}
	profileIndexKey()
}

// primary key starting index..
type ProfilePrimaryKey = ProfileIdIndexKey

type ProfileIdIndexKey struct {
	vs []interface{}
}

func (x ProfileIdIndexKey) id() uint32            { return 0 }
func (x ProfileIdIndexKey) values() []interface{} { return x.vs }
func (x ProfileIdIndexKey) profileIndexKey()      {}

func (this ProfileIdIndexKey) WithId(id string) ProfileIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type profileTable struct {
	table ormtable.Table
}

func (this profileTable) Insert(ctx context.Context, profile *Profile) error {
	return this.table.Insert(ctx, profile)
}

func (this profileTable) Update(ctx context.Context, profile *Profile) error {
	return this.table.Update(ctx, profile)
}

func (this profileTable) Save(ctx context.Context, profile *Profile) error {
	return this.table.Save(ctx, profile)
}

func (this profileTable) Delete(ctx context.Context, profile *Profile) error {
	return this.table.Delete(ctx, profile)
}

func (this profileTable) Has(ctx context.Context, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this profileTable) Get(ctx context.Context, id string) (*Profile, error) {
	var profile Profile
	found, err := this.table.PrimaryKey().Get(ctx, &profile, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &profile, nil
}

func (this profileTable) List(ctx context.Context, prefixKey ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ProfileIterator{it}, err
}

func (this profileTable) ListRange(ctx context.Context, from, to ProfileIndexKey, opts ...ormlist.Option) (ProfileIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ProfileIterator{it}, err
}

func (this profileTable) DeleteBy(ctx context.Context, prefixKey ProfileIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this profileTable) DeleteRange(ctx context.Context, from, to ProfileIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this profileTable) doNotImplement() {}

var _ ProfileTable = profileTable{}

func NewProfileTable(db ormtable.Schema) (ProfileTable, error) {
	table := db.GetTable(&Profile{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Profile{}).ProtoReflect().Descriptor().FullName()))
	}
	return profileTable{table}, nil
}

type StateStore interface {
	AssertionTable() AssertionTable
	ProfileTable() ProfileTable

	doNotImplement()
}

type stateStore struct {
	assertion AssertionTable
	profile   ProfileTable
}

func (x stateStore) AssertionTable() AssertionTable {
	return x.assertion
}

func (x stateStore) ProfileTable() ProfileTable {
	return x.profile
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	assertionTable, err := NewAssertionTable(db)
	if err != nil {
		return nil, err
	}

	profileTable, err := NewProfileTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		assertionTable,
		profileTable,
	}, nil
}
