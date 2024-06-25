-- +goose Up
-- +goose StatementBegin
ALTER TABLE comments
ADD COLUMN has_parent BOOLEAN;
UPDATE comments
SET has_parent = CASE
    WHEN parent_id = 0 THEN FALSE
    ELSE TRUE
END;
ALTER TABLE comments
ALTER COLUMN has_parent SET NOT NULL,
ALTER COLUMN parent_id SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE comments
DROP COLUMN has_parent,
ALTER COLUMN parent_id DROP NOT NULL;
-- +goose StatementEnd
