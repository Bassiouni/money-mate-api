set -xe

docker run --rm -v ./db/migrations:/migrations --network money-mate_backend migrate/migrate $@