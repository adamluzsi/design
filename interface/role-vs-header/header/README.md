# Header interface example 

Notice that in order to "loose couple" the components,
Public/Exported functions had been added to a Header interface.

But in reality, this loose coupling is just virtual,
and in practice, it is unlikely that without much effort,
replacing the initial supplier would be feasible
unless the code and tests are changed of the consumer.

So in practice Consumer not depending on an expected behavior of a supplier, but on a exact supplier. 
