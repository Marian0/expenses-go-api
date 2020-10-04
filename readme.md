# Expenses API

## Project description
Expenses API for learning purposes written in golang using gorm, gin gonic, firebase auth and jwt.

## Installation
- clone this repo
- copy environemnt variables `cp .env.default .env`
- update env variables accordingly `nano .env`
- run `docker-compose up`
- run `go run main.go`

## How to create a new migration
- Create a new file at `database/migrations/YYYYMMDD000X_sample_summary_description.go`
- Note that the 4 ending chars represent an unique name and the description to understand what's the migration created for
- Use the following template :

```

package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var MYYYYMMDD000X = gormigrate.Migration{
	ID: "YYYYMMDD000X",
	Migrate: func(tx *gorm.DB) error {
    //add here schema changes
	},
	Rollback: func(tx *gorm.DB) error {
    //add here reverse schema changes
	},
}

```
- Add this migragion object to the `database/migrate.go` array
