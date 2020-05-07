#Go-microservices example 
## Requirements
- `skaffold` through `brew install skaffold`
- `Docker-desktop`/`minikube`

## How to run
- set up your local cluster with `minikube` or `docker-desktop`. `Docker-desktop` is easiest https://www.techrepublic.com/article/how-to-add-kubernetes-support-to-docker-desktop/
- `skaffold run`
- ctrl+c to delete your pods + namespaces (good cleanup)

## Whats happening
1) kafka + zookeeper are deployed in your local cluster
2) microservices are built and deployed into your local cluster for debugging
3) Look through logs, see we have two services talking to each other through a kafka topic ;)

## Whats next
- use zap.logger
- converting old middlewares
- figure out cleaner way for configs
- All datastore layer (gopg) 

## things to figure out
- how do i set the consumer offset to start somewhere besides latest?
- emitting different events on the same topic (maybe using generic broker.Event to have the interface inside?)

## Credits
https://github.com/d1egoaz/minikube-kafka-cluster

test
