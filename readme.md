## docker rabbit mq cmds
- docker pull rabbitmq:management
- docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
- docker container ls -a (shows all containers)
- docker ps (shows running containers)
- docker stop <name> (stops running container)
- docker rm <name> (removes container)

## run cmdline interface inside a docker container
- docker exec -it rabbitmq bash
### rabbitmq cmds
- rabbitmqctl status (status of rabbit mq instance)
- rabbitmqctl list_queues name type leader members
- rabbitmqctl list_exchanges
- rabbitmqctl list_users

- rabbitmq-plugins enable rabbitmq_management
- rabbitmq-plugins list

### extra plugins
- apt install -y wget

- wget -O /plugins/rabbitmq_delayed_message_exchange.ez \ https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/v4.1.0/rabbitmq_delayed_message_exchange-4.1.0.ez