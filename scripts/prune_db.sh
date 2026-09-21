#!/bin/bash
set -euo pipefail

DB_PATH="${1:-data/dns-toolkit.db}"

if [[ ! -f "$DB_PATH" ]]; then
    echo "No database at $DB_PATH; nothing to prune."
    exit 0
fi

if ! command -v sqlite3 >/dev/null; then
    echo "error: sqlite3 CLI not found" >&2
    exit 1
fi

echo "Pruning tables and compacting $DB_PATH ..."

sqlite3 "$DB_PATH" <<'SQL'
PRAGMA wal_checkpoint(TRUNCATE);
BEGIN;
DELETE FROM dnstk_consolidated_general;
DELETE FROM dnstk_consolidated_group;
DELETE FROM dnstk_consolidated_category;
DELETE FROM dnstk_top_entries;
DELETE FROM dnstk_overlap_results;
COMMIT;
VACUUM;
PRAGMA wal_checkpoint(TRUNCATE);
SQL

echo "Done."
