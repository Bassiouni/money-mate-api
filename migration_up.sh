set -xe

source .env

./migrate.sh -source file:///migrations -database postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@postgresql:5432/$POSTGRES_DB?sslmode=disable up
