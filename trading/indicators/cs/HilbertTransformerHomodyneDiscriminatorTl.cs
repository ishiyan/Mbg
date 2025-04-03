using System;
using System.Runtime.Serialization;
using Mbst.Numerics;

namespace Mbst.Trading.Indicators.JohnEhlers
{
    /// <summary>
    /// A Hilbert Transformer of WMA-smoothed and detrended data with the Homodyne Discriminator applied.
    /// TA-Lib implementation with unrolled loops.
    /// </summary>
    [DataContract]
    public sealed class HilbertTransformerHomodyneDiscriminatorTl : IHilbertTransformerCycleEstimator
    {
        #region Members and accessors
        /// <summary>
        /// The current WMA smoothed value used by underlying Hilbert transformer.
        /// </summary>
        public double SmoothedValue => smoothedValue;

        /// <summary>
        /// The current detrended value.
        /// </summary>
        public double DetrendedValue => detrendedValue;

        /// <summary>
        /// The current Quadrature component value.
        /// </summary>
        public double Quadrature => quadrature;

        /// <summary>
        /// The current InPhase component value.
        /// </summary>
        public double InPhase => inPhase;

        /// <summary>
        /// The current period value.
        /// </summary>
        public double Period => period;

        /// <summary>
        /// The current count value.
        /// </summary>
        public int Count => count;

        /// <summary>
        /// The WMA smoothing length.
        /// </summary>
        public int SmoothingLength => smoothingLength;

        /// <summary>
        /// If the instance is primed.
        /// </summary>
        public bool IsPrimed => isWarmedUp; // && isPrimed;

        /// <summary>
        /// The minimal cycle period.
        /// </summary>
        public int MinPeriod => minPeriod;

        /// <summary>
        /// The maximual cycle period.
        /// </summary>
        public int MaxPeriod => maxPeriod;

        /// <summary>
        /// The value of α (0 &lt; α ≤ 1) used in EMA to smooth the in-phase and quadrature components.
        /// </summary>
        public double AlphaEmaQuadratureInPhase => alphaEmaQuadratureInPhase;

        /// <summary>
        /// The value of α (0 &lt; α ≤ 1) used in EMA to smooth the instantaneous period.
        /// </summary>
        public double AlphaEmaPeriod => alphaEmaPeriod;

        /// <summary>
        /// The value of the number of updates before the indicator is primed (MaxPeriod * 2 = 100).
        /// </summary>
        public int WarmUpPeriod => warmUpPeriod;

        /// <summary>
        /// The default value of α (0 &lt; α ≤ 1) used in EMA to smooth the in-phase and quadrature components.
        /// </summary>
        public const double DefaultAlphaEmaQuadratureInPhase = 0.2;

        /// <summary>
        /// The default value of α (0 &lt; α ≤ 1) used in EMA to smooth the instantaneous period.
        /// </summary>
        public const double DefaultAlphaEmaPeriod = 0.2;

        /// <summary>
        /// The default value of the WMA (linear-Weighted Moving Average) smoothing length.
        /// </summary>
        public const int DefaultSmoothingLength = 4;

        /// <summary>
        /// The default value of the number of updates before the indicator is primed (MaxPeriod * 2 = 100).
        /// </summary>
        public const int DefaultWarmUpPeriod = maxPeriod * 2;

        private const int minPeriod = 6;
        private const int maxPeriod = 50;
        private const double minPreviousPeriodFactor = 0.67;
        private const double maxPreviousPeriodFactor = 1.5;
        private const double a = 0.0962, b = 0.5769;

        // The TA-Lib implementation uses the following lookback value with hardcoded smoothingLength=4.
        //
        // The fixed lookback is 32 and is establish as follows:
        // 12 price bar to be compatible with the implementation of Tradestation found in John Ehlers book,
        // 6 price bars for the Detrender,
        // 6 price bars for Q1,
        // 3 price bars for jI,
        // 3 price bars for jQ,
        // 1 price bar for Re/Im,
        // 1 price bar for the Delta Phase,
        // --------
        // 32 Total.
        //
        // The first 9 bars are not used by TA-Lib, they are just skipped for the compatibility with the Tradestation.
        // We do not skip them. Thus, we use the fixed lookback value 32 - 9 = 23.
        private const int primedCount = 23;

