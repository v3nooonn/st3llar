# st3llar

This is a repo for Cobra experiment, which is aim to create a CLI and follows some best practices rules.

### Principles

Principles for a good CLI: [REFERENCE](https://clig.dev/)

- Human-first design

  If a command is going to be used primarily by humans, it should be designed for humans first.

- Simple parts that work together

  A core tenet of the original UNIX philosophy is the idea that small, simple programs with clean interfaces can be combined to build larger systems.  
  Whatever software you’re building, you can be absolutely certain that people will use it in ways you didn't anticipate.  
  Your software will become a part in a larger system: your only choice is over whether it will be a well-behaved part.

- Consistency across programs

  The terminal’s conventions are hardwired into our fingers.  
  We had to pay an upfront cost by learning about command line syntax, flags, environment variables and so on, but it pays off in long-term efficiency… as long as programs are consistent.  
  Where possible, a CLI should follow patterns that already exist.
 `st3lalr VERB NOUN --ADJECTIVE` or `st3llar COMMAND ARG --FLAG`

- Output is just enough

  The terminal is a world of pure information.  
  You could make an argument that information is the interface—and that, just like with any interface, there’s often too much or too little of it.

- Easy to discover and learn

  Discoverable CLIs have comprehensive help texts, provide lots of examples, suggest what command to run next, suggest what to do when there is an error.  
  There are lots of ideas that can be stolen from GUIs to make CLIs easier to learn and use, even for power users.
- Conversation as the norm
- Robustness

### Cobra Concepts

In Cobra, it's based on these three main concepts:

- Command
- Args
- Flags