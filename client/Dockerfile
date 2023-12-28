FROM node:21.2-bookworm as dev
WORKDIR /app
COPY . ./
RUN npm install -g pnpm
RUN pnpm install
EXPOSE 8000
CMD ["pnpm", "run", "dev"]

