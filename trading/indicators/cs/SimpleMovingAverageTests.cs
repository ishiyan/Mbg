using System;
using System.Collections.Generic;
using Mbs.Trading.Data;
using Mbs.Trading.Indicators2;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace Mbs.UnitTests.Trading.Indicators2
{
    [TestClass]
    public class SimpleMovingAverageTests
    {
        /// <summary>
        /// Perry Kaufman, Trading Systems an Methods, 3rd edition, page 72.
        /// </summary>
        private readonly List<double> kaufmanInput = new List<double>
        {
            64.59, 64.23, 65.26, 65.24, 65.07, 65.14, 64.98, 64.76, 65.11, 65.46,
            65.94, 66.10, 66.87, 66.56, 66.71, 66.19, 66.14, 66.64, 67.33, 68.18,
            67.48, 67.19, 66.46, 67.20, 67.62, 67.66, 67.89, 69.19, 69.68, 69.31,
            69.11, 69.27, 68.97, 69.11, 69.50, 69.70, 69.94, 69.11, 67.64, 67.75,
            67.47, 67.50, 68.18, 67.35, 66.74, 67.00, 67.46, 67.36, 67.37, 67.78,
            67.96,
        };

        /// <summary>
        /// Perry Kaufman, Trading Systems an Methods, 3rd edition, page 72.
        /// </summary>
        private readonly List<double> kaufmanExpected3 = new List<double>
        {
            double.NaN, double.NaN, 64.69, 64.91, 65.19, 65.15, 65.06, 64.96, 64.95, 65.11,
            65.50, 65.83, 66.30, 66.51, 66.71, 66.49, 66.35, 66.32, 66.70, 67.38,
            67.66, 67.62, 67.04, 66.95, 67.09, 67.49, 67.72, 68.25, 68.92, 69.39,
            69.37, 69.23, 69.12, 69.12, 69.19, 69.44, 69.71, 69.58, 68.90, 68.17,
            67.62, 67.57, 67.72, 67.68, 67.42, 67.03, 67.07, 67.27, 67.40, 67.50,
            67.70,
        };

        /// <summary>
        /// Perry Kaufman, Trading Systems an Methods, 3rd edition, page 72.
        /// </summary>
        private readonly List<double> kaufmanExpected5 = new List<double>
        {
            double.NaN, double.NaN, double.NaN, double.NaN, 64.88, 64.99, 65.14, 65.04, 65.01, 65.09,
            65.25, 65.47, 65.90, 66.19, 66.44, 66.49, 66.49, 66.45, 66.60, 66.90,
            67.15, 67.36, 67.33, 67.30, 67.19, 67.23, 67.37, 67.91, 68.41, 68.75,
            69.04, 69.31, 69.27, 69.15, 69.19, 69.31, 69.44, 69.47, 69.18, 68.83,
            68.38, 67.89, 67.71, 67.65, 67.45, 67.35, 67.35, 67.18, 67.19, 67.39,
            67.59,
        };

        /// <summary>
        /// Perry Kaufman, Trading Systems an Methods, 3rd edition, page 72.
        /// </summary>
        private readonly List<double> kaufmanExpected10 = new List<double>
        {
            double.NaN, double.NaN, double.NaN, double.NaN, double.NaN, double.NaN, double.NaN, double.NaN, double.NaN, 64.98,
            65.12, 65.31, 65.47, 65.60, 65.76, 65.87, 65.98, 66.17, 66.39, 66.67,
            66.82, 66.93, 66.89, 66.95, 67.04, 67.19, 67.37, 67.62, 67.86, 67.97,
            68.13, 68.34, 68.59, 68.78, 68.97, 69.17, 69.38, 69.37, 69.17, 69.01,
            68.85, 68.67, 68.59, 68.41, 68.14, 67.87, 67.62, 67.45, 67.42, 67.42,
            67.47,
        };

        [TestMethod]
        public void SimpleMovingAverage_Name_HasCorrectValue()
        {
            var target = new SimpleMovingAverage(5);
            Assert.AreEqual("SMA", target.Name);
        }

        [TestMethod]
        public void SimpleMovingAverage_Moniker_HasCorrectValue()
        {
            var target = new SimpleMovingAverage(5);
            Assert.AreEqual("SMA5", target.Moniker);
        }

        [TestMethod]
        public void SimpleMovingAverage_Description_HasCorrectValue()
        {
            var target = new SimpleMovingAverage(5);
            Assert.AreEqual("Simple Moving Average", target.Description);
        }

        [TestMethod]
        public void SimpleMovingAverage_IsPrimed_HasCorrectValue()
        {
            var target = new SimpleMovingAverage(5);
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

        [TestMethod]
        public void SimpleMovingAverage_Value_HasCorrectValue()
        {
            var target = new SimpleMovingAverage(5);
            Assert.IsTrue(double.IsNaN(target.Value));

            var scalar = new Scalar(DateTime.Now, 1d);
            Assert.IsTrue(double.IsNaN(target.Update(scalar).Value));

            scalar.Value = 2d;
            Assert.IsTrue(double.IsNaN(target.Update(scalar).Value));

            scalar.Value = 3d;
            Assert.IsTrue(double.IsNaN(target.Update(scalar).Value));

            scalar.Value = 6d;
            Assert.IsTrue(double.IsNaN(target.Update(scalar).Value));

            scalar.Value = 3.5;
            Assert.AreEqual(3.1, target.Update(scalar).Value);

            scalar.Value = 0.5;
            Assert.AreEqual(3d, target.Update(scalar).Value);
        }

        [TestMethod]
        public void SimpleMovingAverage_Length_HasCorrectValue()
        {
            var target = new SimpleMovingAverage(11);
            Assert.AreEqual(11, target.Length);

            target = new SimpleMovingAverage(22);
            Assert.AreEqual(22, target.Length);
        }

        [TestMethod]
        public void SimpleMovingAverage_Update_ReturnsCorrectValues()
        {
            var scalar = new Scalar(DateTime.Now, 1d);
            int count = kaufmanInput.Count;

            var target = new SimpleMovingAverage(3);
            for (int i = 0; i < 2; i++)
            {
                scalar.Value = kaufmanInput[i];
                double d = target.Update(scalar).Value;
                Assert.IsTrue(double.IsNaN(d));
            }

            for (int i = 2; i < count; i++)
            {
                scalar.Value = kaufmanInput[i];
                double d = target.Update(scalar).Value;
                Assert.AreEqual(kaufmanExpected3[i], Math.Round(d, 2));
            }

            target = new SimpleMovingAverage(5);
            for (int i = 0; i < 4; i++)
            {
                scalar.Value = kaufmanInput[i];
                double d = target.Update(scalar).Value;
                Assert.IsTrue(double.IsNaN(d));
            }

            for (int i = 4; i < count; i++)
            {
                scalar.Value = kaufmanInput[i];
                double d = target.Update(scalar).Value;
                Assert.AreEqual(kaufmanExpected5[i], Math.Round(d, 2));
            }

            target = new SimpleMovingAverage(10);
            for (int i = 0; i < 9; i++)
            {
                scalar.Value = kaufmanInput[i];
                double d = target.Update(scalar).Value;
                Assert.IsTrue(double.IsNaN(d));
            }

            for (int i = 9; i < count; i++)
            {
                scalar.Value = kaufmanInput[i];
                double d = target.Update(scalar).Value;
                Assert.AreEqual(Math.Round(kaufmanExpected10[i], 1), Math.Round(d, 1));
            }
        }

        [TestMethod]
        public void SimpleMovingAverage_Calculate_ReturnsCorrectValues()
        {
            List<double> actual = SimpleMovingAverage.Calculate(kaufmanInput, 3);
            for (int i = 0; i < 2; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            for (int i = 2; i < kaufmanInput.Count; i++)
            {
                Assert.AreEqual(kaufmanExpected3[i], Math.Round(actual[i], 2));
            }

            actual = SimpleMovingAverage.Calculate(kaufmanInput, 5);
            for (int i = 0; i < 4; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            for (int i = 4; i < kaufmanInput.Count; i++)
            {
                Assert.AreEqual(kaufmanExpected5[i], Math.Round(actual[i], 2));
            }

            actual = SimpleMovingAverage.Calculate(kaufmanInput, 10);
            for (int i = 0; i < 9; i++)
            {
                Assert.IsTrue(double.IsNaN(actual[i]));
            }

            for (int i = 9; i < kaufmanInput.Count; i++)
            {
                Assert.AreEqual(kaufmanExpected10[i], Math.Round(actual[i], 2));
            }
        }

        [TestMethod]
        public void SimpleMovingAverage_Reset_FunctionsCorrectly()
        {
            var target = new SimpleMovingAverage(3);
            Assert.IsTrue(double.IsNaN(target.Update(1d)));
            Assert.IsTrue(double.IsNaN(target.Update(2d)));
            Assert.AreEqual(2d, target.Update(3d));

            target.Reset();
            Assert.IsTrue(double.IsNaN(target.Update(1d)));
            Assert.IsTrue(double.IsNaN(target.Update(2d)));
            Assert.AreEqual(2d, target.Update(3d));
        }

        [TestMethod]
        public void SimpleMovingAverage_Constructor_CreatesAnInstance()
        {
            var target = new SimpleMovingAverage(5);
            Assert.AreEqual(5, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.ClosingPrice);

            target = new SimpleMovingAverage(15, OhlcvComponent.MedianPrice);
            Assert.AreEqual(15, target.Length);
            Assert.IsTrue(double.IsNaN(target.Value));
            Assert.IsFalse(target.IsPrimed);
            Assert.IsTrue(target.OhlcvComponent == OhlcvComponent.MedianPrice);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void SimpleMovingAverage_Constructor_LengthIsOne_ThrowsException()
        {
            var target = new SimpleMovingAverage(1);
            Assert.IsNotNull(target);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void SimpleMovingAverage_Constructor_LengthIsZero_ThrowsException()
        {
            var target = new SimpleMovingAverage(0);
            Assert.IsNotNull(target);
        }

        [TestMethod]
        [ExpectedException(typeof(ArgumentOutOfRangeException))]
        public void SimpleMovingAverage_Constructor_LengthIsNegative_ThrowsException()
        {
            var target = new SimpleMovingAverage(-8);
            Assert.IsNotNull(target);
        }
    }
}