        [DataMember]
        private int count;
        [DataMember]
        private readonly int smoothingLength;
        [DataMember]
        private readonly int warmUpPeriod;
        [DataMember]
        private int index;
        [DataMember]
        private double period = minPeriod;
        [DataMember]
        private double smoothedValue;
        [DataMember]
        private double detrendedValue;
        [DataMember]
        private double inPhase;
        [DataMember]
        private double quadrature;
        [DataMember]
        private readonly double smoothingMultiplier;
        [DataMember]
        private double adjustedPeriod;
        [DataMember]
        private double i2Previous;
        [DataMember]
        private double q2Previous;
        [DataMember]
        private double re;
        [DataMember]
        private double im;
        [DataMember]
        private readonly double alphaEmaQuadratureInPhase;
        [DataMember]
        private readonly double oneMinAlphaEmaQuadratureInPhase;
        [DataMember]
        private readonly double alphaEmaPeriod;
        [DataMember]
        private readonly double oneMinAlphaEmaPeriod;
        [DataMember]
        private bool isPrimed;
        [DataMember]
        private bool isWarmedUp;

        #region WMA smoother private members
        [DataMember]
        private double wmaSum;
        [DataMember]
        private double wmaSub;
        [DataMember]
        private double wmaInput1;
        [DataMember]
        private double wmaInput2;
        [DataMember]
        private double wmaInput3;
        [DataMember]
        private double wmaInput4;
        #endregion

        #region Detrender private members
        [DataMember]
        private double detrenderOdd0;
        [DataMember]
        private double detrenderOdd1;
        [DataMember]
        private double detrenderOdd2;
        [DataMember]
        private double detrenderPreviousOdd;
        [DataMember]
        private double detrenderPreviousInputOdd;
        [DataMember]
        private double detrenderEven0;
        [DataMember]
        private double detrenderEven1;
        [DataMember]
        private double detrenderEven2;
        [DataMember]
        private double detrenderPreviousEven;
        [DataMember]
        private double detrenderPreviousInputEven;
        #endregion

        #region Quadrature (Q1) component private members
        [DataMember]
        private double q1Odd0;
        [DataMember]
        private double q1Odd1;
        [DataMember]
        private double q1Odd2;
        [DataMember]
        private double q1PreviousOdd;
        [DataMember]
        private double q1PreviousInputOdd;
        [DataMember]
        private double q1Even0;
        [DataMember]
        private double q1Even1;
        [DataMember]
        private double q1Even2;
        [DataMember]
        private double q1PreviousEven;
        [DataMember]
        private double q1PreviousInputEven;
        #endregion

        #region InPhase (I1) private members
        [DataMember]
        private double i1Previous1Odd;
        [DataMember]
        private double i1Previous2Odd;
        [DataMember]
        private double i1Previous1Even;
        [DataMember]
        private double i1Previous2Even;
        #endregion

        #region jI private members
        [DataMember]
        private double jiOdd0;
        [DataMember]
        private double jiOdd1;
        [DataMember]
        private double jiOdd2;
        [DataMember]
        private double jiPreviousOdd;
        [DataMember]
        private double jiPreviousInputOdd;
        [DataMember]
        private double jiEven0;
        [DataMember]
        private double jiEven1;
        [DataMember]
        private double jiEven2;
        [DataMember]
        private double jiPreviousEven;
        [DataMember]
        private double jiPreviousInputEven;
        #endregion

