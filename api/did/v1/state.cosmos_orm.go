// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package didv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AuthenticationTable interface {
	Insert(ctx context.Context, authentication *Authentication) error
	Update(ctx context.Context, authentication *Authentication) error
	Save(ctx context.Context, authentication *Authentication) error
	Delete(ctx context.Context, authentication *Authentication) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Authentication, error)
	HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error)
	// GetByControllerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerSubject(ctx context.Context, controller string, subject string) (*Authentication, error)
	List(ctx context.Context, prefixKey AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error)
	ListRange(ctx context.Context, from, to AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error)
	DeleteBy(ctx context.Context, prefixKey AuthenticationIndexKey) error
	DeleteRange(ctx context.Context, from, to AuthenticationIndexKey) error

	doNotImplement()
}

type AuthenticationIterator struct {
	ormtable.Iterator
}

func (i AuthenticationIterator) Value() (*Authentication, error) {
	var authentication Authentication
	err := i.UnmarshalMessage(&authentication)
	return &authentication, err
}

type AuthenticationIndexKey interface {
	id() uint32
	values() []interface{}
	authenticationIndexKey()
}

// primary key starting index..
type AuthenticationPrimaryKey = AuthenticationDidIndexKey

type AuthenticationDidIndexKey struct {
	vs []interface{}
}

func (x AuthenticationDidIndexKey) id() uint32              { return 0 }
func (x AuthenticationDidIndexKey) values() []interface{}   { return x.vs }
func (x AuthenticationDidIndexKey) authenticationIndexKey() {}

func (this AuthenticationDidIndexKey) WithDid(did string) AuthenticationDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type AuthenticationControllerSubjectIndexKey struct {
	vs []interface{}
}

func (x AuthenticationControllerSubjectIndexKey) id() uint32              { return 1 }
func (x AuthenticationControllerSubjectIndexKey) values() []interface{}   { return x.vs }
func (x AuthenticationControllerSubjectIndexKey) authenticationIndexKey() {}

func (this AuthenticationControllerSubjectIndexKey) WithController(controller string) AuthenticationControllerSubjectIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this AuthenticationControllerSubjectIndexKey) WithControllerSubject(controller string, subject string) AuthenticationControllerSubjectIndexKey {
	this.vs = []interface{}{controller, subject}
	return this
}

type authenticationTable struct {
	table ormtable.Table
}

func (this authenticationTable) Insert(ctx context.Context, authentication *Authentication) error {
	return this.table.Insert(ctx, authentication)
}

func (this authenticationTable) Update(ctx context.Context, authentication *Authentication) error {
	return this.table.Update(ctx, authentication)
}

func (this authenticationTable) Save(ctx context.Context, authentication *Authentication) error {
	return this.table.Save(ctx, authentication)
}

func (this authenticationTable) Delete(ctx context.Context, authentication *Authentication) error {
	return this.table.Delete(ctx, authentication)
}

func (this authenticationTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this authenticationTable) Get(ctx context.Context, did string) (*Authentication, error) {
	var authentication Authentication
	found, err := this.table.PrimaryKey().Get(ctx, &authentication, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &authentication, nil
}

func (this authenticationTable) HasByControllerSubject(ctx context.Context, controller string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		subject,
	)
}

func (this authenticationTable) GetByControllerSubject(ctx context.Context, controller string, subject string) (*Authentication, error) {
	var authentication Authentication
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &authentication,
		controller,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &authentication, nil
}

func (this authenticationTable) List(ctx context.Context, prefixKey AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AuthenticationIterator{it}, err
}

func (this authenticationTable) ListRange(ctx context.Context, from, to AuthenticationIndexKey, opts ...ormlist.Option) (AuthenticationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AuthenticationIterator{it}, err
}

func (this authenticationTable) DeleteBy(ctx context.Context, prefixKey AuthenticationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this authenticationTable) DeleteRange(ctx context.Context, from, to AuthenticationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this authenticationTable) doNotImplement() {}

var _ AuthenticationTable = authenticationTable{}

func NewAuthenticationTable(db ormtable.Schema) (AuthenticationTable, error) {
	table := db.GetTable(&Authentication{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Authentication{}).ProtoReflect().Descriptor().FullName()))
	}
	return authenticationTable{table}, nil
}

