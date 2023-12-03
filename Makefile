KAFKA_CONTID = $(shell docker ps -qf "name=kafka_cont")
DB_CONTID=  $(shell docker ps -qf "name=db_cont")
PG_ADMIN_CONTID=	$(shell docker ps -qf "name=pgadmin_cont")

list-topics:
	docker exec  ${KAFKA_CONTID}  kafka-topics.sh --bootstrap-server kafka:9092 --list

consumer: 
	docker exec  ${KAFKA_CONTID} kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic $(topic) --from-beginning

producer:
	docker exec -it ${KAFKA_CONTID} kafka-console-producer.sh --bootstrap-server kafka:9092 --topic $(topic)
#make producer topic=test3

db-ip:
	docker inspect  ${DB_CONTID} | grep IPAddress

pg-ip:
	docker inspect  ${PG_ADMIN_CONTID} | grep IPAddress