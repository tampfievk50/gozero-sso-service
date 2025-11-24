-- ===============================
-- resource_seq definition
-- ===============================

CREATE SEQUENCE IF NOT EXISTS resource_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1
    NO CYCLE;


-- ===============================
-- resource table definition
-- ===============================

CREATE TABLE IF NOT EXISTS "resource" (
    id          BIGINT NOT NULL DEFAULT nextval('resource_seq'::regclass),
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
-- user_resource table definition
-- ===============================

CREATE TABLE IF NOT EXISTS user_resource (
    resource_id BIGINT NOT NULL,
    user_id     BIGINT NOT NULL,

    PRIMARY KEY (resource_id, user_id),
    CONSTRAINT fk_user_resource_user_id
    FOREIGN KEY (user_id) REFERENCES "user" (id),
    CONSTRAINT fk_user_resource_resource_id
    FOREIGN KEY (resource_id) REFERENCES "resource" (id)
);

INSERT INTO resource ("name", description, created_at, updated_at, deleted_at, created_by, updated_by, deleted_by)
VALUES('VMG', 'Công ty cổ phần truyền thông VMG', NULL, NULL, NULL, NULL, NULL, NULL);