        #region jQ private members
        [DataMember]
        private double jqOdd0;
        [DataMember]
        private double jqOdd1;
        [DataMember]
        private double jqOdd2;
        [DataMember]
        private double jqPreviousOdd;
        [DataMember]
        private double jqPreviousInputOdd;
        [DataMember]
        private double jqEven0;
        [DataMember]
        private double jqEven1;
        [DataMember]
        private double jqEven2;
        [DataMember]
        private double jqPreviousEven;
        [DataMember]
        private double jqPreviousInputEven;
        #endregion
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the class.
        /// </summary>
        /// <param name="smoothingLength">The WMA smoothing length. The valid values are 2, 3, 4. The default value is <c>DefaultSmoothingLength</c>.</param>
        /// <param name="warmUpPeriod">The number of updates before the indicator is primed. If less than the real primed period, then will be reset to the real primed period. The default value is <c>DefaultWarmUpPeriod</c>.</param>
        /// <param name="alphaEmaQuadratureInPhase">The value of α (0 &lt; α ≤ 1) used in EMA to smooth the in-phase and quadrature components. The default value is <c>DefaultAlphaEmaQuadratureInPhase</c>.</param>
        /// <param name="alphaEmaPeriod">The value of α (0 &lt; α ≤ 1) used in EMA to smooth the instantaneous period. The default value is <c>DefaultAlphaEmaPeriod</c>.</param>
        internal HilbertTransformerHomodyneDiscriminatorTl(int smoothingLength = DefaultSmoothingLength, int warmUpPeriod = DefaultWarmUpPeriod,
            double alphaEmaQuadratureInPhase = DefaultAlphaEmaQuadratureInPhase, double alphaEmaPeriod = DefaultAlphaEmaPeriod)
        {
            if (2 > smoothingLength || 4 < smoothingLength)
                throw new ArgumentOutOfRangeException(nameof(smoothingLength));
            if (0d > alphaEmaQuadratureInPhase || 1d < alphaEmaQuadratureInPhase)
                throw new ArgumentOutOfRangeException(nameof(alphaEmaQuadratureInPhase));
            if (0d > alphaEmaPeriod || 1d < alphaEmaPeriod)
                throw new ArgumentOutOfRangeException(nameof(alphaEmaPeriod));
            this.smoothingLength = smoothingLength;
            if (2 == smoothingLength)
                smoothingMultiplier = 1d / 3d;
            else if (3 == smoothingLength)
                smoothingMultiplier = 1d / 6d;
            else //if (4 == smoothingLength)
                smoothingMultiplier = 0.1;
            this.warmUpPeriod = warmUpPeriod < primedCount ? primedCount : warmUpPeriod;
            this.alphaEmaQuadratureInPhase = alphaEmaQuadratureInPhase;
            oneMinAlphaEmaQuadratureInPhase = 1d - alphaEmaQuadratureInPhase;
            this.alphaEmaPeriod = alphaEmaPeriod;
            oneMinAlphaEmaPeriod = 1d - alphaEmaPeriod;
        }
        #endregion

        #region Reset
        /// <summary>
        /// Resets the instance.
        /// </summary>
        public void Reset()
        {
            isWarmedUp = false; isPrimed = false;
            count = 0; index = 0; period = minPeriod; adjustedPeriod = 0; re = 0; im = 0; detrendedValue = 0;
            smoothedValue = 0; quadrature = 0; inPhase = 0; i2Previous = 0; q2Previous = 0;
            wmaSum = 0; wmaSub = 0; wmaInput1 = 0; wmaInput2 = 0; wmaInput3 = 0; wmaInput4 = 0;
            detrenderOdd0 = 0; detrenderOdd1 = 0; detrenderOdd2 = 0; detrenderPreviousOdd = 0; detrenderPreviousInputOdd = 0;
            detrenderEven0 = 0; detrenderEven1 = 0; detrenderEven2 = 0; detrenderPreviousEven = 0; detrenderPreviousInputEven = 0;
            q1Odd0 = 0; q1Odd1 = 0; q1Odd2 = 0; q1PreviousOdd = 0; q1PreviousInputOdd = 0;
            q1Even0 = 0; q1Even1 = 0; q1Even2 = 0; q1PreviousEven = 0; q1PreviousInputEven = 0;
            i1Previous1Odd = 0; i1Previous2Odd = 0; i1Previous1Even = 0; i1Previous2Even = 0;
            jiOdd0 = 0; jiOdd1 = 0; jiOdd2 = 0; jiPreviousOdd = 0; jiPreviousInputOdd = 0;
            jiEven0 = 0; jiEven1 = 0; jiEven2 = 0; jiPreviousEven = 0; jiPreviousInputEven = 0;
            jqOdd0 = 0; jqOdd1 = 0; jqOdd2 = 0; jqPreviousOdd = 0; jqPreviousInputOdd = 0;
            jqEven0 = 0; jqEven1 = 0; jqEven2 = 0; jqPreviousEven = 0; jqPreviousInputEven = 0;
        }
        #endregion

