CREATE TABLE IF NOT EXISTS Users (
                                     id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                     login VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL
    );

CREATE TABLE IF NOT EXISTS Catalogs (
                                        user_id BIGINT UNSIGNED NOT NULL,
                                        product_id BIGINT UNSIGNED NOT NULL,
                                        count INT NOT NULL,
                                        FOREIGN KEY (user_id) REFERENCES Users(id)
    );
