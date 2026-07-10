PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS user_roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_active INTEGER NOT NULL DEFAULT 1, -- 1 for true, 0 for false
    role_id INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY(role_id) REFERENCES user_roles(id)
);

CREATE TABLE IF NOT EXISTS request_types (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    is_relevant INTEGER NOT NULL DEFAULT 1,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT NULL,
    closed_at TEXT DEFAULT NULL,
    dispatcher_id INTEGER NOT NULL,
    responder_id INTEGER NOT NULL,
    request_type_id INTEGER NOT NULL,

    FOREIGN KEY (dispatcher_id) REFERENCES users(id),
    FOREIGN KEY (responder_id) REFERENCES users(id),
    FOREIGN KEY (request_type_id) REFERENCES request_types(id)
);
