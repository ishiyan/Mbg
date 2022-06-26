using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Media;
using Mbst;
using Mbst.Charts;
using Mbst.Controls.ControlStyles;
using Mbst.Numerics;
using Mbst.Trading.Indicators;
using Ehlers = Mbst.Trading.Indicators.JohnEhlers;

namespace Numerics
{
    internal sealed partial class FilterResponseWindow
    {
        protected override void OnSourceInitialized(EventArgs e)
        {
            base.OnSourceInitialized(e);
            // Apply aero effect to entire window.
            if (Properties.Settings.Default.Theme == ControlStyles.Glass)
                this.ExtendGlassFrame();
        }

        public FilterResponseWindow()
        {
            InitializeComponent();
            Filters = FilterArray();
            currentFilter = Filters[0];
            foreach (var f in Filters)
                predefinedFilterComboBox.Items.Add(f.Moniker);
        }

        private void WindowLoaded(object sender, RoutedEventArgs e)
        {
            var pen = new Pen(Brushes.Red, 0.5) { DashStyle = DashStyles.Dot };
            pen.Freeze();
            chart.AxisPen = pen;
            FilterParameterChanged(null, null);
            MethodChanged();
            if (autoplotCheckBox.IsChecked.HasValue && autoplotCheckBox.IsChecked.Value)
                AddFilter(null, null);
            chart.InvalidateVisual();
        }

        private enum Method { MagnitudeDb = 0, PowerDb, MagnitudeLinear, PowerLinear, Phase }

        private Method currentMethod = Method.Phase;

        private sealed class Filter
        {
            public string Moniker;
            public string Latex;
            public PredefinedFunctionStyle Style;
            public Grid Grid;
            public Action<Filter> UpdateMoniker, CreateSpectrum;
            public LineIndicatorSpectrum LineIndicatorSpectrum;
        }

        private readonly Filter[] Filters;
        private Filter currentFilter;

