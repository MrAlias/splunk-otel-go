module github.com/signalfx/splunk-otel-go/instrumentation/github.com/jmoiron/sqlx/splunksqlx

go 1.16

require (
	github.com/jmoiron/sqlx v1.3.4
	github.com/signalfx/splunk-otel-go/instrumentation/database/sql/splunksql v0.7.0
	github.com/stretchr/testify v1.7.1
)

replace (
	github.com/signalfx/splunk-otel-go => ../../../../../
	github.com/signalfx/splunk-otel-go/instrumentation/database/sql/splunksql => ../../../../database/sql/splunksql
)
