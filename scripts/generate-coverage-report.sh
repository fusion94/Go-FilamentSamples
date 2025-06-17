#!/bin/bash

# Generate Test Coverage Report in Markdown
# Usage: ./scripts/generate-coverage-report.sh

set -e

echo "🧪 Generating test coverage report..."

# Run tests with coverage
go test -coverprofile=coverage.out ./...

# Get overall coverage percentage
OVERALL_COVERAGE=$(go tool cover -func=coverage.out | grep "total:" | awk '{print $3}')

# Get detailed function coverage
DETAILED_COVERAGE=$(go tool cover -func=coverage.out)

# Get package coverage summary
PACKAGE_COVERAGE=$(go test -cover ./... 2>/dev/null | grep "coverage:" | sort -k4 -nr)

# Generate timestamp
TIMESTAMP=$(date "+%Y-%m-%d %H:%M:%S")

# Create markdown report
cat > COVERAGE_REPORT.md << EOF
# Test Coverage Report

**Generated:** $TIMESTAMP  
**Overall Coverage:** $OVERALL_COVERAGE

## Package Coverage Summary

\`\`\`
$PACKAGE_COVERAGE
\`\`\`

## Detailed Function Coverage

\`\`\`
$DETAILED_COVERAGE
\`\`\`

## Coverage by Package

EOF

# Parse package coverage and add to markdown
echo "$PACKAGE_COVERAGE" | while read line; do
    if [[ $line == *"coverage:"* ]]; then
        PACKAGE=$(echo "$line" | awk '{print $2}' | cut -d'/' -f4-)
        COVERAGE=$(echo "$line" | awk '{print $4}')
        PERCENTAGE=$(echo "$COVERAGE" | tr -d '%' | cut -d'.' -f1)
        
        if [ "$PERCENTAGE" -ge 80 ] 2>/dev/null; then
            STATUS="🟢 Excellent"
        elif [ "$PERCENTAGE" -ge 60 ] 2>/dev/null; then
            STATUS="🟡 Good"
        elif [ "$PERCENTAGE" -ge 40 ] 2>/dev/null; then
            STATUS="🟠 Fair"
        else
            STATUS="🔴 Poor"
        fi
        
        echo "- **$PACKAGE**: $COVERAGE $STATUS" >> COVERAGE_REPORT.md
    fi
done

cat >> COVERAGE_REPORT.md << EOF

## How to View Coverage

### HTML Report
\`\`\`bash
make test-coverage
open coverage.html
\`\`\`

### Command Line
\`\`\`bash
go tool cover -func=coverage.out
\`\`\`

### Package Summary
\`\`\`bash
go test -cover ./...
\`\`\`

## Coverage Goals

| Level | Percentage | Status |
|-------|------------|--------|
| Excellent | 80%+ | 🟢 |
| Good | 60-79% | 🟡 |
| Fair | 40-59% | 🟠 |
| Poor | <40% | 🔴 |

EOF

echo "✅ Coverage report generated: COVERAGE_REPORT.md"
echo "📊 Overall coverage: $OVERALL_COVERAGE"

# Generate HTML report
go tool cover -html=coverage.out -o coverage.html
echo "🌐 HTML report generated: coverage.html"