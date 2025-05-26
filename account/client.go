package account

import (
	"context"
	"log"

	"github.com/muhammadali7768/go-grpc-microservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("account client NewClient URL", url)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
		return nil, err
	}

	service := pb.NewAccountServiceClient(conn)
	return &Client{conn, service}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostAccount(ctx context.Context, name string) (*Account, error) {
	a, err := c.service.PostAccount(ctx, &pb.PostAccountRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return &Account{ID: a.Account.Id, Name: a.Account.Name}, nil
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	log.Println("account client BEFORE:", id)
	r, err := c.service.GetAccount(ctx, &pb.GetAccountRequest{Id: id})
	log.Println("account client:", r)
	if err != nil {
		log.Println("Error in account client:GetAccount")
		return nil, err
	}
	return &Account{ID: r.Account.Id, Name: r.Account.Name}, nil
}
func (c *Client) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	r, err := c.service.GetAccounts(ctx, &pb.GetAccountsRequest{Skip: skip, Take: take})

	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	for _, a := range r.Accounts {
		accounts = append(accounts, Account{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return accounts, nil
}