        private Filter[] FilterArray()
        {
            return new[]
            {
                new Filter
                {
                    Moniker = "sma(ℓ)",
                    Style = PredefinedFunctionStyle.LawnGreenLine,
                    Grid = smaGrid,
                    Latex = @"y = \frac{x + x[1] + \ldots + x[\ell-1]}{\ell}",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "sma(ℓ={0})",
                            smaLengthSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new SimpleMovingAverage(smaLengthSpinner.Value);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "wma(ℓ)",
                    Style = PredefinedFunctionStyle.ForestGreenLine,
                    Grid = wmaGrid,
                    Latex = @"y = \frac{\ell\cdot x + (\ell-1)\cdot x[1] + \ldots + x[\ell-1]}{\ell + (\ell-1) +\ldots + 1}",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "wma(ℓ={0})",
                            wmaLengthSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new WeightedMovingAverage(wmaLengthSpinner.Value);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "trima(ℓ)",
                    Style = PredefinedFunctionStyle.PaleGreenLine,
                    Grid = trimaGrid,
                    //Latex = @"trima(N, x) = \left\{ \begin{array}{ll} sma(N/2+1, sma(N/2, x)) &N=even\\ sma((N+1)/2, sma((N+1)/2, x)) &N=odd\end{array} \right",
                    //Latex = @"trima_{N}(x) = \left\{ \begin{array}{ll} sma_{\frac{N}{2}+1}(sma_{\frac{N}{2}}(x))&N=even\\ sma_{\frac{N+1}{2}}(sma_{\frac{N+1}{2}}(x))&N=odd\end{array} \right",
                    Latex = @"trima_{\ell}(x) = \left\{ \begin{array}{ll} sma_{\ell/2+1}(sma_{\ell/2}(x))&\ell=even\\ sma_{(\ell+1)/2}(sma_{(\ell+1)/2}(x))&\ell=odd\end{array} \right",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "trima(ℓ={0})",
                            trimaLengthSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new TriangularMovingAverage(trimaLengthSpinner.Value);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "ema(α)",
                    Style = PredefinedFunctionStyle.DarkSeaGreenLine,
                    Grid = emaGrid,
                    Latex = @"y = \alpha x + (1 - \alpha)y[1]",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "ema(α={0:#.##})",
                            emaAlphaNumericSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new ExponentialMovingAverage(emaAlphaNumericSpinner.Value, false);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "dema(α)",
                    Style = PredefinedFunctionStyle.SeaGreenLine,
                    Grid = demaGrid,
                    Latex = @"y = 2\cdot ema_{\alpha}(x) - ema_{\alpha}(ema_{\alpha}(x)))",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "dema(α={0:#.##})",
                            demaAlphaNumericSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new DoubleExponentialMovingAverage(demaAlphaNumericSpinner.Value, false);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "tema(α)",
                    Style = PredefinedFunctionStyle.LightSeaGreenLine,
                    Grid = temaGrid,
                    Latex = @"y = 3\cdot (ema_{\alpha}(x) - ema_{\alpha}(ema_{\alpha}(x))) + ema_{\alpha}(ema_{\alpha}(ema_{\alpha}(x)))",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "tema(α={0:#.##})",
                            temaAlphaNumericSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new TripleExponentialMovingAverage(temaAlphaNumericSpinner.Value, false);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "zema(α,gf,ml)",
                    Style = PredefinedFunctionStyle.MediumSeaGreenLine,
                    Grid = zemaGrid,
                    Latex = @"y = \alpha (x + g_{f}(x - x[m_{l}])) + (1 - \alpha)y[1]",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "zema(α={0:0.####},gf={1:0.####},ml={2})",
                            zemaAlphaNumericSpinner.Value, zemaGainFactorNumericSpinner.Value, zemaMomentumSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new Ehlers.ZeroLagExponentialMovingAverage(zemaAlphaNumericSpinner.Value, zemaGainFactorNumericSpinner.Value, zemaMomentumSpinner.Value);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "zecema(α,gl,gs)",
                    Style = PredefinedFunctionStyle.MediumSeaGreenLine,
                    Grid = zecemaGrid,
                    Latex = @"y = \alpha (x + g(x - x[m_{l}])) + (1 - \alpha)y[1]",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "zecema(α={0:0.####},gl={1:0.####},gs={2})",
                            zecemaAlphaNumericSpinner.Value, zecemaGainLimitNumericSpinner.Value, zecemaGainStepNumericSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new Ehlers.ZeroLagErrorCorectingExponentialMovingAverage(zecemaAlphaNumericSpinner.Value, zecemaGainLimitNumericSpinner.Value, zecemaGainStepNumericSpinner.Value);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
                new Filter
                {
                    Moniker = "mom(ℓ)",
                    Style = PredefinedFunctionStyle.AquamarineLine,
                    Grid = momGrid,
                    Latex = "y = x - x[ℓ]",
                    UpdateMoniker = filter =>
                    {
                        filter.Moniker = string.Format(CultureInfo.InvariantCulture, "mom(ℓ={0})",
                            momLengthSpinner.Value);
                    },
                    CreateSpectrum = filter =>
                    {
                        var indicator = new Momentum(momLengthSpinner.Value);
                        filter.LineIndicatorSpectrum = new LineIndicatorSpectrum(indicator);
                    }
                },
            };
        }

        private void MethodChanged()
        {
            string yMin, yMax;
            StringBuilder stringBuilder = new StringBuilder("Filter response: ");
            switch (methodComboBox.SelectedIndex)
            {
                case 0:
                    currentMethod = Method.MagnitudeDb;
                    stringBuilder.Append("magnitude (dB)");
                    yMin = "-60";
                    yMax = "5";
                    break;
                case 1:
                    currentMethod = Method.PowerDb;
                    stringBuilder.Append("power (dB)");
                    yMin = "-80";
                    yMax = "5";
                    break;
                case 2:
                    currentMethod = Method.MagnitudeLinear;
                    stringBuilder.Append("magnitude (linear)");
                    yMin = "0";
                    yMax = "1";
                    break;
                case 3:
                    currentMethod = Method.PowerLinear;
                    stringBuilder.Append("power (linear)");
                    yMin = "0";
                    yMax = "1";
                    break;
                case 4:
                    currentMethod = Method.Phase;
                    stringBuilder.Append("phase (deg)");
                    yMin = "-180";
                    yMax = "180";
                    break;
                default:
                    throw new ArgumentException("Unknown methodComboBox.SelectedIndex");
            }
            if (null != yMinTextBox)
                yMinTextBox.Text = yMin;
            if (null != yMaxTextBox)
                yMaxTextBox.Text = yMax;
            if (null != xMinTextBox)
                xMinTextBox.Text = "0";
            if (null != xMaxTextBox)
                xMaxTextBox.Text = "1";
            stringBuilder.Append(" vs normalized frequency");
            Title = stringBuilder.ToString();
            UpdateRange(null, null);
        }

        private void MethodChanged(object sender, SelectionChangedEventArgs e)
        {
            if (methodComboBox.SelectedIndex != (int)currentMethod && IsInitialized)
            {
                MethodChanged();
                DeleteAllFilters();
                AddFilter(null, null);
            }
        }

        private void StyleSelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            if (e.AddedItems.Count > 0)
                chart.ApplyStyle(e.AddedItems[0].ToString());
        }

