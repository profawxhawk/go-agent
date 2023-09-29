module github.com/newrelic/go-agent/v3/integrations/nrpgx5

go 1.19

require (
	github.com/egon12/pgsnap v0.0.0-20221022154027-2847f0124ed8
	github.com/jackc/pgx/v5 v5.0.3
	github.com/newrelic/go-agent/v3 v3.26.0
	github.com/stretchr/testify v1.8.0
)


replace github.com/newrelic/go-agent/v3 => ../..
