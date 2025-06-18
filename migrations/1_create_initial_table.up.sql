    CREATE TABLE "position"
    (
        "position_id" BIGSERIAL PRIMARY KEY,
        "position_name" VARCHAR(100) NOT NULL,
        "department_id" BIGINT,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

    CREATE TABLE "type_of_work"
    (
        "type_of_work_id" BIGSERIAL PRIMARY KEY,
        "type_of_work_name" VARCHAR(100) NOT NULL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

    CREATE TABLE "location"
    (
        "location_id" BIGSERIAL PRIMARY KEY,
        "location_name" VARCHAR(100) NOT NULL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

    CREATE TABLE "job"
    (
        "job_id" BIGSERIAL PRIMARY KEY,
        "department_id" BIGINT,
        "position_id" BIGINT,
        "type_of_work" VARCHAR(100) NOT NULL,
        "location_id" BIGINT,
        "first_periode" TIMESTAMP WITH TIME ZONE,
        "end_periode" TIMESTAMP WITH TIME ZONE,
        "description" TEXT,
        "requirement" TEXT,
        "is_active" BOOL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );
    ALTER TABLE "job" ADD FOREIGN KEY ("position_id") REFERENCES "position" ("position_id");
    ALTER TABLE "job" ADD FOREIGN KEY ("location_id") REFERENCES "location" ("location_id");

    CREATE TABLE "counter"
    (
        "counter_id" BIGSERIAL PRIMARY KEY,
        "counter_name" VARCHAR(100) NOT NULL,
        "counter_qty" BIGINT,
        "is_active" BOOL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

    CREATE TABLE "client"
    (
        "client_id" BIGSERIAL PRIMARY KEY,
        "client_name" VARCHAR(100) NOT NULL,
        "client_photo" TEXT,
        "is_active" BOOL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

    CREATE TABLE "testimoni"
    (
        "testimoni_id" BIGSERIAL PRIMARY KEY,
        "customer_name" VARCHAR(100) NOT NULL,
        "testimoni_type" VARCHAR(100) NOT NULL,
        "position" TEXT,
        "testimoni_order" BIGINT,
        "testimoni_photo" TEXT,
        "testimoni_desk" TEXT,
        "is_active" BOOL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

        CREATE TABLE "news_and_event"
    (
        "news_and_event_id" BIGSERIAL PRIMARY KEY,
        "title" VARCHAR(100) NOT NULL,
        "photo" TEXT,
        "description" TEXT,
        "is_active" BOOL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

    CREATE TABLE "article"
    (
        "article_id" BIGSERIAL PRIMARY KEY,
        "title" VARCHAR(100) NOT NULL,
        "cover" TEXT,
        "content" TEXT,
        "is_active" BOOL,
        "created_by" VARCHAR(64) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        "updated_by" VARCHAR(64),
        "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );