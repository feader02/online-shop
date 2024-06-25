CREATE TABLE IF NOT EXISTS Product (
                           id INT AUTO_INCREMENT PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           price DECIMAL(10, 2) NOT NULL,
                           description TEXT,
                           height INT,
                           width INT,
                           depth INT,
                           photo VARCHAR(255),
                           type TEXT
);