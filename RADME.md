ICMP message 
+--0-7--+--8-15--+--16-23--+--24-31--+
| type  |  code  |      checksum     |
+-------+--------+-------------------+
|               data                 |         
+------------------------------------+

type/code 
----------------------------
0/0 Эхо-ответ
8/0 Эхо-запрос
3/0 Network unrichable
3/1 Host unrichable
11/0 TTL expired in transit
...

data
----------------------------
Sequence number(2 byte) Используется для сопоставления эхо-запросов с соответствующим ответом
Identifier (2 byte) Может использоваться для сопоставления эхо-запросов со связанным ответом


ping
0.0.0.0 (8/0, seq=1) -> host
0.0.0.0 <- (0/0, seq=1) host

traceroute
0.0.0.0 (8/0, seq=1 ttl 1) -> host
0.0.0.0 <- (11/0, seq=1) host

Identifier, Sequence number BE - Прямой порядок байт 
Identifier, Sequence number LE - Обратный порядок байт 


