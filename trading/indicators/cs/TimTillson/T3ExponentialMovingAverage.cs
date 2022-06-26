using System;
using System.Globalization;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators.Abstractions;

namespace Mbs.Trading.Indicators.TimTillson
{
    /// <summary>
    /// The T3 Exponential Moving Average is a smoothing indicator with less lag than a straight exponential moving average.
    /// <para>In filter theory parlance, T3 is a six-pole non-linear Kalman filter.</para>
    /// <para>The T3 was developed by Tim Tillson and is described in the article:</para>
    /// <para>Technical Analysis of Stocks &amp; Commodities, November 1998, Better Moving Averages.</para>
    /// <para>The calculation is as follows:</para>
    /// <para><c>EMA¹ᵢ = EMA(Pᵢ) = αPᵢ + (1-α)EMA¹ᵢ₋₁ = EMA¹ᵢ₋₁ + α(Pᵢ - EMA¹ᵢ₋₁), 0 &lt; α ≤ 1</c></para>
    /// <para><c>EMA²ᵢ = EMA(EMA¹ᵢ) = αEMA¹ᵢ + (1-α)EMA²ᵢ₋₁ = EMA²ᵢ₋₁ + α(EMA¹ᵢ - EMA²ᵢ₋₁), 0 &lt; α ≤ 1</c></para>
    /// <para><c>GDᵛᵢ = (1+ν)EMA¹ᵢ - νEMA²ᵢ = EMA¹ᵢ + ν(EMA¹ᵢ - EMA²ᵢ), 0 &lt; ν ≤ 1</c></para>
    /// <para><c>T3ᵢ = GDᵛᵢ(GDᵛᵢ(GDᵛᵢ))</c></para>
    /// <para>where <c>GD</c> stands for 'Generalized DEMA' with 'volume' ν. The default value of <c>ν</c> is 0.7.</para>
    /// <para>When <c>ν=0</c>, GD is just an EMA, and when <c>ν=1</c>, GD is DEMA. In between, GD is a cooler DEMA.</para>
    /// </summary>
    /// <remarks>
    /// If <c>x</c> stands for the action of running a time series through an EMA,
    /// <para> <c>ƒ</c> is our formula for Generalized Dema with 'volume' ν: <c>ƒ = (1+ν)x -νx²</c></para>
    /// <para>Running the filter though itself three times is equivalent to cubing <c>ƒ</c>:</para>
    /// <para><c>-ν³x⁶ + (3ν²+3ν³)x⁵ + (-6ν²-3ν-3ν³)x⁴ + (1+3ν+ν³+3ν²)x³</c></para>
    /// <para>The Metastock code for T3 is:</para>
    /// <para>e1=Mov(P,periods,E)</para>
    /// <para>e2=Mov(e1,periods,E)</para>
    /// <para>e3=Mov(e2,periods,E)</para>
    /// <para>e4=Mov(e3,periods,E)</para>
    /// <para>e5=Mov(e4,periods,E)</para>
    /// <para>e6=Mov(e5,periods,E)</para>
    /// <para>c1=-ν³</para>
    /// <para>c2=3ν²+3ν³</para>
    /// <para>c3=-6*ν²-3ν-3ν³</para>
    /// <para>c4=1+3ν+ν³+3ν²</para>
    /// <para>t3=c1*e6+c2*e5+c3*e4+c4*e3</para>
    /// </remarks>
    public sealed class T3ExponentialMovingAverage : Indicator, ILineIndicator
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
        /// The ν-factor of the T3 exponential moving average, <c>0 ≤ ν ≤ 1</c>.
        /// </summary>
        public double VFactor { get; }

        private double value = double.NaN;
        /// <summary>
        /// The current value of the the T3 exponential moving average.
        /// </summary>
        public double Value { get { lock (updateLock) { return value; } } }

        private const double Epsilon = 0.00000001;
        private const string T3Ema = "T3EMA";
        private const string T3EmaFull = "T3 Exponential Moving Average";

        private double e1;
        private double e2;
        private double e3;
        private double e4;
        private double e5;
        private double e6;
        private readonly double c1;
        private readonly double c2;
        private readonly double c3;
        private readonly double c4;
        private readonly double oneMinSmoothingFactor;
        private readonly int length2;
        private readonly int length3;
        private readonly int length4;
        private readonly int length5;
        private readonly int length6;
        private int count;
        #endregion

        #region Construction
        /// <summary>
        /// Constructs a new a new instance of the <see cref="T3ExponentialMovingAverage"/> class.
        /// </summary>
        /// <param name="length">The length (<c>ℓ</c>) of the exponential moving average.
        /// The equivalent smoothing factor (<c>α</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public T3ExponentialMovingAverage(int length, OhlcvComponent ohlcvComponent)
            : this(length, 0.7, ohlcvComponent)
        {
        }

