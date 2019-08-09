CREATE TABLE dbo.temperature (id serial PRIMARY KEY, temperature numeric(6, 2), create_date TIMESTAMP NOT NULL DEFAULT CURRENT_DATE)

INSERT INTO dbo.versions (name)
VALUES ('1.1')