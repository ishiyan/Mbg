using System;
using System.Collections.Generic;
using System.Globalization;

using Mbs.Trading.Data;
using Mbs.Trading.Indicators.Abstractions;

namespace Mbs.Trading.Indicators.PerryKaufman
{
    /// <summary>
    /// Kaufman's adaptive moving average (KAMA) is an EMA with the <c>α</c> (smoothing factor) being changed with each new sample within the fastest and the slowest boundaries.
    /// <para><c>KAMAᵢ = αPᵢ + (1 - α)*KAMAᵢ₋₁</c></para>
    /// where the smoothing factor
    /// <para><c>α = (αs + (αf - αs)ε)²</c></para>
    /// where the <c>αf</c> is the α of the fastest (shortest, default 2 samples) period boundary, the <c>αs</c> is the α of the slowest (longest, default 30 samples) period boundary.
    /// <para>The efficiency ratio</para>
    /// <para><c>ε = |P - Pℓ| / ∑|Pᵢ - Pᵢ₊₁|</c></para>
    /// <para>where <c>i ≤ ℓ-1</c> and <c>ℓ</c> is a number of samples used to calculate the <c>ε</c>.
    /// The indicator is not primed during the first <c>ℓ</c> updates. The recommended values of <c>ℓ</c> are in the range of 8 to 10.</para>
    /// <para>The efficiency ratio has the value of 1 when samples move in the same direction for the full <c>ℓ</c> periods, and a value of 0 when samples are unchanged over the <c>ℓ</c> periods.
    /// When samples move in wide swings within the interval, the sum of the denominator becomes very large compared with the numerator and the <c>ε</c> approaches 0.
    /// Smaller values of <c>ε</c> result in a smaller smoothing constant and a slower trend.</para>
    /// <para>See <c>Perry J. Kaufman, Smarter Trading, McGraw-Hill, Ney York, 1995, pp. 129-153</c> for a complete discussion.</para>
    /// </summary>
    public sealed class KaufmanAdaptiveMovingAverage : Indicator, ILineIndicator
    {
        #region Members and accessors

        /// <summary>
        /// The fastest boundary length, <c>ℓf</c>. The equivalent smoothing factor <c>αf</c> is
        /// <para><c>αf = 2/(ℓf + 1), 2 ≤ ℓf</c></para>
        /// The default value is 2.
        /// </summary>
        public int FastestLength { get; }

        /// <summary>
        /// The fastest boundary smoothing factor, <c>αf</c>. The equivalent length <c>ℓf</c> is
        /// <para><c>αf = 2/(ℓf + 1), 0 ≤ αf ≤ 1</c></para>
        /// The default value is 2/3 (0.66666666666666666666666666666667).
        /// </summary>
        public double FastestSmoothingFactor { get; }

        /// <summary>
        /// The slowest boundary length, <c>ℓs</c>. The equivalent smoothing factor <c>αs</c> is
        /// <para><c>αs = 2/(ℓs + 1), 1 ≤ ℓs</c></para>
        /// The default value is 30.
        /// </summary>
        public int SlowestLength { get; }

        /// <summary>
        /// The slowest boundary smoothing factor, <c>αs</c>. The equivalent length <c>ℓs</c> is
        /// <para><c>αs = 2/(ℓs + 1), 0 ≤ α ≤ 1</c></para>
        /// The default value is 2/31 (0.06451612903225806451612903225806).
        /// </summary>
        public double SlowestSmoothingFactor { get; }

        /// <summary>
        /// The number of last samples, <c>ℓ</c>, used to calculate the efficiency ratio. The default value is 10.
        /// </summary>
        public int EfficiencyRatioLength { get; }

        private double value = double.NaN;

        /// <summary>
        /// The current value of the the Kaufman adaptive moving average, or <c>NaN</c> if not primed.
        /// The indicator is not primed during the first <c>ℓ</c> updates, where <c>ℓ</c> is the number of samples used to calculate the efficiency ratio.
        /// </summary>
        public double Value
        {
            get
            {
                lock (updateLock)
                {
                    return value;
                }
            }
        }

        private readonly double smoothingFactorDiff;
        private readonly int length1;
        private int windowCount;
        private readonly double[] window;
        private readonly double[] absoluteDelta;
        private double sumAbsoluteDelta;