        /// <summary>
        /// Constructs a new a new instance of the <see cref="T3ExponentialMovingAverage"/> class.
        /// </summary>
        /// <param name="length">The length (<c>ℓ</c>) of the exponential moving average.
        /// The equivalent smoothing factor (<c>α</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </param>
        /// <param name="vFactor">The ν-factor the T3 exponential moving average, <c>0 ≤ ν ≤ 1</c>.</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public T3ExponentialMovingAverage(int length, double vFactor = 0.7, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(T3Ema, T3EmaFull, ohlcvComponent)
        {
            if (1 > length)
                throw new ArgumentOutOfRangeException(nameof(length));
            if (0d > vFactor || 1d < vFactor)
                throw new ArgumentOutOfRangeException(nameof(vFactor));
            Length = length;
            length2 = length * 2 - 1;
            length3 = length * 3 - 2;
            length4 = length * 4 - 3;
            length5 = length * 5 - 4;
            length6 = length * 6 - 5;
            VFactor = vFactor;
            double v2 = vFactor * vFactor;
            c1 = -v2 * vFactor;
            c2 = 3d * (v2 - c1);
            c3 = -6d * v2 - 3d * (vFactor - c1);
            c4 = 1d + 3d * vFactor - c1 + 3d * v2;
            SmoothingFactor = 2d / (1d + length);
            oneMinSmoothingFactor = 1d - SmoothingFactor;
            Moniker = string.Concat(T3Ema, length.ToString(CultureInfo.InvariantCulture));
        }

        /// <summary>
        /// Constructs a new a new instance of the <see cref="T3ExponentialMovingAverage"/> class.
        /// </summary>
        /// <param name="smoothingFactor">The smoothing factor (<c>α</c>) of the exponential moving average.
        /// The equivalent length (<c>ℓ</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public T3ExponentialMovingAverage(double smoothingFactor, OhlcvComponent ohlcvComponent)
            : this(smoothingFactor, 0.7, ohlcvComponent)
        {
        }

        /// <summary>
        /// Constructs a new a new instance of the <see cref="T3ExponentialMovingAverage"/> class.
        /// </summary>
        /// <param name="smoothingFactor">The smoothing factor (<c>α</c>) of the exponential moving average.
        /// The equivalent length (<c>ℓ</c>) is
        /// <para><c>α = 2/(ℓ + 1), 0 ≤ α ≤ 1, 1 ≤ ℓ</c></para>
        /// </param>
        /// <param name="vFactor">The ν-factor the T3 exponential moving average, <c>0 ≤ ν ≤ 1</c>.</param>
        /// <param name="ohlcvComponent">The Ohlcv component.</param>
        public T3ExponentialMovingAverage(double smoothingFactor, double vFactor = 0.7, OhlcvComponent ohlcvComponent = OhlcvComponent.ClosingPrice)
            : base(T3Ema, T3EmaFull, ohlcvComponent)
        {
            if (0d > smoothingFactor || 1d < smoothingFactor)
                throw new ArgumentOutOfRangeException(nameof(smoothingFactor));
            if (0d > vFactor || 1d < vFactor)
                throw new ArgumentOutOfRangeException(nameof(vFactor));
            SmoothingFactor = smoothingFactor;
            oneMinSmoothingFactor = 1d - smoothingFactor;
            VFactor = vFactor;
            double v2 = vFactor * vFactor;
            c1 = -v2 * vFactor;
            c2 = 3d * (v2 - c1);
            c3 = -6d * v2 - 3d * (vFactor - c1);
            c4 = 1d + 3d * vFactor - c1 + 3d * v2;
            if (Epsilon > smoothingFactor)
                Length = int.MaxValue;
            else
                Length = (int)Math.Round(2d / smoothingFactor) - 1;
            length2 = Length * 2 - 1;
            length3 = Length * 3 - 2;
            length4 = Length * 4 - 3;
            length5 = Length * 5 - 4;
            length6 = Length * 6 - 5;
            Moniker = string.Concat(T3Ema, Length.ToString(CultureInfo.InvariantCulture));
        }
        #endregion

        #region Reset
        /// <inheritdoc />
        public override void Reset()
        {
            lock (updateLock)
            {
                value = double.NaN;
                e1 = 0d; e2 = 0d; e3 = 0d; e4 = 0d; e5 = 0d; e6 = 0d;
                count = 0;
                primed = false;
            }
        }
        #endregion

