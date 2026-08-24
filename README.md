# Static Links

CLI tool written in Go that generates a Linktree-esque HTML page from a YAML file.

## Installation and Usage

1. Clone this repo.

```bash
git clone https://github.com/dmfcat/static-links
```

2. Run the script.

```bash
cd static-links
go run .
```

3. Output automatically gets generated in `public/`.

## Configuration

You can configure what links are generated and the colour scheme by editing the variables in [`config.yaml`](config.yaml).

If the config file is missing, a default one is automatically generated upon running the script.

## Licence

[MIT](LICENCE)
