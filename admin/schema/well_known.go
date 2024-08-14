package schema

import "github.com/kazhuravlev/just"

// knownOptions is a list of well-known kafka options.
// https://docs.confluent.io/platform/current/installation/configuration/topic-configs.html.
var knownOptions = []string{ //nolint:gochecknoglobals
	"cleanup.policy",
	"compression.type",
	"confluent.cluster.link.allow.legacy.message.format",
	"confluent.key.schema.validation",
	"confluent.key.subject.name.strategy",
	"confluent.placement.constraints",
	"confluent.tier.enable",
	"confluent.tier.local.hotset.bytes",
	"confluent.tier.local.hotset.ms",
	"confluent.value.schema.validation",
	"confluent.value.subject.name.strategy",
	"delete.retention.ms",
	"file.delete.delay.ms",
	"flush.messages",
	"flush.ms",
	"follower.replication.throttled.replicas",
	"index.interval.bytes",
	"leader.replication.throttled.replicas",
	"max.compaction.lag.ms",
	"max.message.bytes",
	"message.downconversion.enable",
	"message.format.version",
	"message.timestamp.difference.max.ms",
	"message.timestamp.type",
	"min.cleanable.dirty.ratio",
	"min.compaction.lag.ms",
	"min.insync.replicas",
	"preallocate",
	"retention.bytes",
	"retention.ms",
	"segment.bytes",
	"segment.index.bytes",
	"segment.jitter.ms",
	"segment.ms",
	"unclean.leader.election.enable",
}

// knownOptionsMap is just a useful wrapper.
var knownOptionsMap = just.Slice2Map(knownOptions) //nolint:gochecknoglobals
