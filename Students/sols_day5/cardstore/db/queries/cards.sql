-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1 LIMIT 1;

-- name: ListCards :many
SELECT * FROM cards
WHERE 
date = sqlc.arg(date1) OR
date = sqlc.arg(date2) OR
date = sqlc.arg(date3) OR
date = sqlc.arg(date4) OR
date = sqlc.arg(date5) OR
date = sqlc.arg(date6) OR
date = sqlc.arg(date7)
ORDER BY date;

-- name: CreateCard :one
INSERT INTO cards (
	time, 
	date, 
	title, 
	client,
	"user" -- user is a reserved keyword, must use ""
) VALUES (
    $1, 
		$2, 
		$3, 
		$4, 
		$5
) RETURNING *;

-- name: DeleteCard :exec
DELETE FROM cards
WHERE id = $1;

-- name: UpdateCards :one
UPDATE cards
SET 
    time = $2,
    date = $3,
    title = $4,
    client = $5,
		"user" = $6
WHERE id = $1
RETURNING *;
