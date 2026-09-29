-- dns-toolkit schema
-- sources: consolidated from all sources_*.json files
create table if not exists dnstk_sources (
  id integer primary key,
  name text not null,
  url text,
  url_per_category text,
  url_per_group text,
  frequency text not null default 'daily' check (frequency in ('hourly', 'daily', 'weekly', 'monthly')),
  license text,
  website text,
  notes text,
  type_count integer not null default 0 check (type_count >= 0),
  count_to_consider integer not null default 0 check (count_to_consider >= 0),
  disabled integer not null default 0 check (disabled in (0, 1)),
  skip_general_consolidation integer not null default 0 check (skip_general_consolidation in (0, 1)),
  skip_groups_consolidation integer not null default 0 check (skip_groups_consolidation in (0, 1)),
  skip_categories_consolidation integer not null default 0 check (skip_categories_consolidation in (0, 1)),
  source_file text,
  definition_checksum text,
  unique (name, source_file)
) strict;

create index if not exists idx_sources_disabled_name on dnstk_sources (disabled, name);

-- type_names: global registry of source processor type strings
create table if not exists dnstk_type_names (id integer primary key, name text not null unique) strict;

-- list_type_names: global registry of list type strings
create table if not exists dnstk_list_type_names (id integer primary key, name text not null unique check (name in ('blocklist', 'allowlist'))) strict;

-- group_names: global registry of size group strings
create table if not exists dnstk_group_names (id integer primary key, name text not null unique) strict;

-- category_names: global registry of category strings
create table if not exists dnstk_category_names (id integer primary key, name text not null unique) strict;

-- source_types: which type names each source uses, with per-source disabled flag
create table if not exists dnstk_source_types (
  id integer primary key,
  source_id integer not null references dnstk_sources (id) on delete cascade,
  type_name_id integer not null references dnstk_type_names (id),
  notes text,
  disabled integer not null default 0 check (disabled in (0, 1)),
  unique (source_id, type_name_id)
) strict;

create index if not exists idx_source_types_source_id on dnstk_source_types (source_id);

create index if not exists idx_source_types_type_name_id on dnstk_source_types (type_name_id);

-- source_list_types: which list type names each source type uses
create table if not exists dnstk_source_list_types (
  id integer primary key,
  source_type_id integer not null references dnstk_source_types (id) on delete cascade,
  list_type_name_id integer not null references dnstk_list_type_names (id),
  disabled integer not null default 0 check (disabled in (0, 1)),
  must_consider integer not null default 0 check (must_consider in (0, 1)),
  unique (source_type_id, list_type_name_id)
) strict;

create index if not exists idx_source_list_types_type_id on dnstk_source_list_types (source_type_id);

create index if not exists idx_source_list_types_list_type_name_id on dnstk_source_list_types (list_type_name_id);

-- source_list_type_notes: optional per-source-list-type notes (only inserted when non-empty)
create table if not exists dnstk_source_list_type_notes (source_list_type_id integer primary key references dnstk_source_list_types (id) on delete cascade, notes text not null) strict;

-- source_list_type_groups: which group names a list type belongs to
create table if not exists dnstk_source_list_type_groups (
  source_list_type_id integer not null references dnstk_source_list_types (id) on delete cascade,
  group_name_id integer not null references dnstk_group_names (id),
  primary key (source_list_type_id, group_name_id)
) strict,
without rowid;

create index if not exists idx_source_list_type_groups_list_type_id on dnstk_source_list_type_groups (source_list_type_id);

create index if not exists idx_source_list_type_groups_group_name_id on dnstk_source_list_type_groups (group_name_id);

-- source_categories: category membership per source, referencing global category_names
create table if not exists dnstk_source_categories (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  category_name_id integer not null references dnstk_category_names (id),
  primary key (source_id, category_name_id)
) strict,
without rowid;

-- source_countries: 2-letter country codes per source
create table if not exists dnstk_source_countries (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  country_code text not null check (length(country_code) = 2),
  primary key (source_id, country_code)
) strict,
without rowid;

