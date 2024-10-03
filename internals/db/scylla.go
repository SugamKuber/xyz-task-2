package db

import (
	"fmt"
	"xyz-task-2/internals/models"

	"github.com/gocql/gocql"
)

type ScyllaClient struct {
	session *gocql.Session
}

type ScyllaConfig struct {
	Hosts    []string
	Keyspace string
}

func NewScyllaClient(cfg ScyllaConfig) (*ScyllaClient, error) {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &ScyllaClient{session: session}, nil
}

func (c *ScyllaClient) Execute(query string, values ...interface{}) error {
	return c.session.Query(query, values...).Exec()
}

func (c *ScyllaClient) Query(query string, values ...interface{}) *gocql.Iter {
	return c.session.Query(query, values...).Iter()
}

func (sc *ScyllaClient) GetSession() *gocql.Session {
	return sc.session
}

func (c *ScyllaClient) GetTopErrors(userID string, limit int) ([]models.Error, error) {
	var errors []models.Error
	query := `
		SELECT error_category, error_subcategory, frequency
		FROM error_frequencies
		WHERE user_id = ?
		ORDER BY frequency DESC
		LIMIT ?
	`
	iter := c.session.Query(query, userID, limit).Iter()
	fmt.Println(iter)
	fmt.Println("::::")

	var category, subcategory string
	var frequency int
	for iter.Scan(&category, &subcategory, &frequency) {
		errors = append(errors, models.Error{
			Category:    category,
			Subcategory: subcategory,
			Frequency:   frequency,
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return errors, nil
}

func (c *ScyllaClient) Close() {
	c.session.Close()
}
