FROM oven/bun:latest

WORKDIR /usr/src/app

COPY . /usr/src/app

RUN bun install -g @angular/cli

RUN bun install

CMD ["ng", "serve", "--host", "0.0.0.0"]

