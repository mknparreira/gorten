FROM kong:latest

USER root

ENV KONG_DATABASE=off
ENV KONG_DECLARATIVE_CONFIG=/usr/local/kong/declarative/kong.yml

COPY kong.yml /usr/local/kong/declarative/kong.yml

CMD ["kong", "docker-start"]
