﻿using System;
using System.Collections.Generic;
using System.IO;
using System.Runtime.Serialization;
using System.Xml;
using Microsoft.VisualStudio.TestTools.UnitTesting;

using Mbst.Trading;
using Mbst.Trading.Indicators;

namespace Tests.Indicators
{
    [TestClass]
    public class WeightedMovingAverageTest
    {
        #region Test data
        /// <summary>
        /// Input test data, length = 5, unbiased = false (popiulation variance).
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_data.c, TA_SREF_close_daily_ref_0_PRIV[252].
        /// </summary>
        private readonly List<double> input = new List<double>
        {
            91.500000,94.815000,94.375000,95.095000,93.780000,94.625000,92.530000,92.750000,90.315000,92.470000,96.125000,
            97.250000,98.500000,89.875000,91.000000,92.815000,89.155000,89.345000,91.625000,89.875000,88.375000,87.625000,
            84.780000,83.000000,83.500000,81.375000,84.440000,89.250000,86.375000,86.250000,85.250000,87.125000,85.815000,
            88.970000,88.470000,86.875000,86.815000,84.875000,84.190000,83.875000,83.375000,85.500000,89.190000,89.440000,
            91.095000,90.750000,91.440000,89.000000,91.000000,90.500000,89.030000,88.815000,84.280000,83.500000,82.690000,
            84.750000,85.655000,86.190000,88.940000,89.280000,88.625000,88.500000,91.970000,91.500000,93.250000,93.500000,
            93.155000,91.720000,90.000000,89.690000,88.875000,85.190000,83.375000,84.875000,85.940000,97.250000,99.875000,
            104.940000,106.000000,102.500000,102.405000,104.595000,106.125000,106.000000,106.065000,104.625000,108.625000,
            109.315000,110.500000,112.750000,123.000000,119.625000,118.750000,119.250000,117.940000,116.440000,115.190000,
            111.875000,110.595000,118.125000,116.000000,116.000000,112.000000,113.750000,112.940000,116.000000,120.500000,
            116.620000,117.000000,115.250000,114.310000,115.500000,115.870000,120.690000,120.190000,120.750000,124.750000,
            123.370000,122.940000,122.560000,123.120000,122.560000,124.620000,129.250000,131.000000,132.250000,131.000000,
            132.810000,134.000000,137.380000,137.810000,137.880000,137.250000,136.310000,136.250000,134.630000,128.250000,
            129.000000,123.870000,124.810000,123.000000,126.250000,128.380000,125.370000,125.690000,122.250000,119.370000,
            118.500000,123.190000,123.500000,122.190000,119.310000,123.310000,121.120000,123.370000,127.370000,128.500000,
            123.870000,122.940000,121.750000,124.440000,122.000000,122.370000,122.940000,124.000000,123.190000,124.560000,
            127.250000,125.870000,128.860000,132.000000,130.750000,134.750000,135.000000,132.380000,133.310000,131.940000,
            130.000000,125.370000,130.130000,127.120000,125.190000,122.000000,125.000000,123.000000,123.500000,120.060000,
            121.000000,117.750000,119.870000,122.000000,119.190000,116.370000,113.500000,114.250000,110.000000,105.060000,
            107.000000,107.870000,107.000000,107.120000,107.000000,91.000000,93.940000,93.870000,95.500000,93.000000,
            94.940000,98.250000,96.750000,94.810000,94.370000,91.560000,90.250000,93.940000,93.620000,97.000000,95.000000,
            95.870000,94.060000,94.620000,93.750000,98.000000,103.940000,107.870000,106.060000,104.500000,105.000000,
            104.190000,103.060000,103.420000,105.270000,111.870000,116.000000,116.620000,118.280000,113.370000,109.000000,
            109.700000,109.250000,107.000000,109.190000,110.000000,109.200000,110.120000,108.000000,108.620000,109.750000,
            109.810000,109.000000,108.750000,107.870000
        };

