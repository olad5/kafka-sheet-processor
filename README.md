
# Kafka Sheets Proccessor

This demo showcases the use of Kafka as Message Broker in Golang




## 1 ) How to start the dependencies
```
git clone https://github.com/olad5/kafka-sheet-processor
cd kafka-sheet-processor

#to start the kafka and zookeeper docker dependencies
docker compose up -d 

```
## 2 ) Start the Consumer
```
#In terminal 1, run
make run.consumer
```

## 3 ) Start the Producer

```
#In terminal 2, run
make run.producer
```

## 4 ) View the results

```
$ cat result.json
```





##  Resources that were helpful 
https://levelup.gitconnected.com/introduction-to-kafka-in-go-2a5755df504c



## Acknowledgements
- [Ryo Kusnadi](https://medium.com/@ryokusnadi_20) I basically learnt kafka using his article above
