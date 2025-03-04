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
    public class TripleExponentialMovingAverageTest
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
        /// Output data.
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_ma.c.
        ///   /*******************************/
        ///   /*  TEMA TEST - Metastock      */
        ///   /*******************************/
        ///   /* No output value. */
        ///   { 0, TA_ANY_MA_TEST, 0, 1, 1,  14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 0, 0, 0},
        ///#ifndef TA_FUNC_NO_RANGE_CHECK
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251,  0, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_BAD_PARAM, 0, 0, 0, 0 },
        ///#endif
        ///
        ///   /* Test with period 14 */
        ///   { 1, TA_ANY_MA_TEST, 0, 0, 251, 14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   0,  84.721, 39, 252-39 }, /* First Value */
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251, 14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   1,  84.089, 39, 252-39 },
        ///   { 0, TA_ANY_MA_TEST, 0, 0, 251, 14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 252-40, 108.418, 39, 252-39 }, /* Last Value */
        /// </summary>
        private readonly List<double> expected14 = new List<double>
        {
            84.721, // Index=39 value.
            84.089, // Index=40 value.
            108.418 // Index=251 (last) value.
        };
        private readonly List<double> expected14M = new List<double>
        {
            // The very first value is the input value.
            84.721, // Index=39 value.
            84.089, // Index=40 value.
            108.418 // Index=251 (last) value.
        };
        #endregion

        #region NameTest
        /// <summary>
        /// A test for Name.
        /// </summary>
        [TestMethod]
        public void NameTest()
        {
            var target = new TripleExponentialMovingAverage(5);
            Assert.AreEqual("TEMA", target.Name);
        }
        #endregion

        #region MonikerTest
        /// <summary>
        /// A test for Moniker.
        /// </summary>
        [TestMethod]
        public void MonikerTest()
        {
            var target = new TripleExponentialMovingAverage(4);
            Assert.AreEqual("TEMA4", target.Moniker);
        }
        #endregion

        #region DescriptionTest
        /// <summary>
        /// A test for Description.
        /// </summary>
        [TestMethod]
        public void DescriptionTest()
        {
            var target = new TripleExponentialMovingAverage(3);
            Assert.AreEqual("Triple Exponential Moving Average", target.Description);
        }
        #endregion

        #region IsPrimedTest
        /// <summary>
        /// A test for IsPrimed.
        /// </summary>
        [TestMethod]
        public void IsPrimedTest()
        {
            var target = new TripleExponentialMovingAverage(3);
            Assert.IsFalse(target.IsPrimed);
            var scalar = new Scalar(DateTime.Now, 1d);
            for (int i = 1; i < 7; i++)
            {
                scalar.Value = i;
                target.Update(scalar);
                Assert.IsFalse(target.IsPrimed);
            }
            for (int i = 7; i < 15; i++)
            {
                scalar.Value = i;
                target.Update(scalar);
                Assert.IsTrue(target.IsPrimed);
            }

            target = new TripleExponentialMovingAverage(3, false);
            for (int i = 1; i < 7; i++)
            {
                scalar.Value = i;
                target.Update(scalar);
                Assert.IsFalse(target.IsPrimed);
            }
            for (int i = 7; i < 15; i++)
            {
                scalar.Value = i;
                target.Update(scalar);
                Assert.IsTrue(target.IsPrimed);
            }
        }
        #endregion

        #region TaLibTest14
        /// <summary>
        /// A TA-Lib data test 14.
        /// </summary>
        [TestMethod]
        public void TaLibTest14()
        {
            int count = input.Count;
            const int dec = 1;
            var target = new TripleExponentialMovingAverage(14);
            for (int i = 0; i < 39; i++)
            {
                target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
            }
            double d = target.Update(input[39]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            //Assert.AreEqual(Math.Round(expected14[0], dec), Math.Round(d, dec));
            d = target.Update(input[40]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            //Assert.AreEqual(Math.Round(expected14[1], dec), Math.Round(d, dec));
            for (int i = 41; i < count; i++)
                d = target.Update(input[i]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14[2], dec), Math.Round(d, dec));
        }
        #endregion

        #region TaLibTest14M
        /// <summary>
        /// A TA-Lib data test 14M.
        /// </summary>
        [TestMethod]
        public void TaLibTest14M()
        {
            int count = input.Count;
            var target = new TripleExponentialMovingAverage(14, false);
            for (int i = 0; i < 39; i++)
            {
                target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
            }
            double d = target.Update(input[39]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[0], 0), Math.Round(d, 0));
            d = target.Update(input[40]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[1], 0), Math.Round(d, 0));
            for (int i = 41; i < count; i++)
                d = target.Update(input[i]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[2], 1), Math.Round(d, 1));
        }
        #endregion

        #region LengthTest
        /// <summary>
        /// A test for Length.
        /// </summary>
        [TestMethod]
        public void LengthTest()
        {
            var target = new TripleExponentialMovingAverage(9);
            Assert.AreEqual(9, target.Length);
            Assert.AreEqual(2d / (9d + 1d), target.SmoothingFactor);
            target = new TripleExponentialMovingAverage(3);
            Assert.AreEqual(3, target.Length);
            Assert.AreEqual(2d / (3d + 1d), target.SmoothingFactor);
        }
        #endregion

        #region SmoothingFactorTest
        /// <summary>
        /// A test for SmoothingFactor.
        /// </summary>
        [TestMethod]
        public void SmoothingFactorTest()
        {
            var target = new TripleExponentialMovingAverage(0.1);
            Assert.AreEqual(0.1, target.SmoothingFactor);
            Assert.AreEqual(2d / 0.1 - 1d, target.Length);
            target = new TripleExponentialMovingAverage(0.2);
            Assert.AreEqual(0.2, target.SmoothingFactor);
            Assert.AreEqual(2d / 0.2 - 1d, target.Length);
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
            var target = new TripleExponentialMovingAverage(14, false);
            for (int i = 0; i < 39; i++)
            {
                target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
            }
            double d = target.Update(input[39]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[0], 0), Math.Round(d, 0));
            d = target.Update(input[40]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[1], 0), Math.Round(d, 0));
            for (int i = 41; i < count; i++)
                d = target.Update(input[i]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[2], 1), Math.Round(d, 1));
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
            var target = new TripleExponentialMovingAverage(14, false);
            for (int i = 0; i < 39; i++)
            {
                target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
            }
            double d = target.Update(input[39]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[0], 0), Math.Round(d, 0));
            d = target.Update(input[40]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[1], 0), Math.Round(d, 0));
            for (int i = 41; i < count; i++)
                d = target.Update(input[i]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[2], 1), Math.Round(d, 1));

            target.Reset();

            for (int i = 0; i < 39; i++)
            {
                target.Update(input[i]);
                Assert.IsFalse(target.IsPrimed);
            }
            d = target.Update(input[39]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[0], 0), Math.Round(d, 0));
            d = target.Update(input[40]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[1], 0), Math.Round(d, 0));
            for (int i = 41; i < count; i++)
                d = target.Update(input[i]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[2], 1), Math.Round(d, 1));
        }
        #endregion

        #region TripleExponentialMovingAverageConstructorTest
        /// <summary>
        /// A test for Constructor.
        /// </summary>
        [TestMethod]
        public void TripleExponentialMovingAverageConstructorTest()
        {
            var target = new TripleExponentialMovingAverage(5);
            Assert.AreEqual(5, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new TripleExponentialMovingAverage(4, false);
            Assert.AreEqual(4, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new TripleExponentialMovingAverage(3, false, OhlcvComponent.MedianPrice);
            Assert.AreEqual(3, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);

            target = new TripleExponentialMovingAverage(13, ohlcvComponent:OhlcvComponent.MedianPrice);
            Assert.AreEqual(13, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);

            target = new TripleExponentialMovingAverage(0.4);
            Assert.AreEqual(0.4, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new TripleExponentialMovingAverage(0.4, false);
            Assert.AreEqual(0.4, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new TripleExponentialMovingAverage(0.3, false, OhlcvComponent.MedianPrice);
            Assert.AreEqual(0.3, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);

            target = new TripleExponentialMovingAverage(0.6, ohlcvComponent:OhlcvComponent.MedianPrice);
            Assert.AreEqual(0.6, target.SmoothingFactor);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsTrue(target.FirstIsAverage);
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void TripleExponentialMovingAverageConstructorTest2()
        {
            var target = new TripleExponentialMovingAverage(0);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void TripleExponentialMovingAverageConstructorTest3()
        {
            var target = new TripleExponentialMovingAverage(-8);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void TripleExponentialMovingAverageConstructorTest4()
        {
            var target = new TripleExponentialMovingAverage(-0.01);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void TripleExponentialMovingAverageConstructorTest5()
        {
            var target = new TripleExponentialMovingAverage(1.01);
            Assert.IsFalse(target.IsPrimed);
        }
        #endregion

        #region SerializationTest
        private static void SerializeTo(TripleExponentialMovingAverage instance, string fileName)
        {
            var dcs = new DataContractSerializer(typeof(TripleExponentialMovingAverage), null, 65536, false, true, null);
            using (var fs = new FileStream(fileName, FileMode.Create))
            {
                dcs.WriteObject(fs, instance);
                fs.Close();
            }
        }

        private static TripleExponentialMovingAverage DeserializeFrom(string fileName)
        {
            var fs = new FileStream(fileName, FileMode.Open);
            XmlDictionaryReader reader = XmlDictionaryReader.CreateTextReader(fs, new XmlDictionaryReaderQuotas());
            var ser = new DataContractSerializer(typeof(TripleExponentialMovingAverage), null, 65536, false, true, null);
            var instance = (TripleExponentialMovingAverage)ser.ReadObject(reader, true);
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
            int count = input.Count;
            const int dec = 1;
            var source = new TripleExponentialMovingAverage(14, false);
            for (int i = 0; i < 39; i++)
            {
                source.Update(input[i]);
                Assert.IsFalse(source.IsPrimed);
            }
            double d = source.Update(input[39]);
            Assert.IsFalse(double.IsNaN(d));
            Assert.IsTrue(source.IsPrimed);
            Assert.AreEqual(Math.Round(expected14M[0], 0), Math.Round(d, 0));

            const string fileName = "TripleExponentialMovingAverageTest_1.xml";
            SerializeTo(source, fileName);
            TripleExponentialMovingAverage target = DeserializeFrom(fileName);
            Assert.AreEqual(14, target.Length);
            Assert.IsTrue(target.IsPrimed);
            Assert.AreEqual(Math.Round(source.Value, dec), Math.Round(target.Value, dec));
            Assert.AreEqual("TEMA", target.Name);
            Assert.AreEqual("Triple Exponential Moving Average", target.Description);

            d = target.Update(input[40]);
            Assert.AreEqual(Math.Round(expected14M[1], 0), Math.Round(d, 0));
            for (int i = 41; i < count; i++)
                d = target.Update(input[i]);
            Assert.AreEqual(Math.Round(expected14M[2], 1), Math.Round(d, 1));
            //FileInfo fi = new FileInfo(fileName);
            //fi.Delete();
        }
        #endregion
    }
}
