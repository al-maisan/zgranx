DROP TABLE IF EXISTS asset;
CREATE TABLE asset (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     name VARCHAR(30) NOT NULL,
     symbol VARCHAR(30) NOT NULL,
     decimals INT UNSIGNED NOT NULL DEFAULT 0,
     asset_type ENUM('crypto', 'fiat') NOT NULL DEFAULT 'crypto',
     PRIMARY KEY (id),
     unique(name),
     unique(symbol)
 );

-- please note: we use the coingecko symbol for the asset/coin in the `name`
-- column
INSERT INTO asset(name, symbol, decimals) VALUES('bitcoin', 'btc', 8);
INSERT INTO asset(name, symbol, decimals) VALUES('ethereum', 'eth', 18);
INSERT INTO asset(name, symbol, decimals) VALUES('bnb', 'bnb', 8);
INSERT INTO asset(name, symbol, decimals) VALUES('cardano', 'ada', 6);
INSERT INTO asset(name, symbol, decimals) VALUES('solana', 'sol', 9);
INSERT INTO asset(name, symbol, decimals) VALUES('polkadot', 'dot', 10);
INSERT INTO asset(name, symbol, decimals) VALUES('avalanche', 'avax', 9);
INSERT INTO asset(name, symbol, decimals) VALUES('polygon', 'matic', 18);
INSERT INTO asset(name, symbol, decimals) VALUES('litecoin', 'ltc', 8);
INSERT INTO asset(name, symbol, decimals) VALUES('tether', 'usdt', 6);
INSERT INTO asset(asset_type, name, symbol, decimals) VALUES('fiat', 'usd', 'usd', 2);
INSERT INTO asset(asset_type, name, symbol, decimals) VALUES('fiat', 'eur', 'eur', 2);
INSERT INTO asset(asset_type, name, symbol, decimals) VALUES('fiat', 'jpy', 'jpy', 0);
INSERT INTO asset(asset_type, name, symbol, decimals) VALUES('fiat', 'chf', 'chf', 2);
INSERT INTO asset(asset_type, name, symbol, decimals) VALUES('fiat', 'cad', 'cad', 2);
INSERT INTO asset(asset_type, name, symbol, decimals) VALUES('fiat', 'krw', 'krw', 0);
INSERT INTO asset(name, symbol, decimals) VALUES('usd-coin', 'usdc', 6);


DROP TABLE IF EXISTS data_source;
CREATE TABLE data_source (
   id MEDIUMINT NOT NULL AUTO_INCREMENT,
   name VARCHAR(128) NOT NULL,
   uri VARCHAR(128) NOT NULL,
   api_key VARCHAR(128),
   api_secret VARCHAR(128),
   PRIMARY KEY (id),
   unique(name),
   unique(uri)
);
INSERT INTO data_source(name, uri) VALUES('coingecko', 'https://www.coingecko.com/');
INSERT INTO data_source(name, uri) VALUES('bitstamp', 'https://www.bitstamp.net/');
INSERT INTO data_source(name, uri) VALUES('binance', 'https://www.binance.com/en');
INSERT INTO data_source(name, uri) VALUES('bitfinex', 'https://www.bitfinex.com/');
INSERT INTO data_source(name, uri) VALUES('ftx', 'https://www.ftx.com/');
INSERT INTO data_source(name, uri) VALUES('gate_io', 'https://www.gate.io/');
INSERT INTO data_source(name, uri) VALUES('huobi', 'https://www.huobi.com/');


DROP TABLE IF EXISTS pair;
CREATE TABLE pair (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     base MEDIUMINT NOT NULL,
     quote MEDIUMINT NOT NULL,
     symbol VARCHAR(30) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (base) REFERENCES asset (id),
     FOREIGN KEY (quote) REFERENCES asset (id),
     unique(symbol)
 );
INSERT INTO pair(base, quote, symbol) VALUES(1,10,'btc-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(2,10,'eth-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(3,10,'bnb-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(4,10,'ada-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(5,10,'sol-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(6,10,'dot-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(7,10,'avax-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(8,10,'matic-usdt');
INSERT INTO pair(base, quote, symbol) VALUES(9,10,'ltc-usdt');

DROP TABLE IF EXISTS price;
CREATE TABLE price (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     ts TIMESTAMP NOT NULL,
     base MEDIUMINT NOT NULL,
     quote MEDIUMINT NOT NULL,
     data_source_id MEDIUMINT NOT NULL,
     price DECIMAL(30,8) NOT NULL,
     q_volume DECIMAL(36,8) NOT NULL,
     q_volume_change DECIMAL(30,8) NOT NULL,
     -- one of: '1M', '3M', '5M', '15M', '30M', '1H', '2H', '3H', '4H', '1d',
     --         '1w', '1m'
     period VARCHAR(8) NOT NULL DEFAULT '5M',
     PRIMARY KEY (id),
     FOREIGN KEY (base) REFERENCES asset (id),
     FOREIGN KEY (quote) REFERENCES asset (id),
     FOREIGN KEY (data_source_id) REFERENCES data_source (id),
     unique(ts, base, quote, period, data_source_id)
 );

DROP TABLE IF EXISTS ohlc;
CREATE TABLE ohlc (
     id MEDIUMINT NOT NULL AUTO_INCREMENT,
     ts TIMESTAMP NOT NULL,
     base MEDIUMINT NOT NULL,
     quote MEDIUMINT NOT NULL,
     data_source_id MEDIUMINT NOT NULL,
     open DECIMAL(30,8) NOT NULL,
     high DECIMAL(30,8) NOT NULL,
     low DECIMAL(30,8) NOT NULL,
     close DECIMAL(30,8) NOT NULL,
     -- one of: '1M', '3M', '5M', '15M', '30M', '1H', '2H', '3H', '4H', '1d',
     --         '1w', '1m'
     period VARCHAR(8) NOT NULL DEFAULT '5M',
     count INT UNSIGNED,
     q_volume DECIMAL(30,8),
     PRIMARY KEY (id),
     FOREIGN KEY (base) REFERENCES asset (id),
     FOREIGN KEY (quote) REFERENCES asset (id),
     FOREIGN KEY (data_source_id) REFERENCES data_source (id),
     unique(ts, base, quote, period, data_source_id)
 );
