using System;
using System.Globalization;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators.Abstractions;

namespace Mbs.Trading.Indicators
{
    /// <summary>
    /// The triangular moving average (TRIMA) is a weighted moving average. Instead of the WMA who put more weight
    /// on the latest price bar, the triangular put more weight on the data in the middle of the specified period.
    /// <para>Using algebra, it can be demonstrated that the TRIMA is equivalent to doing a SMA of a SMA.
    /// The following explain the rules:</para>
    /// <para>➊ When the period π is even, TRIMA(x,π) = SMA(SMA(x,π/2), (π/2)+1).</para>
    /// <para>➋ When the period π is odd, TRIMA(x,π)=SMA(SMA(x,(π+1)/2), (π+1)/2).</para>
    /// <para>The SMA of a SMA is the algorithm generally found in books.</para>
    /// <para>TradeStation deviate from the generally accepted implementation by making the TRIMA to be as follows:</para>
    /// <para>TRIMA(x,π) = SMA(SMA(x, (int)(π/2)+1), (int)(π/2)+1 ).</para>
    /// <para>This formula is done regardless if the period is even or odd.</para>
    /// <para>In other words:</para>
    /// <para>➊ A period of 4 becomes TRIMA(x,4) = SMA(SMA(x,3), 3).</para>
    /// <para>➋ A period of 5 becomes TRIMA(x,5) = SMA(SMA(x,3), 3).</para>
    /// <para>➌ A period of 6 becomes TRIMA(x,6) = SMA(SMA(x,4), 4).</para>
    /// <para>➍ A period of 7 becomes TRIMA(x,7) = SMA(SMA(x,4), 4).</para>
    /// <para>The Metastock implementation is the same as the generally accepted one.</para>
    /// <para>To optimize speed, this implementation uses a better algorithm than the usual SMA of a SMA.</para>
    /// <para>The calculation from one TRIMA value to the next is done by doing 4 little adjustments.</para>
    /// <para>The following show a TRIMA 4-period:</para>
    /// <para>TRIMA at time δ: ((1*α)+(2*β)+(2*γ)+(1*δ)) / 6</para>
    /// <para>TRIMA at time ε: ((1*β)+(2*γ)+(2*δ)+(1*ε)) / 6</para>
    /// <para>To go from TRIMA δ to ε, the following is done:</para>
    /// <para>➊ α and β are subtract from the numerator.</para>
    /// <para>➋ δ is added to the numerator.</para>
    /// <para>➌ ε is added to the numerator.</para>
    /// <para>➍ TRIMA is calculated by doing numerator / 6.</para>
    /// <para>➎ Sequence is repeated for the next output.</para>
    /// </summary>
    public sealed class TriangularMovingAverage : LineIndicator
    {
        #region Members and accessors
        /// <summary>
        /// The length (in periods) of the triangular moving average.
        /// </summary>
        public int Length { get; }

        /// <summary>
        /// The current value of the the triangular moving average, or <c>NaN</c> if not primed.
        /// The indicator is not primed during the first <c>ℓ-1</c> updates, where e <c>ℓ</c> is the length.
        /// </summary>
        public double Value { get { lock (updateLock) { return primed ? value : double.NaN; } } }

        private readonly double factor;
        private double numerator;
        private double numeratorSub;
        private double numeratorAdd;
        private double value = double.NaN;
        private readonly double[] window;
        private readonly int length2;
        private int count;
        private readonly bool isOdd;

        private const string Trima = "trima";
        private const string TrimaFull = "Triangular Moving Average";
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the class.
        /// </summary>
        /// <param name="length">The length (the number of time periods) of the triangular moving average.</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public TriangularMovingAverage(int length, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(Trima, TrimaFull, ohlcvComponent)
        {
            if (2 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            Length = length;
            isOdd = (length % 2) == 1;
            length2 = length >> 1;
            int l = 1 + length2;
            if (isOdd)
            {
                // Let period = 5 and l=(int)(period/2), then  the formula for a "triangular" series is:
                // 1+2+3+2+1 = l*(l+1) + l+1 = (l+1)*(l+1) = 3*3 = 9.
                factor = 1d / (l * l);
            }
            else
            {
                // Let period = 6 and l=(int)(period/2), then  the formula for a "triangular" series is:
                // 1+2+3+3+2+1 = l*(l+1) = 3*4 = 12.
                factor = 1d / (length2 * l);
                --length2;
            }
            window = new double[length];
            Moniker = string.Concat(Trima, length.ToString(CultureInfo.InvariantCulture));
        }
        #endregion

        #region Reset
        /// <inheritdoc />
        public override void Reset()
        {
            lock (updateLock)
            {
                primed = false;
                count = 0;
                numerator = 0d;
                numeratorSub = 0d;
                numeratorAdd = 0d;
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
                    numerator -= numeratorSub;
                    numeratorSub -= window[0];
                    int j = Length - 1;
                    for (int i = 0; i < j; )
                        window[i] = window[++i];
                    window[j] = sample;
                    double temp = window[length2];
                    numeratorSub += temp;
                    if (isOdd) // The logic for an odd length.
                    {
                        numerator += numeratorAdd;
                        numeratorAdd -= temp;
                    }
                    else // The logic for an even length.
                    {
                        numeratorAdd -= temp;
                        numerator += numeratorAdd;
                    }
                    numeratorAdd += sample;
                    numerator += sample;
                    value = numerator * factor;
                }
                else
                {
                    window[count] = sample;
                    if (Length == ++count)
                    {
                        for (int i = length2; i >= 0; --i)
                        {
                            numeratorSub += window[i];
                            numerator += numeratorSub;
                        }
                        for (int i = length2 + 1; i < Length; ++i)
                        {
                            numeratorAdd += window[i];
                            numerator += numeratorAdd;
                        }
                        value = numerator * factor;
                        primed = true;
                    }
                }
                return value;
            }
        }
        #endregion
    }
}
