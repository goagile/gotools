

REDIS
=====
	KV Storage
	Однопоточный сервер
	Riplication
	Dump
	Transactional
	Packet commands
	PubSub
 
Application
	Cache
	Sessions and Users
	Query Server (PubSub)
	Captures, Flags, Search queries Store
	Sharding Dict

Install
	wget http://download.redis.io/releases/redis-4.0.11.tar.gz
	tar xzf redis-4.0.11.tar.gz
	cd redis-4.0.11
	make

Run
	src/redis-server

CLI
	src/redis-cli

SetGet
	SET server:name "fido"	
	GET server:name => "fido"
	SET books 10
	GET books => "10"

IncrDecrDel
	INCR books => (integer) 11
	DEL books => (integer) 1
	INCR books => 1	
	DECR books => 0

Lock
  	SET resource:lock "Redis Demo 1"
    EXPIRE resource:lock 120
    TTL resource:lock => 119
    SET resource:lock "Redis Demo 2"
    TTL resource:lock => -1

Stack
	RPUSH friends "A"
	RPOP friends => "A"
	LPUSH friends "B"
	LPOP friends => "B"

List
	RPUSH friends "Alex"
	RPUSH friends "Bob"
	RPUSH friends "Cris"
	LRANGE 0 -1 => 1) "Alex" 2) "Bob" 3) "Cris"
	LLEN friends => 3

Set
	SADD names "Alex" => 1
	SREM names "Alex" => 1
	SISMEMBER names "Alex" => 1
	SMEMBERS name => 1) "Alex"
	SADD a "A"
	SADD b "B"
	SUNION a b => 1) "A" 2) "B"
 
 Hash
 	HSET user:0 name "A"
 	HGET user:0 name => "A"
 	HGETALL user:0
	=>	1) "name"
		2) "A"
	HMSET user:0 name "A" emali "a@gmail.com"
