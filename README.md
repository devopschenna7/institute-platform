


-------------using docker compose commands-----

Force a completely fresh build:

docker compose build --no-cache

Start everything
docker compose up

run in backeground
docker compose up -d

Frontend  → http://localhost:3000
Backend   → http://localhost:8080
Postgres  → localhost:5432


This stops the containers but does not remove them:

docker compose stop

Start them again:

docker compose start

Stop and remove containers
docker compose down


Stop everything and delete database data

 This will delete your PostgreSQL data.

docker compose down -v


. Enter the PostgreSQL database
docker exec -it institute-postgres psql -U postgres -d institute

Inside PostgreSQL:

\dt

List tables.

\l

List databases.

\d

Show table information.

Exit:

\q


----------------------------------
docker build -t frontend:1.0 .

docker run -d -p 8080:80 --name frontend frontend:1.0

docker save -o frontend.tar frontend:1.0


docker build -t backend:1.0 .

docker run -d -p 8080:8080 --name backend backend:1.0

docker save -o backend.tar backend:1.0

--------------------------------

chmod +x run-local.sh

----------------------------

./run-local.sh

------------------

cd frontend

npm install

npm run dev


backend

 go run ./cmd/api


--------------------------
in api.js change the below

const API_URL = "http://localhost:8080/api/v1";


Connect to the database:

docker exec -it postgres psql -U postgres -d myapp

List databases:

\l

List tables:

\dt

Exit:

\q