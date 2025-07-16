# Builder stage
FROM python:3.12-alpine as builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache build-base

# Create virtual environment
RUN python -m venv /opt/venv

# Install dependencies
COPY requirements.txt .
RUN /opt/venv/bin/pip install --no-cache-dir --upgrade pip && \
    /opt/venv/bin/pip install --no-cache-dir -r requirements.txt

# Runtime stage
FROM python:3.12-alpine

WORKDIR /app

# Copy virtual environment from builder
COPY --from=builder /opt/venv /opt/venv

# Copy application files
COPY . .

# Environment variables
ENV PATH="/opt/venv/bin:$PATH" \
    PYTHONUNBUFFERED=1 \
    PYTHONPATH="/app"

# Create non-root user
RUN adduser -D appuser && \
    chown -R appuser:appuser /app
USER appuser

# Health check (adjust as needed)
HEALTHCHECK --interval=30s --timeout=10s \
    CMD python -c "import requests; requests.get('http://localhost:8000/health', timeout=5)" || exit 1

# Entrypoint
ENTRYPOINT ["python", "bot.py"]