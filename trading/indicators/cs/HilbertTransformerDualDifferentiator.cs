using System;
using System.Runtime.Serialization;
using Mbst.Numerics;

namespace Mbst.Trading.Indicators.JohnEhlers
{
    /// <summary>
    /// A Hilbert transformer of WMA-smoothed and detrended data followed by the Dual Differentiator to determine the instant period.
    /// See John Ehlers, Rocket Science for Traders, Wiley, 2001, 0471405671, pp 70-74.
    /// </summary>
    [DataContract]
    public sealed class HilbertTransformerDualDifferentiator : IHilbertTransformerCycleEstimator
    {
        #region Members and accessors
        /// <summary>
        /// The WMA (linear-Weighted Moving Average) smoothing length.
        /// </summary>
        public int SmoothingLength => smoothingLength;

        /// <summary>
        /// The current WMA-smoothed value used by underlying Hilbert transformer.
        /// <para/>
        /// The linear-Weighted Moving Average has a window size of <c>SmoothingLength</c>.
        /// </summary>
        public double SmoothedValue => wmaSmoothed[0];

        /// <summary>
        /// The current detrended value.
        /// </summary>
        public double DetrendedValue => detrended[0];

        /// <summary>
        /// The current Quadrature component value.
        /// </summary>
        public double Quadrature => quadrature[0];

        /// <summary>
        /// The current InPhase component value.
        /// </summary>
        public double InPhase => inPhase[0];

        /// <summary>
        /// The current period value.
        /// </summary>
        public double Period => period;

        /// <summary>
        /// The current count value.
        /// </summary>
        public int Count => count;

        /// <summary>
        /// If the instance is primed.
        /// </summary>
        public bool IsPrimed => count > warmUpPeriod && isPrimed;

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
        public const double DefaultAlphaEmaQuadratureInPhase = 0.15;

        /// <summary>
        /// The default value of α (0 &lt; α ≤ 1) used in EMA to smooth the instantaneous period.
        /// </summary>
        public const double DefaultAlphaEmaPeriod = 0.15;

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
        private const int htLength = 7;
        private const int quadratureIndex = htLength / 2;

        [DataMember]
        private readonly int smoothingLengthPlusHtLengthMin1;
        [DataMember]
        private readonly int smoothingLengthPlus2HtLengthMin2;
        [DataMember]
        private readonly int smoothingLengthPlus3HtLengthMin3;
        [DataMember]
        private readonly int smoothingLengthPlus3HtLengthMin2;
        [DataMember]
        private readonly int smoothingLengthPlus3HtLengthMin1;

        [DataMember]
        private int count;
        [DataMember]
        private readonly int smoothingLength;
        [DataMember]
        private readonly double[] rawValues;
        [DataMember]
        private readonly double[] wmaFactors;
        [DataMember]
        private readonly double[] wmaSmoothed = new double[htLength];
        [DataMember]
        private readonly double[] detrended = new double[htLength];
        [DataMember]
        private readonly double[] inPhase = new double[htLength];
        [DataMember]
        private readonly double[] quadrature = new double[htLength];
        [DataMember]
        private readonly double[] jInPhase = new double[htLength];
        [DataMember]
        private readonly double[] jQuadrature = new double[htLength];
        [DataMember]
        private double smoothedInPhasePrevious;
        [DataMember]
        private double smoothedQuadraturePrevious;
        [DataMember]
        private double period;
        [DataMember]
        private readonly int warmUpPeriod;
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
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new instance of the class.
        /// </summary>
        /// <param name="smoothingLength">The WMA smoothing length. The valid values are 2, 3, 4. The default value is <c>DefaultSmoothingLength</c>.</param>
        /// <param name="warmUpPeriod">The number of updates before the indicator is primed. If less than the real primed period, then will be reset to the real primed period. The default value is <c>DefaultWarmUpPeriod</c>.</param>
        /// <param name="alphaEmaQuadratureInPhase">The value of α (0 &lt; α ≤ 1) used in EMA to smooth the in-phase and quadrature components. The default value is <c>DefaultAlphaEmaQuadratureInPhase</c>.</param>
        /// <param name="alphaEmaPeriod">The value of α (0 &lt; α ≤ 1) used in EMA to smooth the instantaneous period. The default value is <c>DefaultAlphaEmaPeriod</c>.</param>
        internal HilbertTransformerDualDifferentiator(int smoothingLength = DefaultSmoothingLength, int warmUpPeriod = DefaultWarmUpPeriod,
            double alphaEmaQuadratureInPhase = DefaultAlphaEmaQuadratureInPhase, double alphaEmaPeriod = DefaultAlphaEmaPeriod)
        {
            if (2 > smoothingLength || 4 < smoothingLength)
                throw new ArgumentOutOfRangeException(nameof(smoothingLength));
            if (0d > alphaEmaQuadratureInPhase || 1d < alphaEmaQuadratureInPhase)
                throw new ArgumentOutOfRangeException(nameof(alphaEmaQuadratureInPhase));
            if (0d > alphaEmaPeriod || 1d < alphaEmaPeriod)
                throw new ArgumentOutOfRangeException(nameof(alphaEmaPeriod));
            this.smoothingLength = smoothingLength;
            rawValues = new double[smoothingLength];
            wmaFactors = new double[smoothingLength];
            if (2 == smoothingLength)
            {
                wmaFactors[0] = 2d / 3d;
                wmaFactors[1] = 1d / 3d;
            }
            else if (3 == smoothingLength)
            {
                wmaFactors[0] = 3d / 6d;
                wmaFactors[1] = 2d / 6d;
                wmaFactors[2] = 1d / 6d;
            }
            else //if (4 == smoothingLength)
            {
                wmaFactors[0] = 4d / 10d;
                wmaFactors[1] = 3d / 10d;
                wmaFactors[2] = 2d / 10d;
                wmaFactors[3] = 1d / 10d;
            }
            smoothingLengthPlusHtLengthMin1 = smoothingLength + htLength - 1;
            smoothingLengthPlus2HtLengthMin2 = smoothingLengthPlusHtLengthMin1 + htLength - 1;
            smoothingLengthPlus3HtLengthMin3 = smoothingLengthPlus2HtLengthMin2 + htLength - 1;
            smoothingLengthPlus3HtLengthMin2 = smoothingLengthPlus3HtLengthMin3 + 1;
            smoothingLengthPlus3HtLengthMin1 = smoothingLengthPlus3HtLengthMin2 + 1;
            this.warmUpPeriod = warmUpPeriod < smoothingLengthPlus3HtLengthMin1 ? smoothingLengthPlus3HtLengthMin1 : warmUpPeriod;
            this.alphaEmaQuadratureInPhase = alphaEmaQuadratureInPhase;
            oneMinAlphaEmaQuadratureInPhase = 1d - alphaEmaQuadratureInPhase;
            this.alphaEmaPeriod = alphaEmaPeriod;
            oneMinAlphaEmaPeriod = 1d - alphaEmaPeriod;
            Reset();
        }
        #endregion

