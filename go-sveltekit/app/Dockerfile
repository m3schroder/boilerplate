# Multi part build slims down the final image size
FROM oven/bun
WORKDIR /app
COPY package.json .
COPY bun.lockb .
RUN bun install
COPY . .
RUN bun run build
EXPOSE 3000
ENTRYPOINT ["bun", "./build"]