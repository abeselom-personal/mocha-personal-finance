# services/gateway/Dockerfile
FROM nginx:1.25-alpine

# Install envsubst for template processing
RUN apk add --no-cache gettext

# Remove default config
RUN rm /etc/nginx/conf.d/default.conf

# Copy template configuration
COPY nginx.conf.template /etc/nginx/templates/

# Copy SSL certificates (if available)
COPY certs/ /etc/nginx/certs/

EXPOSE ${GATEWAY_PORT_HTTP} ${GATEWAY_PORT_HTTPS}

CMD ["/bin/sh", "-c", "envsubst '$$AUTH_PORT $$NOTIFICATION_PORT $$RECEIPT_SCRAPER_PORT $$REGEX_PORT $$SMS_PARSER_PORT $$SYNC_PORT' < /etc/nginx/templates/nginx.conf.template > /etc/nginx/nginx.conf && exec nginx -g 'daemon off;'"]
