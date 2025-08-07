#!/bin/bash

set -e

echo "Starting DNS Toolkit pipeline..."

if [ ! -f "bin/dns-toolkit" ]; then
  echo "dns-toolkit binary not found. Building..."
  go build -o bin/dns-toolkit .
  chmod +x bin/dns-toolkit
fi

# Run each step using short forms:
# d  - download
# p  - process
# c  - consolidate
# cg - consolidate groups
# cc - consolidate categories
# t  - top
# o  - overlap
# op - output
# gr - output README
# gor- overlap README
# gsr- summaries README
# gs - stats README
# cp - copy summaries

if [ "$#" -gt 0 ]; then
    IFS=',' read -ra steps <<< "$1"
else
    steps=("d" "p" "c" "cg" "cc" "t" "o" "op" "gr" "gor" "gsr" "gs" "cp")
fi

for step in "${steps[@]}"; do
    case "$step" in
        d)
            echo "Step 1: Downloading data..."
            ./bin/dns-toolkit download
            ;;
        p)
            echo "Step 2: Processing data..."
            ./bin/dns-toolkit process
            ;;
        c)
            echo "Step 3: Consolidating data..."
            ./bin/dns-toolkit consolidate
            ;;
        cg)
            echo "Step 4: Grouping consolidated data..."
            ./bin/dns-toolkit consolidate groups
            ;;
        cc)
            echo "Step 5: Categorizing data..."
            ./bin/dns-toolkit consolidate categories
            ;;
        t)
            echo "Step 6: Generating top entries..."
            ./bin/dns-toolkit top
            ;;
        o)
            echo "Step 7: Finding overlaps..."
            ./bin/dns-toolkit overlap #--cpu-profile --mem-profile
            ;;
        op)
            echo "Step 8: Generating output files..."
            ./bin/dns-toolkit generate output -i
            ;;
        gr)
            echo "Step 9: Generating output README..."
            ./bin/dns-toolkit generate output-readme
            ;;
        gor)
            echo "Step 10: Generating overlap README..."
            ./bin/dns-toolkit generate overlap-readme
            ;;
        gsr)
            echo "Step 11: Generating summaries README..."
            ./bin/dns-toolkit generate summaries-readme
            ;;
        gs)
            echo "Step 12: Generating stats README..."
            ./bin/dns-toolkit generate stats-readme
            ;;
        cp)
            echo "Step 13: Copy summary files to archive..."
            cp data/output/summaries/* data/archive/
            ;;
        *)
            echo "Unknown step: $step"
            exit 1
            ;;
    esac
done

echo "Pipeline completed successfully!"
