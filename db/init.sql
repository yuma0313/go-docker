CREATE DATABASE IF NOT EXISTS test;

USE test;

CREATE TABLE IF NOT EXISTS todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    detail VARCHAR(255) NOT NULL
);

INSERT INTO todos (title, detail) VALUES
    ('買い物', '食料品を買う'),
    ('プロジェクト', '新しいプロジェクトの計画を立てる'),
    ('ジョギング', '公園でジョギングする');
