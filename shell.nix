# To use this file, run `nix-shell` in the same directory.
# This will drop you into a new shell with postgresql available.

with import <nixpkgs> { };

mkShell {
  # The buildInputs attribute lists the packages that will be available
  # in the development shell.
  buildInputs = [ postgresql ];

  # The shellHook is a script that runs every time you enter the shell.
  # We use it here to set up the database environment and provide instructions.
  shellHook = ''
    # Set the PGDATA environment variable to a local directory.
    # This tells PostgreSQL where to store the database files.
    export PGDATA=$(pwd)/.pgdata

    # Set a default database name for convenience.
    export PGDATABASE=gator

    echo "--- PostgreSQL Environment ---"
    echo "Database location: $PGDATA"
    echo "Default database:  $PGDATABASE"
    echo ""
    echo "--- Available Commands ---"
    echo "db-init:    Initialize a new database cluster in $PGDATA"
    echo "db-start:   Start the PostgreSQL server in the background"
    echo "db-stop:    Stop the PostgreSQL server"
    echo "db-connect: Connect to the '$PGDATABASE' database using psql"
    echo "db-create:  Create the '$PGDATABASE' database if it doesn't exist"
    echo "--------------------------"

    # Define aliases for our database management commands.
    alias db-init="initdb -D $PGDATA --no-locale --encoding=UTF8"
    alias db-start="pg_ctl -D $PGDATA -l $PGDATA/logfile start"
    alias db-stop="pg_ctl -D $PGDATA stop"
    alias db-connect="psql $PGDATABASE"

    # This alias is a bit more complex. It tries to create the database
    # and silences the "already exists" error so it can be run safely.
    alias db-create="createdb $PGDATABASE 2>/dev/null || true"
  '';
}
