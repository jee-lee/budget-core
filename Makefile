-include .env

TEST_APP_ENV := TEST
TEST_DB_NAME := budget_test
export TEST_DATABASE_URL := postgres://${LOCAL_DB_USER}:${LOCAL_DB_PASSWORD}@localhost:${LOCAL_DB_PORT}/${TEST_DB_NAME}?sslmode=disable

migrate-test-db:
	DATABASE_URL=$(TEST_DATABASE_URL) dbmate up

rollback-test-db:
	DATABASE_URL=$(TEST_DATABASE_URL) dbmate rollback

drop-test-db:
	DATABASE_URL=$(TEST_DATABASE_URL) dbmate drop

migrate-dev-db:
	dbmate up

rollback-dev-db:
	dbmate rollback

drop-dev-db:
	dbmate drop

run-tests:
	APP_ENV=$(TEST_APP_ENV) go test -v -p 1 ./... -coverpkg=./... -coverprofile ./coverage.out

view-test-coverage:
	APP_ENV=$(TEST_APP_ENV) go tool cover -html=coverage.out
