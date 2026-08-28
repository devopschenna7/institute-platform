#!/bin/sh

if [ -z "$BACKEND_URL" ]; then
    echo "WARNING: BACKEND_URL is not set."
    echo "Backend API requests will not work."

    BACKEND_URL="http://127.0.0.1:9"
fi

echo "Using backend: $BACKEND_URL"

sed "s|\${BACKEND_URL}|${BACKEND_URL}|g" \
    /etc/nginx/conf.d/default.conf.template \
    > /etc/nginx/conf.d/default.conf

#echo "Generated Nginx configuration:"
#cat /etc/nginx/http.d/default.conf

exec nginx -g "daemon off;"
