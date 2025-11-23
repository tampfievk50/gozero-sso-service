-- resource_seq definition

CREATE SEQUENCE IF NOT EXISTS resource_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1
	NO CYCLE;

-- "resource" definition

-- Drop table

CREATE TABLE IF NOT EXISTS "resource"
(
    id
    int8
    NOT
    NULL
    DEFAULT
    nextval
(
    'resource_seq'
    :
    :
    regclass
),
    "name" varchar
(
    255
) NULL,
    description varchar NULL,
    created_at timestamp
(
    6
) NULL,
    updated_at timestamp
(
    6
) NULL,
    deleted_at timestamp
(
    6
) NULL,
    created_by int8 NULL,
    updated_by int8 NULL,
    deleted_by int8 NULL,
    CONSTRAINT resource_pkey PRIMARY KEY
(
    id
)
    );


-- user_role definition

-- Drop table

CREATE TABLE IF NOT EXISTS user_resource
(
    resource_id
    int8
    NOT
    NULL,
    user_id
    int8
    NOT
    NULL,
    CONSTRAINT
    user_resource_pkey
    PRIMARY
    KEY
(
    resource_id,
    user_id
),
    CONSTRAINT "fk_user_resource_user_id" FOREIGN KEY
(
    user_id
) REFERENCES "user"
(
    id
),
    CONSTRAINT "fk_user_resource_resource_id" FOREIGN KEY
(
    resource_id
) REFERENCES "resource"
(
    id
)
    );