type ControllerTable interface {
	Insert(ctx context.Context, controller *Controller) error
	InsertReturningNumber(ctx context.Context, controller *Controller) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, controller *Controller) error
	Save(ctx context.Context, controller *Controller) error
	Delete(ctx context.Context, controller *Controller) error
	Has(ctx context.Context, number uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, number uint64) (*Controller, error)
	HasBySonrAddress(ctx context.Context, sonr_address string) (found bool, err error)
	// GetBySonrAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetBySonrAddress(ctx context.Context, sonr_address string) (*Controller, error)
	HasByEthAddress(ctx context.Context, eth_address string) (found bool, err error)
	// GetByEthAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByEthAddress(ctx context.Context, eth_address string) (*Controller, error)
	HasByBtcAddress(ctx context.Context, btc_address string) (found bool, err error)
	// GetByBtcAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByBtcAddress(ctx context.Context, btc_address string) (*Controller, error)
	HasByDid(ctx context.Context, did string) (found bool, err error)
	// GetByDid returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByDid(ctx context.Context, did string) (*Controller, error)
	List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error
	DeleteRange(ctx context.Context, from, to ControllerIndexKey) error

	doNotImplement()
}

type ControllerIterator struct {
	ormtable.Iterator
}

func (i ControllerIterator) Value() (*Controller, error) {
	var controller Controller
	err := i.UnmarshalMessage(&controller)
	return &controller, err
}

type ControllerIndexKey interface {
	id() uint32
	values() []interface{}
	controllerIndexKey()
}

// primary key starting index..
type ControllerPrimaryKey = ControllerNumberIndexKey

type ControllerNumberIndexKey struct {
	vs []interface{}
}

func (x ControllerNumberIndexKey) id() uint32            { return 0 }
func (x ControllerNumberIndexKey) values() []interface{} { return x.vs }
func (x ControllerNumberIndexKey) controllerIndexKey()   {}

func (this ControllerNumberIndexKey) WithNumber(number uint64) ControllerNumberIndexKey {
	this.vs = []interface{}{number}
	return this
}

type ControllerSonrAddressIndexKey struct {
	vs []interface{}
}

func (x ControllerSonrAddressIndexKey) id() uint32            { return 1 }
func (x ControllerSonrAddressIndexKey) values() []interface{} { return x.vs }
func (x ControllerSonrAddressIndexKey) controllerIndexKey()   {}

func (this ControllerSonrAddressIndexKey) WithSonrAddress(sonr_address string) ControllerSonrAddressIndexKey {
	this.vs = []interface{}{sonr_address}
	return this
}

type ControllerEthAddressIndexKey struct {
	vs []interface{}
}

func (x ControllerEthAddressIndexKey) id() uint32            { return 2 }
func (x ControllerEthAddressIndexKey) values() []interface{} { return x.vs }
func (x ControllerEthAddressIndexKey) controllerIndexKey()   {}

func (this ControllerEthAddressIndexKey) WithEthAddress(eth_address string) ControllerEthAddressIndexKey {
	this.vs = []interface{}{eth_address}
	return this
}

type ControllerBtcAddressIndexKey struct {
	vs []interface{}
}

func (x ControllerBtcAddressIndexKey) id() uint32            { return 3 }
func (x ControllerBtcAddressIndexKey) values() []interface{} { return x.vs }
func (x ControllerBtcAddressIndexKey) controllerIndexKey()   {}

func (this ControllerBtcAddressIndexKey) WithBtcAddress(btc_address string) ControllerBtcAddressIndexKey {
	this.vs = []interface{}{btc_address}
	return this
}

type ControllerDidIndexKey struct {
	vs []interface{}
}

func (x ControllerDidIndexKey) id() uint32            { return 4 }
func (x ControllerDidIndexKey) values() []interface{} { return x.vs }
func (x ControllerDidIndexKey) controllerIndexKey()   {}

func (this ControllerDidIndexKey) WithDid(did string) ControllerDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type controllerTable struct {
	table ormtable.AutoIncrementTable
}

func (this controllerTable) Insert(ctx context.Context, controller *Controller) error {
	return this.table.Insert(ctx, controller)
}

func (this controllerTable) Update(ctx context.Context, controller *Controller) error {
	return this.table.Update(ctx, controller)
}

func (this controllerTable) Save(ctx context.Context, controller *Controller) error {
	return this.table.Save(ctx, controller)
}

func (this controllerTable) Delete(ctx context.Context, controller *Controller) error {
	return this.table.Delete(ctx, controller)
}

func (this controllerTable) InsertReturningNumber(ctx context.Context, controller *Controller) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, controller)
}

func (this controllerTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this controllerTable) Has(ctx context.Context, number uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, number)
}

func (this controllerTable) Get(ctx context.Context, number uint64) (*Controller, error) {
	var controller Controller
	found, err := this.table.PrimaryKey().Get(ctx, &controller, number)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasBySonrAddress(ctx context.Context, sonr_address string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		sonr_address,
	)
}

func (this controllerTable) GetBySonrAddress(ctx context.Context, sonr_address string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &controller,
		sonr_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByEthAddress(ctx context.Context, eth_address string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		eth_address,
	)
}

