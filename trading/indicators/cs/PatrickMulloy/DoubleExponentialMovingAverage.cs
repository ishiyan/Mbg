using System;
using System.Globalization;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators.Abstractions;

namespace Mbs.Trading.Indicators.PatrickMulloy
{
    /// <summary>
    /// The Double Exponential Moving Average (DEMA) is a smoothing indicator with less lag than a straight exponential moving average.
    /// <para>The DEMA was developed by Patrick G. Mulloy and is described in two articles:</para>
    /// <para>❶ Technical Analysis of Stocks &amp; Commodities v.12:1 (11-19), Smoothing Data With Faster Moving Averages.</para>
    /// <para>❷ Technical Analysis of Stocks &amp; Commodities v.12:2 (72-80), Smoothing Data With Less Lag.</para>
    /// <para>The calculation is as follows:</para>
    /// <para><c>EMA¹ᵢ = EMA(Pᵢ) = αPᵢ + (1-α)EMA¹ᵢ₋₁ = EMA¹ᵢ₋₁ + α(Pᵢ - EMA¹ᵢ₋₁), 0 &lt; α ≤ 1</c></para>
    /// <para><c>EMA²ᵢ = EMA(EMA¹ᵢ) = αEMA¹ᵢ + (1-α)EMA²ᵢ₋₁ = EMA²ᵢ₋₁ + α(EMA¹ᵢ - EMA²ᵢ₋₁), 0 &lt; α ≤ 1</c></para>
    /// <para><c>DEMAᵢ = 2EMA¹ᵢ - EMA²ᵢ</c></para>
    /// <para>The very first EMA value (the seed for subsequent values) is calculated differently. This implementation allows for two algorithms for this seed.</para>
    /// <para>❶ Use a simple average of the first 'period'. This is the most widely documented approach.</para>
    /// <para>❷ Use first sample value as a seed. This is used in Metastock.</para>
    /// </summary>
    public sealed class DoubleExponentialMovingAverage : LineIndicator
    {
        #region Members and accessors
        /// <summary>
        /// The length (<c>ℓ</c>) of the exponential moving average. The equivalent smoothing factor (<c>α</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </summary>
        public int Length { get; }

        /// <summary>
        /// The smoothing factor (<c>α</c>) of the exponential moving average. The equivalent length (<c>ℓ</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </summary>
        public double SmoothingFactor { get; }

        /// <summary>
        /// The current value of the the Double Exponential Moving Average.
        /// </summary>
        public double Value { get { lock (updateLock) { return primed ? value : double.NaN; } } }

        /// <summary>
        /// If the very first EMA value is a simple average of the first 'period' (the most widely documented approach) or the first input value (used in Metastock).
        /// </summary>
        public bool FirstIsAverage { get; }

        private readonly int length2;
        private int count1;
        private int count2;
        private double sum1;
        private double sum2;
        private double value = double.NaN;
        private double value1;
        private double value2;

        private const double Epsilon = 0.00000001;
        private const string Dema = "dema";
        private const string DemaFull = "Double Exponential Moving Average";
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the class.
        /// </summary>
        /// <param name="length">The length (<c>ℓ</c>) of the exponential moving average.
        /// The equivalent smoothing factor (<c>α</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </param>
        /// <param name="firstIsAverage">If the very first EMA value is a simple average of the first 'period' (the most widely documented approach) or the first input value (used in Metastock).</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public DoubleExponentialMovingAverage(int length, bool firstIsAverage = true, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(Dema, DemaFull, ohlcvComponent)
        {
            if (1 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            Length = length;
            length2 = length * 2;
            SmoothingFactor = 2d / (1d + length);
            FirstIsAverage = firstIsAverage;
            Moniker = string.Concat(Dema, length.ToString(CultureInfo.InvariantCulture));
        }

        /// <summary>
        /// Constructs a new instance of the class.
        /// </summary>
        /// <param name="smoothingFactor">The smoothing factor (<c>α</c>) of the exponential moving average.
        /// The equivalent length (<c>ℓ</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </param>
        /// <param name="firstIsAverage">If the very first EMA value is a simple average of the first 'period' (the most widely documented approach) or the first input value (used in Metastock).</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public DoubleExponentialMovingAverage(double smoothingFactor, bool firstIsAverage = true, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(Dema, DemaFull, ohlcvComponent)
        {
            if (0d > smoothingFactor || 1d < smoothingFactor)
                throw new ArgumentOutOfRangeException(nameof(smoothingFactor));
            SmoothingFactor = smoothingFactor;
            if (Epsilon > smoothingFactor)
                Length = int.MaxValue;
            else
                Length = (int)Math.Round(2d / smoothingFactor) - 1;
            length2 = Length * 2;
            FirstIsAverage = firstIsAverage;
            Moniker = string.Concat(Dema, Length.ToString(CultureInfo.InvariantCulture));
        }
        #endregion

        #region Reset
        /// <inheritdoc />
        public override void Reset()
        {
            lock (updateLock)
            {
                primed = false;
                count1 = 0;
                sum1 = 0d;
                value1 = 0d;
                count2 = 0;
                sum2 = 0d;
                value2 = 0d;
                value = double.NaN;
            }
        }
        #endregion

        #region Update
        /// <summary>
        /// Updates the value of the indicator.
        /// </summary>
        /// <param name="sample">A new sample.</param>
        /// <returns>The new value of the indicator.</returns>
        public override double Update(double sample)
        {
            if (double.IsNaN(sample))
                return sample;
            lock (updateLock)
            {
                if (primed)
                {
                    value1 += (sample - value1) * SmoothingFactor;
                    value2 += (value1 - value2) * SmoothingFactor;
                    value = 2d * value1 - value2;
                }
                else
                {
                    if (FirstIsAverage)
                    {
                        if (Length > count1)
                        {
                            sum1 += sample;
                            if (Length == ++count1)
                            {
                                value1 = sum1 / Length;
                                sum2 += value1;
                            }
                        }
                        else
                        {
                            value1 += (sample - value1) * SmoothingFactor;
                            sum2 += value1;
                            if (Length == ++count2)
                            {
                                value2 = sum2 / Length;
                                value = 2d * value1 - value2;
                                primed = true;
                            }
                        }
                    }
                    else
                    {
                        if (Length > count1)
                        {
                            if (1 == ++count1)
                                value1 = sample;
                            else
                                value1 += (sample - value1) * SmoothingFactor;
                        }
                        else
                        {
                            value1 += (sample - value1) * SmoothingFactor;
                            if (Length == count1++)
                                value2 = value1;
                            else
                            {
                                value2 += (value1 - value2) * SmoothingFactor;
                                if (length2 == count1)
                                {
                                    value = 2d * value1 - value2;
                                    primed = true;
                                }
                            }
                        }
                    }
                }
                return value;
            }
        }
        #endregion
    }
}