        private const string Kama = "KAMA";
        private const string KamaFull = "Kaufman Adaptive Moving Average";
        private const string ArgumentEfficiencyRatioLength = "efficiencyRatioLength";
        private const string ArgumentFastestSmoothingFactor = "fastestSmoothingFactor";
        private const string ArgumentSlowestSmoothingFactor = "slowestSmoothingFactor";
        private const string ArgumentFastestLength = "fastestLength";
        private const string ArgumentSlowestLength = "slowestLength";
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the <see cref="KaufmanAdaptiveMovingAverage"/> class.
        /// </summary>
        /// <param name="efficiencyRatioLength">The number of last samples used to calculate the efficiency ratio.</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public KaufmanAdaptiveMovingAverage(int efficiencyRatioLength,
            OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : this(efficiencyRatioLength, 2, 30, ohlcvComponent)
        {
        }

        /// <summary>
        /// Constructs a new instance of the <see cref="KaufmanAdaptiveMovingAverage"/> class.
        /// </summary>
        /// <param name="efficiencyRatioLength">The number of last samples used to calculate the efficiency ratio.</param>
        /// <param name="fastestLength">The fastest boundary length, <c>ℓf</c>. The default value is 2.
        /// The equivalent smoothing factor <c>αf</c> is
        /// <para><c>αf = 2/(ℓf + 1), 2 ≤ ℓ</c></para>
        /// </param>
        /// <param name="slowestLength">The slowest boundary length, <c>ℓs</c>. The default value is 30.
        /// The equivalent smoothing factor <c>αs</c> is
        /// <para><c>αs = 2/(ℓs + 1), 2 ≤ ℓ</c></para>
        /// </param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public KaufmanAdaptiveMovingAverage(int efficiencyRatioLength, int fastestLength, int slowestLength,
            OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(Kama, KamaFull, ohlcvComponent)
        {
            if (2 > efficiencyRatioLength)
                throw new ArgumentOutOfRangeException(ArgumentEfficiencyRatioLength);
            if (2 > fastestLength)
                throw new ArgumentOutOfRangeException(ArgumentFastestLength);
            if (2 > slowestLength)
                throw new ArgumentOutOfRangeException(ArgumentSlowestLength);
            this.FastestLength = fastestLength;
            this.SlowestLength = slowestLength;
            FastestSmoothingFactor = 2d / (1d + fastestLength);
            SlowestSmoothingFactor = 2d / (1d + slowestLength);
            smoothingFactorDiff = FastestSmoothingFactor - SlowestSmoothingFactor;
            this.EfficiencyRatioLength = efficiencyRatioLength;
            length1 = efficiencyRatioLength + 1;
            window = new double[length1];
            absoluteDelta = new double[length1];
            Moniker = string.Concat(Kama, efficiencyRatioLength.ToString(CultureInfo.InvariantCulture));
        }

        /// <summary>
        /// Constructs a new instance of the <see cref="KaufmanAdaptiveMovingAverage"/> class.
        /// </summary>
        /// <param name="efficiencyRatioLength">The number of last samples used to calculate the efficiency ratio.</param>
        /// <param name="fastestSmoothingFactor">The fastest boundary smoothing factor, <c>αf</c>. The default value is 2/3 (0.66666666666666666666666666666667).
        /// The equivalent length <c>ℓf</c> is
        /// <para><c>αf = 2/(ℓf + 1), 0 ≤ αf ≤ 1</c></para>
        /// </param>
        /// <param name="slowestSmoothingFactor">The slowest boundary smoothing factor, <c>αs</c>. The default value is 2/31 (0.06451612903225806451612903225806).
        /// The equivalent length <c>ℓs</c> is
        /// <para><c>αs = 2/(ℓs + 1), 0 ≤ αs ≤ 1</c></para>
        /// </param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public KaufmanAdaptiveMovingAverage(int efficiencyRatioLength, double fastestSmoothingFactor,
            double slowestSmoothingFactor, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(Kama, KamaFull, ohlcvComponent)
        {
            if (2 > efficiencyRatioLength)
                throw new ArgumentOutOfRangeException(ArgumentEfficiencyRatioLength);
            if (0d > fastestSmoothingFactor || 1d < fastestSmoothingFactor)
                throw new ArgumentOutOfRangeException(ArgumentFastestSmoothingFactor);
            if (0d > slowestSmoothingFactor || 1d < slowestSmoothingFactor)
                throw new ArgumentOutOfRangeException(ArgumentSlowestSmoothingFactor);
            this.FastestSmoothingFactor = fastestSmoothingFactor;
            this.SlowestSmoothingFactor = slowestSmoothingFactor;
            smoothingFactorDiff = fastestSmoothingFactor - slowestSmoothingFactor;
            if (0.0000001 > fastestSmoothingFactor)
                FastestLength = int.MaxValue;
            else
                FastestLength = (int) Math.Round(2d / fastestSmoothingFactor) - 1;
            if (0.0000001 > slowestSmoothingFactor)
                SlowestLength = int.MaxValue;
            else
                SlowestLength = (int) Math.Round(2d / slowestSmoothingFactor) - 1;
            this.EfficiencyRatioLength = efficiencyRatioLength;
            length1 = efficiencyRatioLength + 1;
            window = new double[length1];
            absoluteDelta = new double[length1];
            Moniker = string.Concat(Kama, efficiencyRatioLength.ToString(CultureInfo.InvariantCulture));
        }
        #endregion

