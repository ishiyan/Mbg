using System;
using System.Collections.Generic;
using System.IO;
using System.Runtime.Serialization;
using System.Xml;
using Mbst.Trading;
using Microsoft.VisualStudio.TestTools.UnitTesting;

using Mbst.Trading.Indicators;

namespace Tests.Indicators
{
    [TestClass]
    public class KaufmanAdaptiveMovingAverageTest
    {
        #region Test data
        /// <summary>
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_KAMA.xsl, Close, C5…C256, 252 entries.
        /// </summary>
        private readonly List<double> rawInput = new List<double>
        {
            91.500, 94.815, 94.375, 95.095, 93.780, 94.625, 92.530, 92.750, 90.315, 92.470,
            96.125, 97.250, 98.500, 89.875, 91.000, 92.815, 89.155, 89.345, 91.625, 89.875,
            88.375, 87.625, 84.780, 83.000, 83.500, 81.375, 84.440, 89.250, 86.375, 86.250,
            85.250, 87.125, 85.815, 88.970, 88.470, 86.875, 86.815, 84.875, 84.190, 83.875,
            83.375, 85.500, 89.190, 89.440, 91.095, 90.750, 91.440, 89.000, 91.000, 90.500,
            89.030, 88.815, 84.280, 83.500, 82.690, 84.750, 85.655, 86.190, 88.940, 89.280,
            88.625, 88.500, 91.970, 91.500, 93.250, 93.500, 93.155, 91.720, 90.000, 89.690,
            88.875, 85.190, 83.375, 84.875, 85.940, 97.250, 99.875,104.940,106.000,102.500,
           102.405,104.595,106.125,106.000,106.065,104.625,108.625,109.315,110.500,112.750,
           123.000,119.625,118.750,119.250,117.940,116.440,115.190,111.875,110.595,118.125,
           116.000,116.000,112.000,113.750,112.940,116.000,120.500,116.620,117.000,115.250,
           114.310,115.500,115.870,120.690,120.190,120.750,124.750,123.370,122.940,122.560,
           123.120,122.560,124.620,129.250,131.000,132.250,131.000,132.810,134.000,137.380,
           137.810,137.880,137.250,136.310,136.250,134.630,128.250,129.000,123.870,124.810,
           123.000,126.250,128.380,125.370,125.690,122.250,119.370,118.500,123.190,123.500,
           122.190,119.310,123.310,121.120,123.370,127.370,128.500,123.870,122.940,121.750,
           124.440,122.000,122.370,122.940,124.000,123.190,124.560,127.250,125.870,128.860,
           132.000,130.750,134.750,135.000,132.380,133.310,131.940,130.000,125.370,130.130,
           127.120,125.190,122.000,125.000,123.000,123.500,120.060,121.000,117.750,119.870,
           122.000,119.190,116.370,113.500,114.250,110.000,105.060,107.000,107.870,107.000,
           107.120,107.000, 91.000, 93.940, 93.870, 95.500, 93.000, 94.940, 98.250, 96.750,
            94.810, 94.370, 91.560, 90.250, 93.940, 93.620, 97.000, 95.000, 95.870, 94.060,
            94.620, 93.750, 98.000,103.940,107.870,106.060,104.500,105.000,104.190,103.060,
           103.420,105.270,111.870,116.000,116.620,118.280,113.370,109.000,109.700,109.250,
           107.000,109.190,110.000,109.200,110.120,108.000,108.620,109.750,109.810,109.000,
           108.750, 107.870
        };

