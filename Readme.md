# Temporal Phantom Worker

Temporal Phantom Worker is a Go-based project designed to facilitate testing Temporal workflows in environments where not all microservices are available. It allows users to create workflow stubs that register workflows and activities, returning predefined responses based on configuration.

**Work in Progress:** Currently, the Phantom Worker only connects to localhost on port 7233, using the default namespace.

## Table of Contents

- [Usage](#usage)
- [Configuration](#configuration)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Usage

To run the project, use the following command:

```bash
go run main.go start -c ./config/sample1.yaml
```

## Configuration
The configuration file should be in YAML format and define the workers, workflows, and activities. Each worker can have multiple workflows and activities, along with their expected results.

### Example Config File
Here is an example configuration file ([sample1.yaml](config/sample1.yaml)):

```yaml
workers:
  - name: SimpleWorker
    task_queue: simple
    workflows:
      - type: HelloWorldWorkflow
        result: "Hello World"
    activities:
      - type: HelloWorldActivity
        result: "Hello World"
  - name: ComplexWorker
    task_queue: complex
    workflows:
      - type: HelloWorldWorkflow
        result:
          supportedLanguagesCount: 2
          supportedLanguages:
            - english
            - maltese
          messages:
            english: "Hello World"
            maltese: "Aw Dinja!"
      - type: GoodbyeWorldWorkflow
        result:
          supportedLanguagesCount: 2
          supportedLanguages:
            - english
            - maltese
          messages:
            english: "Goodbye World"
            maltese: "Ċaw Dinja!"
    activities:
      - type: HelloWorldActivity
        result:
          supportedLanguagesCount: 2
          supportedLanguages:
            - english
            - maltese
          messages:
            english: "Hello World"
            maltese: "Aw Dinja!"
      - type: GoodbyeWorldActivity
        result:
          supportedLanguagesCount: 2
          supportedLanguages:
            - english
            - maltese
          messages:
            english: "Goodbye World"
            maltese: "Ċaw Dinja!"
```

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or encounter any problems.

