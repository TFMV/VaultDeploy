package services

import (
	"context"
	"fmt"

	"github.com/TFMV/VaultDeploy/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateOrUpdateHub(dbpool *pgxpool.Pool, schema string, hub models.Hub) error {
	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (%s)`, schema, hub.TableName, formatColumns(hub.Columns))
	_, err := dbpool.Exec(context.Background(), createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func CreateOrUpdateLink(dbpool *pgxpool.Pool, schema string, link models.Link) error {
	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (%s)`, schema, link.TableName, formatColumns(link.Columns))
	_, err := dbpool.Exec(context.Background(), createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func CreateOrUpdateSatellite(dbpool *pgxpool.Pool, schema string, satellite models.Satellite) error {
	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (%s)`, schema, satellite.TableName, formatColumns(satellite.Columns))
	_, err := dbpool.Exec(context.Background(), createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func formatColumns(columns []string) string {
	return fmt.Sprintf("%s", join(columns, ", "))
}

func join(elements []string, separator string) string {
	result := ""
	for i, element := range elements {
		if i > 0 {
			result += separator
		}
		result += element
	}
	return result
}