        /// <summary>
        /// Taken from TA-Lib (http://ta-lib.org/) tests, test_KAMA.xsl, KAMA, J5…J256, 252 entries.
        /// Efficiency ratio length is 10.
        /// </summary>
        private readonly List<double> kama = new List<double>
        {
                  double.NaN,       double.NaN,       double.NaN,       double.NaN,       double.NaN,       double.NaN,       double.NaN,       double.NaN,       double.NaN,       double.NaN,
            92.6574744421924, 92.7783471257434, 93.0592520064115, 92.9356368995325, 92.9000149644911, 92.8990048732289, 92.8229942018608, 92.7516051928620, 92.7414384525517, 92.6960363223993,
            92.3934372123882, 91.9139380062599, 90.7658162726830, 90.0740111936089, 89.3620815288014, 87.6656280861040, 87.4895131032692, 87.4974604839614, 87.4487997113532, 87.4134797590652,
            87.3586513546248, 87.3571985565411, 87.3428271277309, 87.4342339727455, 87.4790967331831, 87.4478089486627, 87.4341052772180, 87.2779545841798, 87.1866387951289, 87.0799098978843,
            86.9861110535034, 86.9549433796085, 87.0479997922396, 87.0668566957271, 87.2090146571776, 87.4600776240503, 87.8014795040326, 87.8826076877600, 88.2803844203263, 88.5454141018648,
            88.5859031486005, 88.5965040436874, 88.2719621445720, 87.8163354339468, 86.8611444903465, 86.6741610056912, 86.5906930013157, 86.5766752991618, 86.6296450514704, 86.6650208354184,
            86.6783504731998, 86.6895963952268, 87.6981988794437, 88.5095835057360, 89.9508715587081, 90.9585930437125, 91.4794679492180, 91.5092409530174, 91.4856744284233, 91.4717808315536,
            91.4557387469302, 91.1940009725015, 89.4266294004067, 88.8455374050859, 88.3697094609281, 88.5930899916723, 89.1316678888979, 90.8601116442358, 93.2091460910382, 94.0581656977510,
            94.9201636069605, 96.8889752566530, 99.4062425239817,101.1201449462390,102.3769237660390,102.6006738368170,103.3003850710980,103.6578508957870,104.0764855627630,106.4159093020280,
           112.1346727325330,113.5057358502340,114.2548283428500,115.0085673230990,115.3491682211620,115.4744042357010,115.4586954188130,115.4033778968360,115.3819703222920,115.4596680866820,
           115.4927139908920,115.5083211482970,115.3016588863670,115.2382416224770,115.1532481002890,115.1580191296150,115.3257950434630,115.3602912952500,115.4272550190370,115.4236654978450,
           115.4094918992810,115.4100431369950,115.4265778341240,115.7744740794160,116.0930627623780,116.3101967717570,116.6603109196670,117.3487018143020,117.8153888221880,118.4531290804430,
           119.3499419409230,119.8086689971510,120.6175024210070,122.0458817467430,123.9704416533650,125.8138480326600,126.3738969105690,127.6872486354350,129.2393432164220,131.6880947713340,
           133.5239638088170,135.0004207395880,135.6288233403940,135.7374059656390,135.8007904215550,135.7583248045180,135.5543718432480,135.2569852680960,133.6204824276490,131.3192797761920,
           128.7932379609940,128.4062405870340,128.4039316032540,128.0791656483760,127.8414201748350,127.1988985844810,126.5381546649790,125.6607070438540,125.6440698902700,125.6229493897650,
           125.5972771029140,125.1856884028260,125.1156207098550,124.9914050152240,124.9677440635400,125.0508437113440,125.3554407671800,125.3059272985400,125.2940386783170,125.2530757692210,
           125.2419747210570,125.1887237516160,125.1656598262800,125.1342643444030,125.1261708430550,125.0293527295390,125.0082100078360,125.1058124672220,125.1321388339230,125.5284397017590,
           126.2554117345480,126.9803557764160,128.5646940398630,129.8559054638140,130.0995104273400,130.5156892070650,130.6273781337970,130.6136632314180,130.5821372483140,130.5780360175850,
           130.4619826221790,130.2592097652620,129.0901503140520,128.7592330158310,128.3218396854650,127.9194919253990,127.1326782278630,126.7107330400510,126.1909025410680,125.5077119513560,
           125.3652360592940,125.0689417277010,124.6785367307510,123.1715118076970,122.3246069304410,120.4996045001390,118.0226226271800,116.5389084881180,115.7700047414230,114.4762055991300,
           112.8691910705370,111.7330463494810,105.8813879559000,103.7386265802100,101.7705073498860,100.9556429673090,100.0740835866110, 99.5051792798608, 99.4197548401710, 99.2260466472373,
            98.8377738185378, 98.4351675572326, 98.3887252314702, 98.0891751313173, 98.0708172638065, 98.0047820815841, 97.9717872707032, 97.9587393847739, 97.9160266616328, 97.8272391679346,
            97.8109932013579, 97.7811643727499, 97.7968786191168, 98.8421055702164,100.3972096134300,101.1278312905150,101.3486183367770,101.7632588756100,101.9699249107700,102.0803180404650,
           102.2131955779830,102.6495717799380,104.1660350536590,105.9174582846280,107.1295132390960,109.3610815395210,109.7246822740860,109.7071337912410,109.7068748325140,109.6867591775540,
           109.6319778699710,109.6221417907160,109.6271816752350,109.5930223785590,109.6314010730650,109.3937985883840,109.3445353771140,109.3487688924230,109.3510517081720,109.3489501843720,
           109.3310159853090,109.2940150671190
        };
        #endregion

