FROM rabbitmq:management

USER root

RUN apt-get update && apt-get install -y curl python3

RUN ln -s /usr/bin/python3 /usr/bin/python

RUN curl -o /usr/local/bin/rabbitmqadmin https://raw.githubusercontent.com/rabbitmq/rabbitmq-management/v3.8.2/bin/rabbitmqadmin
RUN chmod +x /usr/local/bin/rabbitmqadmin

COPY setup-rabbitmq.sh /setup-rabbitmq.sh
RUN chmod +x /setup-rabbitmq.sh

CMD ["sh", "-c", "/setup-rabbitmq.sh & rabbitmq-server"]