        private void PredefinedFilterChanged(object sender, SelectionChangedEventArgs e)
        {
            Filter f = Filters[predefinedFilterComboBox.SelectedIndex];
            styleComboBox.SelectedIndex = (int)f.Style;
            BeginInit();
            latexBlock.Text = f.Latex;
            currentFilter.Grid.Visibility = Visibility.Collapsed;
            currentFilter = f;
            f.Grid.Visibility = Visibility.Visible;
            FilterParameterChanged(null, null);
            EndInit();
            if (IsLoaded && autoplotCheckBox.IsChecked.HasValue && autoplotCheckBox.IsChecked.Value)
                AddFilter(null, null);
        }

        private void AddFilter(object sender, RoutedEventArgs e)
        {
            BeginInit();
            if (autoclearCheckBox.IsChecked.HasValue && autoclearCheckBox.IsChecked.Value)
                DeleteAllFilters();
            currentFilter.CreateSpectrum(currentFilter);
            Func<double, double> func = currentFilter.LineIndicatorSpectrum.Function(currentMethod);
            string moniker = monikerTextBox.Text;
            if (string.IsNullOrEmpty(moniker))
            {
                moniker = currentFilter.Moniker;
                monikerTextBox.Text = moniker;
            }
            if (Properties.Settings.Default.ExportFilterResponseData)
                currentFilter.LineIndicatorSpectrum.ExportToFile(moniker);
            IFunctionChartSource source = chart.AddSource(func, moniker, styleComboBox.SelectedIndex);
            source.DomainMinimum = 0;
            source.DomainMaximum = 1;
            source.IsDomainMinimumInclusive = true;
            source.IsDomainMaximumInclusive = true;
            functionsComboBox.Items.Add(moniker);
            functionsComboBox.SelectedItem = moniker;
            EndInit();
        }

        private void FilterParameterChanged(object sender, RoutedPropertyChangedEventArgs<int> e)
        {
            if (null == currentFilter)
                return;
            currentFilter.UpdateMoniker(currentFilter);
            monikerTextBox.Text = currentFilter.Moniker;
        }

        private void EmaAlphaChanged(object sender, RoutedPropertyChangedEventArgs<double> e)
        {
            if (null == currentFilter)
                return;
            double length = 2 / emaAlphaNumericSpinner.Value - 1;
            emaLengthTextBlock.Text = string.Format(CultureInfo.InvariantCulture, "{0:#.##}", length);
            FilterParameterChanged(null, null);
        }

        private void DemaAlphaChanged(object sender, RoutedPropertyChangedEventArgs<double> e)
        {
            if (null == currentFilter)
                return;
            double length = 2 / demaAlphaNumericSpinner.Value - 1;
            demaLengthTextBlock.Text = string.Format(CultureInfo.InvariantCulture, "{0:#.##}", length);
            FilterParameterChanged(null, null);
        }

