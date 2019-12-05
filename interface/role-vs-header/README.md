# Interface: Role vs Header

this case study will present which interface type is usable in which situation.

## TL;DR

A role interface aim to define the most minimum expectation of the consumer from a supplier.
This achieve fruits like [ISP](/principle/ISP.md) and [RAP](/principle/RAP.md).
The main goal of the role interface is that you first define the need of the consumer
then you implement it with a supplier.

This is contrast to a Header Interface where the interface defined 
based on the publicly (exported) available functions of a concrete supplier implementation,
and then this is offered to the consumer to try use it to its need.  

## What is Role Interface

Role Interfaces provide a way to specify the behavior of an object: if something can do this, then it can be used here.
We've seen a couple of simple examples already; custom printers can be implemented by a String method while Fprintf can generate output to anything with a Write method. 
Role Interfaces with only one or two methods are common.
A certain supplier can implement multiple interfaces.

The strength of a role interface is that it clearly communicates the actual collaboration between an activity and its supplier.
Often a class doesn't use all the methods of a class, so it's good to show which ones are actually needed.
This can be particularly useful if you need to substitute it later on.

The disadvantage of a role interface is that it's more effort to come up with,
since you need to look at each collaboration to form the role interface,
while a header interface all you have to do is duplicate the public methods, no thought needed.

In dynamic level languages, where Duck typing supported, role interface comes naturally,
since the interface takes form as what methods actually called at runtime need to be implemented on them,
while explicit interfaces require implementation of all methods at compile time.

### Where to use ?

Role Interfaces works best to represent implementation of business logic, and external resources handlers.
These type of dependencies often used within business logic layers (interactors),
where the goal is to describe a business logic rule or action.

### Where to avoid ?

Role interfaces work counter intuitive when you work on low level concrete implementations
in the external interface or external resource handling layer.
For situations where you want to create code that you mainly test with integration test,
and you only want to use interface to test the code for rainy paths trough mocking.

## What is Header Interface

In some languages such as C++ you must define header interfaces.
These constructs define how a public type is exposed in terms of its public interface.
Other languages take a different route.
C# or Java do not require headers but they are still very much in circulation.
This unfortunately brings along some unwanted side effects.

A header interface is an interface that mimics the implicit public interface of a class.
Essentially you take all or most of the public methods of a class and declare them in an interface.
You can then supply an alternative implementation for the class
which is usually takes place in the for of mocking/stubbing.

One of the disadvantage of header interface forces you to implement every method,
even if you're not going to need them; but with a role interface you only need to implement exactly what's needed.

### Where to use ?

When you need to test rainy paths of a implementation that works with a concrete technology to manage external resources.
Such example is an API client.
To naturally create scenarios where during a happy path usage, the API fails unexpectedly,
it can be challenging, unless the concrete technology dependency replaced with a header interface,
and then a mock implementation provided for the rainy path tests.

### Where to avoid ?

When creating business rules (interactor), we must avoid the usage of a header interface,
because dependency on the interface doesn't invert the flow of dependency in reality,
and thus our business rule logic will be coupled together with a certain technology stack.
This will set constrains on the ability to later on evolve the code base depending on the needs.

## Example

* [example to header interface](/interface/role-vs-header/header)
* [example to role interface](/interface/role-vs-header/role)

## Costs of using this pattern

Using Role interfaces enforce additional unit test coverage requirements,
to test business rules without the concrete implementation including rainy paths.

In case the consumer expect that the supplier will be stateful depending on interactions,
additional specification/test for describing the expected behavior is heavily suggested,
to create a contract from the consumer point of view towards the supplier.

In case this spec coverage done in a way where we only test the concrete supplier implementation,
then writing it may be cheaper, but they are less valuable from architecture point of view,
since they cannot be reused, later on.

In most cases, where the project explicitly states, that there is no way, that an specific technology stack will change under the hood,
then the trade-off of writing the test specification bundled together with the supplier is acceptable in a project lifecycle,
else a purely behavior based interface would have to be used, where Create/Read/Delete operations would be necessary,
to setup the test runtime context.

* role based testing conventions to cover requirements from consumer towards suppliers

## Benefits of using this pattern

* Architecture flexibility
* [less risk of Feature Envy Smell](http://wiki.c2.com/?FeatureEnvySmell)
* [less risk of leaky abstraction](https://en.wikipedia.org/wiki/Leaky_abstraction)
* [support Reused Abstractions Principle in the code base](/principle/RAP.md)
* [support Dependency Inversion Principle in the code base](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
    * Harder to switch objects via DI as you are forced to implement all members even if you do not use the whole interface.
* [support Interface Segregation Principle in the code base](https://en.wikipedia.org/wiki/Interface_segregation_principle)
    * Header interfaces tend to break the Interface Segregation Principle.
* Meyer's open/closed principle in the form of Middleware pattern using embedding
* Polymorphic open/closed principle using multiple implementations
* Components become loosely coupled in the system
* removes the overhead of headers in languages that do not need them
