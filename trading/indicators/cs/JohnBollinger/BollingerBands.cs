using System;
using System.Globalization;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators.Abstractions;

namespace Mbs.Trading.Indicators.JohnBollinger
{
    /// <summary>
    /// Bollinger bands are a type of price envelope invented by John Bollinger in the 1980s. Bollinger bands consist of:
    /// <para>❶ a middle band being an <c>ℓ</c>-period moving average (MA)</para>
    /// <para>❷ an upper band at <c>K</c> times an <c>ℓ</c>-period standard deviation <c>σ</c> above the middle band (<c>MA + Kσ</c>)</para>
    /// <para>❸ an lower band at <c>K</c> times an <c>ℓ</c>-period standard deviation <c>σ</c> below the middle band (<c>MA - Kσ</c>)</para>
    /// <para>Typical values for <c>ℓ</c> and <c>Kĸ</c> are 20 and 2, respectively. The default choice for the average is a simple moving average, but other types of averages can be employed as needed.</para>
    /// <para>Exponential moving averages are a common second choice. Usually the same period is used for both the middle band and the calculation of standard deviation.</para>
    /// </summary>
    public sealed class BollingerBands : Indicator, IBandIndicator
    {
        #region Members and accessors
        private double upperValue = double.NaN;
        /// <summary>
        /// The upper value.
        /// </summary>
        public double UpperValue { get { lock (updateLock) { return upperValue; } } }

        private double middleValue = double.NaN;
        /// <summary>
        /// The middle value.
        /// </summary>
        public double MiddleValue { get { lock (updateLock) { return middleValue; } } }

        private double lowerValue = double.NaN;
        /// <summary>
        /// The lower value.
        /// </summary>
        public double LowerValue { get { lock (updateLock) { return lowerValue; } } }

        private readonly double upperMultiplier;
        /// <summary>
        /// The upper multiplier.
        /// </summary>
        public double UpperMultiplier { get { lock (updateLock) { return upperMultiplier; } } }

        private readonly double lowerMultiplier;
        /// <summary>
        /// The lower multiplier.
        /// </summary>
        public double LowerMultiplier { get { lock (updateLock) { return lowerMultiplier; } } }

        /// <summary>
        /// Tells how wide the Bollinger Bands are on a normalized basis: <c>BandWidth = (upper - lower) / middle</c>.
        /// </summary>
        public double BandWidthValue
        {
            get
            {
                lock (updateLock)
                {
                    return (Math.Abs(middleValue) < double.Epsilon) ?
                        1d : (upperValue - lowerValue) / middleValue;
                }
            }
        }

        /// <summary>
        /// Equals 1 at the upper band and 0 at the lower band: <c>PercentBand = MiddleValue - LowerValue) / (UpperValue - LowerValue)</c>.
        /// </summary>
        public double PercentBandValue
        {
            get
            {
                lock (updateLock)
                {
                    double temp = upperValue - lowerValue;
                    return (Math.Abs(temp) < double.Epsilon) ? 1d :
                        (middleValue - lowerValue) / temp;
                }
            }
        }

        /// <summary>
        /// The standard deviation value.
        /// </summary>
        public double StandardDeviationValue
        {
            get
            {
                lock (updateLock)
                {
                    return standardDeviation.Value;
                }
            }
        }

        /// <summary>
        /// The standard deviation look-back length.
        /// </summary>
        public int StandardDeviationLength { get; }

        /// <summary>
        /// If the estimate of the variance is the unbiased sample variance or the population variance. The default value is true.
        /// </summary>
        public bool IsUnbiased => standardDeviation.IsUnbiased;

        /// <summary>
        /// The source line indicator.
        /// </summary>
        public ILineIndicator SourceLineIndicator { get; }

        private readonly bool hasUpperMultiplier;
        private readonly bool hasLowerMultiplier;
        private readonly StandardDeviation standardDeviation;

        private const string Bb = "BBands";
        private const string BbFull = "Bollinger Bands";
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the <see cref="BollingerBands"/> object.
        /// </summary>
        /// <param name="sourceLineIndicator">The line indicator which serves as an input source.</param>
        /// <param name="length">The standard deviation look-back length.</param>
        /// <param name="multiplier">The lower and upper band multiplier.</param>
        /// <param name="unbiased">If the estimate of the variance is the unbiased sample variance or the population variance. The default value is true.</param>
        public BollingerBands(ILineIndicator sourceLineIndicator, int length, double multiplier, bool unbiased = true)
            : this(sourceLineIndicator, length, multiplier, multiplier, unbiased)
        {
        }

        /// <summary>
        /// Constructs a new instance of the <see cref="BollingerBands"/> object.
        /// </summary>
        /// <param name="ohlcvComponent">The ohlcv component.</param>
        /// <param name="length">The standard deviation look-back length.</param>
        /// <param name="multiplier">The lower and upper band multiplier.</param>
        public BollingerBands(OhlcvComponent ohlcvComponent, int length, double multiplier)
            : this(ohlcvComponent, length, multiplier, multiplier)
        {
        }

        /// <summary>
        /// Constructs a new instance of the <see cref="BollingerBands"/> object.
        /// </summary>
        /// <param name="ohlcvComponent">The ohlcv component.</param>
        /// <param name="length">The standard deviation look-back length.</param>
        /// <param name="multiplier">The lower and upper band multiplier.</param>
        /// <param name="unbiased">If the estimate of the variance is the unbiased sample variance or the population variance. The default value is true.</param>
        public BollingerBands(OhlcvComponent ohlcvComponent, int length, double multiplier, bool unbiased)
            : this(ohlcvComponent, length, multiplier, multiplier, unbiased)
        {
        }

