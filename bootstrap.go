package main

func bootstrap(settings *Settings) {
	database := get_database(settings)

	statement := `
    DROP SCHEMA IF EXISTS public CASCADE;

    CREATE SCHEMA IF NOT EXISTS public;

    CREATE TABLE IF NOT EXISTS categories
    (
        id INTEGER NOT NULL,
        title TEXT NOT NULL,
        sub_title TEXT NOT NULL,
        status TEXT NOT NULL DEFAULT 'Off'
    );

    CREATE SEQUENCE categories_id_sequence;

    ALTER TABLE categories ALTER COLUMN id SET DEFAULT NEXTVAL
    ('categories_id_sequence'::REGCLASS);

    ALTER TABLE categories ADD CONSTRAINT categories_id_constraint PRIMARY KEY
    (id);

    CREATE UNIQUE INDEX categories_title_sub_title ON categories
    (title, sub_title);

    CREATE INDEX categories_title ON categories USING BTREE (title);

    CREATE INDEX categories_sub_title ON categories USING BTREE (sub_title);

    CREATE INDEX categories_status ON categories USING BTREE (status);

    CREATE TABLE IF NOT EXISTS skills
    (
        id INTEGER NOT NULL,
        title TEXT NOT NULL,
        status TEXT NOT NULL DEFAULT 'Off'
    );

    CREATE SEQUENCE skills_id_sequence;

    ALTER TABLE skills ALTER COLUMN id SET DEFAULT NEXTVAL
    ('skills_id_sequence'::REGCLASS);

    ALTER TABLE skills ADD CONSTRAINT skills_id_constraint PRIMARY KEY (id);

    CREATE UNIQUE INDEX skills_title ON skills (title);

    CREATE INDEX skills_status ON skills USING BTREE (status);

    CREATE TABLE IF NOT EXISTS jobs
    (
        id INTEGER NOT NULL,
        budget INTEGER NOT NULL,
        category TEXT NOT NULL,
        client_country TEXT NOT NULL,
        client_feedback NUMERIC(3, 2) NOT NULL,
        client_jobs_posted INTEGER NOT NULL,
        client_past_hires INTEGER NOT NULL,
        client_reviews_count INTEGER NOT NULL,
        date_created TIMESTAMP WITHOUT TIME ZONE,
        duration TEXT NOT NULL,
        job_status TEXT NOT NULL,
        job_type TEXT NOT NULL,
        skills JSONB NOT NULL,
        snippet TEXT NOT NULL,
        sub_category TEXT NOT NULL,
        title TEXT NOT NULL,
        url TEXT NOT NULL,
        workload TEXT NOT NULL
    );

    CREATE SEQUENCE jobs_id_sequence;

    ALTER TABLE jobs ALTER COLUMN id SET DEFAULT NEXTVAL
    ('jobs_id_sequence'::REGCLASS);

    CREATE UNIQUE INDEX jobs_url ON jobs (url);

    ALTER TABLE jobs ADD CONSTRAINT jobs_id_constraint PRIMARY KEY (id);

    CREATE INDEX jobs_budget ON jobs USING BTREE (budget);

    CREATE INDEX jobs_category ON jobs USING BTREE (category);

    CREATE INDEX jobs_client_country ON jobs USING BTREE (client_country);

    CREATE INDEX jobs_client_feedback ON jobs USING BTREE (client_feedback);

    CREATE INDEX jobs_client_jobs_posted ON jobs USING BTREE
    (client_jobs_posted);

    CREATE INDEX jobs_client_past_hires ON jobs USING BTREE (client_past_hires);

    CREATE INDEX jobs_client_reviews_count ON jobs USING BTREE
    (client_reviews_count);

    CREATE INDEX jobs_date_created ON jobs USING BTREE (date_created);

    CREATE INDEX jobs_duration ON jobs USING BTREE (duration);

    CREATE INDEX jobs_job_status ON jobs USING BTREE (job_status);

    CREATE INDEX jobs_job_type ON jobs USING BTREE (job_type);

    CREATE INDEX jobs_skills ON jobs USING GIN (skills);

    CREATE INDEX jobs_snippet ON jobs USING BTREE (snippet);

    CREATE INDEX jobs_sub_category ON jobs USING BTREE (sub_category);

    CREATE INDEX jobs_title ON jobs USING BTREE (title);

    CREATE INDEX jobs_workload ON jobs USING BTREE (workload);

    CREATE FUNCTION categories_insert
    (title_value TEXT, sub_title_value TEXT, status_value TEXT) RETURNS
    VOID AS
    $$
        BEGIN
            INSERT INTO categories (title, sub_title, status) VALUES
            (title_value, sub_title_value, status_value) ON CONFLICT
            (title, sub_title) DO NOTHING;
        END;
    $$
    LANGUAGE plpgsql;

    CREATE FUNCTION skills_insert (title_value TEXT, status_value TEXT) RETURNS
    VOID AS
    $$
        BEGIN
            INSERT INTO skills (title, status) VALUES
            (title_value, status_value) ON CONFLICT (title) DO NOTHING;
        END;
    $$
    LANGUAGE plpgsql;
    `

	defer database.Close()

	database.MustExec(statement)
}
