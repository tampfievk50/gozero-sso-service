-- Seed "Authenticated" role (default role for all new users)
INSERT INTO role (name, description, created_at, updated_at, is_deleted)
VALUES ('Authenticated', 'Default role for all authenticated users', NOW(), NOW(), false)
ON CONFLICT DO NOTHING;

-- Seed account permissions (only if not already present)
INSERT INTO permission (name, description, path, action, is_deleted, created_at, updated_at)
SELECT 'View Account', 'View own account profile', '/v1/account/:id', 'GET', false, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM permission WHERE path = '/v1/account/:id' AND action = 'GET');

INSERT INTO permission (name, description, path, action, is_deleted, created_at, updated_at)
SELECT 'Update Account', 'Update own account profile', '/v1/account/:id', 'PUT', false, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM permission WHERE path = '/v1/account/:id' AND action = 'PUT');

-- Add Casbin policies for "Authenticated" role to access account endpoints
-- Policy format: ptype=p, v0=role_id, v1=domain_id, v2=path, v3=action, v4=effect
INSERT INTO role_permission (ptype, v0, v1, v2, v3, v4)
SELECT 'p', CAST(r.id AS VARCHAR), '0', '/v1/account/:id', 'GET', 'allow'
FROM role r WHERE r.name = 'Authenticated'
AND NOT EXISTS (SELECT 1 FROM role_permission WHERE v0 = CAST(r.id AS VARCHAR) AND v2 = '/v1/account/:id' AND v3 = 'GET');

INSERT INTO role_permission (ptype, v0, v1, v2, v3, v4)
SELECT 'p', CAST(r.id AS VARCHAR), '0', '/v1/account/:id', 'PUT', 'allow'
FROM role r WHERE r.name = 'Authenticated'
AND NOT EXISTS (SELECT 1 FROM role_permission WHERE v0 = CAST(r.id AS VARCHAR) AND v2 = '/v1/account/:id' AND v3 = 'PUT');
