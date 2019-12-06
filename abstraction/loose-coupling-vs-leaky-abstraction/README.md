# Loose coupling vs Leaky Abstraction

## What is Decoupling ?

'Coupling' is a term that describes the relationship between two entities in a software system.
When a class uses another class, or communicates with it,
it's said to 'depend' on that other class, and so these classes are 'coupled'.
At least one of them 'knows' about the other.

The idea is that we should try to keep the coupling between classes in our systems as 'loose' as possible: hence 'loose coupling' or sometimes 'decoupling' 
(although in English 'decoupling' would mean 'no coupling at all', people often use it to imply 'loose coupling' between entities).

So: what is loose-coupling versus strong coupling in practice, and why should we make entities loosely-coupled?

Coupling describes the degree of dependency between one entity to another entity. Often classes or objects.

When ClassA depends heavily on ClassB, the chances of ClassA being affected when ClassB is changed are high. This is strong coupling.

However if ClassA depends lightly on ClassB, than the chances of ClassA being affected in any way by a change in the code of ClassB, are low. This is loose coupling, or a 'decoupled' relationship.

Loose coupling is good because we don't want the components of our system to heavily depend on each other. We want to keep our system modular, where we can safely change one part without affecting the other.

When two parts are loosely coupled, they are more independent of each other and are less likely to break when the other changes.

For example, when building a car, you wouldn't want an internal change in the engine to break something in the steering wheel.

While this would never happen by accident when building a car, similar things happen to programmers all the time. Loose coupling is meant to reduce the risk of such things happening.

Strong coupling usually occurs when entity A knows too much about entity B. If entity A makes too many assumptions about how entity B operates or how it is built, than there is a high risk that a change in entity B will affect entity A. This is because one of it's assumptions about entity B are now incorrect.

For example, imagine that as a driver, you would make certain assumptions about how the engine of your car works.

The day you buy a new car with an engine that works differently (or for some reason your engine was replaced), your previous assumptions would be incorrect. If you were code in a computer, you would now be incorrect code that doesn't work properly.

However, if all the assumptions that as a driver you made about cars is that: A- they have steering wheels and B- they have brake and gas pedals, than changes in the car won't affect you, as long as your few assumptions stay correct. This is loose coupling.

An important technique to achieve loose coupling is Encapsulation. The idea is that a class hides it's internal details from other classes, and offers a strictly defined interface for other classes to communicate with it.

So for example, if you were defining a class Car, it's interface (public methods) would probably be drive(), stop(), steerLeft(), steerRight(), getSpeed(). These are the methods other objects can invoke on Car objects.

All of the other details of the Car class: how the engine works, kind of fuel it uses, etc. are hidden from other classes - to prevent them from knowing too much about Car.

The moment class A knows too much about class B: we have a strongly coupled relationship, where class A is too dependent on class B and a change in class B is likely to affect class A. Making the system hard to expand and maintain.

