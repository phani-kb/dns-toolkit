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

-- source_types: processor type per source (e.g., domain, ipv4, adguard)
create table if not exists dnstk_source_types (
  id integer primary key,
  source_id integer not null references dnstk_sources (id) on delete cascade,
  type_name text not null,
  notes text,
  disabled integer not null default 0 check (disabled in (0, 1)),
  unique (source_id, type_name)
) strict;

create index if not exists idx_source_types_source_id on dnstk_source_types (source_id);

-- source_list_types: blocklist/allowlist per source type
create table if not exists dnstk_source_list_types (
  id integer primary key,
  source_type_id integer not null references dnstk_source_types (id) on delete cascade,
  list_type_name text not null check (list_type_name in ('blocklist', 'allowlist')),
  notes text,
  disabled integer not null default 0 check (disabled in (0, 1)),
  must_consider integer not null default 0 check (must_consider in (0, 1)),
  unique (source_type_id, list_type_name)
) strict;

create index if not exists idx_source_list_types_type_id on dnstk_source_list_types (source_type_id);

-- source_list_type_groups: size group membership for a list type
create table if not exists dnstk_source_list_type_groups (
  source_list_type_id integer not null references dnstk_source_list_types (id) on delete cascade,
  group_name text not null,
  primary key (source_list_type_id, group_name)
) strict, without rowid;

create index if not exists idx_source_list_type_groups_list_type_id on dnstk_source_list_type_groups (source_list_type_id);

-- source_categories: category tags per source
create table if not exists dnstk_source_categories (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  category text not null,
  primary key (source_id, category)
) strict, without rowid;

create index if not exists idx_source_categories_source_id on dnstk_source_categories (source_id);

-- source_countries: 2-letter country codes per source
create table if not exists dnstk_source_countries (
  source_id integer not null references dnstk_sources (id) on delete cascade,
  country_code text not null check (length(country_code) = 2),
  primary key (source_id, country_code)
) strict, without rowid;

create index if not exists idx_source_countries_source_id on dnstk_source_countries (source_id);

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
) strict, without rowid;

create index if not exists idx_source_files_source_id on dnstk_source_files (source_id);

-- downloads: download metadata and checksums
create table if not exists dnstk_downloads (
  source_id integer primary key references dnstk_sources (id) on delete cascade,
  url text,
  filepath text,
  frequency text,
  checksum text,
  error text,
  last_download_timestamp text check (
    last_download_timestamp is null or
    last_download_timestamp glob '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]*'
  ),
  last_checked_timestamp text check (
    last_checked_timestamp is null or
    last_checked_timestamp glob '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]*'
  ),
  type_count integer not null default 0 check (type_count >= 0),
  count_to_consider integer not null default 0 check (count_to_consider >= 0),
  skip_general_consolidation integer not null default 0 check (skip_general_consolidation in (0, 1)),
  skip_groups_consolidation integer not null default 0 check (skip_groups_consolidation in (0, 1)),
  skip_categories_consolidation integer not null default 0 check (skip_categories_consolidation in (0, 1))
) strict;

create index if not exists idx_downloads_source_id on dnstk_downloads (source_id);

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

-- entry_groups: group membership for processed file batches
create table if not exists dnstk_entry_groups (
  id integer primary key,
  source_id integer not null references dnstk_sources (id) on delete cascade,
  source_type text not null,
  list_type text not null,
  group_name text not null,
  unique (source_id, source_type, list_type, group_name)
) strict;

create index if not exists idx_entry_groups_source_id on dnstk_entry_groups (source_id);

-- entry_categories: category tags for processed file batches
create table if not exists dnstk_entry_categories (
  id integer primary key,
  source_id integer not null references dnstk_sources (id) on delete cascade,
  source_type text not null,
  list_type text not null,
  category text not null,
  unique (source_id, source_type, list_type, category)
) strict;

create index if not exists idx_entry_categories_source_id on dnstk_entry_categories (source_id);

-- consolidated_entries: deduplicated results after consolidation
create table if not exists dnstk_consolidated_entries (
  id integer primary key,
  entry text not null,
  generic_source_type text not null,
  list_type text not null check (list_type in ('blocklist', 'allowlist')),
  consolidation_type text not null,
  group_name text,
  category text,
  valid integer not null default 1 check (valid in (0, 1)),
  source_count integer not null default 1 check (source_count >= 0),
  check (
    (consolidation_type = 'general' and group_name is null and category is null) or
    (consolidation_type = 'group' and group_name is not null and category is null) or
    (consolidation_type = 'category' and group_name is null and category is not null)
  )
) strict;

create index if not exists idx_consolidated_lookup on dnstk_consolidated_entries (
  entry,
  generic_source_type,
  list_type,
  consolidation_type
);

create index if not exists idx_consolidated_type on dnstk_consolidated_entries (
  consolidation_type,
  generic_source_type,
  list_type
);

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
  overlap_percent real not null default 0.0 check (overlap_percent >= 0.0 and overlap_percent <= 100.0)
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