        /// <summary>
        /// Output data, length = 2.
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_ma.c.
        /// /*******************************/
        /// /*   WMA TEST  - CLASSIC       */
        /// /*******************************/
        /// #ifndef TA_FUNC_NO_RANGE_CHECK
        /// /* No output value. */
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  0, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_BAD_PARAM, 0, 0, 0, 0 },
        /// #endif
        /// /* One value tests. */
        /// { 0, TA_ANY_MA_TEST, 0, 2,   2,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,  94.52,   2, 1 },
        /// /* Misc tests: period 2, 30 */
        /// { 1, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,   93.71,  1,  252-1  }, /* First Value */
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   1,   94.52,  1,  252-1  },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   2,   94.85,  1,  252-1  },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 250,  108.16,  1,  252-1  }, /* Last Value */
        ///
        /// { 1, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,  88.567,  29,  252-29 }, /* First Value */
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   1,  88.233,  29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   2,  88.034,  29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,  29,  87.191,  29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 221, 109.3413, 29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 222, 109.3466, 29,  252-29 }, /* Last Value */
        /// </summary>
        private readonly List<double> expected2 = new List<double>
        {
            // The first value is double.NaN.
            93.71,  // Index=1 value.
            94.52,  // Index=2 value.
            94.855, // Index=3 value.
            108.16  // Index=251 (last) value.
        };

        /// <summary>
        /// Output data, length = 30.
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_ma.c.
        /// /*******************************/
        /// /*   WMA TEST  - CLASSIC       */
        /// /*******************************/
        /// #ifndef TA_FUNC_NO_RANGE_CHECK
        /// /* No output value. */
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  0, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_BAD_PARAM, 0, 0, 0, 0 },
        /// #endif
        /// /* One value tests. */
        /// { 0, TA_ANY_MA_TEST, 0, 2,   2,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,  94.52,   2, 1 },
        /// /* Misc tests: period 2, 30 */
        /// { 1, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,   93.71,  1,  252-1  }, /* First Value */
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   1,   94.52,  1,  252-1  },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   2,   94.85,  1,  252-1  },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251,  2, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 250,  108.16,  1,  252-1  }, /* Last Value */
        ///
        /// { 1, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   0,  88.567,  29,  252-29 }, /* First Value */
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   1,  88.233,  29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,   2,  88.034,  29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,  29,  87.191,  29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 221, 109.3413, 29,  252-29 },
        /// { 0, TA_ANY_MA_TEST, 0, 0, 251, 30, TA_MAType_WMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 222, 109.3466, 29,  252-29 }, /* Last Value */
        /// </summary>
        private readonly List<double> expected30 = new List<double>
        {
            // The first 29 values are double.NaN.
            88.5677,  // Index=29 value.
            88.2337,  // Index=30 value.
            88.034,   // Index=31 value.
            87.191,   // Index=58 value.
            109.3466, // Index=250 value.
            109.3413  // Index=251 (last) value.
        };
        #endregion

        #region NameTest
        /// <summary>
        /// A test for Name.
        /// </summary>
        [TestMethod]
        public void NameTest()
        {
            var target = new WeightedMovingAverage(5);
            Assert.AreEqual("WMA", target.Name);
        }
        #endregion

        #region MonikerTest
        /// <summary>
        /// A test for Moniker.
        /// </summary>
        [TestMethod]
        public void MonikerTest()
        {
            var target = new WeightedMovingAverage(4);
            Assert.AreEqual("WMA4", target.Moniker);
        }
        #endregion

        #region DescriptionTest
        /// <summary>
        /// A test for Description.
        /// </summary>
        [TestMethod]
        public void DescriptionTest()
        {
            var target = new WeightedMovingAverage(3);
            Assert.AreEqual("Weighted Moving Average", target.Description);
        }
        #endregion

        #region IsPrimedTest
        /// <summary>
        /// A test for IsPrimed.
        /// </summary>
        [TestMethod]
        public void IsPrimedTest()
        {
            var target = new WeightedMovingAverage(5);
            Assert.IsFalse(target.IsPrimed);
            var scalar = new Scalar(DateTime.Now, 1d);
            for (int i = 1; i < 5; i++)
            {
                scalar.Value = i;
                target.Update(scalar);
                Assert.IsFalse(target.IsPrimed);
            }
            for (int i = 5; i < 10; i++)
            {
                scalar.Value = i;
                target.Update(scalar);
                Assert.IsTrue(target.IsPrimed);
            }
        }
        #endregion

        #region TaLib2Test
        /// <summary>
        /// A TA-Lib data test, length = 2.
        /// </summary>
        [TestMethod]
        public void TaLib2Test()
        {
            int count = input.Count;
            const int dec = 2;
            double d;
            var target = new WeightedMovingAverage(2);
            for (int i = 0; i < 1; i++)
            {
                d = target.Update(input[i]);
                Assert.IsTrue(double.IsNaN(d));
            }
            d = target.Update(input[1]);
            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(d, dec));
            d = target.Update(input[2]);
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(d, dec));
            d = target.Update(input[3]);
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(d, dec));
            for (int i = 4; i < count; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected2[3], dec), Math.Round(d, dec));
        }
        #endregion

        #region TaLib30Test
        /// <summary>
        /// A TA-Lib data test, length = 30.
        /// </summary>
        [TestMethod]
        public void TaLib30Test()
        {
            const int dec = 3;
            double d;
            var target = new WeightedMovingAverage(30);
            for (int i = 0; i < 29; i++)
            {
                d = target.Update(input[i]);
                Assert.IsTrue(double.IsNaN(d));
            }
            d = target.Update(input[29]);
            Assert.AreEqual(Math.Round(expected30[0], dec), Math.Round(d, dec));
            d = target.Update(input[30]);
            Assert.AreEqual(Math.Round(expected30[1], dec), Math.Round(d, dec));
            d = target.Update(input[31]);
            Assert.AreEqual(Math.Round(expected30[2], dec), Math.Round(d, dec));
            for (int i = 32; i < 59; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected30[3], dec), Math.Round(d, dec));
            for (int i = 59; i < 251; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected30[4], dec), Math.Round(d, dec));
            d = target.Update(input[251]);
            Assert.AreEqual(Math.Round(expected30[5], dec), Math.Round(d, dec));
        }
        #endregion

        #region LengthTest
        /// <summary>
        /// A test for Length.
        /// </summary>
        [TestMethod]
        public void LengthTest()
        {
            var target = new WeightedMovingAverage(11);
            Assert.AreEqual(11, target.Length);
            target = new WeightedMovingAverage(22);
            Assert.AreEqual(22, target.Length);
        }
        #endregion

        #region UpdateTest
        /// <summary>
        /// A test for Update.
        /// </summary>
        [TestMethod]
        public void UpdateTest()
        {
            int count = input.Count;
            const int dec = 2;
            double d;
            var scalar = new Scalar(DateTime.Now, 1d);
            var target = new WeightedMovingAverage(2);
            for (int i = 0; i < 1; i++)
            {
                scalar.Value = input[i];
                d = target.Update(scalar).Value;
                Assert.IsTrue(double.IsNaN(d));
            }
            scalar.Value = input[1];
            d = target.Update(scalar).Value;
            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(d, dec));
            scalar.Value = input[2];
            d = target.Update(scalar).Value;
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(d, dec));
            scalar.Value = input[3];
            d = target.Update(scalar).Value;
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(d, dec));
            for (int i = 4; i < count; i++)
            {
                scalar.Value = input[i];
                d = target.Update(scalar).Value;
            }
            Assert.AreEqual(Math.Round(expected2[3], dec), Math.Round(d, dec));
        }
        #endregion

        #region CalculateTest
        /// <summary>
        /// A test for Calculate.
        /// </summary>
        [TestMethod]
        public void CalculateTest()
        {
            const int dec = 3;
            List<double> actual = WeightedMovingAverage.Calculate(input, 30);
            for (int i = 0; i < 29; i++)
                Assert.IsTrue(double.IsNaN(actual[i]));
            Assert.AreEqual(Math.Round(expected30[0], dec), Math.Round(actual[29], dec));
            Assert.AreEqual(Math.Round(expected30[1], dec), Math.Round(actual[30], dec));
            Assert.AreEqual(Math.Round(expected30[2], dec), Math.Round(actual[31], dec));
            Assert.AreEqual(Math.Round(expected30[3], dec), Math.Round(actual[58], dec));
            Assert.AreEqual(Math.Round(expected30[4], dec), Math.Round(actual[250], dec));
            Assert.AreEqual(Math.Round(expected30[5], dec), Math.Round(actual[251], dec));
        }
        #endregion

        #region ResetTest
        /// <summary>
        /// A test for Reset.
        /// </summary>
        [TestMethod]
        public void ResetTest()
        {
            int count = input.Count;
            const int dec = 2;
            double d;
            var target = new WeightedMovingAverage(2);
            for (int i = 0; i < 1; i++)
            {
                d = target.Update(input[i]);
                Assert.IsTrue(double.IsNaN(d));
            }
            d = target.Update(input[1]);
            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(d, dec));
            d = target.Update(input[2]);
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(d, dec));
            d = target.Update(input[3]);
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(d, dec));
            for (int i = 4; i < count; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected2[3], dec), Math.Round(d, dec));
            target.Reset();
            for (int i = 0; i < 1; i++)
            {
                d = target.Update(input[i]);
                Assert.IsTrue(double.IsNaN(d));
            }
            d = target.Update(input[1]);
            Assert.AreEqual(Math.Round(expected2[0], dec), Math.Round(d, dec));
            d = target.Update(input[2]);
            Assert.AreEqual(Math.Round(expected2[1], dec), Math.Round(d, dec));
            d = target.Update(input[3]);
            Assert.AreEqual(Math.Round(expected2[2], dec), Math.Round(d, dec));
            for (int i = 4; i < count; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected2[3], dec), Math.Round(d, dec));
        }
        #endregion

        #region WeightedMovingAverageConstructorTest
        /// <summary>
        /// A test for constructor.
        /// </summary>
        [TestMethod]
        public void WeightedMovingAverageConstructorTest()
        {
            var target = new WeightedMovingAverage(5);
            Assert.AreEqual(5, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for a constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void WeightedMovingAverageConstructorTest2()
        {
            var target = new WeightedMovingAverage(1);
            Assert.IsNotNull(target);
        }

        /// <summary>
        /// A test for a constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void WeightedMovingAverageConstructorTest3()
        {
            var target = new WeightedMovingAverage(0);
            Assert.IsNotNull(target);
        }

        /// <summary>
        /// A test for a constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void WeightedMovingAverageConstructorTest4()
        {
            var target = new WeightedMovingAverage(-8);
            Assert.IsNotNull(target);
        }
        #endregion

        #region SerializationTest
        private static void SerializeTo(WeightedMovingAverage instance, string fileName)
        {
            var dcs = new DataContractSerializer(typeof(WeightedMovingAverage), null, 65536, false, true, null);
            using (var fs = new FileStream(fileName, FileMode.Create))
            {
                dcs.WriteObject(fs, instance);
                fs.Close();
            }
        }

        private static WeightedMovingAverage DeserializeFrom(string fileName)
        {
            var fs = new FileStream(fileName, FileMode.Open);
            XmlDictionaryReader reader = XmlDictionaryReader.CreateTextReader(fs, new XmlDictionaryReaderQuotas());
            var ser = new DataContractSerializer(typeof(WeightedMovingAverage), null, 65536, false, true, null);
            var instance = (WeightedMovingAverage)ser.ReadObject(reader, true);
            reader.Close();
            fs.Close();
            return instance;
        }

        /// <summary>
        /// A test for the serialization.
        /// </summary>
        [TestMethod]
        public void SerializationTest()
        {
            const int dec = 3;
            double d;
            var source = new WeightedMovingAverage(30);
            for (int i = 0; i < 29; i++)
            {
                d = source.Update(input[i]);
                Assert.IsTrue(double.IsNaN(d));
            }
            d = source.Update(input[29]);
            Assert.AreEqual(Math.Round(expected30[0], dec), Math.Round(d, dec));
            d = source.Update(input[30]);
            Assert.AreEqual(Math.Round(expected30[1], dec), Math.Round(d, dec));
            d = source.Update(input[31]);
            Assert.AreEqual(Math.Round(expected30[2], dec), Math.Round(d, dec));
            const string fileName = "WeightedMovingAverageTest_1.xml";
            SerializeTo(source, fileName);
            WeightedMovingAverage target = DeserializeFrom(fileName);
            Assert.AreEqual(30, target.Length);
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected30[2], dec), Math.Round(target.Value, dec));
            Assert.AreEqual("WMA", target.Name);
            Assert.AreEqual("Weighted Moving Average", target.Description);
            for (int i = 32; i < 59; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected30[3], dec), Math.Round(d, dec));
            for (int i = 59; i < 251; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected30[4], dec), Math.Round(d, dec));
            d = target.Update(input[251]);
            Assert.AreEqual(Math.Round(expected30[5], dec), Math.Round(d, dec));
            //FileInfo fi = new FileInfo(fileName);
            //fi.Delete();
        }
        #endregion
    }
}