func (this controllerTable) GetByEthAddress(ctx context.Context, eth_address string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &controller,
		eth_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByBtcAddress(ctx context.Context, btc_address string) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		btc_address,
	)
}

func (this controllerTable) GetByBtcAddress(ctx context.Context, btc_address string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &controller,
		btc_address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByDid(ctx context.Context, did string) (found bool, err error) {
	return this.table.GetIndexByID(4).(ormtable.UniqueIndex).Has(ctx,
		did,
	)
}

func (this controllerTable) GetByDid(ctx context.Context, did string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(4).(ormtable.UniqueIndex).Get(ctx, &controller,
		did,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this controllerTable) DeleteRange(ctx context.Context, from, to ControllerIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this controllerTable) doNotImplement() {}

var _ ControllerTable = controllerTable{}

func NewControllerTable(db ormtable.Schema) (ControllerTable, error) {
	table := db.GetTable(&Controller{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Controller{}).ProtoReflect().Descriptor().FullName()))
	}
	return controllerTable{table.(ormtable.AutoIncrementTable)}, nil
}

type VerificationTable interface {
	Insert(ctx context.Context, verification *Verification) error
	Update(ctx context.Context, verification *Verification) error
	Save(ctx context.Context, verification *Verification) error
	Delete(ctx context.Context, verification *Verification) error
	Has(ctx context.Context, did string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, did string) (*Verification, error)
	HasByIssuerSubject(ctx context.Context, issuer string, subject string) (found bool, err error)
	// GetByIssuerSubject returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByIssuerSubject(ctx context.Context, issuer string, subject string) (*Verification, error)
	HasByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (found bool, err error)
	// GetByControllerDidMethodIssuer returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (*Verification, error)
	HasByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (found bool, err error)
	// GetByVerificationTypeSubjectIssuer returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (*Verification, error)
	List(ctx context.Context, prefixKey VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error)
	ListRange(ctx context.Context, from, to VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error)
	DeleteBy(ctx context.Context, prefixKey VerificationIndexKey) error
	DeleteRange(ctx context.Context, from, to VerificationIndexKey) error

	doNotImplement()
}

type VerificationIterator struct {
	ormtable.Iterator
}

func (i VerificationIterator) Value() (*Verification, error) {
	var verification Verification
	err := i.UnmarshalMessage(&verification)
	return &verification, err
}

type VerificationIndexKey interface {
	id() uint32
	values() []interface{}
	verificationIndexKey()
}

// primary key starting index..
type VerificationPrimaryKey = VerificationDidIndexKey

type VerificationDidIndexKey struct {
	vs []interface{}
}

func (x VerificationDidIndexKey) id() uint32            { return 0 }
func (x VerificationDidIndexKey) values() []interface{} { return x.vs }
func (x VerificationDidIndexKey) verificationIndexKey() {}

func (this VerificationDidIndexKey) WithDid(did string) VerificationDidIndexKey {
	this.vs = []interface{}{did}
	return this
}

type VerificationIssuerSubjectIndexKey struct {
	vs []interface{}
}

func (x VerificationIssuerSubjectIndexKey) id() uint32            { return 1 }
func (x VerificationIssuerSubjectIndexKey) values() []interface{} { return x.vs }
func (x VerificationIssuerSubjectIndexKey) verificationIndexKey() {}

func (this VerificationIssuerSubjectIndexKey) WithIssuer(issuer string) VerificationIssuerSubjectIndexKey {
	this.vs = []interface{}{issuer}
	return this
}

func (this VerificationIssuerSubjectIndexKey) WithIssuerSubject(issuer string, subject string) VerificationIssuerSubjectIndexKey {
	this.vs = []interface{}{issuer, subject}
	return this
}

type VerificationControllerDidMethodIssuerIndexKey struct {
	vs []interface{}
}

func (x VerificationControllerDidMethodIssuerIndexKey) id() uint32            { return 2 }
func (x VerificationControllerDidMethodIssuerIndexKey) values() []interface{} { return x.vs }
func (x VerificationControllerDidMethodIssuerIndexKey) verificationIndexKey() {}

func (this VerificationControllerDidMethodIssuerIndexKey) WithController(controller string) VerificationControllerDidMethodIssuerIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this VerificationControllerDidMethodIssuerIndexKey) WithControllerDidMethod(controller string, did_method string) VerificationControllerDidMethodIssuerIndexKey {
	this.vs = []interface{}{controller, did_method}
	return this
}

func (this VerificationControllerDidMethodIssuerIndexKey) WithControllerDidMethodIssuer(controller string, did_method string, issuer string) VerificationControllerDidMethodIssuerIndexKey {
	this.vs = []interface{}{controller, did_method, issuer}
	return this
}