        #region Reset
        /// <inheritdoc />
        public override void Reset()
        {
            lock (updateLock)
            {
                primed = false;
                value = double.NaN;
                windowCount = 0;
                sumAbsoluteDelta = 0d;
            }
        }
        #endregion

        #region Update
        /// <summary>
        /// Updates the value of the indicator.
        /// </summary>
        /// <param name="sample">A new sample.</param>
        /// <returns>The new value of the indicator.</returns>
        public double Update(double sample)
        {
            if (double.IsNaN(sample))
                return sample;
            lock (updateLock)
            {
                double temp, range;
                if (primed)
                {
                    temp = Math.Abs(sample - window[EfficiencyRatioLength]);
                    sumAbsoluteDelta += temp - absoluteDelta[1];
                    for (int i = 0, j = 1; i < EfficiencyRatioLength; i++, j++)
                    {
                        window[i] = window[j];
                        absoluteDelta[i] = absoluteDelta[j];
                    }

                    window[EfficiencyRatioLength] = sample;
                    absoluteDelta[EfficiencyRatioLength] = temp;
                    range = Math.Abs(sample - window[0]);
                    if (sumAbsoluteDelta <= range || sumAbsoluteDelta < 0.00000001)
                        temp = 1d;
                    else
                        temp = range / sumAbsoluteDelta;
                    temp = SlowestSmoothingFactor + temp * smoothingFactorDiff;
                    value += (sample - value) * temp * temp;
                }
                else // Not primed.
                {
                    window[windowCount] = sample;
                    if (0 < windowCount)
                    {
                        temp = Math.Abs(sample - window[windowCount - 1]);
                        absoluteDelta[windowCount] = temp;
                        sumAbsoluteDelta += temp;
                    }

                    if (length1 == ++windowCount)
                    {
                        primed = true;
                        range = Math.Abs(sample - window[0]);
                        if (sumAbsoluteDelta <= range || sumAbsoluteDelta < 0.00000001)
                            temp = 1d;
                        else
                            temp = range / sumAbsoluteDelta;
                        temp = SlowestSmoothingFactor + temp * smoothingFactorDiff;
                        value = window[EfficiencyRatioLength - 1];
                        value += (sample - value) * temp * temp;
                    }
                }

                return value;
            }
        }

        /// <summary>
        /// Updates the value of the indicator.
        /// </summary>
        /// <param name="sample">A new sample.</param>
        /// <param name="dateTime">A date-time of the new sample.</param>
        /// <returns>The new value of the indicator.</returns>
        public Scalar Update(double sample, DateTime dateTime)
        {
            return new Scalar(dateTime, Update(sample));
        }

        /// <summary>
        /// Updates the value of the indicator.
        /// </summary>
        /// <param name="scalar">A new scalar.</param>
        /// <returns>The new value of the indicator.</returns>
        public Scalar Update(Scalar scalar)
        {
            return new Scalar(scalar.Time, Update(scalar.Value));
        }

