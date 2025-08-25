#!/bin/bash

FMT_TOOLS=("goimports" "golines" "gofmt")
INSTALL_TOOLS=("goimports" "golines")
NO_PATH_TOOLS=("gofmt")
GOBIN=$(go env GOPATH)/bin

install_tools() {
  for tool in "${INSTALL_TOOLS[@]}"; do
    if ! command -v "$GOBIN/$tool" &>/dev/null; then
      echo "Installing $tool..."
      if [ 0"$tool" = "golines" ]; then
        go install github.com/segmentio/golines@latest
      else
        go install golang.org/x/tools/cmd/"$tool"@latest
      fi
    fi
  done
  export PATH="$GOBIN:$PATH"
}

fmt() {
  install_tools
  for tool in "${FMT_TOOLS[@]}"; do
    echo "Running $tool..."
    if [ "$tool" = "golines" ]; then
      find . -name '*.go' | xargs "$GOBIN/$tool" --max-len=120 --base-formatter=gofumpt -w
    elif [ "$tool" = "gofmt" ]; then
      $tool -s -w .
    else
      "$GOBIN/$tool" -w .
    fi
  done
}

clean_tools() {
  for tool in "${INSTALL_TOOLS[@]}"; do
    echo "Removing $tool..."
    rm -f "$GOBIN/$tool"
  done
}

sort_config_sources() {
  for file in data/config/sources_*.json; do
    if [[ "$file" == *"sources_schema.json" ]]; then
      continue
    fi
    echo "Sorting $file..."
    jq '{
      sources: (
        .sources
        | sort_by(.name | ascii_downcase)
        | map(
            {name} + (del(.name) | to_entries | sort_by(.key) | from_entries | 
            with_entries(
              if .key == "categories" or .key == "types" or .key == "list_types" then
                if .value | type == "string" then
                  .value = (.value | 
                    split(", ") | 
                    sort | 
                    join(", ")
                  )
                else .
                end
              else .
              end
            ))
          )
      )
    }' "$file" >"${file}.sorted"
    mv "${file}.sorted" "$file"
  done
}

print_types_names() {
  echo "Printing types names..."
  find data/config -name "sources_[i|d|l]*.json" -type f | xargs jq -r '.sources[] | select(.types != null) | .types[] | select(.name != null) | .name' | sort -u
}

print_types_with_counts() {
  echo "Source types and their counts:"
  echo "============================="
  
  # Get all type names and count occurrences
  find data/config -name "sources_*.json" -type f ! -name "*schema*" | \
  xargs jq -r '.sources[] | select(.types != null) | .types[] | select(.name != null) | .name' | \
  sort | uniq -c | sort -nr | \
  while read count type; do
    printf "%-30s %3d\n" "$type" "$count"
  done
  
  echo "Total unique types: $(find data/config -name "sources_*.json" -type f ! -name "*schema*" | xargs jq -r '.sources[] | select(.types != null) | .types[] | select(.name != null) | .name' | sort -u | wc -l)"
  echo "Total type instances: $(find data/config -name "sources_*.json" -type f ! -name "*schema*" | xargs jq -r '.sources[] | select(.types != null) | .types[] | select(.name != null) | .name' | wc -l)"
}

print_source_names() {
  find data/config -name "sources_*.json" -type f ! -name "*schema*" | \
    xargs jq -r '.sources[] | select(.name != null) | .name' | \
    sort | \
    awk '{printf "%s%s", sep, $0; sep=", "}' && echo
}

print_duplicate_source_names() {
  find data/config -name "sources_*.json" -type f ! -name "*schema*" | \
    xargs jq -r '.sources[] | select(.name != null) | .name' | \
    sort | uniq -c | awk '$1 > 1 {print $2 " (" $1 " times)"}'
}

print_source_urls() {
  echo "Blocklist URLs:"
  print_blocklist_urls
  echo ""

  echo "Allowlist URLs:"
  print_allowlist_urls
}



print_blocklist_urls() {
  find data/config -name "sources_*.json" -type f ! -name "*schema*" | \
    xargs jq -r '
      .sources[] |
      select((.types[]? | (.list_types == null or (.list_types[]?.name == "blocklist"))) ) |
      select(.url != null) | .url
    ' | sort -u
}

print_allowlist_urls() {
  find data/config -name "sources_*.json" -type f ! -name "*schema*" | \
    xargs jq -r '
      .sources[] |
      select((.types[]? | (.list_types[]?.name == "allowlist"))) |
      select(.url != null) | .url
    ' | sort -u
}

case "$1" in
fmt)
  fmt
  ;;
install-tools)
  install_tools
  ;;
clean-tools)
  clean_tools
  ;;
sort-config-sources)
  sort_config_sources
  ;;
print-source-types-names)
  print_types_names
  ;;
print-source-types-counts)
  print_types_with_counts
  ;;
print-source-names)
  print_source_names
  ;;
print-duplicate-source-names)
  print_duplicate_source_names
  ;;
print-source-urls)
  print_source_urls
  ;;

*)
  echo "Usage: $0 {fmt|install-tools|clean-tools|sort-config-sources|print-source-types-names|print-source-types-counts|print-source-names|print-duplicate-source-names|print-source-urls}"
  exit 1
  ;;
esac