create index if not exists idx_source_countries_country_code on dnstk_source_countries (country_code);

-- source_content: inline content entries (for sources without URLs)
create table if not exists dnstk_source_content (
  id integer primary key,
  source_id integer not null references dnstk_sources (id) on delete cascade,
  content_type text not null default 'content',
  entry text not null,
  unique (source_id, content_type, entry)
) strict;

create index if not exists idx_source_content_source_id on dnstk_source_content (source_id);

-- source_files: archive file references
create table if not exists dnstk_source_files (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  filename text not null,
  primary key (source_id, filename)
) strict,
without rowid;

-- downloads: download metadata and checksums
create table if not exists dnstk_downloads (
  source_id integer primary key references dnstk_sources (id) on delete cascade,
  url text,
  filepath text,
  frequency text,
  checksum text,
  last_processed_checksum text,
  error text,
  last_download_timestamp text check (
    last_download_timestamp is null
    or last_download_timestamp glob '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]*'
  ),
  last_checked_timestamp text check (
    last_checked_timestamp is null
    or last_checked_timestamp glob '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]*'
  ),
  last_processed_timestamp text check (
    last_processed_timestamp is null
    or last_processed_timestamp glob '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]*'
  ),
  type_count integer not null default 0 check (type_count >= 0),
  count_to_consider integer not null default 0 check (count_to_consider >= 0),
  skip_general_consolidation integer not null default 0 check (skip_general_consolidation in (0, 1)),
  skip_groups_consolidation integer not null default 0 check (skip_groups_consolidation in (0, 1)),
  skip_categories_consolidation integer not null default 0 check (skip_categories_consolidation in (0, 1))
) strict;

-- entries: the main table storing all processed domains/IPs
create table if not exists dnstk_entries (
  id integer primary key,
  source_id integer not null references dnstk_sources (id) on delete cascade,
  entry text not null,
  generic_source_type text not null,
  actual_source_type text not null,
  list_type text not null check (list_type in ('blocklist', 'allowlist')),
  valid integer not null default 1 check (valid in (0, 1)),
  must_consider integer not null default 0 check (must_consider in (0, 1)),
  unique (source_id, entry, generic_source_type, actual_source_type, list_type)
) strict;

create index if not exists idx_entries_lookup on dnstk_entries (entry, generic_source_type, list_type);

create index if not exists idx_entries_source on dnstk_entries (source_id, generic_source_type);

create index if not exists idx_entries_source_type_list on dnstk_entries (source_id, actual_source_type, list_type);

create index if not exists idx_entries_consolidation on dnstk_entries (generic_source_type, list_type, valid, entry, source_id, must_consider);

-- entry_groups: group membership for processed file batches
create table if not exists dnstk_entry_groups (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  source_type text not null,
  list_type text not null,
  group_name text not null,
  primary key (source_id, source_type, list_type, group_name)
) strict,
without rowid;

create index if not exists idx_entry_groups_join on dnstk_entry_groups (source_id, source_type, list_type);

create index if not exists idx_entry_groups_scope on dnstk_entry_groups (group_name, source_type, list_type, source_id);

-- entry_categories: category tags for processed file batches
create table if not exists dnstk_entry_categories (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  source_type text not null,
  list_type text not null,
  category text not null,
  primary key (source_id, source_type, list_type, category)
) strict,
without rowid;

create index if not exists idx_entry_categories_join on dnstk_entry_categories (source_id, source_type, list_type);

create index if not exists idx_entry_categories_scope on dnstk_entry_categories (category, source_type, list_type, source_id);

-- source_entry_counts: per-source processed-entry
create table if not exists dnstk_source_entry_counts (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  generic_source_type text not null,
  actual_source_type text not null,
  list_type text not null check (list_type in ('blocklist', 'allowlist')),
  valid integer not null default 1 check (valid in (0, 1)),
  must_consider integer not null default 0 check (must_consider in (0, 1)),
  entry_count integer not null default 0 check (entry_count >= 0),
  primary key (source_id, generic_source_type, actual_source_type, list_type, valid)
) strict,
without rowid;

