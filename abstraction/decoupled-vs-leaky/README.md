# Leaky Abstraction

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

