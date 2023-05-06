mkdir -p ./migrations
cat << EOF > ./migrations/$1.up.sql
CREATE TABLE IF NOT EXISTS ${2%%.*} (
	id SERIAL PRIMARY KEY,
	column1 TEXT NOT NULL,
	column2 INTEGER NOT NULL,
	column3 TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
EOF

cat << EOF > ./migrations/$1.down.sql
DROP TABLE IF EXISTS ${2%%.*};
EOF