        private void TemaAlphaChanged(object sender, RoutedPropertyChangedEventArgs<double> e)
        {
            if (null == currentFilter)
                return;
            double length = 2 / temaAlphaNumericSpinner.Value - 1;
            temaLengthTextBlock.Text = string.Format(CultureInfo.InvariantCulture, "{0:#.##}", length);
            FilterParameterChanged(null, null);
        }

        private void ZemaAlphaChanged(object sender, RoutedPropertyChangedEventArgs<double> e)
        {
            if (null == currentFilter)
                return;
            double length = 2 / zemaAlphaNumericSpinner.Value - 1;
            zemaLengthTextBlock.Text = string.Format(CultureInfo.InvariantCulture, "{0:#.##}", length);
            FilterParameterChanged(null, null);
        }

        private void ZecemaAlphaChanged(object sender, RoutedPropertyChangedEventArgs<double> e)
        {
            if (null == currentFilter)
                return;
            double length = 2 / zecemaAlphaNumericSpinner.Value - 1;
            zecemaLengthTextBlock.Text = string.Format(CultureInfo.InvariantCulture, "{0:#.##}", length);
            FilterParameterChanged(null, null);
        }

        private void UpdateRange(object sender, RoutedEventArgs e)
        {
            try
            {
                chart.AbscissaMinimum = double.Parse(xMinTextBox.Text, NumberStyles.Any, CultureInfo.InvariantCulture);
                chart.AbscissaMaximum = double.Parse(xMaxTextBox.Text, NumberStyles.Any, CultureInfo.InvariantCulture);
                chart.OrdinateMinimum = double.Parse(yMinTextBox.Text, NumberStyles.Any, CultureInfo.InvariantCulture);
                chart.OrdinateMaximum = double.Parse(yMaxTextBox.Text, NumberStyles.Any, CultureInfo.InvariantCulture);
            }
            catch (Exception)
            {
                // Ignored.
            }
        }

        private void DeleteAllFilters()
        {
            List<IFunctionChartSource> sources = chart.Sources.ToList();
            int count = sources.Count;
            for (int i = 0; i < count; i++)
                chart.RemoveSource(sources[i]);
            functionsComboBox.Items.Clear();
        }

        private void DeleteFilter(object sender, RoutedEventArgs e)
        {
            if (functionsComboBox.SelectedIndex < 0)
                return;
            var name = (string)functionsComboBox.SelectedItem;
            IFunctionChartSource source = chart.Sources.FirstOrDefault(s => s.Name == name);
            if (null != source)
                chart.RemoveSource(source);
            functionsComboBox.Items.RemoveAt(functionsComboBox.SelectedIndex);
            if (!functionsComboBox.Items.IsEmpty)
                functionsComboBox.SelectedIndex = 0;
        }

        private class LineIndicatorSpectrum
        {
            private const int signalLength = 4096;
            private const int spectrumLength = signalLength / 2 - 1;
            private static readonly double[] frequency = new double[spectrumLength];
            private readonly ILineIndicator lineIndicator;
            private readonly int warmUp;
            private readonly double[] signal = new double[signalLength];
            private readonly double[] powerDb = new double[spectrumLength];
            private readonly double[] powerLinear = new double[spectrumLength];
            private readonly double[] amplitudeDb = new double[spectrumLength];
            private readonly double[] amplitudeLinear = new double[spectrumLength];
            private readonly double[] phase = new double[spectrumLength];

            private static void PopulateStringBuilder(StringBuilder stringBuilder, string name, double[] array)
            {
                stringBuilder.Append("// ");
                stringBuilder.Append(name);
                stringBuilder.Append(" min = ");
                stringBuilder.Append(array.Min().ToString(CultureInfo.InvariantCulture));
                stringBuilder.Append(", max = ");
                stringBuilder.Append(array.Max().ToString(CultureInfo.InvariantCulture));
                stringBuilder.Append(", length = ");
                stringBuilder.AppendLine(array.Length.ToString(CultureInfo.InvariantCulture));
                stringBuilder.Append("var ");
                stringBuilder.Append(name);
                stringBuilder.Append(" = [");
                stringBuilder.Append(array[0].ToString(CultureInfo.InvariantCulture));
                for (int i = 1; i < spectrumLength; ++i)
                {
                    stringBuilder.Append(", ");
                    stringBuilder.Append(array[i].ToString(CultureInfo.InvariantCulture));
                }
                stringBuilder.AppendLine("];");
            }

