// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SplunkAggregationQueueRatio:                 MetricConfig{Enabled: true},
					SplunkAuthzRolesStatuscode:                  MetricConfig{Enabled: true},
					SplunkBucketsSearchableStatus:               MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedBucketCount:        MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedBucketEventCount:   MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedBucketHotCount:     MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedBucketWarmCount:    MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedEventCount:         MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedRawSize:            MetricConfig{Enabled: true},
					SplunkDataIndexesExtendedTotalSize:          MetricConfig{Enabled: true},
					SplunkHealth:                                MetricConfig{Enabled: true},
					SplunkIndexerAvgRate:                        MetricConfig{Enabled: true},
					SplunkIndexerCPUTime:                        MetricConfig{Enabled: true},
					SplunkIndexerQueueRatio:                     MetricConfig{Enabled: true},
					SplunkIndexerRawWriteTime:                   MetricConfig{Enabled: true},
					SplunkIndexerThroughput:                     MetricConfig{Enabled: true},
					SplunkIndexesAvgSize:                        MetricConfig{Enabled: true},
					SplunkIndexesAvgUsage:                       MetricConfig{Enabled: true},
					SplunkIndexesBucketCount:                    MetricConfig{Enabled: true},
					SplunkIndexesMedianDataAge:                  MetricConfig{Enabled: true},
					SplunkIndexesSize:                           MetricConfig{Enabled: true},
					SplunkIoAvgIops:                             MetricConfig{Enabled: true},
					SplunkKvstoreBackupStatus:                   MetricConfig{Enabled: true},
					SplunkKvstoreReplicationStatus:              MetricConfig{Enabled: true},
					SplunkKvstoreStatus:                         MetricConfig{Enabled: true},
					SplunkLicenseIndexUsage:                     MetricConfig{Enabled: true},
					SplunkParseQueueRatio:                       MetricConfig{Enabled: true},
					SplunkPipelineSetCount:                      MetricConfig{Enabled: true},
					SplunkSchedulerAvgExecutionLatency:          MetricConfig{Enabled: true},
					SplunkSchedulerAvgRunTime:                   MetricConfig{Enabled: true},
					SplunkSchedulerCompletionRatio:              MetricConfig{Enabled: true},
					SplunkServerIntrospectionQueuesCurrent:      MetricConfig{Enabled: true},
					SplunkServerIntrospectionQueuesCurrentBytes: MetricConfig{Enabled: true},
					SplunkServerSearchartifactsAdhoc:            MetricConfig{Enabled: true},
					SplunkServerSearchartifactsCompleted:        MetricConfig{Enabled: true},
					SplunkServerSearchartifactsIncomplete:       MetricConfig{Enabled: true},
					SplunkServerSearchartifactsInvalid:          MetricConfig{Enabled: true},
					SplunkServerSearchartifactsJobCacheCount:    MetricConfig{Enabled: true},
					SplunkServerSearchartifactsJobCacheSize:     MetricConfig{Enabled: true},
					SplunkServerSearchartifactsSavedsearches:    MetricConfig{Enabled: true},
					SplunkServerSearchartifactsScheduled:        MetricConfig{Enabled: true},
					SplunkTypingQueueRatio:                      MetricConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					SplunkAggregationQueueRatio:                 MetricConfig{Enabled: false},
					SplunkAuthzRolesStatuscode:                  MetricConfig{Enabled: false},
					SplunkBucketsSearchableStatus:               MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedBucketCount:        MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedBucketEventCount:   MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedBucketHotCount:     MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedBucketWarmCount:    MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedEventCount:         MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedRawSize:            MetricConfig{Enabled: false},
					SplunkDataIndexesExtendedTotalSize:          MetricConfig{Enabled: false},
					SplunkHealth:                                MetricConfig{Enabled: false},
					SplunkIndexerAvgRate:                        MetricConfig{Enabled: false},
					SplunkIndexerCPUTime:                        MetricConfig{Enabled: false},
					SplunkIndexerQueueRatio:                     MetricConfig{Enabled: false},
					SplunkIndexerRawWriteTime:                   MetricConfig{Enabled: false},
					SplunkIndexerThroughput:                     MetricConfig{Enabled: false},
					SplunkIndexesAvgSize:                        MetricConfig{Enabled: false},
					SplunkIndexesAvgUsage:                       MetricConfig{Enabled: false},
					SplunkIndexesBucketCount:                    MetricConfig{Enabled: false},
					SplunkIndexesMedianDataAge:                  MetricConfig{Enabled: false},
					SplunkIndexesSize:                           MetricConfig{Enabled: false},
					SplunkIoAvgIops:                             MetricConfig{Enabled: false},
					SplunkKvstoreBackupStatus:                   MetricConfig{Enabled: false},
					SplunkKvstoreReplicationStatus:              MetricConfig{Enabled: false},
					SplunkKvstoreStatus:                         MetricConfig{Enabled: false},
					SplunkLicenseIndexUsage:                     MetricConfig{Enabled: false},
					SplunkParseQueueRatio:                       MetricConfig{Enabled: false},
					SplunkPipelineSetCount:                      MetricConfig{Enabled: false},
					SplunkSchedulerAvgExecutionLatency:          MetricConfig{Enabled: false},
					SplunkSchedulerAvgRunTime:                   MetricConfig{Enabled: false},
					SplunkSchedulerCompletionRatio:              MetricConfig{Enabled: false},
					SplunkServerIntrospectionQueuesCurrent:      MetricConfig{Enabled: false},
					SplunkServerIntrospectionQueuesCurrentBytes: MetricConfig{Enabled: false},
					SplunkServerSearchartifactsAdhoc:            MetricConfig{Enabled: false},
					SplunkServerSearchartifactsCompleted:        MetricConfig{Enabled: false},
					SplunkServerSearchartifactsIncomplete:       MetricConfig{Enabled: false},
					SplunkServerSearchartifactsInvalid:          MetricConfig{Enabled: false},
					SplunkServerSearchartifactsJobCacheCount:    MetricConfig{Enabled: false},
					SplunkServerSearchartifactsJobCacheSize:     MetricConfig{Enabled: false},
					SplunkServerSearchartifactsSavedsearches:    MetricConfig{Enabled: false},
					SplunkServerSearchartifactsScheduled:        MetricConfig{Enabled: false},
					SplunkTypingQueueRatio:                      MetricConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
