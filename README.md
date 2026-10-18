# TarsOWASPAgentic

MCP integration for OWASP Agentic AI security toolkit.

## Tools

- `owasp_scan` — Run a security scan on a target
- `owasp_status` — Get scan status and results

## Quick Start

```bash
go build -o tarsowaspagentic
TARS_MCP_ENDPOINT=https://api.example.com TARS_MCP_API_KEY=xxx ./tarsowaspagentic
```

## Environment Variables

- `TARS_MCP_ENDPOINT` — upstream API endpoint
- `TARS_MCP_API_KEY` — upstream API key
- `TARS_MCP_MODE` — `stdio` (default) or `http`

## License

MIT
