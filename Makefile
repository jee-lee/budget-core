-include .env

TEST_APP_ENV := TEST
TEST_DB_NAME := budget_test
export TEST_DATABASE_URL := postgres://${LOCAL_DB_USER}:${LOCAL_DB_PASSWORD}@localhost:${LOCAL_DB_PORT}/${TEST_DB_NAME}?sslmode=disable

# Database commands
.PHONY: test_db rollback_test_db drop_test_db
test_db:
	DATABASE_URL=$(TEST_DATABASE_URL) dbmate up

rollback_test_db:
	DATABASE_URL=$(TEST_DATABASE_URL) dbmate rollback

drop_test_db:
	DATABASE_URL=$(TEST_DATABASE_URL) dbmate drop

.PHONY: dev_db rollback_dev_db drop_dev_db
dev_db:
	dbmate up

rollback_dev_db:
	dbmate rollback

drop_dev_db:
	dbmate drop

# Test commands
.PHONY: test coverage_report
test coverage.out:
	APP_ENV=$(TEST_APP_ENV) go test -v -p 1 ./... -coverpkg=./... -coverprofile ./coverage.out

coverage_report: coverage.out
	APP_ENV=$(TEST_APP_ENV) go tool cover -html=coverage.out