            public void ExportToFile(string moniker)
            {
                var sb = new StringBuilder();
                sb.AppendLine("// Filter response ------------------------------------------------");
                sb.AppendLine("var moniker = \"" + moniker + "\";");
                PopulateStringBuilder(sb, "normFreq", frequency);
                PopulateStringBuilder(sb, "powerDb", powerDb);
                PopulateStringBuilder(sb, "powerLinear", powerLinear);
                PopulateStringBuilder(sb, "magnitudeDb", amplitudeDb);
                PopulateStringBuilder(sb, "magnitudeLinear", amplitudeLinear);
                PopulateStringBuilder(sb, "phaseDeg", phase);
                File.AppendAllText("filter_response_export.txt", sb.ToString(), Encoding.UTF8);
            }

            static LineIndicatorSpectrum()
            {
#if false
                DirectFftTest();
                FftTest();
#endif
                for (int i = 0; i < spectrumLength; ++i)
                    frequency[i] = (1.0 + i) / spectrumLength;
            }

            public LineIndicatorSpectrum(ILineIndicator lineIndicator, int warmUp = 0)
            {
                this.lineIndicator = lineIndicator;
                this.warmUp = warmUp;
                Calculate();
            }

            private static double Interpolate(double freq, double[] spectrum)
            {
                if (freq < frequency[0] || freq > frequency[spectrumLength - 1])
                    return double.NaN;
                int index = Array.BinarySearch(frequency, freq);
                if (index < 0)
                {
                    index = ~index;
                    if (index == spectrumLength)
                        return spectrum[spectrumLength - 1];
                    if (index == 0)
                            return spectrum[0];
                    double fraction = frequency[index];
                    fraction = (fraction - freq) / (fraction - frequency[index - 1]);
                    return spectrum[index - 1] * fraction + spectrum[index] * (1 - fraction);
                }
                return spectrum[index];
            }

            private double PowerDb(double freq)
            {
                return Interpolate(freq, powerDb);
            }

            private double PowerLinear(double freq)
            {
                return Interpolate(freq, powerLinear);
            }

            private double AmplitudeLinear(double freq)
            {
                return Interpolate(freq, amplitudeLinear);
            }

            private double AmplitudeDb(double freq)
            {
                return Interpolate(freq, amplitudeDb);
            }

            private double Phase(double freq)
            {
                return Interpolate(freq, phase);
            }

            public Func<double, double> Function(Method method)
            {
                switch (method)
                {
                    case Method.MagnitudeDb:
                        return AmplitudeDb;
                    case Method.PowerDb:
                        return PowerDb;
                    case Method.MagnitudeLinear:
                        return AmplitudeLinear;
                    case Method.PowerLinear:
                        return PowerLinear;
                    case Method.Phase:
                        return Phase;
                    default:
                        throw new ArgumentException(nameof(method));
                }
            }

            private void Calculate()
            {
                while (!lineIndicator.IsPrimed)
                    lineIndicator.Update(0);
                for (int i = 0; i < warmUp; ++i)
                    lineIndicator.Update(0);
                signal[0] = lineIndicator.Update(1000);
                for (int i = 1; i < signalLength; ++i)
                    signal[i] = lineIndicator.Update(0);

                DirectRealFastFourierTransform(signal, signalLength);

                int k = 1;
                double powerLinearMax = double.MinValue;
                double amplitudeLinearMax = double.MinValue;
                for (int i = 0; i < spectrumLength; ++i)
                {
                    double re = signal[++k];
                    double im = signal[++k];
                    phase[i] = -Math.Atan2(im, re) * Constants.Rad2Deg;
                    re = re * re + im * im;
                    if (powerLinearMax < re)
                        powerLinearMax = re;
                    powerLinear[i] = re;
                    re = Math.Sqrt(re);
                    if (amplitudeLinearMax < re)
                        amplitudeLinearMax = re;
                    amplitudeLinear[i] = re;
                }

                NormalizeLinear(spectrumLength, powerLinear, powerLinearMax);
                NormalizeLinear(spectrumLength, amplitudeLinear, amplitudeLinearMax);

                NormalizedLinearToDb(spectrumLength, powerLinear, powerDb);
                NormalizedLinearToDb(spectrumLength, amplitudeLinear, amplitudeDb);
            }

