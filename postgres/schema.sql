CREATE SCHEMA pratik;

CREATE TABLE users (
   user_id serial PRIMARY KEY,
   name VARCHAR (150) NOT NULL,
   mobile_number VARCHAR (20) NOT NULL,
   password VARCHAR (60) NOT NULL,
   email VARCHAR (100) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL
);

ALTER TABLE users SET SCHEMA pratik;


-- Enables pgcrypto to be used when encrypting password
-- https://x-team.com/blog/storing-secure-passwords-with-postgresql/
CREATE EXTENSION pgcrypto;

-- Procedure to register a new user
CREATE OR REPLACE PROCEDURE registerUser(
  pName VARCHAR (150),
  pNumber VARCHAR (20),
  pPass VARCHAR (60),
  pEmail VARCHAR (100)
) LANGUAGE plpgsql
AS $$
BEGIN
  INSERT INTO users (
    name, mobile_number, password, email, created_at
  ) VALUES (
    pName,
    pNumber,
    crypt(pPass, gen_salt('bf')),
    pEmail,
    NOW()
  );

END;
$$;

-- Procedure to check if user exists in DB
CREATE OR REPLACE FUNCTION loginUser(
  pNumber VARCHAR (20),
  pPass VARCHAR (60)
)
RETURNS VARCHAR (150) AS $$
DECLARE userName VARCHAR (150);
BEGIN
  SELECT
    users.name
  INTO
    userName
  FROM
    users
  WHERE
    mobile_number = pNumber AND password = crypt(pPass, password);

  RETURN userName;
END;
$$  LANGUAGE plpgsql
