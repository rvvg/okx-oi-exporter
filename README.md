# OKX Open Interest Exporter

A Prometheus exporter that collects open interest data from OKX cryptocurrency exchange.

## Features

- Collects SWAP open interest data from OKX API
- Exposes metrics in Prometheus format
- Auto-updates metrics every 5 seconds
- Simple HTTP server on port 8080

## Metrics
```
okx_open_interest{instId="<instrument>",instType="<type>"} <value>
```
- `instId`: The instrument identifier
- `instType`: The instrument type (SWAP)
- `value`: Open interest value in USD

## Usage

```
# Clone repository
git clone git@github.com:rvvg/okx-oi-exporter.git

# Build the binary
cd okx-oi-exporter
go build -o okx-oi-exporter

# Run the exporter
./okx-oi-exporter
```

## Configuration

The exporter uses these default settings:
- Port: 8080
- Endpoint: `/metrics`
- Update interval: 5 seconds
- OKX API endpoint: `https://www.okx.com/api/v5/public/open-interest?instType=SWAP`

## License

MIT License

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.