            /// <summary>
            /// Normalize to [0,1] range.
            /// </summary>
            private static void NormalizeLinear(int length, double[] linear, double max)
            {
                if (max > double.Epsilon)
                {
                    for (int i = 0; i < length; ++i)
                        linear[i] /= max;
                }
            }

            private static void NormalizedLinearToDb(int length, double[] normalizedLinear, double[] db)
            {
                double minDb = double.MaxValue;
                for (int i = 0; i < length; ++i)
                {
                    double d = 20 * Math.Log10(normalizedLinear[i]);
                    if (minDb < d)
                        minDb = d;
                    db[i] = d;
                }
                for (int i = 10; i > 0; --i)
                {
                    if (minDb >= -i * 10 && minDb < -(i - 1) * 10)
                        minDb = -i * 10;
                }
                if (minDb < -100)
                {
                    for (int i = 0; i < length; ++i)
                        if (db[i] < -100)
                            db[i] = -100;
                    //minDb = -100;
                }
            }

            /// <summary>
            /// Performs a direct real fast Fourier transform.
            /// </summary>
            /// <param name="array">A data array containing real data on input and {re,im} pairs on return.</param>
            /// <param name="arrayLength">A length of the array. Should be a power of 2, the minimal value is 2.</param>
            private static void DirectRealFastFourierTransform(double[] array, int arrayLength)
            {
                const double c1 = 0.5;
                const double c2 = -0.5;
                double ttheta = Constants.TwoPi / arrayLength;
                int nn = arrayLength / 2;
                int j = 1;
                for (int ii = 1; ii <= nn; ++ii)
                {
                    int i = 2 * ii - 1;
                    if (j > i)
                    {
                        double tempR = array[j - 1];
                        double tempI = array[j];
                        array[j - 1] = array[i - 1];
                        array[j] = array[i];
                        array[i - 1] = tempR;
                        array[i] = tempI;
                    }
                    int m = nn;
                    while (m >= 2 && j > m)
                    {
                        j -= m;
                        m /= 2;
                    }
                    j += m;
                }
                int mMax = 2;
                var n = arrayLength;
                while (n > mMax)
                {
                    int istep = 2 * mMax;
                    double theta = Constants.TwoPi / mMax;
                    double wpR = Math.Sin(0.5 * theta);
                    wpR = -2 * wpR * wpR;
                    var wpI = Math.Sin(theta);
                    var wR = 1.0;
                    var wI = 0.0;
                    for (int ii = 1; ii <= mMax / 2; ++ii)
                    {
                        int m = 2 * ii - 1;
                        for (int jj = 0; jj <= (n - m) / istep; ++jj)
                        {
                            int i = m + jj * istep;
                            j = i + mMax;
                            double tempR = wR * array[j - 1] - wI * array[j];
                            double tempI = wR * array[j] + wI * array[j - 1];
                            array[j - 1] = array[i - 1] - tempR;
                            array[j] = array[i] - tempI;
                            array[i - 1] = array[i - 1] + tempR;
                            array[i] = array[i] + tempI;
                        }
                        var wtemp = wR;
                        wR = wR * wpR - wI * wpI + wR;
                        wI = wI * wpR + wtemp * wpI + wI;
                    }
                    mMax = istep;
                }
                double twpR = Math.Sin(0.5 * ttheta);
                twpR = -2 * twpR * twpR;
                double twpI = Math.Sin(ttheta);
                double twR = 1 + twpR;
                double twI = twpI;
                n = arrayLength / 4 + 1;
                for (int i = 2; i <= n; ++i)
                {
                    int i1 = i + i - 2;
                    int i2 = i1 + 1;
                    int i3 = arrayLength + 1 - i2;
                    int i4 = i3 + 1;
                    double wRs = twR;
                    double wIs = twI;
                    double h1R = c1 * (array[i1] + array[i3]);
                    double h1I = c1 * (array[i2] - array[i4]);
                    double h2R = -c2 * (array[i2] + array[i4]);
                    double h2I = c2 * (array[i1] - array[i3]);
                    array[i1] = h1R + wRs * h2R - wIs * h2I;
                    array[i2] = h1I + wRs * h2I + wIs * h2R;
                    array[i3] = h1R - wRs * h2R + wIs * h2I;
                    array[i4] = -h1I + wRs * h2I + wIs * h2R;
                    double twTemp = twR;
                    twR = twR * twpR - twI * twpI + twR;
                    twI = twI * twpR + twTemp * twpI + twI;
                }
                twR = array[0];
                array[0] = twR + array[1];
                array[1] = twR - array[1];
            }

#if false
            private static void DirectFftTest()
            {
                double[] test = {1d, 1d, 1d, 1d};
                DirectRealFastFourierTransform(test, 4);
                bool passed =
                    Math.Abs(test[0] - 4) < double.Epsilon &&
                    Math.Abs(test[1]) < double.Epsilon &&
                    Math.Abs(test[2]) < double.Epsilon &&
                    Math.Abs(test[3]) < double.Epsilon;
                if (!passed)
                    throw new ApplicationException("DirectFftTest: expecting {1,1,1,1} -> {4,0,0,0}");
            }

