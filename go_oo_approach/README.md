## Encapsulation

- Uppercase letters make fields, types, and functions public/exported.
- Lowercase letters make fields, types, and functions private/unexported.
- Encapsulation protects the internal state of a type and exposes controlled methods to access or modify it.

## Composition

- Go does not use traditional class inheritance.
- Instead, Go uses composition to build complex types by combining simpler types.
- Struct embedding allows one type to reuse the fields and methods of another type.

## Polymorphism

- Polymorphism allows different types to be treated through the same interface.
- In Go, polymorphism is achieved using interfaces.
- Any type that implements the required methods satisfies the interface automatically.

## Interfaces

- An interface defines a contract that specifies the methods a type must implement.
- Interfaces describe behavior, not data.
- Types implement interfaces implicitly; no explicit declaration is required.
