# mean reversion bot

* entry condition: SMA
* exit conditions:
   * take profit
   * stoploss

# parameters

* fund allocation
* base order limit
* extra orders
* trading frequency (period, e.g. 1H or 5M)
* take profit percentage
* stoploss percentage

# commands

* start
* stop
* backtest(look-back-interval)

# design

The bot operates a main loop and wakes up every `period` -- it will then check whether

1. any of the exit conditions are met for the open positions, and place orders to close these if/as needed
1. an entry condition is met and place an order to open a position if it is still inside the configured fund allocation budget

# references

1. [Mean Reversion Trading Strategy Using Python](https://medium.com/coinmonks/mean-reversion-trading-strategy-using-python-4cfecb51859e)
1. [How to Build your First Mean Reversion Trading Strategy in Python](https://raposa.trade/blog/how-to-build-your-first-mean-reversion-trading-strategy-in-python/)
