# money-tracking-app

This app can track your income and expenses.

To run this project:
- Clone
- Run the migration with goose `cd database/schema` and then `goose postgres {your_postgres_url} up`
- Run the project with seeder `go run main.go -seeder true`

Endpoints:
1. /api/categories
2. /api/users
3. /api/wallets
4. /api/transactions

See the code for more details or ask me directly.