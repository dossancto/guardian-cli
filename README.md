# Guardian CLI

Guardian CLI, used to geneate and help using guardian.

> Just for Fun

# Run

```sh
go run
```

# Build

```sh
go build -o ./bin/guardian .
```

# TODO

- Use Cases

	- [ ] Generate Select Use Case

	- [ ] Generate Update Use Case

	- [ ] Generate Delete Use Case

- Controllers

	- [ ] Generate Dependency Injection Controller

	- [ ] Generate Management Controller

	- [ ] Generate Get Controller

- Editing existing code

	- [ ] Add depedency injection methods (Use Cases and Repositories)

	- [ ] Add DbSet Entity in ApplicationDbContext

- Init

	- [ ] Add Customize project creation, as choosing the ORM, database, repository pattern, use graphql and/or REST

- Configuration

    - [ ] Create a config option file that defines a default rule for scaffolding.

# FIX

- Repositories

	- [ ] Fix ApplicationDbContext Class Name. In some locations, the classname has a typo.