create index if not exists idx_source_entry_counts_source on dnstk_source_entry_counts (source_id);

-- consolidated_general: consolidated results for general scope
create table if not exists dnstk_consolidated_general (
  entry text not null,
  generic_source_type text not null,
  list_type text not null check (list_type in ('blocklist', 'allowlist')),
  valid integer not null default 1 check (valid in (0, 1)),
  source_count integer not null default 1 check (source_count >= 0),
  primary key (entry, generic_source_type, list_type, valid)
) strict,
without rowid;

create index if not exists idx_consolidated_general_type on dnstk_consolidated_general (generic_source_type, list_type, valid, entry);

-- consolidated_group: consolidated results for group scope
create table if not exists dnstk_consolidated_group (
  entry text not null,
  generic_source_type text not null,
  list_type text not null check (list_type in ('blocklist', 'allowlist')),
  group_name text not null,
  valid integer not null default 1 check (valid in (0, 1)),
  source_count integer not null default 1 check (source_count >= 0),
  primary key (entry, generic_source_type, list_type, group_name, valid)
) strict,
without rowid;

create index if not exists idx_consolidated_group_scope on dnstk_consolidated_group (group_name, generic_source_type, list_type, valid, entry);

create index if not exists idx_consolidated_group_type on dnstk_consolidated_group (generic_source_type, list_type, valid, entry);

-- consolidated_category: consolidated results for category scope
create table if not exists dnstk_consolidated_category (
  entry text not null,
  generic_source_type text not null,
  list_type text not null check (list_type in ('blocklist', 'allowlist')),
  category text not null,
  valid integer not null default 1 check (valid in (0, 1)),
  source_count integer not null default 1 check (source_count >= 0),
  primary key (entry, generic_source_type, list_type, category, valid)
) strict,
without rowid;

create index if not exists idx_consolidated_category_scope on dnstk_consolidated_category (category, generic_source_type, list_type, valid, entry);

create index if not exists idx_consolidated_category_type on dnstk_consolidated_category (generic_source_type, list_type, valid, entry);

-- overlap_results: overlap analysis between sources
create table if not exists dnstk_overlap_results (
  id integer primary key,
  source_name text not null,
  target_name text not null,
  generic_source_type text not null,
  source_list_type text not null,
  target_list_type text not null,
  overlap_count integer not null default 0 check (overlap_count >= 0),
  source_count integer not null default 0 check (source_count >= 0),
  target_count integer not null default 0 check (target_count >= 0),
  overlap_percent real not null default 0.0 check (
    overlap_percent >= 0.0
    and overlap_percent <= 100.0
  )
) strict;

create index if not exists idx_overlap_source on dnstk_overlap_results (source_name, generic_source_type);

-- top_entries: entries appearing in N+ sources
create table if not exists dnstk_top_entries (
  id integer primary key,
  entry text not null,
  generic_source_type text not null,
  list_type text not null,
  source_count integer not null check (source_count >= 0),
  min_sources integer not null check (min_sources >= 0),
  unique (entry, generic_source_type, list_type, min_sources)
) strict;

create index if not exists idx_top_entries_lookup on dnstk_top_entries (generic_source_type, list_type, min_sources);

-- resolved_allow: entries that are allowed after resolution
create table if not exists dnstk_resolved_allow (
  generic_source_type text not null,
  entry text not null,
  must_consider integer not null default 0 check (must_consider in (0, 1)),
  primary key (generic_source_type, entry)
) strict;

-- scoped_allow: small per-scope allow set for group/category
create table if not exists dnstk_scoped_allow (
  consolidation_type text not null,
  scope_value text not null,
  entry text not null,
  must_consider integer not null default 0 check (must_consider in (0, 1)),
  primary key (consolidation_type, scope_value, entry)
) strict,
without rowid;
