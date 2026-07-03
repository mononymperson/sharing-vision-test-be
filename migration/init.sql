CREATE TABLE IF NOT EXISTS posts (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(200)  NOT NULL,
    content     TEXT          NOT NULL,
    category    VARCHAR(100)  NOT NULL,
    created_date DATETIME     NOT NULL,
    updated_date DATETIME     NOT NULL,
    status      VARCHAR(100)  NOT NULL DEFAULT 'draft'
);
