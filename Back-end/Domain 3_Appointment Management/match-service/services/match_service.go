package services

import (
	"context"
	"log"

	"github.com/Elein69/match-service/models"
	"github.com/Elein69/match-service/neo4j"
	neo4jdriver "github.com/neo4j/neo4j-go-driver/v5/neo4j" // 👈 Renombrado para evitar conflicto
)

type MatchResult struct {
	ID           string
	Name         string
	Specialty    string
	Availability string
}

func FindBestVeterinarian(req models.MatchRequest) (*MatchResult, error) {
	ctx := context.Background()

	session := neo4j.Driver.NewSession(ctx, neo4jdriver.SessionConfig{
		AccessMode: neo4jdriver.AccessModeRead,
	})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4jdriver.ManagedTransaction) (any, error) {
		query := `
			MATCH (v:Veterinarian)-[:OFFERS]->(s:Service {name: $service})
			MATCH (v)-[:AVAILABLE_ON]->(t:TimeSlot {day: $day, time: $time})
			RETURN v.id AS id, v.name AS name, s.name AS specialty, t.day + " " + t.time AS availability
			LIMIT 1
		`
		params := map[string]any{
			"service": req.Reason,
			"day":     req.PreferredDay,
			"time":    req.PreferredTime,
		}

		record, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		if record.Next(ctx) {
			r := record.Record()
			return &MatchResult{
				ID:           r.Values[0].(string),
				Name:         r.Values[1].(string),
				Specialty:    r.Values[2].(string),
				Availability: r.Values[3].(string),
			}, nil
		}

		return nil, nil
	})

	if err != nil {
		log.Println("❌ Neo4j query error:", err)
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result.(*MatchResult), nil
}
