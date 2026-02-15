-- ===============================
-- Migration: Change all ID columns from BIGINT to UUID
-- ===============================

-- 1. Drop junction tables (they have FK constraints referencing columns we need to alter)
DROP TABLE IF EXISTS user_role;
DROP TABLE IF EXISTS user_resource;

-- 2. Alter id columns from BIGINT to UUID with gen_random_uuid() default

-- resource
ALTER TABLE "resource"
    ALTER COLUMN id DROP DEFAULT,
    ALTER COLUMN id SET DATA TYPE UUID USING lpad(to_hex(id), 32, '0')::uuid,
    ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- permission
ALTER TABLE "permission"
    ALTER COLUMN id DROP DEFAULT,
    ALTER COLUMN id SET DATA TYPE UUID USING lpad(to_hex(id), 32, '0')::uuid,
    ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- role
ALTER TABLE "role"
    ALTER COLUMN id DROP DEFAULT,
    ALTER COLUMN id SET DATA TYPE UUID USING lpad(to_hex(id), 32, '0')::uuid,
    ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- user
ALTER TABLE "user"
    ALTER COLUMN id DROP DEFAULT,
    ALTER COLUMN id SET DATA TYPE UUID USING lpad(to_hex(id), 32, '0')::uuid,
    ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- role_permission: DO NOT convert id to UUID
-- The role_permission table is managed by casbin gorm-adapter which
-- hardcodes ID as uint/bigint in its LoadPolicy method.

-- 3. Alter created_by, updated_by, deleted_by columns from BIGINT to UUID

ALTER TABLE "resource"
    ALTER COLUMN created_by SET DATA TYPE UUID USING CASE WHEN created_by IS NOT NULL THEN lpad(to_hex(created_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN updated_by SET DATA TYPE UUID USING CASE WHEN updated_by IS NOT NULL THEN lpad(to_hex(updated_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN deleted_by SET DATA TYPE UUID USING CASE WHEN deleted_by IS NOT NULL THEN lpad(to_hex(deleted_by), 32, '0')::uuid ELSE NULL END;

ALTER TABLE "permission"
    ALTER COLUMN created_by SET DATA TYPE UUID USING CASE WHEN created_by IS NOT NULL THEN lpad(to_hex(created_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN updated_by SET DATA TYPE UUID USING CASE WHEN updated_by IS NOT NULL THEN lpad(to_hex(updated_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN deleted_by SET DATA TYPE UUID USING CASE WHEN deleted_by IS NOT NULL THEN lpad(to_hex(deleted_by), 32, '0')::uuid ELSE NULL END;

ALTER TABLE "role"
    ALTER COLUMN created_by SET DATA TYPE UUID USING CASE WHEN created_by IS NOT NULL THEN lpad(to_hex(created_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN updated_by SET DATA TYPE UUID USING CASE WHEN updated_by IS NOT NULL THEN lpad(to_hex(updated_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN deleted_by SET DATA TYPE UUID USING CASE WHEN deleted_by IS NOT NULL THEN lpad(to_hex(deleted_by), 32, '0')::uuid ELSE NULL END;

ALTER TABLE "user"
    ALTER COLUMN created_by SET DATA TYPE UUID USING CASE WHEN created_by IS NOT NULL THEN lpad(to_hex(created_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN updated_by SET DATA TYPE UUID USING CASE WHEN updated_by IS NOT NULL THEN lpad(to_hex(updated_by), 32, '0')::uuid ELSE NULL END,
    ALTER COLUMN deleted_by SET DATA TYPE UUID USING CASE WHEN deleted_by IS NOT NULL THEN lpad(to_hex(deleted_by), 32, '0')::uuid ELSE NULL END;

-- 4. Recreate junction tables with UUID columns

CREATE TABLE IF NOT EXISTS user_resource (
    resource_id UUID NOT NULL,
    user_id     UUID NOT NULL,
    PRIMARY KEY (resource_id, user_id)
);

CREATE TABLE IF NOT EXISTS user_role (
    user_id     UUID NOT NULL,
    role_id     UUID NOT NULL,
    resource_id UUID NOT NULL,
    PRIMARY KEY (role_id, user_id, resource_id)
);

-- 5. Drop old sequences
DROP SEQUENCE IF EXISTS resource_seq;
DROP SEQUENCE IF EXISTS permission_seq;
DROP SEQUENCE IF EXISTS role_seq;
DROP SEQUENCE IF EXISTS user_seq;
