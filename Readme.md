# Sprinckle

Sprinckle is a cmd util that takes a keyword as parameter from Stdin and generates pseudo-random outputs by either prefixing or postfixing it with eligible words contained in a mysql db.

## Getting Started

Start a mysql instance with docker:

```
$ docker run -p 3306:3306 --name mysql-instance -e MYSQL_ROOT_PASSWORD=<PASSWORD> -d mysql:latest
```

Log-In to the db instance and create a database named "sprinckle"

```
$ docker exec -it mysql-instance bash
mysql> mysql -uroot -p<PASSWORD>
mysql> create database sprinckle;
```

Create a table named "complements"

```
mysql> use sprinckle;
mysql> create table complements (ID INT, value VARCHAR(255), prefix BOOLEAN, suffix BOOLEAN );
```

Insert some data to the table

```
mysql> insert into complements (value, prefix, suffix) VALUES ('hq', TRUE, FALSE);
```

build and run

```
$ cd <sprinckle folder>
$ go build -o sprincle
$ echo "test" | ./sprinckle
```

### Prerequisites

MySQL  
### Installing

```
go get github.com/rbroggi/sprinckle
```

## Dependencies/Built With

* [MySQL](https://www.mysql.com/) - Database
* [Golang](https://golang.org/) - Golang
* [Docker](https://www.docker.com/) - Docker, container

## Contributing

## Authors

* **Rodrigo Broggi** - *Initial work* - [Rodrigo Broggi](https://github.com/rbroggi)

## Acknowledgments

This small cmd util is a small enhancement of the one found in:

* [GO BLUEPRINTS](https://github.com/matryer/goblueprints) - (https://www.amazon.com/Go-Programming-Blueprints-Mat-Ryer-ebook/dp/B01GQCQ8OW)