A relationship between two entities, where they know little about each other (only what's necessary) - is a loosely coupled, or decoupled relationship.

## Here's a [meatspace](https://www.urbandictionary.com/define.php?term=meatspace) example from Mark E. Haase:

Automobiles have abstractions for drivers.
In its purest form, there's a steering wheel, accelerator and brake.
This abstraction hides a lot of detail about what's under the hood: engine, cams, timing belt, spark plugs, radiator, etc.
The neat thing about this abstraction is that we can replace parts of the implementation with improved parts without retraining the user.
Let's say we replace the distributor cap with electronic ignition, and we replace the fixed cam with a variable cam.
These changes improve performance but the user still steers with the wheel and uses the pedals to start and stop.
It's actually quite remarkable...
a 16 year old or an 80 year old can operate this complicated piece of machinery without really knowing much about how it works inside!

But there are leaks.
The transmission is a small leak.
In an automatic transmission you can feel the car lose power for a moment as it switches gears, 
whereas in CVT you feel smooth torque all the way up.
There are bigger leaks, too.
If you rev the engine too fast, you may do damage to it.
If the engine block is too cold, the car may not start or it may have poor performance.
And if you crank the radio, headlights, and AC all at the same time, you'll see your gas mileage go down.



### What is Role Interface

A role interface is defined by looking at a specific interaction between suppliers and consumers.
A supplier component will usually implement several role interfaces, one for each of these patterns of interaction. 
This contrasts to a HeaderInterface, where the supplier will only have a single interface.
In your role interface, you define behavior you expect from that object,
you keep your interface limited to the things you only need for your current goal.
In the interface you should try to avoid function signature's which defines the implementation details,
and focus purely on the behavior aspect about business entities or business interactions. 

By doing this we can achieve architecture flexibility which allows the implementer to decide,
if certain roles interfaces belong to the same topic or to a separate one,
and this allow to structure the code more naturally as well.
Also the focus on the behavior allow us to change implementation easily,
without the need to change code that is irrelevant from the implementation details.

We can think of a successor as a role that a collaborating object plays with respect to this object. 
This approach of thinking about objects and the roles that they play in collaborating with others has a long history in the object-oriented world.
The strength of a role interface is that it clearly communicates the actual collaboration between an activity and its successors. 
Often a class doesn't use all the methods of a class, so it's good to show which ones are actually needed. 
This can be particularly useful if you need to substitute it later on. 
A header interface forces you to implement every method, even if you're not going to need them; but with a role interface you only need to implement exactly what's needed.
The disadvantage of a role interface is that it's more effort to come up with, since you need to look at each collaboration to form the role interface. With a header interface all you have to do is duplicate the public methods, no thought needed. There is also a sense of dependence on your consumers. I say a "sense" because there isn't a formal dependency, but it's still enough to make many people uncomfortable. They prefer header interfaces because they believe that you shouldn't care who uses your service, or how. You publish an interface and they can use it if they find it useful.
On the whole I much prefer role interfaces, so I suggest pushing towards them as much as you can. There is work involved in doing it, but my belief has always been that you should only use interfaces when you really need substitutability, and if you do need interfaces you should think hard about what the consumer of that interface needs.
There is an interesting twist if you think of this in a remoting context using something like web services. If we were to ask a remote service for details on predecessors, what should we expect back? Some might argue that for it to be a role interface, it should only return a document with earliest finish data. I would disagree - I think it's perfectly valid for it to return a document with more data than I asked for. The point is that any type checking involved should only check that the earliest finish data is present. If extra data can be ignored, then it's no crime to supply it; just as a class may implement multiple interfaces. This thinking matches the philosophy of Consumer Driven Contracts, which is one of the reasons I find consumer driven contracts so compelling.
As I've indicated, this notion has been around for a long time. Trygve Reenskaug wrote a methodology book based around analyzing roles and synthesizing them into classes. Robert Martin talks about this topic as the Interface Segregation Principle: role interfaces follow that principle but header interfaces do not.

#### Where to use ?

Role Interfaces works best to represent implementation of business logic, and external resources handlers.
These type of dependencies often used within business logic layers (interactors),
where the goal is to describe a business logic rule or action.  

#### Where to avoid ?

Role interfaces work counter intuitive when you work on low level concrete implementations 
in the external interface or external resource handling layer.
For situations where you want to create code that you mainly test with integration test,
and you only want to use interface to test the code for rainy paths trough mocking.  

### What is Header Interface

A header interface is an explicit interface that mimics the implicit public interface of a class. 
Essentially you take all or most of the public methods of a class and declare them in an interface. 
You can then supply an alternative implementation for the class 
which is usually takes place in the for of mocking/stubbing. 

#### Where to use ?

When you need to test rainy paths of a implementation that works with a concrete technology to manage external resources.
Such example is an API client. 
To naturally create scenarios where during a happy path usage, the API fails unexpectedly,
it can be challenging, unless the concrete technology dependency replaced with a header interface,
and then a mock implementation provided for the rainy path tests.

#### Where to avoid ?

When creating business rules (interactor), we must avoid the usage of a header interface,
because dependency on the interface doesn't invert the flow of dependency in reality,
and thus our business rule logic will be coupled together with a certain technology stack.
This will set constrains on the ability to later on evolve the code base depending on the needs. 

### Example

In example, we will examine a misuse and then a  use first Header Role Interface To R

* [example to header interface misuse](/interface/role-vs-header/misuse.go)
* [example to role interface usage](/interface/role-vs-header/usage.go)
 
### Costs of using this pattern

Using Role interfaces enforce additional unit test coverage requirements, 
to test business rules without the concrete implementation,
And in the case of working with stateful dependencies,
additional test for behavior specification is suggested from the business rule point of view towards the implementor type,
that described the expected edge cases regarding changes with a resource state.   

This extra spec coverage usually done in the form of implicit specification,
where the concrete implementation has it's own test coverage,
that defines the expectations.
They are cheaper to write, but less valuable from architecture point of view,
since they cannot be reused, later on.
In most of the case, this trade-off price is acceptable in a project lifecycle.  

Also if the system is decoupled, it is more difficult to provide transactional integrity; 
additional coordination protocols are required.
Such solution is when the coordinator object passed along encapsulated in some of request environment/context object,
and transactional integrity ensured outside of the business rules.
This becomes concern for cases where Crate/Delete/Update operations are involved on external a resources,
and transactional rollback must be guaranteed.
Nonetheless, it is far from being impossible, patterns like pipeline pattern can solve this problem. 

* role based testing conventions to cover requirements from consumer towards suppliers
* the need of decoupled transaction coordinator
* possibly the need of decoupled transaction coordinator

### Benefits of using this pattern

By Using 

* Architecture flexibility
* [lower]
* [less risk of Feature Envy Smell](http://wiki.c2.com/?FeatureEnvySmell)
* [less risk of leaky abstraction](http://wiki.c2.com/?FeatureEnvySmell)
* [applied usage of Dependency Inversion Principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
* [applied usage of Interface Segregation Principle](https://en.wikipedia.org/wiki/Interface_segregation_principle)
* Meyer's open/closed principle in the form of Middleware pattern using embedding
* Polymorphic open/closed principle using multiple implementations
* Components become loosely coupled in the system

