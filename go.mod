module github.com/prolifera/kit

go 1.21.5

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/google/uuid v1.3.0
	github.com/redis/go-redis/v9 v9.1.0
	github.com/robfig/cron v1.2.0
	github.com/rs/zerolog v1.27.0
	github.com/tristan-club/kit v0.3.9
	github.com/trustwallet/go-primitives v0.1.23
	golang.org/x/crypto v0.9.0
	golang.org/x/time v0.5.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

//replace (
//	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 => github.com/OvyFlash/telegram-bot-api master
//)

replace github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 => github.com/mgintoki/telegram-bot-api/v5 v5.6.0
