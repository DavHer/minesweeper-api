package main

/*
Create game
curl http://localhost:8080/minesweeper \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"rows": 5, "cols": 5, "mines": 5}'

Get game
curl http://localhost:8080/minesweeper/e8a3f0be-2915-44f0-a920-fd8f71ae19e0


Start game
curl http://localhost:8080/minesweeper/e8a3f0be-2915-44f0-a920-fd8f71ae19e0/start \
    --include \
    --header "Content-Type: application/json" \
    --request "POST"

Reveal cell
curl http://localhost:8080/minesweeper/e8a3f0be-2915-44f0-a920-fd8f71ae19e0/reveal \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"row": 0, "col": 0}'

Flag cell
curl http://localhost:8080/minesweeper/e8a3f0be-2915-44f0-a920-fd8f71ae19e0/flag \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"row": 0, "col": 1}'
*/

func main() {
	StartApi()
}
