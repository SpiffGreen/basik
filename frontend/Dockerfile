FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package.json package-lock.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Expose default Vite preview port
EXPOSE 4173

# Start the Vite preview server
CMD ["npm", "run", "preview", "--", "--host"]
