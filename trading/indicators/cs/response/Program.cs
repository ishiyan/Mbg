using System;

namespace fft
{
    static class Program
    {
        static void Main(string[] args)
        {
            double[] inpData = new double[1024];
            inpData.Initialize();
            for (int i = 0; i < 32; ++i) inpData[i] = 1;
            RealFastFourierTransform(inpData, 1024, false);
            //RealFastFourierTransform(inpData, 4, true);

            double[] spectrum = new double[512], spectrum2 = new double[512];
            for (int i = 1; i < spectrum.Length; ++i)
            {
                int k = i + i;
                spectrum[i] = inpData[k] * inpData[k] + inpData[k + 1] * inpData[k + 1];
            }
            int q = 1;
            for (int i = 1; i < spectrum.Length; ++i)
            {
                //int k = i + i;
                double re = inpData[++q], im = inpData[++q];
                spectrum2[i] = re * re + im * im;
            }
            spectrum[0] = 0;
        }

        private static void CalculateAmplitudeSpectrum(double[] array, int arrayLength, double[] powerLinear, double[] powerDb, double[] amplitudeLinear, double[] amplitudeDb)
        {
            int spectrumLength = arrayLength / 2;
            int k = 1;
            double max = double.MinValue;
            for (int i = 1; i < spectrumLength; ++i)
            {
                double re = array[++k];
                double im = array[++k];
                re = re * re + im * im;
                if (max < re)
                    max = re;
                powerLinear[i] = re;
            }
            // Clear constant component level.
            powerLinear[0] = 0;
            amplitudeLinear[0] = 0;
            // Normalize to [0,1] range.
            if (max > double.Epsilon)
            {
                for (int i = 1; i < spectrumLength; ++i)
                {
                    powerLinear[i] /= max;
                    amplitudeLinear[i] = Math.Sqrt(powerLinear[i]);
                }
            }
            else
            {
                for (int i = 1; i < spectrumLength; ++i)
                {
                    powerLinear[i] = 0.0015;
                    amplitudeLinear[i] = 0.0015;
                }
            }
            double minPower = double.MaxValue;
            double minAmplitude = double.MaxValue;
            for (int i = 1; i < spectrumLength; ++i)
            {
                double d = 20 * Math.Log10(powerLinear[i]);
                if (minPower < d)
                    minPower = d;
                powerDb[i] = d;
                d = 20 * Math.Log10(amplitudeLinear[i]);
                if (minAmplitude < d)
                    minAmplitude = d;
                amplitudeDb[i] = d;
            }
            for (int i = 10; i > 0; --i)
            {
                if (minPower >= -i * 10 && minPower < -(i - 1) * 10)
                    minPower = -i * 10;
                if (minAmplitude >= -i * 10 && minAmplitude < -(i - 1) * 10)
                    minAmplitude = -i * 10;
            }

        }


        /// <summary>
        /// Performs a real fast Fourier transform.
        /// </summary>
        /// <param name="array">A data array containing real data on input and {re,im} pairs on return if transform is direct. If transform is inverse, contains {re,im} pairs on input and real data on return.</param>
        /// <param name="arrayLength">A length of the array. Should be a power of 2, the minimal value is 2.</param>
        /// <param name="inverseFft">Specifies whether to perform an inverse transform.</param>
        private static void RealFastFourierTransform(double[] array, int arrayLength, bool inverseFft)
        {
            int n;
            double c1 = 0.5, c2, ttheta = 2 * Math.PI / arrayLength;

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
            int isign = inverseFft ? - 1 : 1;
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
                    j -=m;
                    m /= 2;
                }
                j += m;
            }
            int mMax = 2;
            n = arrayLength;
            while (n > mMax)
            {
                int istep = 2 * mMax;
                double theta = 2 * Math.PI / (isign * mMax);
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
    }
}