        /// <summary>
        /// Constructs a new instance of the <see cref="BollingerBands"/> object.
        /// </summary>
        /// <param name="sourceLineIndicator">The line indicator which serves as an input source.</param>
        /// <param name="length">The standard deviation look-back length.</param>
        /// <param name="lowerMultiplier">The lower band multiplier.</param>
        /// <param name="upperMultiplier">The upper band multiplier.</param>
        /// <param name="unbiased">If the estimate of the variance is the unbiased sample variance or the population variance. The default value is true.</param>
        public BollingerBands(ILineIndicator sourceLineIndicator, int length, double lowerMultiplier = 1d, double upperMultiplier = 1d, bool unbiased = true)
            : base(Bb, BbFull)
        {
            if (2 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            StandardDeviationLength = length;
            this.lowerMultiplier = lowerMultiplier;
            this.upperMultiplier = upperMultiplier;
            hasLowerMultiplier = Math.Abs(1d - lowerMultiplier) > double.Epsilon;
            hasUpperMultiplier = Math.Abs(1d - upperMultiplier) > double.Epsilon;
            SourceLineIndicator = sourceLineIndicator;
            standardDeviation = new StandardDeviation(length, unbiased);
            Moniker = string.Concat(Bb, length.ToString(CultureInfo.InvariantCulture), "[", sourceLineIndicator.Moniker, "]");
        }

        /// <summary>
        /// Constructs a new instance of the <see cref="BollingerBands"/> object.
        /// </summary>
        /// <param name="ohlcvComponent">The ohlcv component.</param>
        /// <param name="length">The standard deviation look-back length.</param>
        /// <param name="lowerMultiplier">The lower band multiplier.</param>
        /// <param name="upperMultiplier">The upper band multiplier.</param>
        /// <param name="unbiased">If the estimate of the variance is the unbiased sample variance or the population variance. The default value is true.</param>
        public BollingerBands(OhlcvComponent ohlcvComponent, int length, double lowerMultiplier = 1d, double upperMultiplier = 1d, bool unbiased = true)
            : base(Bb, BbFull, ohlcvComponent)
        {
            if (2 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            StandardDeviationLength = length;
            this.lowerMultiplier = lowerMultiplier;
            this.upperMultiplier = upperMultiplier;
            hasLowerMultiplier = Math.Abs(1d - lowerMultiplier) > double.Epsilon;
            hasUpperMultiplier = Math.Abs(1d - upperMultiplier) > double.Epsilon;
            standardDeviation = new StandardDeviation(length, unbiased);
            Moniker = string.Concat(Bb, length.ToString(CultureInfo.InvariantCulture));
        }
        #endregion

        #region Reset
        /// <inheritdoc />
        public override void Reset()
        {
            lock (updateLock)
            {
                primed = false;
                lowerValue = double.NaN;
                middleValue = double.NaN;
                upperValue = double.NaN;
                standardDeviation.Reset();
                SourceLineIndicator?.Reset();
            }
        }
        #endregion

        #region Update
        private Band DoUpdate(double sample, DateTime dateTime)
        {
            lock (updateLock)
            {
                if (double.IsNaN(sample))
                    return new Band(dateTime);
                middleValue = sample;
                double temp = standardDeviation.Update(sample);
                if (!primed)
                {
                    primed = standardDeviation.IsPrimed;
                    if (null != SourceLineIndicator)
                    {
                        if (primed && !SourceLineIndicator.IsPrimed)
                            primed = false;
                    }
                    if (!primed)
                        return new Band(dateTime);
                }
                lowerValue = sample - (hasLowerMultiplier ? temp * lowerMultiplier : temp);
                upperValue = sample + (hasUpperMultiplier ? temp * upperMultiplier : temp);
                return new Band(dateTime, upperValue, lowerValue);
            }
        }

        /// <summary>
        /// Updates the indicator.
        /// </summary>
        /// <param name="sample">A new sample.</param>
        /// <param name="dateTime">A date-time of the new sample.</param>
        /// <returns>The updated value of the indicator.</returns>
        public Band Update(double sample, DateTime dateTime)
        {
            if (null != SourceLineIndicator)
                sample = SourceLineIndicator.Update(sample);
            return DoUpdate(sample, dateTime);
        }

        /// <summary>
        /// Updates the indicator.
        /// </summary>
        /// <param name="sample">A new <c>scalar</c> sample.</param>
        /// <returns>The updated value of the indicator.</returns>
        public Band Update(Scalar sample)
        {
            if (null != SourceLineIndicator)
                sample = SourceLineIndicator.Update(sample);
            return DoUpdate(sample.Value, sample.Time);
        }

        /// <summary>
        /// Updates the indicator.
        /// </summary>
        /// <param name="sample">A new <c>ohlcv</c> sample.</param>
        /// <returns>The updated value of the indicator.</returns>
        public Band Update(Ohlcv sample)
        {
            if (null != SourceLineIndicator)
            {
                Scalar scalar = SourceLineIndicator.Update(sample);
                return DoUpdate(scalar.Value, scalar.Time);
            }
            return DoUpdate(sample.Component(OhlcvComponent), sample.Time);
        }
        #endregion
    }
}
