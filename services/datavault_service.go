package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/TFMV/VaultDeploy/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func tableExists(dbpool *pgxpool.Pool, schema, tableName string) (bool, error) {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = '%s'
		AND    table_name   = '%s'
	);`, schema, tableName)
	err := dbpool.QueryRow(context.Background(), query).Scan(&exists)
	return exists, err
}

func CreateOrUpdateHub(dbpool *pgxpool.Pool, schema string, hub models.Hub) error {
	exists, err := tableExists(dbpool, schema, hub.TableName)
	if err != nil {
		return err
	}

	if !exists {
		createTableQuery := fmt.Sprintf(`CREATE TABLE %s.%s (%s)`, schema, hub.TableName, formatColumns(hub.Columns))
		_, err := dbpool.Exec(context.Background(), createTableQuery)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateOrUpdateLink(dbpool *pgxpool.Pool, schema string, link models.Link) error {
	exists, err := tableExists(dbpool, schema, link.TableName)
	if err != nil {
		return err
	}

	if !exists {
		createTableQuery := fmt.Sprintf(`CREATE TABLE %s.%s (%s)`, schema, link.TableName, formatColumns(link.Columns))
		_, err := dbpool.Exec(context.Background(), createTableQuery)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateOrUpdateSatellite(dbpool *pgxpool.Pool, schema string, satellite models.Satellite) error {
	exists, err := tableExists(dbpool, schema, satellite.TableName)
	if err != nil {
		return err
	}

	if !exists {
		createTableQuery := fmt.Sprintf(`CREATE TABLE %s.%s (%s)`, schema, satellite.TableName, formatColumns(satellite.Columns))
		_, err := dbpool.Exec(context.Background(), createTableQuery)
		if err != nil {
			return err
		}
	}
	return nil
}

func formatColumns(columns []string) string {
	return strings.Join(columns, ", ")
}

func InsertRecord(dbpool *pgxpool.Pool, schema string, tableName string, columns []string, record []string) error {
	values := make([]string, len(record))
	for i, value := range record {
		values[i] = fmt.Sprintf("'%s'", value)
	}

	query := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)",
		schema, tableName, strings.Join(columns, ", "), strings.Join(values, ", "))

	_, err := dbpool.Exec(context.Background(), query)
	return err
}
