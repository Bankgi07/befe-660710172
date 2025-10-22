-- 2. Roles Table

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    is_system BOOLEAN DEFAULT false,  -- role ที่ลบไม่ได้ (admin, user)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_roles_name ON roles(name);

-- 3. User-Role Assignment

CREATE TABLE user_roles (
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    assigned_by INTEGER REFERENCES users(id),  -- ใครเป็นคนมอบหมาย
    PRIMARY KEY (user_id, role_id)
);

CREATE INDEX idx_user_roles_user ON user_roles(user_id);
CREATE INDEX idx_user_roles_role ON user_roles(role_id);

-- Seed Roles
INSERT INTO roles (name, description, is_system) VALUES
('admin', 'Administrator with full system access', true),
('editor', 'Can create and edit content', false),
('viewer', 'Read-only access', false),
('user', 'Default role for new users', true);

-- Assign Roles to Users
-- admin user >> admin role
INSERT INTO user_roles (user_id, role_id)
SELECT
    (SELECT id FROM users WHERE username = 'admin'),
    (SELECT id FROM roles WHERE name = 'admin');

-- editor user >> editor role
INSERT INTO user_roles (user_id, role_id)
SELECT
    (SELECT id FROM users WHERE username = 'poohkan'),
    (SELECT id FROM roles WHERE name = 'editor');

-- regular user >> user role
INSERT INTO user_roles (user_id, role_id)
SELECT
    (SELECT id FROM users WHERE username = 'nuttachot'),
    (SELECT id FROM roles WHERE name = 'user');