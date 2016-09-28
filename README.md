# Grover

Grover is a tool to generate version numbers for your packages.

## How does Grover determine the version number for my package?

Grover looks at the exported profile of your package and compares it to the
previous profile of your package. Depending on what changed it will then update
the different parts of the version number accordingly.

## Version Number.


The version number that is generated by Grover follows [Semantic Versioning](http://semver.org).
The version number will take the following form: *MAJOR*.*MINOR*.*PATCH*.

The Major version is incremented for a breaking change; that make the API of
the package incompatible.
The Minor version is incremented for when additional functionally is added in a
backward=compatible manner.
The Patch version is incremented for all other reasons.

If the Major is incremented, then the Minor and Patch is set to zero.
If the Minor is incremented, then the Patch is set to zero.

## Breaking Changes.

As there are various types of things that can be exported by a package, we will
look at each of them separately, and explain how that affects the version number.


---

Note, that renaming something is considered, adding a new thing, and removing the
old thing. If one were to rename a function from “a” to “b,” that would be
considered by the system as adding a function called “b” and removing a function
called “a.”

---

### Interfaces.


Interfaces have a set of methods that a type must implement to satisfy that
interface. These methods tend to be public, and so any changes to them cause
breaking changes to the API of the package.

#### The following modifications are considered breaking changes.

Removing an exported function from the interface is a breaking change — as other
relying on that exported function will need to change their code.

Adding a new exported function to the interface will cause a breaking change.

Changing the return values, or the parameters of an exported function is a
breaking change.

Adding a private function to an interface that does not already have private
functions. By adding a private function to an Interface, this makes it so that
others can not create types that will satisfy the interface. If they had a type
previously that met the interface, then they have to change their design to work
with the new API.


#### These changes will cause the Minor version to be updated.

Removing a private function from the interface.

Adding a private function to an interface that already has at least one private
function.


### Functions.


#### Breaking Changes.

Removing a function from the package.
Changing the signature of the function.

#### Minor Changes.

Adding a new function to the package.


### Variables and Constants

#### Breaking Changes


Removing a Variable or Constant.
Changing the type of a Variable or Constant.

#### Minor Changes

Adding a new Variable or Constant

### Types

#### Breaking Changes

Changing the underlining type of a custom type.
Removing a custom type.
Removing exported function of a custom type.
Changing the signature of a function of a custom type.

#### Minor Changes

Adding new custom type.
Adding a new exported function of a custom type.


### Type Struct

Also, the struct base type has the following additional rules.

#### Breaking changes

Removing a exported field.
Changing the type of an exported field.
Adding an exported atomic.Mutex field, as old software may need to be updated to
so that it properly locks and unlocks the field.

#### Minor changes

Adding a new field, exported or private.