        #region Update
        /// <summary>
        /// Updates the value of the T3 exponential moving average.
        /// </summary>
        /// <param name="sample">A new sample.</param>
        /// <returns>The new value of the indicator.</returns>
        public double Update(double sample)
        {
            if (double.IsNaN(sample))
                return sample;
            lock (updateLock)
            {
                if (primed)
                {
                    e1 = SmoothingFactor * sample + oneMinSmoothingFactor * e1;
                    e2 = SmoothingFactor * e1 + oneMinSmoothingFactor * e2;
                    e3 = SmoothingFactor * e2 + oneMinSmoothingFactor * e3;
                    e4 = SmoothingFactor * e3 + oneMinSmoothingFactor * e4;
                    e5 = SmoothingFactor * e4 + oneMinSmoothingFactor * e5;
                    e6 = SmoothingFactor * e5 + oneMinSmoothingFactor * e6;
                    value = c1 * e6 + c2 * e5 + c3 * e4 + c4 * e3;
                }
                else
                {
                    if (Length > count)
                    {
                        e1 += sample;
                        if (Length == ++count)
                        {
                            e1 /= Length;
                            e2 = e1;
                        }
                    }
                    else if (length2 > count)
                    {
                        e1 = SmoothingFactor * sample + oneMinSmoothingFactor * e1;
                        e2 += e1;
                        if (length2 == ++count)
                        {
                            e2 /= Length;
                            e3 = e2;
                        }
                    }
                    else if (length3 > count)
                    {
                        e1 = SmoothingFactor * sample + oneMinSmoothingFactor * e1;
                        e2 = SmoothingFactor * e1 + oneMinSmoothingFactor * e2;
                        e3 += e2;
                        if (length3 == ++count)
                        {
                            e3 /= Length;
                            e4 = e3;
                        }
                    }
                    else if (length4 > count)
                    {
                        e1 = SmoothingFactor * sample + oneMinSmoothingFactor * e1;
                        e2 = SmoothingFactor * e1 + oneMinSmoothingFactor * e2;
                        e3 = SmoothingFactor * e2 + oneMinSmoothingFactor * e3;
                        e4 += e3;
                        if (length4 == ++count)
                        {
                            e4 /= Length;
                            e5 = e4;
                        }
                    }
                    else if (length5 > count)
                    {
                        e1 = SmoothingFactor * sample + oneMinSmoothingFactor * e1;
                        e2 = SmoothingFactor * e1 + oneMinSmoothingFactor * e2;
                        e3 = SmoothingFactor * e2 + oneMinSmoothingFactor * e3;
                        e4 = SmoothingFactor * e3 + oneMinSmoothingFactor * e4;
                        e5 += e4;
                        if (length5 == ++count)
                        {
                            e5 /= Length;
                            e6 = e5;
                        }
                    }
                    else
                    {
                        e1 = SmoothingFactor * sample + oneMinSmoothingFactor * e1;
                        e2 = SmoothingFactor * e1 + oneMinSmoothingFactor * e2;
                        e3 = SmoothingFactor * e2 + oneMinSmoothingFactor * e3;
                        e4 = SmoothingFactor * e3 + oneMinSmoothingFactor * e4;
                        e5 = SmoothingFactor * e4 + oneMinSmoothingFactor * e5;
                        e6 += e5;
                        if (length6 == ++count)
                        {
                            e6 /= Length;
                            value = c1 * e6 + c2 * e5 + c3 * e4 + c4 * e3;
                            primed = true;
                        }
                    }
                }
                return value;
            }
        }

        /// <summary>
        /// Updates the value of the T3 exponential moving average.
        /// </summary>
        /// <param name="price">A new price.</param>
        /// <param name="dateTime">A date-time of the new price.</param>
        /// <returns>The new value of the indicator.</returns>
        public Scalar Update(double price, DateTime dateTime)
        {
            return new Scalar(dateTime, Update(price));
        }

        /// <summary>
        /// Updates the value of the T3 exponential moving average.
        /// </summary>
        /// <param name="scalar">A new scalar.</param>
        /// <returns>The new value of the indicator.</returns>
        public Scalar Update(Scalar scalar)
        {
            return new Scalar(scalar.Time, Update(scalar.Value));
        }

        /// <summary>
        /// Updates the value of the T3 exponential moving average.
        /// </summary>
        /// <param name="ohlcv">A new ohlcv.</param>
        /// <returns>The new value of the indicator.</returns>
        public Scalar Update(Ohlcv ohlcv)
        {
            return new Scalar(ohlcv.Time, Update(ohlcv.Component(OhlcvComponent)));
        }
        #endregion
    }
}
