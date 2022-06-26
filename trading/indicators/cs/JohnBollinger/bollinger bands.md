# Bollinger Bands

https://stockcharts.com/h-mem/tascredirect.html?artid=\V20\C01\015BOLL.pdf
Bollinger Bands by Amy Wu
Dec 2001 - Stocks & Commodities V. 20:1 (78-79)

https://stockcharts.com/h-mem/tascredirect.html?artid=\V10\C02\USINGBO.pdf
Using Bollinger Bands by John Bollinger
Jan 1992 - Stocks & Commodities V. 10:2 (47-51)

https://school.stockcharts.com/doku.php?id=technical_indicators:bollinger_bands

Bollinger Bands® is a registered trademark of John Bollinger.
Bollinger Bands® are volatility bands placed above and below a moving average.
Volatility is based on the standard deviation, which changes as volatility increases and decreases.
The bands automatically widen when volatility increases and contract when volatility decreases.
Their dynamic nature allows them to be used on different securities with the standard settings.

As such, Bollinger Bands® can be used to determine if prices are relatively high or low.
According to Bollinger, the bands should contain 88-89% of price action, which makes a move outside the bands significant.

The default settings recommended by John Bollinger are:
- simple moving average (SMA), because it matches the standard deviation formula
- the moving average period is 20, the look-back period for the standard deviation is the same
- the outer bands are set 2 standard deviations above and below

The default settings can be adjusted to suit the characteristics of particular securities or trading styles.
Bollinger recommends making small incremental adjustments to the standard deviation multiplier: 2.1 for a 50-period moving average and 1.9 for a 10-period one.
 
For signals, Bollinger Bands can be used to identify M-Tops and W-Bottoms or to determine the strength of the trend.



# Bollinger BandWidth

https://school.stockcharts.com/doku.php?id=technical_indicators:bollinger_band_width

In his book, Bollinger on Bollinger Bands, John Bollinger refers to Bollinger BandWidth as one of two indicators that can be derived from Bollinger Bands (the other being %B).

BandWidth measures the percentage difference between the upper band and the lower band.
BandWidth decreases as Bollinger Bands narrow and increases as Bollinger Bands widen.
Because Bollinger Bands are based on the standard deviation, falling BandWidth reflects decreasing volatility and rising BandWidth reflects increasing volatility.

Calculation: ( (Upper Band - Lower Band) / Middle Band) * 100%

When calculating BandWidth, the first step is to subtract the value of the lower band from the value of the upper band.
This shows the absolute difference.
This difference is then divided by the middle band, which normalizes the value.
This normalized Bandwidth can then be compared across different timeframes or with the BandWidth values for other securities.

Narrow BandWidth is relative. BandWidth values should be gauged relative to prior BandWidth values over a period of time.
It is important to get a good look-back period to define BandWidth range for a particular ETF, index or stock.
For example, an eight- to twelve-month chart will show BandWidth highs and lows over a significant timeframe.
BandWidth is considered narrow as it approaches the lows of this range and wide as it approaches the high end.
[img/bbw-2-a.png]
[img/bbw-2-b.png]

# PercentB (%B) Indicator

https://school.stockcharts.com/doku.php?id=technical_indicators:bollinger_band_perce

%B quantifies a security's price relative to the upper and lower Bollinger Band.
There are six basic relationship levels:

- %B is below 0 when price is below the lower band
- %B equals 0 when price is at the lower band
- %B is between 0 and .50 when price is between the lower and middle band (20-day SMA)
- %B is between .50 and 1 when price is between the upper and middle band (20-day SMA)
- %B equals 1 when price is at the upper band
- %B is above 1 when price is above the upper band

Calculation: %B = (Price - Lower Band)/(Upper Band - Lower Band)


# BB Trend Indicator

https://www.mql5.com/en/market/product/30223#description

BBTrend is a relatively new indicator developed by John Bollinger to work with Bollinger Bands™. It is one of only a few indicators that can signal both strength and direction making it a very valuable tool for traders.

## Calculations

The calculations are fairly simple. The default periods of 20 and 50 are shown, but these can be changed through the parameters.

- Lower = MathAbs(lowerBB(20) - lowerBB(50))
- Upper = MathAbs(upperBB(20) - upperBB(50))
- BBTrend = (lower - upper) / middleBB(20)

## Interpretation

- Bullish trend = BBTrend above zero
- Bearish trend = BBTrend below zero

The degree above or below zero determines the strength or momentum behind the trend.

## Parameters

- Shorter Bands Period – the shorter period for the Bollinger Bands used in the calculation.
- Longer Bands Period – the longer period for the Bollinger Bands used in the calculation.
- iBands Shift – “shift” used by both the longer and shorter Bollinger Bands (iBands) calculations.
- iBands Deviation – “deviation” used by both the longer and shorter Bollinger Bands (iBands) calculations.

[img/bb-trend-indicator-screen.png]