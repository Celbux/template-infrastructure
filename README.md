# Celbux Template Infrastructure

## Further Reading

For detailed documentation on the web framework see [web](./docs/web.md).

## Project Setup

This project has a couple of top level layers:

| Dir              | Contains                                                     |
| ---------------- | ------------------------------------------------------------ |
| `services`       | Executables running background tasks                         |
| `business`       | domain logic, frameworks and platform dependencies           |
| `infrastructure` | third party dependencies, mostly CRUD logic                  |
| `foundation`     | reusable tooling, generally this code should have nothing to do with Celbux |
| `configs`        | configuration files for Docker, App Engine, Cloud Run etc.   |

Other than that, the top level of the project has a `go.mod` file, this `README` file as well as a `makefile`. The idea with the `makefile` is to add any project specific commands in there and to avoid having local scripts and aliases that hide complexity from the team. That being said, the `makefile` is not a free-for-all. It can get messy very fast so there should be structure to the file and discussion around any changes made.