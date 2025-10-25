# Stage 1: Builder
FROM node:22 AS builder

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Final Runtime
FROM node:22-alpine
WORKDIR /app
COPY package*.json ./

COPY --from=builder /app/dist ./dist

RUN npm install --production

EXPOSE 3000
CMD ["npm", "start"]
