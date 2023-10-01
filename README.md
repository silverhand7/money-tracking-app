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

# Run in Docker 
1. Build image `docker compose build`.
2. Create containers `docker compose create`.
3. Run everything `docker compose up -d`.
4. Go to the project bash/sh `docker compose exec app sh`
5. Run migration `cd database/schema` and run `goose postgres postgres://{db_user}:{db_password}@database:5432/money-tracking-app up`.
6. Go back to the root folder and populate the database with seeder `go run main.go -seeder=true` and then press ctr/cmd+c to exit (please ignore the error address already in use, the goal is just to populate the database with seeder).
7. Exit and the project will run in localhost:8080
