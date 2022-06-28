package entity

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// Token represents a row from 'public.tokens'.
type Token struct {
	Name                sql.NullString  `json:"name"`                  // name
	Symbol              sql.NullString  `json:"symbol"`                // symbol
	TotalSupply         sql.NullFloat64 `json:"total_supply"`          // total_supply
	Decimals            sql.NullFloat64 `json:"decimals"`              // decimals
	Type                string          `json:"type"`                  // type
	Cataloged           sql.NullBool    `json:"cataloged"`             // cataloged
	ContractAddressHash []byte          `json:"contract_address_hash"` // contract_address_hash
	InsertedAt          time.Time       `json:"inserted_at"`           // inserted_at
	UpdatedAt           time.Time       `json:"updated_at"`            // updated_at
	HolderCount         sql.NullInt64   `json:"holder_count"`          // holder_count
	SkipMetadata        sql.NullBool    `json:"skip_metadata"`         // skip_metadata
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Token exists in the database.
func (t *Token) Exists() bool {
	return t._exists
}

// Deleted returns true when the Token has been marked for deletion from
// the database.
func (t *Token) Deleted() bool {
	return t._deleted
}

// Insert inserts the Token to the database.
func (t *Token) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.tokens (` +
		`name, symbol, total_supply, decimals, type, cataloged, contract_address_hash, inserted_at, updated_at, holder_count, skip_metadata` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)`
	// run
	logf(sqlstr, t.Name, t.Symbol, t.TotalSupply, t.Decimals, t.Type, t.Cataloged, t.ContractAddressHash, t.InsertedAt, t.UpdatedAt, t.HolderCount, t.SkipMetadata)
	if _, err := db.ExecContext(ctx, sqlstr, t.Name, t.Symbol, t.TotalSupply, t.Decimals, t.Type, t.Cataloged, t.ContractAddressHash, t.InsertedAt, t.UpdatedAt, t.HolderCount, t.SkipMetadata); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Update updates a Token in the database.
func (t *Token) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.tokens SET ` +
		`name = $1, symbol = $2, total_supply = $3, decimals = $4, type = $5, cataloged = $6, inserted_at = $7, updated_at = $8, holder_count = $9, skip_metadata = $10 ` +
		`WHERE contract_address_hash = $11`
	// run
	logf(sqlstr, t.Name, t.Symbol, t.TotalSupply, t.Decimals, t.Type, t.Cataloged, t.InsertedAt, t.UpdatedAt, t.HolderCount, t.SkipMetadata, t.ContractAddressHash)
	if _, err := db.ExecContext(ctx, sqlstr, t.Name, t.Symbol, t.TotalSupply, t.Decimals, t.Type, t.Cataloged, t.InsertedAt, t.UpdatedAt, t.HolderCount, t.SkipMetadata, t.ContractAddressHash); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Token to the database.
func (t *Token) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Token.
func (t *Token) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.tokens (` +
		`name, symbol, total_supply, decimals, type, cataloged, contract_address_hash, inserted_at, updated_at, holder_count, skip_metadata` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)` +
		` ON CONFLICT (contract_address_hash) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, symbol = EXCLUDED.symbol, total_supply = EXCLUDED.total_supply, decimals = EXCLUDED.decimals, type = EXCLUDED.type, cataloged = EXCLUDED.cataloged, inserted_at = EXCLUDED.inserted_at, updated_at = EXCLUDED.updated_at, holder_count = EXCLUDED.holder_count, skip_metadata = EXCLUDED.skip_metadata `
	// run
	logf(sqlstr, t.Name, t.Symbol, t.TotalSupply, t.Decimals, t.Type, t.Cataloged, t.ContractAddressHash, t.InsertedAt, t.UpdatedAt, t.HolderCount, t.SkipMetadata)
	if _, err := db.ExecContext(ctx, sqlstr, t.Name, t.Symbol, t.TotalSupply, t.Decimals, t.Type, t.Cataloged, t.ContractAddressHash, t.InsertedAt, t.UpdatedAt, t.HolderCount, t.SkipMetadata); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Token from the database.
func (t *Token) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.tokens ` +
		`WHERE contract_address_hash = $1`
	// run
	logf(sqlstr, t.ContractAddressHash)
	if _, err := db.ExecContext(ctx, sqlstr, t.ContractAddressHash); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TokenByContractAddressHash retrieves a row from 'public.tokens' as a Token.
//
// Generated from index 'tokens_pkey'.
func TokenByContractAddressHash(ctx context.Context, db DB, contractAddressHash []byte) (*Token, error) {
	// query
	const sqlstr = `SELECT ` +
		`name, symbol, total_supply, decimals, type, cataloged, contract_address_hash, inserted_at, updated_at, holder_count, skip_metadata ` +
		`FROM public.tokens ` +
		`WHERE contract_address_hash = $1`
	// run
	logf(sqlstr, contractAddressHash)
	t := Token{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, contractAddressHash).Scan(&t.Name, &t.Symbol, &t.TotalSupply, &t.Decimals, &t.Type, &t.Cataloged, &t.ContractAddressHash, &t.InsertedAt, &t.UpdatedAt, &t.HolderCount, &t.SkipMetadata); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// TokensBySymbol retrieves a row from 'public.tokens' as a Token.
//
// Generated from index 'tokens_symbol_index'.
func TokensBySymbol(ctx context.Context, db DB, symbol sql.NullString) ([]*Token, error) {
	// query
	const sqlstr = `SELECT ` +
		`name, symbol, total_supply, decimals, type, cataloged, contract_address_hash, inserted_at, updated_at, holder_count, skip_metadata ` +
		`FROM public.tokens ` +
		`WHERE symbol = $1`
	// run
	logf(sqlstr, symbol)
	rows, err := db.QueryContext(ctx, sqlstr, symbol)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Token
	for rows.Next() {
		t := Token{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&t.Name, &t.Symbol, &t.TotalSupply, &t.Decimals, &t.Type, &t.Cataloged, &t.ContractAddressHash, &t.InsertedAt, &t.UpdatedAt, &t.HolderCount, &t.SkipMetadata); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// TokensByType retrieves a row from 'public.tokens' as a Token.
//
// Generated from index 'tokens_type_index'.
func TokensByType(ctx context.Context, db DB, typ string) ([]*Token, error) {
	// query
	const sqlstr = `SELECT ` +
		`name, symbol, total_supply, decimals, type, cataloged, contract_address_hash, inserted_at, updated_at, holder_count, skip_metadata ` +
		`FROM public.tokens ` +
		`WHERE type = $1`
	// run
	logf(sqlstr, typ)
	rows, err := db.QueryContext(ctx, sqlstr, typ)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Token
	for rows.Next() {
		t := Token{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&t.Name, &t.Symbol, &t.TotalSupply, &t.Decimals, &t.Type, &t.Cataloged, &t.ContractAddressHash, &t.InsertedAt, &t.UpdatedAt, &t.HolderCount, &t.SkipMetadata); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Address returns the Address associated with the Token's (ContractAddressHash).
//
// Generated from foreign key 'tokens_contract_address_hash_fkey'.
func (t *Token) Address(ctx context.Context, db DB) (*Address, error) {
	return AddressByHash(ctx, db, t.ContractAddressHash)
}
