# Single Responsibility Principle (SRP)

The single responsibility principle is a computer programming principle that states that every module,
class, or function should have responsibility over a single part of the functionality provided by the software,
and that responsibility should be entirely encapsulated by the class, module or function. 

The responsibility meaning in this context represent an action for a certain domain actor,
or in case it can't be related with an actor, then regarding a certain domain role.
An example to this is when you create a business rule (interactor),
then you should not focus on many actors like the Database

All its services should be narrowly aligned with that responsibility. Robert C. Martin expresses the principle as, "A class should have only one reason to change,"[1] although, because of confusion around the word "reason" he more recently stated "This principle is about people.(Actor)"[2]
