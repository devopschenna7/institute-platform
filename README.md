

----------------------------------
docker build -t frontend:1.0 .

docker run -d -p 8080:80 --name frontend frontend:1.0

docker save -o frontend.tar frontend:1.0


docker build -t backend:1.0 .

docker run -d -p 8080:8080 --name backend backend:1.0

docker save -o backend.tar backend:1.0

--------------------------------

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
