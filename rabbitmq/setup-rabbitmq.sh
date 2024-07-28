#!/bin/bash
# setup-rabbitmq.sh

# Config variable
RABBITMQ_USER=${RABBITMQ_USER:-rabbitmq_admin}
RABBITMQ_PASSWORD=${RABBITMQ_PASSWORD:-my_password}
RABBITMQ_HOST=${RABBITMQ_HOST:-localhost}
RABBITMQ_PORT=${RABBITMQ_PORT:-15672}
WAIT_TIME=${WAIT_TIME:-10}
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'

# Functions
log() {
    echo -e "${GREEN}$(date +'%Y-%m-%d %H:%M:%S') - $1${CYAN}"
}

log_error() {
    echo -e "${RED}$(date +'%Y-%m-%d %H:%M:%S') - ERROR: $1${CYAN}"
}

check_success() {
    if [ $? -ne 0 ]; then
        log_error "$1"
        exit 1
    fi
}

log "Waiting for RabbitMQ to start on ${RABBITMQ_HOST}:${RABBITMQ_PORT}..."
while ! curl -s http://${RABBITMQ_HOST}:${RABBITMQ_PORT}/api/overview > /dev/null; do
    log "RabbitMQ not yet available. Waiting..."
    sleep ${WAIT_TIME}
done

log "RabbitMQ is up and running."

# Create admin user
log "Creating RabbitMQ admin user..."
rabbitmqctl add_user $RABBITMQ_USER $RABBITMQ_PASSWORD
check_success "Failed to create RabbitMQ user."

rabbitmqctl set_user_tags $RABBITMQ_USER administrator
check_success "Failed to set user tags."

rabbitmqctl set_permissions -p / $RABBITMQ_USER ".*" ".*" ".*"
check_success "Failed to set permissions."

# Create exchanges
log "Declaring exchanges..."
rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare exchange name=order_exchange type=direct
check_success "Failed to declare exchange order_exchange."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare exchange name=product_exchange type=topic
check_success "Failed to declare exchange product_exchange."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare exchange name=notification_exchange type=fanout
check_success "Failed to declare exchange notification_exchange."

# Create queues
log "Declaring queues..."
rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare queue name=order_created
check_success "Failed to declare queue order_created."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare queue name=order_paid
check_success "Failed to declare queue order_paid."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare queue name=inventory_update
check_success "Failed to declare queue inventory_update."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare queue name=email_notifications
check_success "Failed to declare queue email_notifications."

# Create bindings
log "Declaring bindings..."
rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare binding source=order_exchange destination=order_created routing_key=order.created
check_success "Failed to declare binding for order_created."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare binding source=order_exchange destination=order_paid routing_key=order.paid
check_success "Failed to declare binding for order_paid."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare binding source=product_exchange destination=inventory_update routing_key=product.added
check_success "Failed to declare binding for product.added."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare binding source=product_exchange destination=inventory_update routing_key=product.updated
check_success "Failed to declare binding for product.updated."

rabbitmqadmin -u $RABBITMQ_USER -p $RABBITMQ_PASSWORD declare binding source=notification_exchange destination=email_notifications
check_success "Failed to declare binding for email_notifications."

log "The RabbitMQ setup was completed successfully."
