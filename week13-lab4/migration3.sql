-- 4. Permissions Table
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    resource VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_permissions_name ON permissions(name);
CREATE INDEX idx_permissions_resource ON permissions(resource);

-- 5. Role-Permission Assignment
CREATE TABLE role_permissions (
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    granted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id, permission_id)
);

CREATE INDEX idx_role_perms_role ON role_permissions(role_id);
CREATE INDEX idx_role_perms_perm ON role_permissions(permission_id);

-- Seed Permissions
INSERT INTO permissions (name, description, resource, action) VALUES
-- Books permissions
('books:read', 'Can view books', 'books', 'read'),
('books:create', 'Can create new books', 'books', 'create'),
('books:update', 'Can update books', 'books', 'update'),
('books:delete', 'Can delete books', 'books', 'delete'),
('books:publish', 'Can publish books', 'books', 'publish'),

-- Users permissions
('users:read', 'Can view users', 'users', 'read'),
('users:create', 'Can create users', 'users', 'create'),
('users:update', 'Can update users', 'users', 'update'),
('users:delete', 'Can delete users', 'users', 'delete'),

-- Roles permissions
('roles:read', 'Can view roles', 'roles', 'read'),
('roles:assign', 'Can assign roles to users', 'roles', 'assign'),
('roles:create', 'Can create new roles', 'roles', 'create'),
('roles:delete', 'Can delete roles', 'roles', 'delete'),

-- Reports permissions
('reports:financial', 'Can view financial reports', 'reports', 'financial'),
('reports:analytics', 'Can view analytics', 'reports', 'analytics');

-- Assign Permissions to Roles

-- Admin: ทุก permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT
    (SELECT id FROM roles WHERE name = 'admin'),
    id
FROM permissions;

-- Editor: books permissions (ยกเว้น delete) + read users
INSERT INTO role_permissions (role_id, permission_id)
SELECT
    (SELECT id FROM roles WHERE name = 'editor'),
    id
FROM permissions
WHERE name IN (
    'books:read', 'books:create', 'books:update', 'books:publish',
    'users:read'
);

-- Viewer: read-only
INSERT INTO role_permissions (role_id, permission_id)
SELECT
    (SELECT id FROM roles WHERE name = 'viewer'),
    id
FROM permissions
WHERE action = 'read';

-- User books:read
INSERT INTO role_permissions (role_id, permission_id)
SELECT
    (SELECT id FROM roles WHERE name = 'user'),
    id
FROM permissions
WHERE name = 'books:read';