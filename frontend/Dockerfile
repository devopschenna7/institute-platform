# Use Node.js
FROM node:22-alpine

# Application directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy application source code
COPY . .

# Build the React application
RUN npm run build

# Install Nginx
RUN apk add --no-cache nginx

# Remove default Nginx configuration
RUN rm -f /etc/nginx/http.d/default.conf

# Copy our Nginx configuration
COPY nginx.conf.template /etc/nginx/http.d/default.conf.template

# Copy React build files to Nginx web directory
RUN mkdir -p /usr/share/nginx/html && \
    cp -r dist/* /usr/share/nginx/html/

# Copy startup script
COPY entrypoint.sh /entrypoint.sh

# Give execute permission
RUN chmod +x /entrypoint.sh

# Application listens on port 80
EXPOSE 80

# Start the application
CMD ["/entrypoint.sh"]