type VerificationVerificationTypeSubjectIssuerIndexKey struct {
	vs []interface{}
}

func (x VerificationVerificationTypeSubjectIssuerIndexKey) id() uint32            { return 3 }
func (x VerificationVerificationTypeSubjectIssuerIndexKey) values() []interface{} { return x.vs }
func (x VerificationVerificationTypeSubjectIssuerIndexKey) verificationIndexKey() {}

func (this VerificationVerificationTypeSubjectIssuerIndexKey) WithVerificationType(verification_type string) VerificationVerificationTypeSubjectIssuerIndexKey {
	this.vs = []interface{}{verification_type}
	return this
}

func (this VerificationVerificationTypeSubjectIssuerIndexKey) WithVerificationTypeSubject(verification_type string, subject string) VerificationVerificationTypeSubjectIssuerIndexKey {
	this.vs = []interface{}{verification_type, subject}
	return this
}

func (this VerificationVerificationTypeSubjectIssuerIndexKey) WithVerificationTypeSubjectIssuer(verification_type string, subject string, issuer string) VerificationVerificationTypeSubjectIssuerIndexKey {
	this.vs = []interface{}{verification_type, subject, issuer}
	return this
}

type verificationTable struct {
	table ormtable.Table
}

func (this verificationTable) Insert(ctx context.Context, verification *Verification) error {
	return this.table.Insert(ctx, verification)
}

func (this verificationTable) Update(ctx context.Context, verification *Verification) error {
	return this.table.Update(ctx, verification)
}

func (this verificationTable) Save(ctx context.Context, verification *Verification) error {
	return this.table.Save(ctx, verification)
}

func (this verificationTable) Delete(ctx context.Context, verification *Verification) error {
	return this.table.Delete(ctx, verification)
}

func (this verificationTable) Has(ctx context.Context, did string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, did)
}

func (this verificationTable) Get(ctx context.Context, did string) (*Verification, error) {
	var verification Verification
	found, err := this.table.PrimaryKey().Get(ctx, &verification, did)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) HasByIssuerSubject(ctx context.Context, issuer string, subject string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		issuer,
		subject,
	)
}

func (this verificationTable) GetByIssuerSubject(ctx context.Context, issuer string, subject string) (*Verification, error) {
	var verification Verification
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &verification,
		issuer,
		subject,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) HasByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		controller,
		did_method,
		issuer,
	)
}

func (this verificationTable) GetByControllerDidMethodIssuer(ctx context.Context, controller string, did_method string, issuer string) (*Verification, error) {
	var verification Verification
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &verification,
		controller,
		did_method,
		issuer,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) HasByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		verification_type,
		subject,
		issuer,
	)
}

func (this verificationTable) GetByVerificationTypeSubjectIssuer(ctx context.Context, verification_type string, subject string, issuer string) (*Verification, error) {
	var verification Verification
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &verification,
		verification_type,
		subject,
		issuer,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &verification, nil
}

func (this verificationTable) List(ctx context.Context, prefixKey VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return VerificationIterator{it}, err
}

func (this verificationTable) ListRange(ctx context.Context, from, to VerificationIndexKey, opts ...ormlist.Option) (VerificationIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return VerificationIterator{it}, err
}

func (this verificationTable) DeleteBy(ctx context.Context, prefixKey VerificationIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this verificationTable) DeleteRange(ctx context.Context, from, to VerificationIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this verificationTable) doNotImplement() {}

var _ VerificationTable = verificationTable{}

func NewVerificationTable(db ormtable.Schema) (VerificationTable, error) {
	table := db.GetTable(&Verification{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Verification{}).ProtoReflect().Descriptor().FullName()))
	}
	return verificationTable{table}, nil
}

type StateStore interface {
	AuthenticationTable() AuthenticationTable
	ControllerTable() ControllerTable
	VerificationTable() VerificationTable

	doNotImplement()
}

type stateStore struct {
	authentication AuthenticationTable
	controller     ControllerTable
	verification   VerificationTable
}

func (x stateStore) AuthenticationTable() AuthenticationTable {
	return x.authentication
}

func (x stateStore) ControllerTable() ControllerTable {
	return x.controller
}

func (x stateStore) VerificationTable() VerificationTable {
	return x.verification
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	authenticationTable, err := NewAuthenticationTable(db)
	if err != nil {
		return nil, err
	}

	controllerTable, err := NewControllerTable(db)
	if err != nil {
		return nil, err
	}

	verificationTable, err := NewVerificationTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		authenticationTable,
		controllerTable,
		verificationTable,
	}, nil
}
