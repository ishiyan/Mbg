using System;
using System.Collections.Generic;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators2;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace Mbs.UnitTests.Trading.Indicators2
{
    [TestClass]
    public class ExponentialMovingAverageTests
    {
        /// <summary>
        /// Input test data, length = 5, unbiased = false (population variance).
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_data.c, TA_SREF_close_daily_ref_0_PRIV[252].
        /// </summary>
        private readonly List<double> input = new List<double>
        {
            91.500000, 94.815000, 94.375000, 95.095000, 93.780000, 94.625000, 92.530000, 92.750000, 90.315000, 92.470000, 96.125000,
            97.250000, 98.500000, 89.875000, 91.000000, 92.815000, 89.155000, 89.345000, 91.625000, 89.875000, 88.375000, 87.625000,
            84.780000, 83.000000, 83.500000, 81.375000, 84.440000, 89.250000, 86.375000, 86.250000, 85.250000, 87.125000, 85.815000,
            88.970000, 88.470000, 86.875000, 86.815000, 84.875000, 84.190000, 83.875000, 83.375000, 85.500000, 89.190000, 89.440000,
            91.095000, 90.750000, 91.440000, 89.000000, 91.000000, 90.500000, 89.030000, 88.815000, 84.280000, 83.500000, 82.690000,
            84.750000, 85.655000, 86.190000, 88.940000, 89.280000, 88.625000, 88.500000, 91.970000, 91.500000, 93.250000, 93.500000,
            93.155000, 91.720000, 90.000000, 89.690000, 88.875000, 85.190000, 83.375000, 84.875000, 85.940000, 97.250000, 99.875000,
            104.940000, 106.000000, 102.500000, 102.405000, 104.595000, 106.125000, 106.000000, 106.065000, 104.625000, 108.625000,
            109.315000, 110.500000, 112.750000, 123.000000, 119.625000, 118.750000, 119.250000, 117.940000, 116.440000, 115.190000,
            111.875000, 110.595000, 118.125000, 116.000000, 116.000000, 112.000000, 113.750000, 112.940000, 116.000000, 120.500000,
            116.620000, 117.000000, 115.250000, 114.310000, 115.500000, 115.870000, 120.690000, 120.190000, 120.750000, 124.750000,
            123.370000, 122.940000, 122.560000, 123.120000, 122.560000, 124.620000, 129.250000, 131.000000, 132.250000, 131.000000,
            132.810000, 134.000000, 137.380000, 137.810000, 137.880000, 137.250000, 136.310000, 136.250000, 134.630000, 128.250000,
            129.000000, 123.870000, 124.810000, 123.000000, 126.250000, 128.380000, 125.370000, 125.690000, 122.250000, 119.370000,
            118.500000, 123.190000, 123.500000, 122.190000, 119.310000, 123.310000, 121.120000, 123.370000, 127.370000, 128.500000,
            123.870000, 122.940000, 121.750000, 124.440000, 122.000000, 122.370000, 122.940000, 124.000000, 123.190000, 124.560000,
            127.250000, 125.870000, 128.860000, 132.000000, 130.750000, 134.750000, 135.000000, 132.380000, 133.310000, 131.940000,
            130.000000, 125.370000, 130.130000, 127.120000, 125.190000, 122.000000, 125.000000, 123.000000, 123.500000, 120.060000,
            121.000000, 117.750000, 119.870000, 122.000000, 119.190000, 116.370000, 113.500000, 114.250000, 110.000000, 105.060000,
            107.000000, 107.870000, 107.000000, 107.120000, 107.000000, 91.000000, 93.940000, 93.870000, 95.500000, 93.000000,
            94.940000, 98.250000, 96.750000, 94.810000, 94.370000, 91.560000, 90.250000, 93.940000, 93.620000, 97.000000, 95.000000,
            95.870000, 94.060000, 94.620000, 93.750000, 98.000000, 103.940000, 107.870000, 106.060000, 104.500000, 105.000000,
            104.190000, 103.060000, 103.420000, 105.270000, 111.870000, 116.000000, 116.620000, 118.280000, 113.370000, 109.000000,
            109.700000, 109.250000, 107.000000, 109.190000, 110.000000, 109.200000, 110.120000, 108.000000, 108.620000, 109.750000,
            109.810000, 109.000000, 108.750000, 107.870000,
        };

        /// <summary>
        /// Output data.
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_ma.c.
        /// <code>
        ///   /*******************************/
        ///   /*   EMA TEST - Classic        */
        ///   /*******************************/
        ///   /* No output value. */
        ///   { 0, TA_ANY_MA_TEST, 0, 1, 1,  14, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 0, 0, 0},
        ///   #ifndef TA_FUNC_NO_RANGE_CHECK
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  0, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_BAD_PARAM, 0, 0, 0, 0 },
        ///   #endif
        ///   /* Misc tests: period 2, 10 */
        ///   { 1, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,  93.15, 1, 251 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   1,  93.96, 1, 251 },
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 250, 108.21, 1, 251 }, /* Last Value */
        ///
        ///   { 1, TA_ANY_MA_TEST, 0, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,    0,  93.22,  9, 243 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,    1,  93.75,  9, 243 },
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   20,  86.46,  9, 243 },
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,  242, 108.97,  9, 243 }, /* Last Value */
        ///   /*******************************/
        ///   /*   EMA TEST - Metastock      */
        ///   /*******************************/
        ///   /* No output value. */
        ///   { 0, TA_ANY_MA_TEST, 0, 1, 1,  14, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 0, 0, 0},
        ///   #ifndef TA_FUNC_NO_RANGE_CHECK
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  0, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_BAD_PARAM, 0, 0, 0, 0 },
        ///   #endif
        ///   /* Test with 1 unstable price bar. Test for period 2, 10 */
        ///   { 1, TA_ANY_MA_TEST, 1, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   0,  94.15, 1+1, 251-1 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 1, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   1,  94.78, 1+1, 251-1 },
        ///   { 0, TA_ANY_MA_TEST, 1, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 250-1, 108.21, 1+1, 251-1 }, /* Last Value */
        ///
        ///   { 1, TA_ANY_MA_TEST, 1, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,    0,  93.24,  9+1, 243-1 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 1, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,    1,  93.97,  9+1, 243-1 },
        ///   { 0, TA_ANY_MA_TEST, 1, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   20,  86.23,  9+1, 243-1 },
        ///   { 0, TA_ANY_MA_TEST, 1, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 242-1, 108.97,  9+1, 243-1 }, /* Last Value */
        ///
        ///   /* Test with 2 unstable price bar. Test for period 2, 10 */
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   0,  94.78, 1+2, 251-2 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   1,  94.11, 1+2, 251-2 },
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  2, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 250-2, 108.21, 1+2, 251-2 }, /* Last Value */
        ///
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,    0,  93.97,  9+2, 243-2 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,    1,  94.79,  9+2, 243-2 },
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   20,  86.39,  9+2, 243-2 },
        ///   { 0, TA_ANY_MA_TEST, 2, 0, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,  242-2, 108.97,  9+2, 243-2 }, /* Last Value */
        ///
        ///   /* Last 3 value with 1 unstable, period 10 */
        ///   { 0, TA_ANY_MA_TEST, 1, 249, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   1, 109.22, 249, 3 },
        ///   { 0, TA_ANY_MA_TEST, 1, 249, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   2, 108.97, 249, 3 },
        ///
        ///   /* Last 3 value with 2 unstable, period 10 */
        ///   { 0, TA_ANY_MA_TEST, 2, 249, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   2, 108.97, 249, 3 },
        ///
        ///   /* Last 3 value with 3 unstable, period 10 */
        ///   { 0, TA_ANY_MA_TEST, 3, 249, 251,  10, TA_MAType_EMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   2, 108.97, 249, 3 }
        /// </code>
        /// </summary>
        private readonly List<double> expected2 = new List<double>
        {
            93.15, // Index=2 value.
            93.96, // Index=3 value.
            108.21, // Index=251 (last) value.
        };

        private readonly List<double> expected10 = new List<double>
        {
            93.22, // Index=9 value.
            93.75, // Index=10 value.
            86.46, // Index=29 value.
            108.97, // Index=251 (last) value.
        };

        private readonly List<double> expected2M = new List<double>
        {
            // The very first value is the input value.
            94.15, // Index=2 value.
            94.78, // Index=3 value.
            108.21, // Index=251 (last) value.
        };

        private readonly List<double> expected10M = new List<double>
        {
            // The very first value is the input value.
            93.24, // Index=10 value.
            93.97, // Index=11 value.
            86.23, // Index=30 value.
            108.97, // Index=251 (last) value.
        };

        [TestMethod]
        public void ExponentialMovingAverage_Name_HasCorrectValue()
        {
            var target = new ExponentialMovingAverage(5);
            Assert.AreEqual("ema", target.Name);
        }

        [TestMethod]
        public void ExponentialMovingAverage_Moniker_HasCorrectValue()
        {
            var target = new ExponentialMovingAverage(4);
            Assert.AreEqual("ema4", target.Moniker);
        }

        [TestMethod]
        public void ExponentialMovingAverage_Description_HasCorrectValue()
        {
            var target = new ExponentialMovingAverage(3);
            Assert.AreEqual("Exponential Moving Average", target.Description);
        }

        [TestMethod]
        public void ExponentialMovingAverage_IsPrimed_HasCorrectValue()
        {
            var target = new ExponentialMovingAverage(5);
            Assert.IsFalse(target.IsPrimed);
            for (int i = 1; i < 5; i++)
            {
                target.Update(new Scalar(DateTime.Now, i));
                Assert.IsFalse(target.IsPrimed);
            }

            for (int i = 5; i < 10; i++)
            {
                target.Update(new Scalar(DateTime.Now, i));
                Assert.IsTrue(target.IsPrimed);
            }

            target = new ExponentialMovingAverage(5, false);
            Assert.IsFalse(target.IsPrimed);
            for (int i = 1; i < 5; i++)
            {
                target.Update(new Scalar(DateTime.Now, i));
                Assert.IsFalse(target.IsPrimed);
            }

            for (int i = 5; i < 10; i++)
            {
                target.Update(new Scalar(DateTime.Now, i));
                Assert.IsTrue(target.IsPrimed);
            }
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest2_Update_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            double d;

            var target = new ExponentialMovingAverage(2);
            for (int i = 0; i < 1; i++)
            {
                d = target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
                Assert.IsTrue(double.IsNaN(d));
            }

            d = target.Update(input[1]);
            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(d, dec));

            d = target.Update(input[2]);
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(d, dec));

            for (int i = 3; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(d, dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest2_Calculate_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            List<double> list = ExponentialMovingAverage.Calculate(input, 2);

            for (int i = 0; i < 1; i++)
            {
                Assert.IsTrue(double.IsNaN(list[i]));
            }

            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(list[1], dec));
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(list[2], dec));
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(list[count - 1], dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest10_Update_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            double d;
            var target = new ExponentialMovingAverage(10);
            for (int i = 0; i < 9; i++)
            {
                d = target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
                Assert.IsTrue(double.IsNaN(d));
            }

            d = target.Update(input[9]);
            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(d, dec));

            d = target.Update(input[10]);
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(d, dec));

            for (int i = 11; i < 30; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(d, dec));

            for (int i = 30; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(d, dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest10_Calculate_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            List<double> list = ExponentialMovingAverage.Calculate(input, 10);

            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(list[i]));
            }

            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(list[9], dec));
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(list[10], dec));
            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(list[29], dec));
            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(list[count - 1], dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest2M_Update_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            double d;
            var target = new ExponentialMovingAverage(2, false);
            for (int i = 0; i < 1; i++)
            {
                d = target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
                Assert.IsTrue(double.IsNaN(d));
            }

            d = target.Update(input[1]);
            Assert.IsFalse(double.IsNaN(d));

            d = target.Update(input[2]);
            Assert.AreEqual(Math.Round(expected2M[0], dec), Math.Round(d, dec));

            d = target.Update(input[3]);
            Assert.AreEqual(Math.Round(expected2M[1], dec), Math.Round(d, dec));

            for (int i = 4; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected2M[2], dec), Math.Round(d, dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest2M_Calculate_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            List<double> list = ExponentialMovingAverage.Calculate(input, 2, false);
            for (int i = 0; i < 1; i++)
            {
                Assert.IsTrue(double.IsNaN(list[i]));
            }

            Assert.IsFalse(double.IsNaN(list[1]));
            Assert.AreEqual(Math.Round(expected2M[0], dec), Math.Round(list[2], dec));
            Assert.AreEqual(Math.Round(expected2M[1], dec), Math.Round(list[3], dec));
            Assert.AreEqual(Math.Round(expected2M[2], dec), Math.Round(list[count - 1], dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest10M_Update_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            double d;
            var target = new ExponentialMovingAverage(10, false);
            for (int i = 0; i < 9; i++)
            {
                d = target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
                Assert.IsTrue(double.IsNaN(d));
            }

            d = target.Update(input[9]);
            Assert.IsFalse(double.IsNaN(d));

            d = target.Update(input[10]);
            Assert.AreEqual(Math.Round(expected10M[0], dec), Math.Round(d, dec));

            d = target.Update(input[11]);
            Assert.AreEqual(Math.Round(expected10M[1], dec), Math.Round(d, dec));

            for (int i = 11; i < 31; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10M[2], dec), Math.Round(d, dec));

            for (int i = 31; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10M[3], dec), Math.Round(d, dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_TaLibTest10M_Calculate_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            List<double> list = ExponentialMovingAverage.Calculate(input, 10, false);
            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(list[i]));
            }

            Assert.IsFalse(double.IsNaN(list[9]));
            Assert.AreEqual(Math.Round(expected10M[0], dec), Math.Round(list[10], dec));
            Assert.AreEqual(Math.Round(expected10M[1], dec), Math.Round(list[11], dec));
            Assert.AreEqual(Math.Round(expected10M[2], dec), Math.Round(list[30], dec));
            Assert.AreEqual(Math.Round(expected10M[3], dec), Math.Round(list[count - 1], dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_Length_HasCorrectValue()
        {
            var target = new ExponentialMovingAverage(9);
            Assert.AreEqual(9, target.Length);
            Assert.AreEqual(2d / (9d + 1d), target.SmoothingFactor);

            target = new ExponentialMovingAverage(3);
            Assert.AreEqual(3, target.Length);
            Assert.AreEqual(2d / (3d + 1d), target.SmoothingFactor);
        }

        [TestMethod]
        public void ExponentialMovingAverage_SmoothingFactor_HasCorrectValue()
        {
            var target = new ExponentialMovingAverage(0.1);
            Assert.AreEqual(0.1, target.SmoothingFactor);
            Assert.AreEqual(2d / 0.1 - 1d, target.Length);

            target = new ExponentialMovingAverage(0.2);
            Assert.AreEqual(0.2, target.SmoothingFactor);
            Assert.AreEqual(2d / 0.2 - 1d, target.Length);
        }

        [TestMethod]
        public void ExponentialMovingAverage_Update_ReturnsCorrectValues()
        {
            int count = input.Count;
            const int dec = 1;
            var target = new ExponentialMovingAverage(10);
            for (int i = 0; i < 9; i++)
            {
                target.Update(input[i]);
            }

            var d = target.Update(input[9]);
            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(d, dec));

            d = target.Update(input[10]);
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(d, dec));

            for (int i = 11; i < 30; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(d, dec));

            for (int i = 30; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(d, dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_Calculate_ReturnsCorrectValues()
        {
            const int dec = 1;
            List<double> actual = ExponentialMovingAverage.Calculate(input, 2);
            for (int i = 0; i < 1; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(actual[1], dec));
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(actual[2], dec));
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 2, false);
            for (int i = 0; i < 1; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected2M[0], dec), Math.Round(actual[2], dec));
            Assert.AreEqual(Math.Round(expected2M[1], dec), Math.Round(actual[3], dec));
            Assert.AreEqual(Math.Round(expected2M[2], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 10);
            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(actual[9], dec));
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(actual[10], dec));
            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(actual[29], dec));
            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 10, false);
            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected10M[0], dec), Math.Round(actual[10], dec));
            Assert.AreEqual(Math.Round(expected10M[1], dec), Math.Round(actual[11], dec));
            Assert.AreEqual(Math.Round(expected10M[2], dec), Math.Round(actual[30], dec));
            Assert.AreEqual(Math.Round(expected10M[3], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 2d / 3d);
            for (int i = 0; i < 1; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(actual[1], dec));
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(actual[2], dec));
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 2d / 3d, false);
            for (int i = 0; i < 1; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected2M[0], dec), Math.Round(actual[2], dec));
            Assert.AreEqual(Math.Round(expected2M[1], dec), Math.Round(actual[3], dec));
            Assert.AreEqual(Math.Round(expected2M[2], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 2d / 11d);
            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(actual[9], dec));
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(actual[10], dec));
            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(actual[29], dec));
            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(actual[251], dec));

            actual = ExponentialMovingAverage.Calculate(input, 2d / 11d, false);
            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            Assert.AreEqual(Math.Round(expected10M[0], dec), Math.Round(actual[10], dec));
            Assert.AreEqual(Math.Round(expected10M[1], dec), Math.Round(actual[11], dec));
            Assert.AreEqual(Math.Round(expected10M[2], dec), Math.Round(actual[30], dec));
            Assert.AreEqual(Math.Round(expected10M[3], dec), Math.Round(actual[251], dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_Reset_FunctionsCorrectly()
        {
            int count = input.Count;
            const int dec = 1;
            double d;
            var target = new ExponentialMovingAverage(10);
            for (int i = 0; i < 9; i++)
            {
                d = target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
                Assert.IsTrue(double.IsNaN(d));
            }

            d = target.Update(input[9]);
            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(d, dec));

            d = target.Update(input[10]);
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(d, dec));

            for (int i = 11; i < 30; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(d, dec));

            for (int i = 30; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(d, dec));

            target.Reset();

            for (int i = 0; i < 9; i++)
            {
                d = target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
                Assert.IsTrue(double.IsNaN(d));
            }

            d = target.Update(input[9]);
            Assert.AreEqual(Math.Round(expected10[0], dec), Math.Round(d, dec));

            d = target.Update(input[10]);
            Assert.AreEqual(Math.Round(expected10[1], dec), Math.Round(d, dec));

            for (int i = 11; i < 30; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[2], dec), Math.Round(d, dec));

            for (int i = 30; i < count; i++)
            {
                d = target.Update(input[i]);
            }

            Assert.AreEqual(Math.Round(expected10[3], dec), Math.Round(d, dec));
        }

        [TestMethod]
        public void ExponentialMovingAverage_Constructor_CreatesAnInstance()
        {
            var target = new ExponentialMovingAverage(5);
            Assert.AreEqual(5, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new ExponentialMovingAverage(4, false);
            Assert.AreEqual(4, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new ExponentialMovingAverage(3, false, OhlcvComponent.MedianPrice);
            Assert.AreEqual(3, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);

            target = new ExponentialMovingAverage(13, ohlcvComponent: OhlcvComponent.MedianPrice);
            Assert.AreEqual(13, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);

            target = new ExponentialMovingAverage(0.4);
            Assert.AreEqual(0.4, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new ExponentialMovingAverage(0.4, false);
            Assert.AreEqual(0.4, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new ExponentialMovingAverage(0.3, false, OhlcvComponent.MedianPrice);
            Assert.AreEqual(0.3, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);

            target = new ExponentialMovingAverage(0.6, ohlcvComponent: OhlcvComponent.MedianPrice);
            Assert.AreEqual(0.6, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void ExponentialMovingAverage_Constructor_LengthIsNegative_ThrowsException()
        {
            var target = new ExponentialMovingAverage(-8);
            Assert.IsFalse(target.IsPrimed);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void ExponentialMovingAverage_Constructor_LengthIsZero_ThrowsException()
        {
            var target = new ExponentialMovingAverage(0);
            Assert.IsFalse(target.IsPrimed);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void ExponentialMovingAverage_Constructor_SmoothingFactorIsNegative_ThrowsException()
        {
            var target = new ExponentialMovingAverage(-0.01);
            Assert.IsFalse(target.IsPrimed);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void ExponentialMovingAverage_Constructor_SmoothingFactorIsGreaterThanOne_ThrowsException()
        {
            var target = new ExponentialMovingAverage(1.01);
            Assert.IsFalse(target.IsPrimed);
        }
    }
}