        #region Reset
        /// <summary>
        /// Resets the instance.
        /// </summary>
        public void Reset()
        {
            count = 0;
            rawValues.Initialize();
            wmaSmoothed.Initialize();
            detrended.Initialize();
            inPhase.Initialize();
            quadrature.Initialize();
            jInPhase.Initialize();
            jQuadrature.Initialize();
            smoothedInPhasePrevious = 0;
            smoothedQuadraturePrevious = 0;
            period = 0;
            isPrimed = false;
            isWarmedUp = false;
        }
        #endregion

        #region Update
        /// <summary>
        /// Updates the indicator.
        /// </summary>
        /// <param name="value">A new value.</param>
        public void Update(double value)
        {
            if (double.IsNaN(value))
                return;
            Push(rawValues, value);
            if (isPrimed)
            {
                if (!isWarmedUp && warmUpPeriod < ++count)
                    isWarmedUp = true;

                // The WMA is used to remove some high-frequency components before detrending the signal.
                Push(wmaSmoothed, Wma(rawValues));

                double amplitudeCorrectionFactor = CorrectAmplitude(period);

                // Since we have an amplitude-corrected Hilbert Transformer, and since we want to detrend
                // over its length, we simply use the Hilbert Transformer itself as the detrender.
                Push(detrended, Ht(wmaSmoothed) * amplitudeCorrectionFactor);

                // Compute both the in-phase and quadrature components of the detrended signal.
                Push(quadrature, Ht(detrended) * amplitudeCorrectionFactor);
                Push(inPhase, detrended[quadratureIndex]);

                // Complex averaging: apply the Hilbert Transformer to both the in-phase and quadrature components.
                // This advances the phase of each component by 90°.
                Push(jInPhase, Ht(inPhase) * amplitudeCorrectionFactor);
                Push(jQuadrature, Ht(quadrature) * amplitudeCorrectionFactor);

                // Phasor addition for 3 bar averaging followed by exponential moving average smoothing.
                double smoothedInPhase = EmaQuadratureInPhase(inPhase[0] - jQuadrature[0], smoothedInPhasePrevious);
                double smoothedQuadrature = EmaQuadratureInPhase(quadrature[0] + jInPhase[0], smoothedQuadraturePrevious);

                // Dual Differential discriminator.
                double discriminator = smoothedQuadrature * (smoothedInPhase - smoothedInPhasePrevious) -  smoothedInPhase * (smoothedQuadrature - smoothedQuadraturePrevious);
                smoothedInPhasePrevious = smoothedInPhase;
                smoothedQuadraturePrevious = smoothedQuadrature;

                double periodPrevious = period;
                period = Constants.TwoPi * (smoothedInPhase * smoothedInPhase + smoothedQuadrature * smoothedQuadrature) / discriminator;
                if (double.IsNaN(period) || double.IsInfinity(period))
                    period = periodPrevious;
                period = AdjustPeriod(period, periodPrevious);

                // Exponential moving average smoothing of the period.
                period = EmaPeriod(period, periodPrevious);
            }
            else
            {
                ++count;
                // On (smoothingLength)-th sample we calculate the first WMA smoothed value and begin with the detrender.
                if (smoothingLength > count)
                    return;

                Push(wmaSmoothed, Wma(rawValues));
                if (smoothingLengthPlusHtLengthMin1 > count)
                    return;

                double amplitudeCorrectionFactor = CorrectAmplitude(period);

                Push(detrended, Ht(wmaSmoothed) * amplitudeCorrectionFactor);
                if (smoothingLengthPlus2HtLengthMin2 > count)
                    return;

                Push(quadrature, Ht(detrended) * amplitudeCorrectionFactor);
                Push(inPhase, detrended[quadratureIndex]);
                if (smoothingLengthPlus3HtLengthMin3 > count)
                    return;

                Push(jInPhase, Ht(inPhase) * amplitudeCorrectionFactor);
                Push(jQuadrature, Ht(quadrature) * amplitudeCorrectionFactor);

                if (smoothingLengthPlus3HtLengthMin3 == count)
                {
                    smoothedInPhasePrevious = inPhase[0] - jQuadrature[0];
                    smoothedQuadraturePrevious = quadrature[0] + jInPhase[0];
                    return;
                }
                double smoothedInPhase = EmaQuadratureInPhase(inPhase[0] - jQuadrature[0], smoothedInPhasePrevious);
                double smoothedQuadrature = EmaQuadratureInPhase(quadrature[0] + jInPhase[0], smoothedQuadraturePrevious);

                double discriminator = smoothedQuadrature * (smoothedInPhase - smoothedInPhasePrevious) - smoothedInPhase * (smoothedQuadrature - smoothedQuadraturePrevious);
                smoothedInPhasePrevious = smoothedInPhase;
                smoothedQuadraturePrevious = smoothedQuadrature;

                if (smoothingLengthPlus3HtLengthMin2 == count)
                {
                    period = Constants.TwoPi * (smoothedInPhase * smoothedInPhase +
                        smoothedQuadrature * smoothedQuadrature) / discriminator;
                    if (double.IsNaN(period) || double.IsInfinity(period))
                        period = minPeriod;
                    period = AdjustPeriod(period, period);
                    return;
                }

                double periodPrevious = period;
                period = Constants.TwoPi * (smoothedInPhase * smoothedInPhase + smoothedQuadrature * smoothedQuadrature) / discriminator;
                if (double.IsNaN(period) || double.IsInfinity(period))
                    period = periodPrevious;
                period = AdjustPeriod(period, periodPrevious);

                period = EmaPeriod(period, periodPrevious);
                isPrimed = true;
            }
        }
        #endregion

