SELECT CAST(0.85 * COUNT(5p_ratio) AS INTEGER)  FROM m5_data WHERE 5p_ratio < 1.0;
SELECT ts, close, 5p_ratio from m5_data where 5p_ratio < 1 order by 5p_ratio limit 245153,1;
SELECT CAST(0.15 * COUNT(5p_ratio) AS INTEGER)  FROM m5_data WHERE 5p_ratio > 1.0;
SELECT ts, close, 5p_ratio from m5_data where 5p_ratio > 1 order by 5p_ratio limit 45428,1;

DROP TABLE IF EXISTS m5_data;
CREATE TABLE m5_data (
     ts TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
     base MEDIUMINT NOT NULL,
     quote MEDIUMINT NOT NULL,
     data_source_id MEDIUMINT NOT NULL,
     close DECIMAL(30,8),
     5p_avg DECIMAL(30,8),
     8p_avg DECIMAL(30,8),
     13p_avg DECIMAL(30,8),
     5p_ratio DECIMAL(30,8),
     8p_ratio DECIMAL(30,8),
     13p_ratio DECIMAL(30,8),
     period VARCHAR(8) NOT NULL DEFAULT '5M',
     FOREIGN KEY (base) REFERENCES asset (id),
     FOREIGN KEY (quote) REFERENCES asset (id),
     FOREIGN KEY (data_source_id) REFERENCES data_source (id),
     INDEX(ts, period),
     INDEX(base, quote, period, ts),
     unique(ts, base, quote, period, data_source_id)
 );

INSERT INTO m5_data(ts, base, quote, data_source_id, close, 5p_avg, 8p_avg, 13p_avg, period)
SELECT ts, base, quote, data_source_id, close, 
       avg(close) over (order by ts rows between 4 preceding and current row) as 5p_avg,
       avg(close) over (order by ts rows between 7 preceding and current row) as 8p_avg,
       avg(close) over (order by ts rows between 12 preceding and current row) as 13p_avg,
      '5M'
from ohlc where base=1 and data_source_id=2 and period='1M' and UNIX_TIMESTAMP(ts)%300=0 ;

UPDATE m5_data
SET
5p_ratio = close/5p_avg,
8p_ratio = close/8p_avg,
13p_ratio = close/13p_avg;

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
