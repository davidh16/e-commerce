mkdir -p ./migrations
cat << EOF > ./migrations/$1.up.sql
CREATE TABLE IF NOT EXISTS ${2%%.*} (
	id SERIAL PRIMARY KEY,
	uuid uuid DEFAULT uuid_generate_v4(),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	----------------------------
    -- extras

);
EOF

cat << EOF > ./migrations/$1.down.sql
DROP TABLE IF EXISTS ${2%%.*};
EOF