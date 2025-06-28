DROP TABLE IF EXISTS `user_settings`;

CREATE TABLE
    IF NOT EXISTS `user_settings` (
        user_id VARCHAR(36),
        daily_reminder TIME,
        reminder_hour TIME
    );

DROP TABLE IF EXISTS `categories`;

CREATE TABLE
    IF NOT EXISTS `categories` (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id VARCHAR(36),
        name VARCHAR(255) NOT NULL
    );

DROP TABLE IF EXISTS `budgets`;

CREATE TABLE
    IF NOT EXISTS `budgets` (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id VARCHAR(36) NOT NULL,
        category_id INT NOT NULL,
        amount INT NOT NULL,
        current_expense INT NOT NULL DEFAULT 0,
        time_period ENUM ('WEEKLY', 'MONTHLY') DEFAULT 'WEEKLY'
    );

DROP TABLE IF EXISTS `moneys`;

CREATE TABLE
    IF NOT EXISTS `moneys` (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id VARCHAR(36) NOT NULL,
        amount INT NOT NULL,
        category_id INT NOT NULL,
        type ENUM ('INCOME', 'EXPENSE') NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

DROP TABLE IF EXISTS `daily_assignments`;

CREATE TABLE
    IF NOT EXISTS `daily_assignments` (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id VARCHAR(36) NOT NULL,
        name VARCHAR(255) NOT NULL,
        is_done TINYINT (1) DEFAULT 0
    );

DROP TABLE IF EXISTS `assignment_categories`;

CREATE TABLE
    IF NOT EXISTS `assignment_categories` (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        name VARCHAR(255) NOT NULL
    );

DROP TABLE IF EXISTS `assignments`;

CREATE TABLE
    IF NOT EXISTS `assignments` (
        id INT PRIMARY KEY NOT NULL,
        user_id VARCHAR(36) NOT NULL,
        name VARCHAR(36) NOT NULL,
        deadline DATETIME NOT NULL,
        status ENUM ('ONGOING', 'DONE') DEFAULT 'ONGOING',
        priority ENUM ('HIGH', 'MEDIUM', 'LOW') DEFAULT 'LOW',
        category_id INT NOT NULL
    );

DROP TABLE IF EXISTS `history_budgets`;

CREATE TABLE
    IF NOT EXISTS `history_budgets` (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        budget_id INT NOT NULL,
        category_id INT NOT NULL,
        spent INT NOT NULL,
        planned INT NOT NULL,
        type ENUM ('WEEKLY', 'MONTHLY'),
        period_start DATE NOT NULL,
        period_end DATE NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );