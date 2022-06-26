using System;
using System.Collections.Generic;
using System.Globalization;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators.Abstractions;

namespace Mbs.Trading.Indicators
{
    /// <summary>
    /// A weighted moving average is a moving average that has multiplying factors to give arithmetically decreasing weights to look back samples.
    /// <para><c>WMAᵢ = (ℓPᵢ + (ℓ-1)Pᵢ₋₁ + ... + Pᵢ₋ℓ) / (ℓ + (ℓ-1) + ... + 2 + 1)</c>, where <c>ℓ</c> is the length.</para>
    /// The denominator is a triangle number and can be computed as <c>½ℓ(ℓ-1)</c>.
    /// In other words, this is an unweighted mean (gives equal weight to each sample) of the previous <c>ℓ</c> samples.
    /// <para>
    /// When calculating the WMA across successive values,
    /// <para><c>WMAᵢ₊₁ - WMAᵢ = ℓPᵢ₊₁ - Pᵢ - ... - Pᵢ₋ℓ₊₁</c></para>
    /// If we denote the sum <c>Pᵢ + ... + Pᵢ₋ℓ₊₁</c> by <c>Totalᵢ</c>, then
    /// <para><c>Totalᵢ₊₁ = Totalᵢ + Pᵢ₊₁ - Pᵢ₋ℓ₊₁</c></para>
    /// <para><c>Numeratorᵢ₊₁ = Numeratorᵢ + ℓPᵢ₊₁ - Totalᵢ</c></para>
    /// <para><c>WMAᵢ₊₁ = Numeratorᵢ₊₁ / ½ℓ(ℓ-1)</c></para>
    /// The WMA indicator is not primed during the first <c>ℓ-1</c> updates.</para>
    /// </summary>
    public sealed class WeightedMovingAverage : LineIndicator
    {
        #region Members and accessors
        /// <summary>
        /// The length, <c>ℓ</c>, (the number of time periods) of the weighted moving average.
        /// </summary>
        public int Length { get; }

        /// <summary>
        /// The current value of the the weighted moving average, or <c>NaN</c> if not primed.
        /// The indicator is not primed during the first <c>ℓ-1</c> updates, where <c>ℓ</c> is the length.
        /// </summary>
        public double Value { get { lock (updateLock) { return value; } } }

        private readonly int lastIndex;
        private readonly int divider;
        private int windowCount;
        private double periodSub;
        private double periodSum;
        private double value = double.NaN;
        private readonly double[] window;

        private const string Wma = "WMA";
        private const string WmaFull = "Weighted Moving Average";
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the class.
        /// </summary>
        /// <param name="length">The number of time periods of the weighted moving average.</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public WeightedMovingAverage(int length, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(Wma, WmaFull, ohlcvComponent)
        {
            if (2 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            Length = length;
            lastIndex = length - 1;
            divider = length * (length + 1) / 2;
            window = new double[length];
            Moniker = string.Concat(Wma, length.ToString(CultureInfo.InvariantCulture));
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
                periodSub = 0d;
                periodSum = 0d;
            }
        }
        #endregion

        #region Update
        /// <inheritdoc />
        public override double Update(double sample)
        {
            if (double.IsNaN(sample))
                return sample;
            lock (updateLock)
            {
                if (primed)
                {
                    periodSum -= periodSub;
                    periodSum += sample * Length;
                    value = periodSum / divider;
                    periodSub -= window[0];
                    periodSub += sample;
                    //Array.Copy(window, 0, window, 1, lastIndex);
                    for (int i = 0; i < lastIndex; )
                        window[i] = window[++i];
                    window[lastIndex] = sample;
                }
                else // Not primed.
                {
                    window[windowCount] = sample;
                    periodSub += sample;
                    periodSum += sample * (windowCount + 1);
                    if (Length == ++windowCount)
                    {
                        primed = true;
                        value = periodSum / divider;
                    }
                }
                return value;
            }
        }
        #endregion

        #region Calculate
        /// <summary>
        /// Calculates the weighted moving average from the input array.
        /// The indicator is not primed during the first <c>ℓ-1</c> updates, where <c>ℓ</c> is the length.
        /// </summary>
        /// <param name="sampleList">The sample list.</param>
        /// <param name="length">The number of time periods of the weighted moving average.</param>
        /// <returns>A list of the weighted moving average values.</returns>
        public static List<double> Calculate(List<double> sampleList, int length)
        {
            if (2 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            int i, count = sampleList.Count;
            var resultList = new List<double>(count);
            if (count < length)
            {
                for (i = 0; i < count; i++)
                    resultList.Add(double.NaN);
            }
            else
            {
                double sample, periodSum = 0d, periodSub = 0d;
                int divider = length * (length + 1) / 2;
                int trailingIndex = 0;
                for (i = 0; i < length;)
                {
                    sample = sampleList[i++];
                    periodSub += sample;
                    periodSum += sample * i;
                }
                for (i = 1; i < length; i++)
                    resultList.Add(double.NaN);
                resultList.Add(periodSum / divider);
                for (i = length; i < count;)
                {
                    sample = sampleList[i++];
                    periodSum -= periodSub;
                    periodSum += sample * length;
                    resultList.Add(periodSum / divider);
                    periodSub -= sampleList[trailingIndex++];
                    periodSub += sample;
                }
            }
            return resultList;
        }
        #endregion
    }
}