        #region MonikerTest
        /// <summary>
        /// A test for Moniker.
        /// </summary>
        [TestMethod]
        public void MonikerTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(10);
            Assert.AreEqual("KAMA10", target.Moniker);
            target = new KaufmanAdaptiveMovingAverage(9);
            Assert.AreEqual("KAMA9", target.Moniker);
            target = new KaufmanAdaptiveMovingAverage(8);
            Assert.AreEqual("KAMA8", target.Moniker);
        }
        #endregion

        #region EfficiencyRatioLengthTest
        /// <summary>
        /// A test for EfficiencyRatioLength.
        /// </summary>
        [TestMethod]
        public void EfficiencyRatioLengthTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(4);
            Assert.AreEqual(4, target.EfficiencyRatioLength);
            target = new KaufmanAdaptiveMovingAverage(5);
            Assert.AreEqual(5, target.EfficiencyRatioLength);
            target = new KaufmanAdaptiveMovingAverage(6);
            Assert.AreEqual(6, target.EfficiencyRatioLength);
        }
        #endregion

        #region FastestLengthTest
        /// <summary>
        /// A test for FastestLength.
        /// </summary>
        [TestMethod]
        public void FastestLengthTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(20, 9, 15);
            Assert.AreEqual(9, target.FastestLength);
            Assert.AreEqual(2d / (9d + 1d), target.FastestSmoothingFactor);
            target = new KaufmanAdaptiveMovingAverage(20, 3, 15);
            Assert.AreEqual(3, target.FastestLength);
            Assert.AreEqual(2d / (3d + 1d), target.FastestSmoothingFactor);
        }
        #endregion

        #region FastestSmoothingFactorTest
        /// <summary>
        /// A test for FastestSmoothingFactor.
        /// </summary>
        [TestMethod]
        public void FastestSmoothingFactorTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(20, 0.1, 0.01);
            Assert.AreEqual(0.1, target.FastestSmoothingFactor);
            Assert.AreEqual(2d / 0.1 - 1d, target.FastestLength);
            target = new KaufmanAdaptiveMovingAverage(20, 0.2, 0.01);
            Assert.AreEqual(0.2, target.FastestSmoothingFactor);
            Assert.AreEqual(2d / 0.2 - 1d, target.FastestLength);
        }
        #endregion

        #region SlowestLengthTest
        /// <summary>
        /// A test for SlowestLength.
        /// </summary>
        [TestMethod]
        public void SlowestLengthTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(20, 7, 9);
            Assert.AreEqual(9, target.SlowestLength);
            Assert.AreEqual(2d / (9d + 1d), target.SlowestSmoothingFactor);
            target = new KaufmanAdaptiveMovingAverage(20, 2, 3);
            Assert.AreEqual(3, target.SlowestLength);
            Assert.AreEqual(2d / (3d + 1d), target.SlowestSmoothingFactor);
        }
        #endregion

        #region SlowestSmoothingFactorTest
        /// <summary>
        /// A test for SlowestSmoothingFactor.
        /// </summary>
        [TestMethod]
        public void SlowestSmoothingFactorTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(20, 0.9, 0.1);
            Assert.AreEqual(0.1, target.SlowestSmoothingFactor);
            Assert.AreEqual(2d / 0.1 - 1d, target.SlowestLength);
            target = new KaufmanAdaptiveMovingAverage(20, 0.9, 0.2);
            Assert.AreEqual(0.2, target.SlowestSmoothingFactor);
            Assert.AreEqual(2d / 0.2 - 1d, target.SlowestLength);
        }
        #endregion

        #region UpdateTest
        /// <summary>
        /// A test for Update.
        /// </summary>
        [TestMethod]
        public void UpdateTest()
        {
            const int digits = 9;
            var target = new KaufmanAdaptiveMovingAverage(10);
            for (int i = 0; i < rawInput.Count; ++i)
            {
                target.Update(rawInput[i]);
                double d = Math.Round(target.Value, digits);
                double u = Math.Round(kama[i], digits);
                Assert.AreEqual(u, d);
            }
        }
        #endregion

        #region CalculateTest
        /// <summary>
        /// A test for Calculate.
        /// </summary>
        [TestMethod]
        public void CalculateTest()
        {
            var priceList = new List<double>(rawInput);
            List<double> calculatedList = KaufmanAdaptiveMovingAverage.Calculate(priceList, 10);
            var target = new KaufmanAdaptiveMovingAverage(10);
            var updatedList = new List<double>();
            // ReSharper disable once LoopCanBeConvertedToQuery
            foreach (double t in rawInput)
                updatedList.Add(target.Update(t));
            for (int i = 0; i < rawInput.Count; ++i)
                Assert.AreEqual(calculatedList[i], updatedList[i]);
        }
        #endregion

        #region ResetTest
        /// <summary>
        /// A test for Reset.
        /// </summary>
        [TestMethod]
        public void ResetTest()
        {
            double d, u; const int digits = 9;
            var target = new KaufmanAdaptiveMovingAverage(10);
            for (int i = 0; i < rawInput.Count; ++i)
            {
                target.Update(rawInput[i]);
                d = Math.Round(target.Value, digits);
                u = Math.Round(kama[i], digits);
                if (Math.Abs(u - d) > double.Epsilon)
                    System.Diagnostics.Debug.WriteLine("{0}: F {1} => {2} -> {3} | {4} <- {5}", i, rawInput[i], target.Value, d, u, kama[i]);
                Assert.AreEqual(u, d);
            }
            target.Reset();
            for (int i = 0; i < rawInput.Count; ++i)
            {
                target.Update(rawInput[i]);
                d = Math.Round(target.Value, digits);
                u = Math.Round(kama[i], digits);
                if (Math.Abs(u - d) > double.Epsilon)
                    System.Diagnostics.Debug.WriteLine("{0}: F {1} => {2} -> {3} | {4} <- {5}", i, rawInput[i], target.Value, d, u, kama[i]);
                Assert.AreEqual(u, d);
            }
        }
        #endregion

        #region KaufmanAdaptiveMovingAverageConstructorTest
        /// <summary>
        /// A test for Constructor.
        /// </summary>
        [TestMethod]
        public void KaufmanAdaptiveMovingAverageConstructorTest()
        {
            var target = new KaufmanAdaptiveMovingAverage(5);
            Assert.AreEqual(5, target.EfficiencyRatioLength);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.AreEqual(2, target.FastestLength);
            Assert.AreEqual(30, target.SlowestLength);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new KaufmanAdaptiveMovingAverage(4, OhlcvComponent.WeightedPrice);
            Assert.AreEqual(4, target.EfficiencyRatioLength);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.AreEqual(2, target.FastestLength);
            Assert.AreEqual(30, target.SlowestLength);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.WeightedPrice);

            target = new KaufmanAdaptiveMovingAverage(3, 5, 50);
            Assert.AreEqual(3, target.EfficiencyRatioLength);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.AreEqual(5, target.FastestLength);
            Assert.AreEqual(50, target.SlowestLength);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new KaufmanAdaptiveMovingAverage(2, 4, 40, OhlcvComponent.WeightedPrice);
            Assert.AreEqual(2, target.EfficiencyRatioLength);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.AreEqual(4, target.FastestLength);
            Assert.AreEqual(40, target.SlowestLength);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.WeightedPrice);

            target = new KaufmanAdaptiveMovingAverage(13, 0.4, 0.1);
            Assert.AreEqual(13, target.EfficiencyRatioLength);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.AreEqual(0.4, target.FastestSmoothingFactor);
            Assert.AreEqual(0.1, target.SlowestSmoothingFactor);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new KaufmanAdaptiveMovingAverage(14, 0.5, 0.2, OhlcvComponent.TypicalPrice);
            Assert.AreEqual(14, target.EfficiencyRatioLength);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.AreEqual(0.5, target.FastestSmoothingFactor);
            Assert.AreEqual(0.2, target.SlowestSmoothingFactor);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.TypicalPrice);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest2()
        {
            var target = new KaufmanAdaptiveMovingAverage(1);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest3()
        {
            var target = new KaufmanAdaptiveMovingAverage(0);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest4()
        {
            var target = new KaufmanAdaptiveMovingAverage(-7);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest5()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 1, 8);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest6()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 0, 7);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest7()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, -3, 6);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest8()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 3, 1);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest9()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 3, 0);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest10()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 3, -9);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest11()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, -0.01, 0.5);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest12()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 1.01, 0.5);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest13()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 0.8, -0.01);
            Assert.IsFalse(target.IsPrimed);
        }

        /// <summary>
        /// A test for constructor exception.
        /// </summary>
        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void KaufmanAdaptiveMovingAverageConstructorTest14()
        {
            var target = new KaufmanAdaptiveMovingAverage(5, 0.8, 1.01);
            Assert.IsFalse(target.IsPrimed);
        }
        #endregion

        #region SerializationTest
        private static void SerializeTo(KaufmanAdaptiveMovingAverage instance, string fileName)
        {
            var dcs = new DataContractSerializer(typeof(KaufmanAdaptiveMovingAverage), null, 65536, false, true, null);
            using (var fs = new FileStream(fileName, FileMode.Create))
            {
                dcs.WriteObject(fs, instance);
                fs.Close();
            }
        }

        private static KaufmanAdaptiveMovingAverage DeserializeFrom(string fileName)
        {
            var fs = new FileStream(fileName, FileMode.Open);
            XmlDictionaryReader reader = XmlDictionaryReader.CreateTextReader(fs, new XmlDictionaryReaderQuotas());
            var ser = new DataContractSerializer(typeof(KaufmanAdaptiveMovingAverage), null, 65536, false, true, null);
            var instance = (KaufmanAdaptiveMovingAverage)ser.ReadObject(reader, true);
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
            double d, u; const int digits = 9;
            var source = new KaufmanAdaptiveMovingAverage(10);
            for (int i = 0; i < 111; ++i)
            {
                source.Update(rawInput[i]);
                d = Math.Round(source.Value, digits);
                u = Math.Round(kama[i], digits);
                if (Math.Abs(u - d) > double.Epsilon)
                    System.Diagnostics.Debug.WriteLine("{0}: F {1} => {2} -> {3} | {4} <- {5}", i, rawInput[i], source.Value, d, u, kama[i]);
                Assert.AreEqual(u, d);
            }
            const string fileName = "KaufmanAdaptiveMovingAverage_1.xml";
            SerializeTo(source, fileName);
            KaufmanAdaptiveMovingAverage target = DeserializeFrom(fileName);
            Assert.AreEqual(source.Value, target.Value);
            for (int i = 111; i < rawInput.Count; ++i)
            {
                target.Update(rawInput[i]);
                d = Math.Round(target.Value, digits);
                u = Math.Round(kama[i], digits);
                if (Math.Abs(u - d) > double.Epsilon)
                    System.Diagnostics.Debug.WriteLine("{0}: F {1} => {2} -> {3} | {4} <- {5}", i, rawInput[i], target.Value, d, u, kama[i]);
                Assert.AreEqual(u, d);
            }
            //FileInfo fi = new FileInfo(fileName);
            //fi.Delete();
        }
        #endregion
    }
}