        #region Implementation
        private static void Push(double[] array, double value)
        {
            Array.Copy(array, 0, array, 1, array.Length - 1);
            array[0] = value;
        }

        private double Wma(double[] array)
        {
            double value = 0;
            for (int i = 0; i < smoothingLength; ++i)
                value += wmaFactors[i] * array[i];
            return value;
        }

        private static double Ht(double[] array)
        {
            const double a = 0.0962;
            const double b = 0.5769;
            double value = 0;
            value += a * array[0];
            value += b * array[2];
            value -= b * array[4];
            value -= a * array[6];
            return value;
        }

        private double EmaPeriod(double value, double valuePrevious)
        {
            return alphaEmaPeriod * value + oneMinAlphaEmaPeriod * valuePrevious;
        }

        private double EmaQuadratureInPhase(double value, double valuePrevious)
        {
            return alphaEmaQuadratureInPhase * value + oneMinAlphaEmaQuadratureInPhase * valuePrevious;
        }

        private static double CorrectAmplitude(double previousPeriod)
        {
            const double a = 0.54;
            const double b = 0.075;
            return a + b * previousPeriod;
        }

        private static double AdjustPeriod(double period, double periodPrevious)
        {
            const double minPreviousPeriodFactor = 0.67;
            const double maxPreviousPeriodFactor = 1.5;
            double temp = maxPreviousPeriodFactor * periodPrevious;
            if (period > temp)
                period = temp;
            else
            {
                temp = minPreviousPeriodFactor * periodPrevious;
                if (period < temp)
                    period = temp;
            }
            if (period < minPeriod)
                period = minPeriod;
            else if (period > maxPeriod)
                period = maxPeriod;
            return period;
        }
        #endregion
    }
}
