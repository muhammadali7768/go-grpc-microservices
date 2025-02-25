package main

import "context"

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	accounts, err := r.server.accountClient.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id *string) ([]*Product, error) {
	accounts, err := r.server.accountClient.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
