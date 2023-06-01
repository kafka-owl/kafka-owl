// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

//go:build integration

package console

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kfake"
	"github.com/twmb/franz-go/pkg/kmsg"
	"go.uber.org/zap"

	"github.com/redpanda-data/console/backend/pkg/config"
	"github.com/redpanda-data/console/backend/pkg/kafka"
	"github.com/redpanda-data/console/backend/pkg/kafka/mocks"
	"github.com/redpanda-data/console/backend/pkg/testutil"
)

func (s *ConsoleIntegrationTestSuite) TestListMessages() {
	ctx := context.Background()
	t := s.T()
	assert := assert.New(t)
	require := require.New(t)

	logCfg := zap.NewDevelopmentConfig()
	logCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	log, err := logCfg.Build()
	require.NoError(err)

	testTopicName := testutil.TopicNameForTest("list_messages")

	testutil.CreateTestData(t, ctx, s.kafkaClient, s.kafkaAdminClient, testTopicName)

	t.Run("empty topic", func(t *testing.T) {
		_, err := s.kafkaAdminClient.CreateTopic(ctx, 1, 1, nil,
			testutil.TopicNameForTest("list_messages_empty_topic"))
		require.NoError(err)

		defer func() {
			s.kafkaAdminClient.DeleteTopics(ctx, testutil.TopicNameForTest("list_messages_empty_topic"))
		}()

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnComplete(gomock.Any(), false)

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:    testutil.TopicNameForTest("list_messages_empty_topic"),
			PartitionID:  -1,
			StartOffset:  -2,
			MessageCount: 100,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("all messages in a topic", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		var msg *kafka.TopicMessage
		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(gomock.AssignableToTypeOf(msg)).Times(20)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(20)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:    testTopicName,
			PartitionID:  -1,
			StartOffset:  -2,
			MessageCount: 100,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("single messages in a topic", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("10")).Times(1)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(1)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:    testTopicName,
			PartitionID:  -1,
			StartOffset:  10,
			MessageCount: 1,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("five messages in a topic", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("10")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("11")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("12")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("13")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("14")).Times(1)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(5)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:    testTopicName,
			PartitionID:  -1,
			StartOffset:  10,
			MessageCount: 5,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("time stamp in future get last record", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("19")).Times(1)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(1)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:      testTopicName,
			PartitionID:    -1,
			MessageCount:   5,
			StartTimestamp: time.Date(2010, time.November, 11, 13, 0, 0, 0, time.UTC).UnixMilli(),
			StartOffset:    StartOffsetTimestamp,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("time stamp in middle get 5 records", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("11")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("12")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("13")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("14")).Times(1)
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("15")).Times(1)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(5)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:      testTopicName,
			PartitionID:    -1,
			MessageCount:   5,
			StartTimestamp: time.Date(2010, time.November, 10, 13, 10, 30, 0, time.UTC).UnixMilli(),
			StartOffset:    StartOffsetTimestamp,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("unknown topic", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		mockProgress.EXPECT().OnPhase("Get Partitions")

		svc := createNewTestService(t, log, t.Name(), s.testSeedBroker)

		input := ListMessageRequest{
			TopicName:    "console_list_messages_topic_test_unknown_topic",
			PartitionID:  -1,
			StartOffset:  -2,
			MessageCount: 100,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.Error(err)
		assert.Equal("failed to get partitions: UNKNOWN_TOPIC_OR_PARTITION: This server does not host this topic-partition.",
			err.Error())
	})

	// This is essentially the same test as "single messages in a topic" test
	// but using kfake package to demonstrate a full test of how to use the package
	// to verify functionality.
	t.Run("single messages in a topic fake", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		fakeCluster, err := kfake.NewCluster(kfake.NumBrokers(1))
		require.NoError(err)

		defer fakeCluster.Close()

		fakeClient, fakeAdminClient := testutil.CreateClients(t, fakeCluster.ListenAddrs())

		testutil.CreateTestData(t, ctx, fakeClient, fakeAdminClient, testTopicName)

		svc := createNewTestService(t, log, t.Name(), fakeCluster.ListenAddrs()[0])

		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("10")).Times(1)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(1)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		var fetchCalls int32
		mdCalls := atomic.Int32{}

		fakeCluster.Control(func(req kmsg.Request) (kmsg.Response, error, bool) {
			fakeCluster.KeepControl()

			switch v := req.(type) {
			case *kmsg.ApiVersionsRequest:
				return nil, nil, false
			case *kmsg.MetadataRequest:
				mdCalls.Add(1)

				assert.Len(v.Topics, 1)
				assert.Equal(testTopicName, *(v.Topics[0].Topic))

				return nil, nil, false
			case *kmsg.ListOffsetsRequest:
				loReq, ok := req.(*kmsg.ListOffsetsRequest)
				assert.True(ok, "request is not a list offset request: %+T", req)

				assert.Len(loReq.Topics, 1)
				assert.Equal(testTopicName, loReq.Topics[0].Topic)

				assert.Len(loReq.Topics[0].Partitions, 1)
				assert.Equal(int32(1), loReq.Topics[0].Partitions[0].MaxNumOffsets)

				return nil, nil, false
			case *kmsg.FetchRequest:
				fetchReq, ok := req.(*kmsg.FetchRequest)
				assert.True(ok, "request is not a list offset request: %+T", req)

				assert.Len(fetchReq.Topics, 1)
				assert.Equal(testTopicName, fetchReq.Topics[0].Topic)

				if atomic.LoadInt32(&fetchCalls) == 0 {
					atomic.StoreInt32(&fetchCalls, 1)

					assert.Len(fetchReq.Topics[0].Partitions, 1)
					assert.Equal(int64(10), fetchReq.Topics[0].Partitions[0].FetchOffset)
				} else if atomic.LoadInt32(&fetchCalls) == 1 {
					atomic.StoreInt32(&fetchCalls, 2)

					assert.Len(fetchReq.Topics[0].Partitions, 1)
					assert.Equal(int64(20), fetchReq.Topics[0].Partitions[0].FetchOffset)
				} else {
					assert.Fail("unexpected call to fake fetch request")
				}

				return nil, nil, false
			default:
				assert.Fail(fmt.Sprintf("unexpected call to fake kafka request %+T", v))

				return nil, nil, false
			}
		})

		input := ListMessageRequest{
			TopicName:    testTopicName,
			PartitionID:  -1,
			StartOffset:  10,
			MessageCount: 1,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})

	t.Run("metadata topic error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		fakeCluster, err := kfake.NewCluster(kfake.NumBrokers(1))
		require.NoError(err)

		defer fakeCluster.Close()

		fakeClient, fakeAdminClient := testutil.CreateClients(t, fakeCluster.ListenAddrs())

		testutil.CreateTestData(t, ctx, fakeClient, fakeAdminClient, testTopicName)

		svc := createNewTestService(t, log, t.Name(), fakeCluster.ListenAddrs()[0])

		mockProgress.EXPECT().OnPhase("Get Partitions")

		fakeCluster.Control(func(req kmsg.Request) (kmsg.Response, error, bool) {
			fakeCluster.KeepControl()

			switch v := req.(type) {
			case *kmsg.ApiVersionsRequest:
				return nil, nil, false
			case *kmsg.MetadataRequest:
				assert.Len(v.Topics, 1)
				assert.Equal(testTopicName, *(v.Topics[0].Topic))

				mdRes := v.ResponseKind().(*kmsg.MetadataResponse)
				mdRes.Topics = make([]kmsg.MetadataResponseTopic, 1)
				mdRes.Topics[0] = kmsg.NewMetadataResponseTopic()
				mdRes.Topics[0].Topic = kmsg.StringPtr(testTopicName)
				mdRes.Topics[0].Partitions = make([]kmsg.MetadataResponseTopicPartition, 1)
				mdRes.Topics[0].Partitions[0].Partition = 0
				mdRes.Topics[0].Partitions[0].Leader = 0
				mdRes.Topics[0].Partitions[0].LeaderEpoch = 1
				mdRes.Topics[0].Partitions[0].Replicas = make([]int32, 1)
				mdRes.Topics[0].Partitions[0].Replicas[0] = 0
				mdRes.Topics[0].Partitions[0].ISR = make([]int32, 1)
				mdRes.Topics[0].Partitions[0].ISR[0] = 0

				mdRes.Topics[0].ErrorCode = kerr.LeaderNotAvailable.Code

				return mdRes, nil, true

			default:
				assert.Fail(fmt.Sprintf("unexpected call to fake kafka request %+T", v))

				return nil, nil, false
			}
		})

		input := ListMessageRequest{
			TopicName:    testTopicName,
			PartitionID:  -1,
			StartOffset:  15,
			MessageCount: 1,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.Error(err)
		assert.Equal("failed to get partitions: LEADER_NOT_AVAILABLE: There is no leader for this topic-partition as we are in the middle of a leadership election.", err.Error())
	})

	t.Run("list offset error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockProgress := mocks.NewMockIListMessagesProgress(mockCtrl)

		fakeCluster, err := kfake.NewCluster(kfake.NumBrokers(1))
		require.NoError(err)

		defer fakeCluster.Close()

		fakeClient, fakeAdminClient := testutil.CreateClients(t, fakeCluster.ListenAddrs())

		testutil.CreateTestData(t, ctx, fakeClient, fakeAdminClient, testTopicName)

		svc := createNewTestService(t, log, t.Name(), fakeCluster.ListenAddrs()[0])

		var int64Type int64

		mockProgress.EXPECT().OnPhase("Get Partitions")
		mockProgress.EXPECT().OnPhase("Get Watermarks and calculate consuming requests")
		mockProgress.EXPECT().OnPhase("Consuming messages")
		mockProgress.EXPECT().OnMessage(testutil.MatchesOrder("16")).Times(1)
		mockProgress.EXPECT().OnMessageConsumed(gomock.AssignableToTypeOf(int64Type)).Times(1)
		mockProgress.EXPECT().OnComplete(gomock.AssignableToTypeOf(int64Type), false)

		fakeCluster.Control(func(req kmsg.Request) (kmsg.Response, error, bool) {
			fakeCluster.KeepControl()

			switch v := req.(type) {
			case *kmsg.ApiVersionsRequest:
				return nil, nil, false
			case *kmsg.MetadataRequest:
				assert.Len(v.Topics, 1)
				assert.Equal(testTopicName, *(v.Topics[0].Topic))

				return nil, nil, false

			case *kmsg.ListOffsetsRequest:
				loReq, ok := req.(*kmsg.ListOffsetsRequest)
				assert.True(ok, "request is not a list offset request: %+T", req)

				assert.Len(loReq.Topics, 1)
				assert.Equal(testTopicName, loReq.Topics[0].Topic)

				assert.Len(loReq.Topics[0].Partitions, 1)
				assert.Equal(int32(1), loReq.Topics[0].Partitions[0].MaxNumOffsets)

				loRes := v.ResponseKind().(*kmsg.ListOffsetsResponse)
				loRes.Topics = make([]kmsg.ListOffsetsResponseTopic, 1)
				loRes.Topics[0] = kmsg.NewListOffsetsResponseTopic()
				loRes.Topics[0].Topic = testTopicName
				loRes.Topics[0].Partitions = make([]kmsg.ListOffsetsResponseTopicPartition, 1)
				loRes.Topics[0].Partitions[0].Partition = 0
				loRes.Topics[0].Partitions[0].LeaderEpoch = 1
				loRes.Topics[0].Partitions[0].Timestamp = -1
				loRes.Topics[0].Partitions[0].Offset = 0

				loRes.Topics[0].Partitions[0].ErrorCode = kerr.NotLeaderForPartition.Code

				return loRes, nil, true

			default:
				assert.Fail(fmt.Sprintf("unexpected call to fake kafka request %+T", v))

				return nil, nil, false
			}
		})

		input := ListMessageRequest{
			TopicName:    testTopicName,
			PartitionID:  -1,
			StartOffset:  16,
			MessageCount: 1,
		}

		err = svc.ListMessages(ctx, input, mockProgress)
		assert.NoError(err)
	})
}

func createNewTestService(t *testing.T, log *zap.Logger,
	testName string, seedBrokers string,
) Servicer {
	metricName := testutil.MetricNameForTest(strings.ReplaceAll(testName, " ", ""))

	cfg := config.Config{}
	cfg.SetDefaults()
	cfg.MetricsNamespace = metricName
	cfg.Kafka.Brokers = []string{seedBrokers}

	svc, err := NewService(&cfg, log, nil, nil)
	assert.NoError(t, err)

	return svc
}
