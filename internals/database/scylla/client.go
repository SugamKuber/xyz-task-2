package scylla

import (
	"xyz-task-2/internals/models"

	"github.com/gocql/gocql"
)

type Client struct {
	session *gocql.Session
}

type Config struct {
	Hosts    []string
	Keyspace string
}

func NewClient(cfg Config) (*Client, error) {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &Client{session: session}, nil
}

func (c *Client) GetTopErrors(userID string, limit int) ([]models.Error, error) {
	var errors []models.Error
	iter := c.session.Query(`
		SELECT category, subcategory, count
		FROM error_counts
		WHERE user_id = ?
		ORDER BY count DESC
		LIMIT ?
	`, userID, limit).Iter()

	var category, subcategory string
	var count int
	for iter.Scan(&category, &subcategory, &count) {
		errors = append(errors, models.Error{
			Category:    category,
			Subcategory: subcategory,
			Count:       count,
		})
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return errors, nil
}

func (c *Client) Close() {
	c.session.Close()
}