        #region Update
        /// <summary>
        /// Updates the indicator.
        /// </summary>
        /// <param name="value">A new value.</param>
        public void Update(double value)
        {
            double val;
            #region WMA
            // We need (smoothingLength - 1) bars to accumulate the WMA sub and sum.
            // On (smoothingLength)-th bar we calculate the first WMA smoothed value and begin with detrender.
            if (smoothingLength >= ++count)
            {
                if (1 == count)
                {
                    wmaSub = value; wmaInput1 = value; wmaSum = value;
                }
                else if (2 == count)
                {
                    wmaSub += value; wmaInput2 = value; wmaSum += value * 2d;
                    if (2 == smoothingLength)
                    {
                        val = wmaSum * smoothingMultiplier;
                        goto DetrendLabel;
                    }
                }
                else if (3 == count)
                {
                    wmaSub += value; wmaInput3 = value; wmaSum += value * 3d;
                    if (3 == smoothingLength)
                    {
                        val = wmaSum * smoothingMultiplier;
                        goto DetrendLabel;
                    }
                }
                else //if (4 == inputCount)
                {
                    wmaSub += value; wmaInput4 = value; wmaSum += value * 4d;
                    val = wmaSum * smoothingMultiplier;
                    goto DetrendLabel;
                }
                return;
            }
            wmaSum -= wmaSub;
            wmaSum += value * smoothingLength;
            val = wmaSum * smoothingMultiplier;
            wmaSub += value; wmaSub -= wmaInput1;
            wmaInput1 = wmaInput2;
            if (4 == smoothingLength)
            {
                wmaInput2 = wmaInput3; wmaInput3 = wmaInput4; wmaInput4 = value;
            }
            else if (3 == smoothingLength)
            {
                wmaInput2 = wmaInput3; wmaInput3 = value;
            }
            else //if (2 == smoothingLength)
            {
                wmaInput2 = value;
            }
        DetrendLabel:
            smoothedValue = val;
            #endregion
            if (!isWarmedUp)
            {
                isWarmedUp = count > warmUpPeriod;
                if (!isPrimed)
                    isPrimed = count > primedCount;
            }
            double detrender, ji, jq;
            double temp = a * smoothedValue; adjustedPeriod = 0.075 * period + 0.54;
            if (0 == count % 2) // Even value count.
            {
                #region explicitely expanded index
                if (0 == index)
                {
                    index = 1;
                    detrender = -detrenderEven0; detrenderEven0 = temp; detrender += temp; detrender -= detrenderPreviousEven;
                    detrenderPreviousEven = b * detrenderPreviousInputEven; detrenderPreviousInputEven = val;
                    detrender += detrenderPreviousEven; detrender *= adjustedPeriod;
                    // Quadrature component.
                    temp = a * detrender; quadrature = -q1Even0; q1Even0 = temp; quadrature += temp; quadrature -= q1PreviousEven;
                    q1PreviousEven = b * q1PreviousInputEven; q1PreviousInputEven = detrender;
                    quadrature += q1PreviousEven; quadrature *= adjustedPeriod;
                    // Advance the phase of the InPhase component by 90°.
                    temp = a * i1Previous2Even; ji = -jiEven0; jiEven0 = temp; ji += temp; ji -= jiPreviousEven;
                    jiPreviousEven = b * jiPreviousInputEven; jiPreviousInputEven = i1Previous2Even;
                    ji += jiPreviousEven; ji *= adjustedPeriod;
                    // Advance the phase of the Quadrature component by 90°.
                    temp = a * quadrature; jq = -jqEven0; jqEven0 = temp;
                }
                else if (1 == index)
                {
                    index = 2;
                    detrender = -detrenderEven1; detrenderEven1 = temp; detrender += temp; detrender -= detrenderPreviousEven;
                    detrenderPreviousEven = b * detrenderPreviousInputEven; detrenderPreviousInputEven = val;
                    detrender += detrenderPreviousEven; detrender *= adjustedPeriod;
                    // Quadrature component.
                    temp = a * detrender; quadrature = -q1Even1; q1Even1 = temp; quadrature += temp; quadrature -= q1PreviousEven;
                    q1PreviousEven = b * q1PreviousInputEven; q1PreviousInputEven = detrender;
                    quadrature += q1PreviousEven; quadrature *= adjustedPeriod;
                    // Advance the phase of the InPhase component by 90°.
                    temp = a * i1Previous2Even; ji = -jiEven1; jiEven1 = temp; ji += temp; ji -= jiPreviousEven;
                    jiPreviousEven = b * jiPreviousInputEven; jiPreviousInputEven = i1Previous2Even;
                    ji += jiPreviousEven; ji *= adjustedPeriod;
                    // Advance the phase of the Quadrature component by 90°.
                    temp = a * quadrature; jq = -jqEven1; jqEven1 = temp;
                }
                else //if (2 == index)
                {
                    index = 0;
                    detrender = -detrenderEven2; detrenderEven2 = temp; detrender += temp; detrender -= detrenderPreviousEven;
                    detrenderPreviousEven = b * detrenderPreviousInputEven; detrenderPreviousInputEven = val;
                    detrender += detrenderPreviousEven; detrender *= adjustedPeriod;
                    // Quadrature component.
                    temp = a * detrender; quadrature = -q1Even2; q1Even2 = temp; quadrature += temp; quadrature -= q1PreviousEven;
                    q1PreviousEven = b * q1PreviousInputEven; q1PreviousInputEven = detrender;
                    quadrature += q1PreviousEven; quadrature *= adjustedPeriod;
                    // Advance the phase of the InPhase component by 90°.
                    temp = a * i1Previous2Even; ji = -jiEven2; jiEven2 = temp; ji += temp; ji -= jiPreviousEven;
                    jiPreviousEven = b * jiPreviousInputEven; jiPreviousInputEven = i1Previous2Even;
                    ji += jiPreviousEven; ji *= adjustedPeriod;
                    // Advance the phase of the Quadrature component by 90°.
                    temp = a * quadrature; jq = -jqEven2; jqEven2 = temp;
                }
                #endregion
                // Advance the phase of the Quadrature component by 90° (continued).
                jq += temp; jq -= jqPreviousEven;
                jqPreviousEven = b * jqPreviousInputEven; jqPreviousInputEven = quadrature;
                jq += jqPreviousEven; jq *= adjustedPeriod;
                // InPhase component.
                inPhase = i1Previous2Even;
                // The current detrender value will be used by the "odd" logic later.
                i1Previous2Odd = i1Previous1Odd;
                i1Previous1Odd = detrender;
            }
            else // Odd value count.
            {
                #region explicitely expanded index
                if (0 == index)
                {
                    index = 1;
                    detrender = -detrenderOdd0; detrenderOdd0 = temp; detrender += temp; detrender -= detrenderPreviousOdd;
                    detrenderPreviousOdd = b * detrenderPreviousInputOdd; detrenderPreviousInputOdd = val;
                    detrender += detrenderPreviousOdd; detrender *= adjustedPeriod;
                    // Quadrature component.
                    temp = a * detrender; quadrature = -q1Odd0; q1Odd0 = temp; quadrature += temp; quadrature -= q1PreviousOdd;
                    q1PreviousOdd = b * q1PreviousInputOdd; q1PreviousInputOdd = detrender;
                    quadrature += q1PreviousOdd; quadrature *= adjustedPeriod;
                    // Advance the phase of the InPhase component by 90°.
                    temp = a * i1Previous2Odd; ji = -jiOdd0; jiOdd0 = temp; ji += temp; ji -= jiPreviousOdd;
                    jiPreviousOdd = b * jiPreviousInputOdd; jiPreviousInputOdd = i1Previous2Odd;
                    ji += jiPreviousOdd; ji *= adjustedPeriod;
                    // Advance the phase of the Quadrature component by 90°.
                    temp = a * quadrature; jq = -jqOdd0; jqOdd0 = temp;
                }
                else if (1 == index)
                {
                    index = 2;
                    // Quadrature component.
                    detrender = -detrenderOdd1; detrenderOdd1 = temp; detrender += temp; detrender -= detrenderPreviousOdd;
                    detrenderPreviousOdd = b * detrenderPreviousInputOdd; detrenderPreviousInputOdd = val;
                    detrender += detrenderPreviousOdd; detrender *= adjustedPeriod;
                    temp = a * detrender; quadrature = -q1Odd1; q1Odd1 = temp; quadrature += temp; quadrature -= q1PreviousOdd;
                    q1PreviousOdd = b * q1PreviousInputOdd; q1PreviousInputOdd = detrender;
                    quadrature += q1PreviousOdd; quadrature *= adjustedPeriod;
                    // Advance the phase of the InPhase component by 90°.
                    temp = a * i1Previous2Odd; ji = -jiOdd1; jiOdd1 = temp; ji += temp; ji -= jiPreviousOdd;
                    jiPreviousOdd = b * jiPreviousInputOdd; jiPreviousInputOdd = i1Previous2Odd;
                    ji += jiPreviousOdd; ji *= adjustedPeriod;
                    // Advance the phase of the Quadrature component by 90°.
                    temp = a * quadrature; jq = -jqOdd1; jqOdd1 = temp;
                }
                else //if (2 == index)
                {
                    index = 0;
                    detrender = -detrenderOdd2; detrenderOdd2 = temp; detrender += temp; detrender -= detrenderPreviousOdd;
                    detrenderPreviousOdd = b * detrenderPreviousInputOdd; detrenderPreviousInputOdd = val;
                    detrender += detrenderPreviousOdd; detrender *= adjustedPeriod;
                    // Quadrature component.
                    temp = a * detrender; quadrature = -q1Odd2; q1Odd2 = temp; quadrature += temp; quadrature -= q1PreviousOdd;
                    q1PreviousOdd = b * q1PreviousInputOdd; q1PreviousInputOdd = detrender;
                    quadrature += q1PreviousOdd; quadrature *= adjustedPeriod;
                    // Advance the phase of the InPhase component by 90°.
                    temp = a * i1Previous2Odd; ji = -jiOdd2; jiOdd2 = temp; ji += temp; ji -= jiPreviousOdd;
                    jiPreviousOdd = b * jiPreviousInputOdd; jiPreviousInputOdd = i1Previous2Odd;
                    ji += jiPreviousOdd; ji *= adjustedPeriod;
                    // Advance the phase of the Quadrature component by 90°.
                    temp = a * quadrature; jq = -jqOdd2; jqOdd2 = temp;
                }
                #endregion
                // Advance the phase of the Quadrature component by 90° (continued).
                jq += temp; jq -= jqPreviousOdd;
                jqPreviousOdd = b * jqPreviousInputOdd; jqPreviousInputOdd = quadrature;
                jq += jqPreviousOdd; jq *= adjustedPeriod;
                // InPhase component.
                inPhase = i1Previous2Odd;
                // The current detrender value will be used by the "even" logic later.
                i1Previous2Even = i1Previous1Even;
                i1Previous1Even = detrender;
            }
            detrendedValue = detrender;
            // Phasor addition for 3 bar averaging.
            double i2 = inPhase - jq;
            double q2 = quadrature + ji;
            // Smooth the InPhase and the Quadrature components before applying the discriminator.
            i2 = alphaEmaQuadratureInPhase * i2 + oneMinAlphaEmaQuadratureInPhase * i2Previous;
            q2 = alphaEmaQuadratureInPhase * q2 + oneMinAlphaEmaQuadratureInPhase * q2Previous;
            // Homodyne discriminator.
            // Homodyne means we are multiplying the signal by itself.
            // We multiply the signal of the current bar with the complex conjugate of the signal 1 bar ago.
            re = alphaEmaQuadratureInPhase * (i2 * i2Previous + q2 * q2Previous) + oneMinAlphaEmaQuadratureInPhase * re;
            im = alphaEmaQuadratureInPhase * (i2 * q2Previous - q2 * i2Previous) + oneMinAlphaEmaQuadratureInPhase * im;
            q2Previous = q2;
            i2Previous = i2;
            temp = period;
            period = Constants.TwoPi / Math.Atan(im / re);//TODO: Math.Atan2(im, re)
            if (double.IsNaN(period) || double.IsInfinity(period))
                period = temp;
            val = maxPreviousPeriodFactor * temp;
            if (period > val)
                period = val;
            else
            {
                val = minPreviousPeriodFactor * temp;
                if (period < val)
                    period = val;
            }
            if (period < minPeriod)
                period = minPeriod;
            else if (period > maxPeriod)
                period = maxPeriod;
            period = alphaEmaPeriod * period + oneMinAlphaEmaPeriod * temp;
        }
        #endregion
    }
}
