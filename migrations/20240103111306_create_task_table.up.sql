-- TASKS
CREATE TABLE IF NOT EXISTS tasks(
    id serial NOT NULL,
    name VARCHAR (255) NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL,
    PRIMARY KEY(id)
);

COMMENT ON COLUMN tasks.name IS '任務名稱';

COMMENT ON COLUMN tasks.status IS '任務狀態';