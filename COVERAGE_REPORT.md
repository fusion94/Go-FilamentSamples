# Test Coverage Report

**Generated:** 2025-06-17 15:31:47  
**Overall Coverage:** 80.3%

## Package Coverage Summary

```
ok  	github.com/guntharp/go-filamentsamples/pkg/models	(cached)	coverage: 100.0% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/openscad	(cached)	coverage: 74.4% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/generator	(cached)	coverage: 95.8% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/csv	(cached)	coverage: 95.5% of statements
ok  	github.com/guntharp/go-filamentsamples/internal/config	(cached)	coverage: 96.4% of statements
ok  	github.com/guntharp/go-filamentsamples/cmd/filament-samples	(cached)	coverage: 0.0% of statements
```

## Detailed Function Coverage

```
github.com/guntharp/go-filamentsamples/cmd/filament-samples/main.go:18:		main			0.0%
github.com/guntharp/go-filamentsamples/cmd/filament-samples/main.go:80:		showHelp		0.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:21:		LoadConfig		100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:35:		loadFromFile		100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:44:		SaveToFile		85.7%
github.com/guntharp/go-filamentsamples/internal/config/config.go:58:		Validate		100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:74:		DefaultConfigPath	100.0%
github.com/guntharp/go-filamentsamples/internal/config/config.go:82:		ExampleConfig		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:17:		NewParser		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:23:		ParseFile		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:33:		Parse			90.9%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:73:		isHeaderRow		100.0%
github.com/guntharp/go-filamentsamples/internal/csv/parser.go:82:		parseRecord		100.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:24:	Validate		100.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:52:	NewGenerator		81.8%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:78:	Generate		100.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:110:	processParallel		96.4%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:157:	worker			100.0%
github.com/guntharp/go-filamentsamples/internal/generator/generator.go:169:	generateSample		100.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:15:	NewExecutor		100.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:27:	GenerateSTL		100.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:43:	findOpenSCADPath	44.4%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:82:	CheckAvailable		100.0%
github.com/guntharp/go-filamentsamples/internal/openscad/executor.go:89:	GetVersion		100.0%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:20:		Validate		100.0%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:47:		validateTemperature	100.0%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:77:		Filename		100.0%
github.com/guntharp/go-filamentsamples/pkg/models/filament.go:82:		OpenSCADArgs		100.0%
total:										(statements)		80.3%
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

