-- ===============================
-- Sequence Definitions
-- ===============================

CREATE SEQUENCE IF NOT EXISTS permission_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1
    NO CYCLE;

CREATE SEQUENCE IF NOT EXISTS role_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1
    NO CYCLE;

CREATE SEQUENCE IF NOT EXISTS user_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1
    NO CYCLE;


-- ===============================
-- Permission Table
-- ===============================

CREATE TABLE IF NOT EXISTS "permission" (
    id          BIGINT NOT NULL DEFAULT nextval('permission_seq'::regclass),
    name        VARCHAR(255),
    description VARCHAR(255),
    path        VARCHAR(255),
    action      VARCHAR(255),
    created_at  TIMESTAMP(6),
    updated_at  TIMESTAMP(6),
    deleted_at  TIMESTAMP(6),
    created_by  BIGINT,
    updated_by  BIGINT,
    deleted_by  BIGINT,
    is_deleted    BOOLEAN DEFAULT false,

    PRIMARY KEY (id)
);


-- ===============================
-- Role Table
-- ===============================

CREATE TABLE IF NOT EXISTS "role" (
    id          BIGINT NOT NULL DEFAULT nextval('role_seq'::regclass),
    name        VARCHAR(255),
    description VARCHAR(255),
    created_at  TIMESTAMP(6),
    updated_at  TIMESTAMP(6),
    deleted_at  TIMESTAMP(6),
    created_by  BIGINT,
    updated_by  BIGINT,
    deleted_by  BIGINT,
    is_deleted    BOOLEAN DEFAULT false,
    PRIMARY KEY (id)
);


-- ===============================
-- User Table
-- ===============================

CREATE TABLE IF NOT EXISTS "user" (
    id            BIGINT NOT NULL DEFAULT nextval('user_seq'::regclass),
    username      VARCHAR(255),
    email         VARCHAR(255),
    password      VARCHAR(255),
    is_supper     BOOLEAN DEFAULT false,
    last_ip       VARCHAR(255),
    last_login    TIMESTAMP(6),
    avatar        VARCHAR(255),
    created_at    TIMESTAMP(6),
    updated_at    TIMESTAMP(6),
    deleted_at    TIMESTAMP(6),
    created_by    BIGINT,
    updated_by    BIGINT,
    deleted_by    BIGINT,
    date_of_birth TIMESTAMP(6),
    age           INT,
    gender        INT NOT NULL DEFAULT 0, -- 1=Male, 2=Female, 3=Other
    is_deleted    BOOLEAN DEFAULT false,
    first_name    VARCHAR(255),
    last_name     VARCHAR(255),
    is_active     BOOLEAN DEFAULT true,

    PRIMARY KEY (id)
);

INSERT INTO "role" ("name", description, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by)
VALUES('Administrator', 'Quản trị hệ thống', NULL, NULL, NULL, NULL, NULL, NULL);