            private static void FftTest()
            {
                double[] test = { 1d, 1d, 1d, 1d };
                RealFastFourierTransform(test, 4);
                bool passed =
                    Math.Abs(test[0] - 4) < double.Epsilon &&
                    Math.Abs(test[1]) < double.Epsilon &&
                    Math.Abs(test[2]) < double.Epsilon &&
                    Math.Abs(test[3]) < double.Epsilon;
                if (!passed)
                    throw new ApplicationException("FftTest direct: expecting {1,1,1,1} -> {4,0,0,0}");
                RealFastFourierTransform(test, 4, true);
                passed =
                    Math.Abs(test[0] - 1) < double.Epsilon &&
                    Math.Abs(test[1] - 1) < double.Epsilon &&
                    Math.Abs(test[2] - 1) < double.Epsilon &&
                    Math.Abs(test[3] - 1) < double.Epsilon;
                if (!passed)
                    throw new ApplicationException("FftTest inverse: expecting {4,0,0,0} -> {1,1,1,1}");
            }

            /// <summary>
            /// Performs a real fast Fourier transform.
            /// </summary>
            /// <param name="array">A data array containing real data on input and {re,im} pairs on return if transform is direct. If transform is inverse, contains {re,im} pairs on input and real data on return.</param>
            /// <param name="arrayLength">A length of the array. Should be a power of 2, the minimal value is 2.</param>
            /// <param name="inverseFft">Specifies whether to perform an inverse transform.</param>
            private static void RealFastFourierTransform(double[] array, int arrayLength, bool inverseFft = false)
            {
                int n;
                const double c1 = 0.5;
                double c2, ttheta = Constants.TwoPi / arrayLength;

                if (inverseFft)
                {
                    c2 = 0.5;
                    ttheta = -ttheta;
                    double twpR = Math.Sin(0.5 * ttheta);
                    twpR = -2 * twpR * twpR;
                    double twpI = Math.Sin(ttheta);
                    double twR = 1.0 + twpR;
                    double twI = twpI;
                    n = arrayLength / 4 + 1;
                    for (int i = 2; i <= n; ++i)
                    {
                        int i1 = i + i - 2;
                        int i2 = i1 + 1;
                        int i3 = arrayLength + 1 - i2;
                        int i4 = i3 + 1;
                        double wRs = twR;
                        double wIs = twI;
                        double h1R = c1 * (array[i1] + array[i3]);
                        double h1I = c1 * (array[i2] - array[i4]);
                        double h2R = -c2 * (array[i2] + array[i4]);
                        double h2I = c2 * (array[i1] - array[i3]);
                        array[i1] = h1R + wRs * h2R - wIs * h2I;
                        array[i2] = h1I + wRs * h2I + wIs * h2R;
                        array[i3] = h1R - wRs * h2R + wIs * h2I;
                        array[i4] = -h1I + wRs * h2I + wIs * h2R;
                        double twTemp = twR;
                        twR = twR * twpR - twI * twpI + twR;
                        twI = twI * twpR + twTemp * twpI + twI;
                    }
                    double tempR = array[0];
                    array[0] = c1 * (tempR + array[1]);
                    array[1] = c1 * (tempR - array[1]);
                }
                else
                    c2 = -0.5;
                int isign = inverseFft ? -1 : 1;
                int nn = arrayLength / 2;
                int j = 1;
                for (int ii = 1; ii <= nn; ++ii)
                {
                    int i = 2 * ii - 1;
                    if (j > i)
                    {
                        double tempR = array[j - 1];
                        double tempI = array[j];
                        array[j - 1] = array[i - 1];
                        array[j] = array[i];
                        array[i - 1] = tempR;
                        array[i] = tempI;
                    }
                    int m = nn;
                    while (m >= 2 && j > m)
                    {
                        j -= m;
                        m /= 2;
                    }
                    j += m;
                }
                int mMax = 2;
                n = arrayLength;
                while (n > mMax)
                {
                    int istep = 2 * mMax;
                    double theta = Constants.TwoPi / (isign * mMax);
                    double wpR = Math.Sin(0.5 * theta);
                    wpR = -2 * wpR * wpR;
                    var wpI = Math.Sin(theta);
                    var wR = 1.0;
                    var wI = 0.0;
                    for (int ii = 1; ii <= mMax / 2; ++ii)
                    {
                        int m = 2 * ii - 1;
                        for (int jj = 0; jj <= (n - m) / istep; ++jj)
                        {
                            int i = m + jj * istep;
                            j = i + mMax;
                            double tempR = wR * array[j - 1] - wI * array[j];
                            double tempI = wR * array[j] + wI * array[j - 1];
                            array[j - 1] = array[i - 1] - tempR;
                            array[j] = array[i] - tempI;
                            array[i - 1] = array[i - 1] + tempR;
                            array[i] = array[i] + tempI;
                        }
                        var wtemp = wR;
                        wR = wR * wpR - wI * wpI + wR;
                        wI = wI * wpR + wtemp * wpI + wI;
                    }
                    mMax = istep;
                }
                if (inverseFft)
                {
                    n = 2 * nn;
                    for (int i = 0; i < n; ++i)
                        array[i] /= nn;
                }
                else
                {
                    double twpR = Math.Sin(0.5 * ttheta);
                    twpR = -2 * twpR * twpR;
                    double twpI = Math.Sin(ttheta);
                    double twR = 1 + twpR;
                    double twI = twpI;
                    n = arrayLength / 4 + 1;
                    for (int i = 2; i <= n; ++i)
                    {
                        int i1 = i + i - 2;
                        int i2 = i1 + 1;
                        int i3 = arrayLength + 1 - i2;
                        int i4 = i3 + 1;
                        double wRs = twR;
                        double wIs = twI;
                        double h1R = c1 * (array[i1] + array[i3]);
                        double h1I = c1 * (array[i2] - array[i4]);
                        double h2R = -c2 * (array[i2] + array[i4]);
                        double h2I = c2 * (array[i1] - array[i3]);
                        array[i1] = h1R + wRs * h2R - wIs * h2I;
                        array[i2] = h1I + wRs * h2I + wIs * h2R;
                        array[i3] = h1R - wRs * h2R + wIs * h2I;
                        array[i4] = -h1I + wRs * h2I + wIs * h2R;
                        double twTemp = twR;
                        twR = twR * twpR - twI * twpI + twR;
                        twI = twI * twpR + twTemp * twpI + twI;
                    }
                    double tempR = array[0];
                    array[0] = tempR + array[1];
                    array[1] = tempR - array[1];
                }
            }
#endif
        }
    }
}
