# The Interface Segregation Principle (ISP)

The Interface Segregation Principle (ISP) states that clients should not be forced to depend on methods that they do not use.
Interfaces should belong to clients, not to libraries or hierarchies. 
Application developers should favor thin, focused interfaces to “fat” interfaces that offer more functionality than a particular class or method needs.

In many languages, such as C# or Go, interfaces can inherit from multiple other interfaces. 
Thus, if you need a larger interface in some parts of the application,
but not in others, you may be able to compose it from two or more other interfaces. 
This is also a good approach to keep in mind if you find yourself refactoring a legacy codebase,
which already has large interfaces that you can’t break.

On high level, an Interfaces provide a way to specify the behavior of an object: 
if something can do this, then it can be used here.
Interfaces with only one or two methods are common in ISP, 
and are usually given a name derived from the method, 
such as Writer for something that implements Write.
