# Temporal Phantom Worker

Temporal Phantom Worker is a Go-based project designed to facilitate testing Temporal workflows in environments where
not all microservices are available. It allows users to create worker stubs that register workflows and activities,
returning predefined responses or errors based on the provided configuration.

**Work in Progress:** Currently, the Phantom Worker only connects to localhost on port 7233, using the default
namespace.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Download the Binary**

   Go to the [Releases](https://github.com/brianzammit/temporal-phantom-worker/releases/latest) page and download the
   appropriate tar file for your operating system.


2. **Extract the Tar File**

   Use the following commands to extract the tar file:

```bash
# Linux/macOS
tar -xvf <tar-file>.tar.gz -C /desired/directory

# Windows (PowerShell)
tar -xvf <tar-file>.tar.gz -C C:\desired\directory

```

3. **Run the Application**

   After extraction, navigate to the directory and run the binary:

```bash
# Linux/macOS
./temporal-phantom-worker --help

# Windows
.\temporal-phantom-worker.exe --help
```

## Usage

To run the project, use the following commands:

### Validating configuration

```bash
./temporal-phantom-worker validate -c ./config/sample.yaml
```

### Starting Phantom Worker

```bash
./temporal-phantom-worker start -c ./config/sample.yaml
```

## Configuration

The configuration file should be in YAML format and define the workers, workflows, and activities. Each worker can have
multiple workflows and activities, along with their expected results.

### Configuration Options

| Field                | Type              | Description                                                                                                                     |
|----------------------|-------------------|---------------------------------------------------------------------------------------------------------------------------------|
| `workers`            | Array of Worker   | A list of workers to run, each defined by a name, task queue, workflows, and activities.                                        |
| ├─ `name`            | String            | The unique name of the worker.                                                                                                  |
| ├─ `task_queue`      | String            | The task queue to which the worker is polling for tasks on (should be unique per worker).                                       |
| ├─ `workflows`       | Array of Workflow | A list of workflows handled by the worker, defined by type, and a result or error.                                              |
| │   ├─ `type`        | String            | The unique identifier of the workflow type (name).                                                                              |
| │   ├─ `result`      | Any               | The expected result of the workflow, which may be a string, object, or number. Setting this denotes a successful Workflow stub. |
| │   ├─ `error`       | Error             | The expected error thrown by the workflow. Setting this denotes an error Workflow stub.                                         |
| │   │   ├─ `type`    | String            | The error type to be thrown. Example for simulating a java file not found error use: `java.io.FileNotFoundException`            |
| │   │   ├─ `message` | String            | A message to be included in the error                                                                                           |
| │   │   ├─ `details` | Any               | Any additional details to be included with the error (any valid yaml accepted)                                                  |
| ├─ `activities`      | Array of Activity | A list of activities handled by the worker, defined by type, and result or error.                                               |
| │   ├─ `type`        | String            | The unique identifier of the activity type (name).                                                                              |
| │   ├─ `result`      | Any               | The expected result of the activity, which may be a string, object, or number. Setting this denotes a successful Activity stub. |
| │   ├─ `error`       | Error             | The expected error thrown by the activity. Setting this denotes an error Activity stub.                                         |
| │   │   ├─ `type`    | String            | The error type to be thrown. Example for simulating a java file not found error use: `java.io.FileNotFoundException`            |
| │   │   ├─ `message` | String            | A message to be included in the error                                                                                           |
| │   │   ├─ `details` | Any               | Any additional details to be included with the error (any valid yaml accepted)                                                  |

### Example Config File

Here is an example configuration file ([sample.yaml](config/sample.yaml)):

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

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or
encounter any problems.

