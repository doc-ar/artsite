---------------------------------------------------
-- Art Queries
---------------------------------------------------

-- name: GetArt :one
SELECT * FROM art
WHERE id = $1 LIMIT 1;

-- name: ListArt :many
SELECT a.*, s.name AS series, s.id AS series_id FROM art a
JOIN series s ON s.id = a.series_id
ORDER BY a.created_at;

-- name: CreateArt :one
INSERT INTO art (title, width, height, category, medium, price, image_url, description, series_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9)
RETURNING *;

-- name: UpdateArt :one
UPDATE art
SET title = $2, width = $3, height = $4, category = $5, medium = $6, price = $7, image_url = $8, description = $9, series_id = $10
WHERE id = $1
RETURNING *;

-- name: IncrementViewCount :exec
UPDATE art
SET view_count = view_count + 1
WHERE id = $1;

-- name: DeleteArt :one
DELETE FROM art
WHERE id = $1
RETURNING *;

---------------------------------------------------
-- Portfolio Queries
---------------------------------------------------

-- name: ListPortfolio :many
SELECT * FROM art a
JOIN portfolio p ON p.art_id = a.id
ORDER BY p.added_at;

-- name: PostPortfolio :one
INSERT INTO portfolio (art_id)
VALUES ($1)
RETURNING *;

-- name: DeletePortfolio :one
DELETE FROM portfolio
WHERE art_id = $1
RETURNING *;

---------------------------------------------------
-- Series Queries
---------------------------------------------------

-- name: GetSeries :one
SELECT * FROM series
WHERE id = $1 LIMIT 1;

-- name: GetSeriesFromName :one
SELECT * FROM series
WHERE LOWER(name) = LOWER($1) LIMIT 1;

-- name: ListSeries :many
SELECT * FROM series
ORDER BY created_at;

-- name: ListSeriesDetails :many
SELECT a.* FROM art a
JOIN series s ON s.id = a.series_id
WHERE LOWER(s.name) = LOWER($1)
ORDER BY a.created_at;

-- name: PostSeries :one
INSERT INTO series (name, description, cover_img)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateSeries :one
UPDATE series
SET name = $2, description = $3, cover_img = $4
WHERE id = $1
RETURNING *;

-- name: DeleteSeries :one
DELETE FROM series
WHERE id = $1
RETURNING *;

---------------------------------------------------
-- Admin Queries
---------------------------------------------------

-- name: GetAdmin :one
SELECT * FROM admins
WHERE username = $1 LIMIT 1;

-- name: UpdateAdmin :one
UPDATE admins
SET username = $2, password_hash = $3
WHERE id = $1
RETURNING *;
