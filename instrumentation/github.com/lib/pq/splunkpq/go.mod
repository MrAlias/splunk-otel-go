module github.com/signalfx/splunk-otel-go/instrumentation/github.com/lib/pq/splunkpq

go 1.16

require (
	github.com/lib/pq v1.10.4
	github.com/signalfx/splunk-otel-go/instrumentation/database/sql/splunksql v0.7.0
	github.com/stretchr/testify v1.7.1
)

replace (
	github.com/signalfx/splunk-otel-go => ../../../../..
	github.com/signalfx/splunk-otel-go/instrumentation/database/sql/splunksql => ../../../../database/sql/splunksql
)
