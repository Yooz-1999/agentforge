CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(128) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    nickname VARCHAR(64) NOT NULL,
    status TINYINT NOT NULL DEFAULT 1,
    last_login_at DATETIME NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    UNIQUE KEY uk_users_email (email)
);

CREATE TABLE IF NOT EXISTS agents (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    name VARCHAR(128) NOT NULL,
    description VARCHAR(500) NOT NULL DEFAULT '',
    system_prompt TEXT NOT NULL,
    model_provider VARCHAR(64) NOT NULL,
    model_name VARCHAR(128) NOT NULL,
    temperature DECIMAL(3,2) NOT NULL DEFAULT 1.00,
    max_tokens INT NULL,
    status TINYINT NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    KEY idx_agents_user_id (user_id)
);

CREATE TABLE IF NOT EXISTS conversations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    agent_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL DEFAULT '',
    status TINYINT NOT NULL DEFAULT 1,
    last_message_at DATETIME NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    KEY idx_conversations_user_id (user_id),
    KEY idx_conversations_agent_id (agent_id)
);

CREATE TABLE IF NOT EXISTS messages (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    conversation_id BIGINT NOT NULL,
    user_id BIGINT NULL,
    agent_id BIGINT NOT NULL,
    role VARCHAR(32) NOT NULL,
    content LONGTEXT NOT NULL,
    model_provider VARCHAR(64) NULL,
    model_name VARCHAR(128) NULL,
    prompt_tokens INT NULL,
    completion_tokens INT NULL,
    status TINYINT NOT NULL DEFAULT 1,
    error_message VARCHAR(500) NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    KEY idx_messages_conversation_id (conversation_id),
    KEY idx_messages_agent_id (agent_id)
);

CREATE TABLE IF NOT EXISTS user_api_keys (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    provider VARCHAR(64) NOT NULL,
    api_key_ciphertext TEXT NOT NULL,
    api_base_url VARCHAR(255) NULL,
    status TINYINT NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    KEY idx_user_api_keys_user_id (user_id)
);
