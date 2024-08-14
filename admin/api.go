package admin

import (
	"context"
	"fmt"
	"github.com/kazhuravlev/just"
	"github.com/segmentio/kafka-go"
)

type ISchema interface {
	KvPairs() map[string]string
}

type CreateTopicReq struct {
	Name              string
	Partitions        uint
	ReplicationFactor uint
	Schema            ISchema
}

// CreateTopic will create a topic in cluster. Note: it will not return an error when topic already exists.
func (c *Client) CreateTopic(ctx context.Context, req CreateTopicReq) error {
	topicCfg := req.Schema.KvPairs()
	cfgEntries := make([]kafka.ConfigEntry, 0, len(topicCfg))
	for key, value := range topicCfg {
		cfgEntries = append(cfgEntries, kafka.ConfigEntry{
			ConfigName:  key,
			ConfigValue: value,
		})
	}

	conn, err := c.getConnection(ctx)
	if err != nil {
		return fmt.Errorf("get connection to controller: %w", err)
	}

	kfkReq := kafka.TopicConfig{ //nolint:exhaustruct
		Topic:             req.Name,
		NumPartitions:     int(req.Partitions),
		ReplicationFactor: int(req.ReplicationFactor),
		ConfigEntries:     cfgEntries,
	}
	if err := conn.CreateTopics(kfkReq); err != nil {
		return fmt.Errorf("create topic: %w", err)
	}

	return nil
}

// ListTopics returns list of topics, available in cluster. Actually, this method will read all partitions, not topics.
// Be careful.
func (c *Client) ListTopics(ctx context.Context) ([]string, error) {
	conn, err := c.getConnection(ctx)
	if err != nil {
		return nil, fmt.Errorf("get connection to controller: %w", err)
	}

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return nil, fmt.Errorf("read partitions: %w", err)
	}

	topicNames := make(map[string]struct{})
	for i := range partitions {
		topicNames[partitions[i].Topic] = struct{}{}
	}

	return just.MapGetKeys(topicNames), nil
}
