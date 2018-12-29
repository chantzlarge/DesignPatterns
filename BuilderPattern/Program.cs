using System;

namespace BuilderPattern
{
    class Program
    {
        static void Main(string[] args)
        {
            Builder builder = new ConcreteBuilder();
            Director director = new Director(builder);

            director.Construct();
            Product product = builder.GetResult();

            product.Show();

            Console.ReadKey();
        }
    }

    // <summary>
    // Constructs an object using the Builder interface.
    // </summary>
    public class Director
    {
        Builder _builder;

        public Director(Builder builder)
        {
            _builder = builder;
        }

        public void Construct()
        {
            _builder.BuildPart();
        }
    }

    // <summary>
    // Specifies an abstract interface for creating parts of a Product object.
    // </summary>
    public abstract class Builder
    {
        public abstract void BuildPart();
        public abstract Product GetResult();
    }

    // <summary>
    // Constructs and assembles parts of the product by implementing the Builder interface.
    // Defines and keeps track of the representation it creates.
    // Provides an interface for retrieving the product
    // </summary>
    public class ConcreteBuilder : Builder
    {
        private Product _product;

        public override void BuildPart()
        {
            _product = new Product();
        }

        public override Product GetResult()
        {
            return _product;
        }
    }

    // <summary>
    // Represents the complex object under construction. 
    // ConcreteBuilder builds the product's internal representation and defines the process by which it's assembled.
    // Includes classes that define the constituent parts, including interfaces for assembling the parts into the final result.
    // </summary>
    public class Product {
        public void Show()
        {
            Console.WriteLine("Product");
        }
    }
}
