using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;
using System.IO;

namespace AbstractFactoryPattern.Test
{
    [TestClass]
    public class ClientTest
    {
        [TestMethod]
        public void Run_ConcreteFactory1_PrintMessage()
        {
            // arrange
            AbstractFactory mockAbstractFactory = new ConcreteFactory1();
            Client client = new Client(mockAbstractFactory);
            StringWriter writer = new StringWriter();
            Console.SetOut(writer);
            
            // act
            client.Run();

            // assert
            Assert.AreEqual(String.Format("ProductB1 interacts with ProductA1{0}", Environment.NewLine), writer.ToString());
        }

        [TestMethod]
        public void Run_ConcreteFactory2_PrintMessage()
        {
            // arrange
            AbstractFactory mockAbstractFactory = new ConcreteFactory2();
            Client client = new Client(mockAbstractFactory);
            StringWriter writer = new StringWriter();
            Console.SetOut(writer);

            // act
            client.Run();

            // assert
            Assert.AreEqual(String.Format("ProductB2 interacts with ProductA2{0}", Environment.NewLine), writer.ToString());
        }
    }
}