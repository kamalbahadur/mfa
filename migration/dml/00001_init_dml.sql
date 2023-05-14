-- +goose Up
DELETE FROM "user" WHERE id = 1;
INSERT INTO "user" (id, first_name, last_name) VALUES (1, 'Kamal', 'Giri');

-- +goose Down