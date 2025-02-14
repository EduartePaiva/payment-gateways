##@ Testing

test:
	godotenv -f .env go test ./... -race -v

##@Live reload
ifeq ($(OS),Windows_NT)
    AIR_CONFIG=.\\.air.windows.conf
else
    AIR_CONFIG=.air.conf
endif

air:
	air -c $(AIR_CONFIG)
