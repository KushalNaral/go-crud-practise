CREATE TABLE accounts (
    account_number VARCHAR(20) PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,               
    name VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) DEFAULT 0.00,    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
