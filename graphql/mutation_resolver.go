package main

import "context"

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) createAccount(ctx context.Context, ai AccountInput) (*Account, error) {
	account, err := server.accountClient.CreateAccount(ctx, ai)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *mutationResolver) createProduct(ctx context.Context, ai ProductInput) (*Product, error) {
	product, err := server.productClient.CreateProduct(ctx, ai)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (r *mutationResolver) createOrder(ctx context.Context, ai OrderInput) (*Order, error) {
	account, err := server.accountClient.CreateAccount(ctx, ai)
	if err != nil {
		return nil, err
	}
	return account, nil
}
