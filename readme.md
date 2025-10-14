## docker rabbit mq cmds
- docker pull rabbitmq:management
- docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
- docker container ls -a (shows all containers)
- docker ps (shows running containers)
- docker stop <name> (stops running container)
- docker rm <name> (removes container)