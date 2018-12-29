using System;

namespace AbstractFactoryPattern
{
    public class Program
    {
        static void Main(string[] args)
        {
            AbstractFactory abstractFactory1 = new ConcreteFactory1();
            AbstractFactory abstractFactory2 = new ConcreteFactory2();

            Client client1 = new Client(abstractFactory1);
            Client client2 = new Client(abstractFactory2);

            client1.Run();
            client2.Run();

            Console.ReadKey();
        }
    }

    // <summary>
    // Declares an interfcae for operations that create abstract product objects.
    // </summary>
    public abstract class AbstractFactory
    {
        public abstract AbstractProductA CreateProductA();
        public abstract AbstractProductB CreateProductB();
    }

    // <summary>
    // Implememts the operations to create concrete product objects.
    // </summary>
    public class ConcreteFactory1 : AbstractFactory
    {
        public override AbstractProductA CreateProductA()
        {
            return new ProductA1();
        }

        public override AbstractProductB CreateProductB()
        {
            return new ProductB1();
        }
    }

    // <summary>
    // Implememts the operations to create concrete product objects.
    // </summary>
    public class ConcreteFactory2 : AbstractFactory
    {
        public override AbstractProductA CreateProductA()
        {
            return new ProductA2();
        }

        public override AbstractProductB CreateProductB()
        {
            return new ProductB2();
        }
    }

    // <summary>
    // Declares an interface for a type of product object.
    // </summary>
    public abstract class AbstractProductA { }

    // <summary>
    // Defines a product object to be created by the corresponding concrete factory.
    // Implements the AbstractProduct interface.
    // </summary>
    public class ProductA1 : AbstractProductA { }

    // <summary>
    // Defines a product object to be created by the corresponding concrete factory.
    // Implements the AbstractProduct interface.
    // </summary>
    public class ProductA2 : AbstractProductA { }

    // <summary>
    // Declares an interface for a type of product object.
    // </summary>
    public abstract class AbstractProductB
    {
        public abstract void Interact(AbstractProductA abstractProductA);
    }

    // <summary>
    // Defines a product object to be created by the corresponding concrete factory.
    // Implements the AbstractProduct interface.
    // </summary>
    public class ProductB1 : AbstractProductB
    {
        public override void Interact(AbstractProductA abstractProductA)
        {
            Console.WriteLine(@"{0} interacts with {1}", this.GetType().Name, abstractProductA.GetType().Name);
        }
    }

    // <summary>
    // Defines a product object to be created by the corresponding concrete factory.
    // Implements the AbstractProduct interface.
    // </summary>
    public class ProductB2 : AbstractProductB
    {
        public override void Interact(AbstractProductA abstractProductA)
        {
            Console.WriteLine(@"{0} interacts with {1}", this.GetType().Name, abstractProductA.GetType().Name);
        }
    }

    // <summary>
    // Uses only the interfaces declared by AbstractFactory and AbstractProduct classes.
    // </summary>
    public class Client
    {
        private AbstractProductA _abstractProductA;
        private AbstractProductB _abstractProductB;

        public Client(AbstractFactory abstractFactory)
        {
            _abstractProductA = abstractFactory.CreateProductA();
            _abstractProductB = abstractFactory.CreateProductB();
        }

        public void Run()
        {
            _abstractProductB.Interact(_abstractProductA);
        }
    }
}
