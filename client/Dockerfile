FROM node:20 as dev

WORKDIR /app
COPY . ./
RUN npm install -g pnpm@8
RUN pnpm install

EXPOSE 3000
CMD ["pnpm", "run", "dev"]


FROM node:20 as builder

WORKDIR /app
COPY --chown=node:node package.json pnpm-lock.yaml ./
RUN npm install -g pnpm@8
RUN pnpm install --prod --frozen-lockfile


FROM gcr.io/distroless/nodejs20-debian11:latest
ENV NODE_ENV=production

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /dist/app.js ./app.js

USER nonroot
EXPOSE 3000

CMD ["app.js"]

