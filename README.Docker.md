### Run Container

```bash
# Run with environment variables
docker run -d \
  --name vkr-backend \
  -p 8081:8081 \
  -e POSTGRES_HOST=your-postgres-host \
  -e POSTGRES_USER=your-user \
  -e POSTGRES_PASSWORD=your-password \
  -e POSTGRES_DB=NIR \
  -e PORT=8081 \
  vkr-backend:latest
```
