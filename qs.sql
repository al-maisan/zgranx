SELECT ts, ba.symbol, qa.symbol, ds.name, open, close, period, q_volume
FROM ohlc, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
ORDER BY ts DESC
LIMIT 10;

SELECT ba.symbol, qa.symbol, ds.name, period, year(ts) as year, count(*) as count
FROM ohlc, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
GROUP BY ba.symbol, qa.symbol, ds.name, period, year(ts)
ORDER BY 6 DESC

SELECT ts, ba.symbol, qa.symbol, ds.name, price, q_volume
FROM price, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
ORDER BY ts DESC
LIMIT 10;

SELECT ds.name, year(ts) as year, count(*) as count
FROM ohlc, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
GROUP BY ds.name, year(ts)
ORDER BY 3 DESC


SELECT close from ohlc WHERE data_source_id=3 and base=1 and quote=10 order by ts desc limit 10;

SELECT ba.symbol, ds.name, period, year(ts) as year, count(*) as count
FROM ohlc, asset ba, data_source ds
WHERE ba.id=base AND ds.id=data_source_id
GROUP BY ba.symbol, ds.name, period, year(ts)
ORDER BY 5 DESC

SELECT ts, ba.symbol, quote, ds.name, open, close, period, q_volume
FROM ohlc, asset ba, data_source ds
WHERE ba.id=base AND ds.id=data_source_id
ORDER BY ts DESC
LIMIT 10;
