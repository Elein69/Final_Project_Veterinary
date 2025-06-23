 # Use the official Node.js base image with Alpine for smaller size
FROM node:18-alpine

# Set working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install only production dependencies (optional: use 'npm install' for dev too)
RUN npm install

# Copy the rest of the application source code
COPY . .

# Build TypeScript (optional if you precompile)
# RUN npm run build

# Expose the port your app listens on
EXPOSE 8090

# Start the application using ts-node-dev or node (adjust for your entry point)
CMD ["npx", "ts-node-dev", "src/app.ts"]
