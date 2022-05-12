migrate_staging:
	APP_ENV=staging go run ./cmd/migration/main.go

seed_staging:
	APP_ENV=staging go run ./cmd/seed/main.go

.PHONY: migrate_staging seed_staging