        /// <summary>
        /// Updates the value of the indicator.
        /// </summary>
        /// <param name="ohlcv">A new ohlcv.</param>
        /// <returns>The new value of the indicator.</returns>
        public Scalar Update(Ohlcv ohlcv)
        {
            return new Scalar(ohlcv.Time, Update(ohlcv.Component(OhlcvComponent)));
        }
        #endregion

        #region Calculate
        /// <summary>
        /// Calculates the Kaufman adaptive moving average from the input array.
        /// </summary>
        /// <param name="sampleList">The sample list.</param>
        /// <param name="efficiencyRatioLength">The number of last samples used to calculate the efficiency ratio.</param>
        /// <param name="fastestSmoothingFactor">The fastest boundary smoothing factor, <c>αf</c>. The default value is 2/3 (0.66666666666666666666666666666667).
        /// The equivalent length <c>ℓf</c> is
        /// <para><c>αf = 2/(ℓf + 1), 0 ≤ αf ≤ 1</c></para>
        /// </param>
        /// <param name="slowestSmoothingFactor">The slowest boundary smoothing factor, <c>αs</c>. The default value is 2/31 (0.06451612903225806451612903225806).
        /// The equivalent length <c>ℓs</c> is
        /// <para><c>αs = 2/(ℓs + 1), 0 ≤ αs ≤ 1</c></para>
        /// </param>
        /// <returns>A list of the indicator values.</returns>
        public static List<double> Calculate(List<double> sampleList, int efficiencyRatioLength,
            double fastestSmoothingFactor = 0.66666666666666666666666666666667,
            double slowestSmoothingFactor = 0.06451612903225806451612903225806)
        {
            if (2 > efficiencyRatioLength)
                throw new ArgumentOutOfRangeException(ArgumentEfficiencyRatioLength);
            if (0d > fastestSmoothingFactor || 1d < fastestSmoothingFactor)
                throw new ArgumentOutOfRangeException(ArgumentFastestSmoothingFactor);
            if (0d > slowestSmoothingFactor || 1d < slowestSmoothingFactor)
                throw new ArgumentOutOfRangeException(ArgumentSlowestSmoothingFactor);
            int count = sampleList.Count;
            var resultList = new List<double>(count);
            if (count <= efficiencyRatioLength)
            {
                for (int i = 0; i < count; i++)
                    resultList.Add(double.NaN);
            }
            else
            {
                double smoothingFactorDiff = fastestSmoothingFactor - slowestSmoothingFactor;
                int firstIndex = 0;
                double range = Math.Abs(sampleList[efficiencyRatioLength] - sampleList[0]);
                double temp, sumAbsoluteDelta = 0;
                for (int i = 0, j = 1; i < efficiencyRatioLength; i++, j++)
                {
                    temp = Math.Abs(sampleList[j] - sampleList[i]);
                    sumAbsoluteDelta += temp;
                }

                for (int i = 0; i < efficiencyRatioLength; i++)
                    resultList.Add(double.NaN);
                if (sumAbsoluteDelta <= range || sumAbsoluteDelta < 0.00000001)
                    temp = 1d;
                else
                    temp = range / sumAbsoluteDelta;
                temp = slowestSmoothingFactor + temp * smoothingFactorDiff;
                double value = sampleList[efficiencyRatioLength - 1];
                value += (sampleList[efficiencyRatioLength] - value) * temp * temp;
                resultList.Add(value);
                while (++efficiencyRatioLength < count)
                {
                    temp = sampleList[firstIndex++];
                    range = sampleList[firstIndex];
                    sumAbsoluteDelta -= Math.Abs(range - temp);
                    double lastSample = sampleList[efficiencyRatioLength];
                    sumAbsoluteDelta += Math.Abs(lastSample - sampleList[efficiencyRatioLength - 1]);
                    range = Math.Abs(lastSample - range);
                    if (sumAbsoluteDelta <= range || sumAbsoluteDelta < 0.00000001)
                        temp = 1d;
                    else
                        temp = range / sumAbsoluteDelta;
                    temp = slowestSmoothingFactor + temp * smoothingFactorDiff;
                    value += (lastSample - value) * temp * temp;
                    resultList.Add(value);
                }
            }

            return resultList;
        }
        #endregion
    }
}
