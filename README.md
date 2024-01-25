# Guardian CLI

Guardian CLI, used to geneate and help using guardian.

> Just for Fun


# Build

```sh
make build
```

# Instalation

> Required [guardian dotnet template](https://github.com/lu-css/Guardian-Template)

- Buld the application

    ```sh
    make build
    ```

- Install

    ```sh
    make install
    ```

# TODO

- WRITE TESTS

- Use Cases

	- [ ] Generate Select Use Case

	- [x] Generate Update Use Case

	- [ ] Generate Delete Use Case

- Controllers

	- [ ] Generate Dependency Injection Controller

	- [ ] Generate Management Controller

	- [ ] Generate Get Controller

	- [ ] Generate DTO Controller

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
