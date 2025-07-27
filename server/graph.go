package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Graph struct {
	Driver neo4j.DriverWithContext
}

func NewGraph() (*Graph, error) {
	d, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		return nil, err
	}
	return &Graph{Driver: d}, nil
}

func (g *Graph) VerifyConnectivity(ctx context.Context) error {
	return g.Driver.VerifyConnectivity(ctx)
}

func (g *Graph) AddFriend(ctx context.Context, from, to string) error {
	session := g.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func() {
		err := session.Close(ctx)
		if err != nil {
			log.Printf("failed to close session: %v\n", err)
		}
	}()
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MERGE (a:User {name: $from})
			MERGE (b:User {name: $to})
			MERGE (a)-[:FRIEND]->(b)`
		_, err := tx.Run(ctx, query, map[string]any{"from": from, "to": to})
		return nil, err
	})
	return err
}

func (g *Graph) GetMutualFriends(ctx context.Context, user1, user2 string) ([]string, error) {
	session := g.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer func() {
		err := session.Close(ctx)
		if err != nil {
			log.Printf("failed to close session: %v\n", err)
		}
	}()

	const mutualFriendKey = "mutualFriend"
	res, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := fmt.Sprintf(`MATCH (a:User {name: $user1})-[:FRIEND]->(f:User)<-[:FRIEND]-(b:User {name: $user2})
			RETURN f.name AS %s`, mutualFriendKey)
		records, err := tx.Run(ctx, query, map[string]any{"user1": user1, "user2": user2})
		if err != nil {
			return nil, err
		}

		var mututalFriends []string
		for records.Next(ctx) {
			friend, _ := records.Record().Get(mutualFriendKey)
			mututalFriends = append(mututalFriends, friend.(string))
		}
		return mututalFriends, nil
	})
	return res.([]string), err
}
