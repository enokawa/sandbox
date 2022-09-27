## Launch MySQL

```bash
docker compose up -d
```

## Connect to MySQL

```bash
docker compose exec db sh
mysql -u root -pexample -D example
```

or

```bash
docker compose exec db mysql -u root -pexample -D example
```

## Shutdown MySQL

```bash
docker compose down
```