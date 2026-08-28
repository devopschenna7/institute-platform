```bash
#!/bin/bash

set -e

echo "======================================"
echo " Starting PostgreSQL"
echo "======================================"

# Start existing PostgreSQL container if it exists,
# otherwise create a new one.
if docker ps -a --format '{{.Names}}' | grep -q '^postgres$'; then
    echo "PostgreSQL container already exists."

    if [ "$(docker inspect -f '{{.State.Running}}' postgres)" = "false" ]; then
        echo "Starting PostgreSQL..."
        docker start postgres
    else
        echo "PostgreSQL is already running."
    fi
else
    echo "Creating PostgreSQL container..."

    docker run -d \
        --name postgres \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_PASSWORD=postgres \
        -e POSTGRES_DB=myapp \
        -p 5432:5432 \
        postgres:17
fi

echo ""
echo "======================================"
echo " PostgreSQL Status"
echo "======================================"

docker ps --filter "name=postgres"

echo ""
echo "Waiting for PostgreSQL..."

until docker exec postgres pg_isready \
    -U postgres \
    -d myapp > /dev/null 2>&1
do
    sleep 1
done

echo "PostgreSQL is ready!"

echo ""
echo "======================================"
echo " Setting Database URL"
echo "======================================"

export DATABASE_URL="postgres://postgres:postgres@localhost:5432/myapp?sslmode=disable"

echo "DATABASE_URL=$DATABASE_URL"

echo ""
echo "======================================"
echo " Starting Go Backend"
echo "======================================"

cd backend

echo "Downloading Go dependencies..."
go mod download

echo ""
echo "Starting backend on port 8080..."
echo ""

go run ./cmd/api
```
