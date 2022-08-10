SELECT ts, ba.symbol, qa.symbol, ds.name, open, close, period, q_volume
FROM ohlc, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
ORDER BY ts DESC
LIMIT 10;

SELECT ba.symbol, qa.symbol, ds.name, period, year(ts), count(*)
FROM ohlc, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
GROUP BY ba.symbol, qa.symbol, ds.name, period, year(ts)
ORDER BY 6 DESC

SELECT ts, ba.symbol, qa.symbol, ds.name, price, q_volume
FROM price, asset ba, asset qa, data_source ds
WHERE ba.id=base AND qa.id=quote AND ds.id=data_source_id
ORDER BY ts DESC
LIMIT 10;
