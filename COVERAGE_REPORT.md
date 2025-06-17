# Test Coverage Report

**Generated:** 2025-06-17 15:05:10  
**Overall Coverage:** 50.6%

## Package Coverage Summary

```
ok  	github.com/guntharp/go-filamentsamples/pkg/models	(cached)	coverage: 75.6% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/openscad	(cached)	coverage: 64.1% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/generator	(cached)	coverage: 22.2% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/csv	(cached)	coverage: 77.3% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/config	(cached)	coverage: 89.3% of statements
ok  	github.com/guntharp/go-filamentsamples/cmd/filament-samples	(cached)	coverage: 0.0% of statements
```

## Detailed Function Coverage

```
github.com/guntharp/go-filamentsamples/cmd/filament-samples/main.go:18:		main			0.0%
github.com/guntharp/go-filamentsamples/cmd/filament-samples/main.go:80:		showHelp		0.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:21:		LoadConfig		100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:35:		loadFromFile		100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:44:		SaveToFile		71.4%
github.com/guntharp/go-filamentsamples/internal/config/config.go:58:		Validate		100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:74:		DefaultConfigPath	75.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:82:		ExampleConfig		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:17:		NewParser		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:23:		ParseFile		0.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:33:		Parse			90.9%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:73:		isHeaderRow		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:82:		parseRecord		75.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:24:	Validate		100.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:52:	NewGenerator		81.8%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:78:	Generate		0.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:110:	processParallel		0.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:157:	worker			0.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:169:	generateSample		0.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:15:	NewExecutor		75.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:27:	GenerateSTL		88.9%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:43:	findOpenSCADPath	38.9%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:82:	CheckAvailable		100.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:89:	GetVersion		80.0%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:20:		Validate		66.7%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:47:		validateTemperature	81.2%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:77:		Filename		100.0%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:82:		OpenSCADArgs		75.0%
total:										(statements)		50.6%
```

## Coverage by Package

- **pkg/models**: coverage: 🔴 Poor
- **internal/openscad**: coverage: 🔴 Poor
- **internal/generator**: coverage: 🔴 Poor
- **internal/csv**: coverage: 🔴 Poor
- **internal/config**: coverage: 🔴 Poor
- **cmd/filament-samples**: coverage: 🔴 Poor

## How to View Coverage

### HTML Report
```bash
make test-coverage
open coverage.html
```

### Command Line
```bash
go tool cover -func=coverage.out
```

### Package Summary
```bash
go test -cover ./...
```

## Coverage Goals

| Level | Percentage | Status |
|-------|------------|--------|
| Excellent | 80%+ | 🟢 |
| Good | 60-79% | 🟡 |
| Fair | 40-59% | 🟠 |
| Poor | <40% | 🔴 |

