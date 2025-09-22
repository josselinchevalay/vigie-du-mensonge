# --- Stage 1: Build with Node ---
FROM node:24.8.0-alpine AS builder

# Avoid running as root where possible (builder still needs perms to write node_modules)
WORKDIR /app

# Install only what is needed for build in a clean environment
COPY package.json package-lock.json ./
RUN npm ci --no-audit --no-fund

# Copy the rest of the source
COPY . .

# Build-time configuration that Vite will embed (override via --build-arg)
# Use a relative API base so frontend and backend can share the same domain, e.g. https://example.com/api
ARG VITE_API_BASE="/api"
ENV VITE_API_BASE="$VITE_API_BASE"

# Build the production assets
RUN npm run build

# --- Stage 2: Run with unprivileged nginx ---
# nginxinc/nginx-unprivileged listens on 8080 by default and runs as non-root (uid 101)
FROM nginxinc/nginx-unprivileged:1.27-alpine

# Security: ensure a clean image and minimal packages
# (Base image is already minimal; avoid adding extra packages.)

# Copy build output
COPY --from=builder /app/dist /usr/share/nginx/html

# Provide hardened nginx configuration (includes SPA fallback and headers)
COPY ops/nginx/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

# Healthcheck hits a cheap endpoint served by nginx
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -qO- http://127.0.0.1:8080/healthz || exit 1

# Final user is already non-root (uid 101) in this base image.