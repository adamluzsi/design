# Reused Abstractions Principle (RAP)

Just time to revisit an OO design principle that I include in my courses - the Reused Abstractions Principle (RAP).

Planned reuse has been widely discredited in software development, especially after the Big Push For Reuse that we saw in the 90's.

Writing code on the basis that it may be applicable in multiple future scenarios is speculation, and speculative generality is considered a code smell. The danger is that when the future comes, the reusable generalisation we planned turns out to be not quite what is actually needed, or doesn't quite fit all the scenarios, or - more usually - that the future eventuality never comes up.

Software development is hard enough when we have to write code that satisfies today's requirements. Adding in a whole bunch of "what if?" code makes it even harder, we have found.

The Agile approach to software design is more scientific. When we see two examples of something, we refactor to remove that duplication. This may involve introducing a common base class, or a common interface. In short, we discover these abstractions during refactoring. They are then guaranteed to be reused, because they demonstrably already have been.

Reuse is a bit of a red herring economically. The real killer is the not the cost of writing new code. It's the cost of changing old code. So we should focus on combatting duplication as soon as it emerges, and not speculate that it might emerge later and on what form it might take.

If I was looking for signs of this discipline of discovering abstractions, then there's a clear indicator. Just as TDD implies very high test coverage (if your coverage is not high, you must have written code that wasn't required to pass a failing test), discovery of abstractions implies that abstractions are used more than once. Super classes have more than one subclass. Interfaces are implemented more than once. Abstract methods are overridden multiple times. And so on.

I introduced this principle 10 years ago when discussing OO design to counter the trend for speculative abstractions in many of the projects I encountered. Most particularly, the trend for putting interfaces everywhere, regardless of whether that interface represented some commonality betweren two or more types of object - which is what they're really for.

So I will stress that "programming to interfaces" is not good OO design, any more than "don't use getters and setters" is. These are overly simplistic formulas that mask more fundamental and nuanced principles. If the only class that ever implements the Customer interface is CustomerImpl, you don't really have polymorphism and substitutability because there is nothing in practice to substitute at runtime. It's fake generality. All you have is indirection and code clutter, which just makes the code harder to understand.