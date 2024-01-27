package modulesapi

import (
	"encoding/json"
	"net/http"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/highway/middleware"
)

// BankHandler is a handler for the bank module
type BankHandler struct{}

// GetAllBalances returns all balances for an address
func (h BankHandler) GetAllBalances(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).AllBalances(r.Context(), &bankv1beta1.QueryAllBalancesRequest{
		Address: address,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetBalance returns a balance for an address and denom
func (h BankHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).Balance(r.Context(), &bankv1beta1.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetTotalSupply returns the total supply
func (h BankHandler) GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).TotalSupply(r.Context(), &bankv1beta1.QueryTotalSupplyRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSupplyOf returns the supply of a denom
func (h BankHandler) GetSupplyOf(w http.ResponseWriter, r *http.Request) {
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).SupplyOf(r.Context(), &bankv1beta1.QuerySupplyOfRequest{Denom: denom})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSpendableBalances returns the spendable balances for an address
func (h BankHandler) GetSpendableBalances(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).SpendableBalances(r.Context(), &bankv1beta1.QuerySpendableBalancesRequest{Address: address})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSpendableBalancesByDenom returns the spendable balances for an address and denom
func (h BankHandler) GetSpendableBalancesByDenom(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).SpendableBalanceByDenom(r.Context(), &bankv1beta1.QuerySpendableBalanceByDenomRequest{Address: address, Denom: denom})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}
