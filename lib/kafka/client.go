package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

// ListGroup 获取所有 group
func ListGroup(client sarama.ClusterAdmin) []string {

	groups, err := client.ListConsumerGroups()
	if err != nil {
		fmt.Println(err.Error())
		return []string{}
	}

	result := make([]string, 0, len(groups))
	for g := range groups {
		result = append(result, g)
	}
	return result
}

// Groups 获取所有 group
func Groups() []string {
	client, err := NewCAClient()
	defer client.Close()

	if err != nil {
		return []string{}
	}

	return ListGroup(client)
}

type GroupTopicOffset struct {
	Topic        string `json:"topic"`
	Partition    int32  `json:"partition"`
	Offset       int64  `json:"offset"`
	LogEndOffset int64  `json:"log_end_offset"`
	Lag          int64  `json:"lag"`
}

// ListGroupOffset group offset 信息
func ListGroupOffset(client sarama.ClusterAdmin, group string) []GroupTopicOffset {
	// var topicPartitions map[string][]int32
	rsp, err := client.ListConsumerGroupOffsets(group, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var result []GroupTopicOffset
	for topic, v := range rsp.Blocks {
		for p, i := range v {
			result = append(result, GroupTopicOffset{
				topic,
				p,
				i.Offset,
				0,
				0,
			})
		}
	}
	return result
}

// GroupOffsets 返回 group offset 信息
func GroupOffsets(group string) []GroupTopicOffset {
	adminClient, err := NewCAClient()
	if err != nil {
		return nil
	}
	defer adminClient.Close()

	gOffsets := ListGroupOffset(adminClient, group)

	client, err := NewClient()
	if err != nil {
		return nil
	}
	defer client.Close()

	for idx, o := range gOffsets {
		lef := GetPartitionOffset(client, o.Topic, o.Partition)
		gOffsets[idx].LogEndOffset = lef
		gOffsets[idx].Lag = lef - o.Offset

	}
	return gOffsets
}

// GetPartitionOffset 获取 topic 每个分区的 offset
func GetPartitionOffset(client sarama.Client, topic string, partition int32) int64 {
	offset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
	if err != nil {
		return 0
	}
	return offset
}

// NewCAClient 创建客户端
func NewCAClient() (sarama.ClusterAdmin, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	client, err := sarama.NewClusterAdmin(viper.GetStringSlice("kafka.brokers"), config)
	return client, err
}

// NewClient 创建客户端
func NewClient() (sarama.Client, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	client, err := sarama.NewClient(viper.GetStringSlice("kafka.brokers"), config)
	return client, err
}
