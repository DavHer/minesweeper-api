package main

/*
Create game
curl http://localhost:8080/minesweeper \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"rows": 5, "cols": 5, "mines": 5}'

Get game
curl http://localhost:8080/minesweeper/00e572f8-603e-4f40-82dd-0083f6717bbc


Start game
curl http://localhost:8080/minesweeper/00e572f8-603e-4f40-82dd-0083f6717bbc/start \
    --include \
    --header "Content-Type: application/json" \
    --request "POST"

Reveal cell
curl http://localhost:8080/minesweeper/00e572f8-603e-4f40-82dd-0083f6717bbc/reveal \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"row": 0, "col": 0}'

Flag cell
curl http://localhost:8080/minesweeper/00e572f8-603e-4f40-82dd-0083f6717bbc/flag \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"row": 0, "col": 1}'
*/

func main() {
	StartApi()
}
