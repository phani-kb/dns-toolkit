#!/bin/bash

set -e

if [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
    echo "Usage: $0 [step1,step2,step3...]"
    echo ""
    echo "Available steps:"
    echo "  ga  - generate allowlist"
    echo "  d   - download"
    echo "  p   - process"
    echo "  c   - consolidate (general with conflict resolution)"
    echo "  cg  - consolidate groups"
    echo "  cc  - consolidate categories"
    echo "  cf  - consolidate fast (skip conflicts and checksums)"
    echo "  cgf - consolidate groups (same as cg)"
    echo "  ccf - consolidate categories (same as cc)"
    echo "  t   - top"
    echo "  o   - overlap"
    echo "  op  - output"
    echo "  gr  - output README"
    echo "  gor - overlap README"
    echo "  gsr - summaries README"
    echo "  gs  - stats README"
    echo "  gc  - credits README"
    echo "  cp  - copy summaries"
    echo ""
    echo "Environment variables:"
    echo "  SKIP_CHECKSUMS=true   - Skip checksum calculation for general consolidation"
    echo "  SKIP_CONFLICTS=true   - Skip conflict resolution for general consolidation"
    echo "  INCLUDE_INVALID=true  - Include invalid entries in general consolidation"
    echo ""
    echo "Examples:"
    echo "  $0                    - Run full pipeline"
    echo "  $0 d,p,c             - Run download, process, consolidate"
    echo "  SKIP_CHECKSUMS=true $0 c  - Fast general consolidation without checksums"
    echo "  $0 cf,cgf,ccf        - Fast consolidation steps"
    exit 0
fi

echo "Starting DNS Toolkit pipeline..."

# Environment variables for pipeline configuration
SKIP_CHECKSUMS=${SKIP_CHECKSUMS:-false}
SKIP_CONFLICTS=${SKIP_CONFLICTS:-false}
INCLUDE_INVALID=${INCLUDE_INVALID:-false}

# Build consolidation flags based on environment
CONSOLIDATION_FLAGS=""
if [ "$SKIP_CHECKSUMS" = "false" ]; then
    CONSOLIDATION_FLAGS="$CONSOLIDATION_FLAGS --calculate-checksum"
fi
if [ "$SKIP_CONFLICTS" = "false" ]; then
    CONSOLIDATION_FLAGS="$CONSOLIDATION_FLAGS --gen-conflicts"
fi
if [ "$INCLUDE_INVALID" = "true" ]; then
    CONSOLIDATION_FLAGS="$CONSOLIDATION_FLAGS --include-invalid"
fi

echo "Consolidation flags: $CONSOLIDATION_FLAGS"

# Track pipeline timing
PIPELINE_START=$(date +%s)

if [ ! -f "bin/dns-toolkit" ]; then
  echo "dns-toolkit binary not found. Building..."
  go build -o bin/dns-toolkit .
  chmod +x bin/dns-toolkit
fi

if [ "$#" -gt 0 ]; then
    IFS=',' read -ra steps <<< "$1"
else
    steps=("ga" "d" "p" "c" "cg" "cc" "t" "o" "op" "gr" "gor" "gsr" "gs" "gc" "cp")
fi

for step in "${steps[@]}"; do
    STEP_START=$(date +%s)
    case "$step" in
        ga)
            echo "Step 1: Generating allowlist..."
            ./bin/dns-toolkit generate allowlist --overwrite
            ;;
        d)
            echo "Step 2: Downloading data..."
            ./bin/dns-toolkit download
            ;;
        p)
            echo "Step 3: Processing data..."
            ./bin/dns-toolkit process
            ;;
        c)
            echo "Step 4: Consolidating data with conflict resolution..."
            ./bin/dns-toolkit consolidate all $CONSOLIDATION_FLAGS
            ;;
        cg)
            echo "Step 5: Grouping consolidated data..."
            ./bin/dns-toolkit consolidate groups
            ;;
        cc)
            echo "Step 6: Categorizing data..."
            ./bin/dns-toolkit consolidate categories
            ;;
        cf)
            echo "Step 4 (Fast): Consolidating data without conflicts/checksums..."
            ./bin/dns-toolkit consolidate all
            ;;
        cgf)
            echo "Step 5: Grouping consolidated data..."
            ./bin/dns-toolkit consolidate groups
            ;;
        ccf)
            echo "Step 6: Categorizing data..."
            ./bin/dns-toolkit consolidate categories
            ;;
        t)
            echo "Step 7: Generating top entries..."
            ./bin/dns-toolkit top
            ;;
        o)
            echo "Step 8: Finding overlaps..."
            ./bin/dns-toolkit overlap #--cpu-profile --mem-profile
            ;;
        op)
            echo "Step 9: Generating output files..."
            ./bin/dns-toolkit generate output -i
            ;;
        gr)
            echo "Step 10: Generating output README..."
            ./bin/dns-toolkit generate output-readme
            ;;
        gor)
            echo "Step 11: Generating overlap README..."
            ./bin/dns-toolkit generate overlap-readme
            ;;
        gsr)
            echo "Step 12: Generating summaries README..."
            ./bin/dns-toolkit generate summaries-readme
            ;;
        gs)
            echo "Step 13: Generating stats README..."
            ./bin/dns-toolkit generate stats-readme
            ;;
        gc)
            echo "Step 14: Generating credits README..."
            ./bin/dns-toolkit generate credits
            ;;
        cp)
            echo "Step 15: Copy summary files to archive..."
            cp data/output/summaries/* data/archive/
            ;;
        *)
            echo "Unknown step: $step"
            exit 1
            ;;
    esac
    STEP_END=$(date +%s)
    STEP_DURATION=$((STEP_END - STEP_START))
    echo "Step $step completed in ${STEP_DURATION}s"
done

PIPELINE_END=$(date +%s)
PIPELINE_DURATION=$((PIPELINE_END - PIPELINE_START))
echo "Pipeline completed successfully in ${PIPELINE_DURATION}s!"
