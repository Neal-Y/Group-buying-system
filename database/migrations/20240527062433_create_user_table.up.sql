-- migrate create -ext sql -dir database/migrations create_user_table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    line_id VARCHAR(255) UNIQUE NOT NULL,
    line_token VARCHAR(255) UNIQUE,
    phone VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_member BOOLEAN DEFAULT FALSE
);
