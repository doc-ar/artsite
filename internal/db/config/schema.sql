CREATE TABLE series (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
  name varchar(50) UNIQUE NOT NULL,
  description varchar(4000) DEFAULT '' NOT NULL,
  cover_img varchar(80) DEFAULT '' NOT NULL,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE art (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
  series_id uuid REFERENCES series (id) NOT NULL,
  title varchar(150) NOT NULL,
  description varchar(4000) NOT NULL,
  image_url varchar(80) NOT NULL,
  width int NOT NULL check (width >= 0),
  height int NOT NULL check (height >= 0),
  category varchar(15) NOT NULL CHECK (category IN ('Original', 'Print', 'Canvas Print')),
  medium varchar(15) NOT NULL CHECK (
    medium IN (
      'Water Color',
      'Oil',
      'Digital',
      'Mix Media',
      'Graphite',
      'Ink',
      'Acrylic'
    )
  ),
  price int NOT NULL check (price >= 0),
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  view_count int NOT NULL DEFAULT 0 check (view_count >= 0)
);

CREATE TABLE portfolio (
  art_id uuid NOT NULL REFERENCES art (id) UNIQUE,
  added_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE admins (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid () NOT NULL,
  username varchar(20) UNIQUE NOT NULL,
  password_hash varchar(60) NOT NULL,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
  series (name)
VALUES
  ('None') RETURNING *;
