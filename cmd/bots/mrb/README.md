# mean reversion bot

* entry condition: SMA
* exit conditions:
   * take profit
   * stoploss

# parameters

* SMA length
* fund allocation (in quote asset (usdt))
* base order limit
* extra orders
* trading frequency (period, e.g. 1H or 5M)
* take profit percentage
* stoploss percentage

# commands

* start(paper-trading?): starts the bot -- if the `paper-trading` is set the bot will not place actual orders but merely log what actions would be taken
* stop: stops the bot
* backtest(lookback-interval); prints what actual would be take in the `lookback-interval` and what profits/losses would be generated

# design

The bot operates a main loop and wakes up every `period` -- it will then check whether

1. any of the exit conditions are met for the open positions, and place orders to close these if/as needed
1. an entry condition is met and place an order to open a position if it is still inside the configured fund allocation budget

alternatively: if the entry condition is met and the bot has funds to trade: place the following orders:

1. open positions/buy order
1. take profit/sell order
1. stoploss/sell order

when the bot wakes up it just needs to check whether it has funds left to trade and repeat the procedure above

# references

1. [Mean Reversion Trading Strategy Using Python](https://medium.com/coinmonks/mean-reversion-trading-strategy-using-python-4cfecb51859e)
1. [How to Build your First Mean Reversion Trading Strategy in Python](https://raposa.trade/blog/how-to-build-your-first-mean-reversion-trading-strategy-in-python/)
