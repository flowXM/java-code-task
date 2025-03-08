# Тестовое задание компании JavaCode на позицию Golang разработчика

## Запуск сервисов (Docker Compose)
```bash
docker compose up -d
```

## Запуск тестов
```bash
go test -v ./...
```

## API
Используемый порт `3333`

### GET Wallet
`127.0.0.1:3333/api/v1/wallets/{id}`

### POST Wallet
`127.0.0.1:3333/api/v1/wallet`

Тело запроса:
```json
{
    "walletId": "09e711a7-7812-497e-9661-7aea014a8a56",
    "operationType": "DEPOSIT",
    "amount": 500.50
}
```
