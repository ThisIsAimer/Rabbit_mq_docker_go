## go packages
- github.com/rabbitmq/amqp091-go

## docker rabbit mq cmds
- docker pull rabbitmq:management
- docker run -d --name rabbitmq --hostname rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
- docker restart <name>
- docker container ls -a (shows all containers)
- docker ps (shows running containers)
- docker stop <name> (stops running container)
- docker rm <name> (removes container)

## run cmdline interface inside a docker container
- docker exec -it rabbitmq bash

## creating docker networks
- docker network ls
- docker network create rabbitmq-cluster
- docker network connect <network name> <container name> (connect)
- docker run -d --name rabbitmq1 --hostname rabbitmq1 --network rabbitmq-cluster -p 5673:5672 -p 15673:15672 rabbitmq:management

- docker network inspect <network name>

### rabbitmq cookies
- docker exec -it <name> cat /var/lib/rabbitmq/.erlang.cookie
- docker cp rabbitmq:/var/lib/rabbitmq/.erlang.cookie /tmp/.erlang.cookie (copies the cookie into temp file)
- docker cp /tmp/.erlang.cookie rabbitmq1:/var/lib/rabbitmq/.erlang.cookie

- docker exec -it rabbitmq chmod 400 /var/lib/rabbitmq/.erlang.cookie (precautionary measure) (after copting cookie)

- docker exec -it rabbitmq1 rabbitmqctl join_cluster rabbit@8e27993dac6 (joins the actual cluster hostname)

- docker exec -it rabbitmq1 rabbitmqctl cluster_status

### make ram node

- docker exec -it rabbitmq1 rabbitmqctl stop_app

- docker exec -it rabbitmq1 rabbitmqctl reset

- docker exec -it rabbitmq1 rabbitmqctl join_cluster --ram rabbit@8e27993dac6

- docker exec -it rabbitmq1 rabbitmqctl start_app

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