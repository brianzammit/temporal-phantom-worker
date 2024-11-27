# Temporal Phantom Worker

Temporal Phantom Worker is a Go-based project designed to facilitate testing Temporal workflows and activities in environments where
parts of the system are unavailable. Functionality includes:
1. Creation of Temporal Worker stubs that register workflows and activities returning predefined responses or errors based on provided configuration
2. Testing Temporal Activities in isolation, without the need to trigger specific parent workflows

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

### Stub

#### Validating configuration

```bash
./temporal-phantom-worker stub validate -c ./config/sample.yaml
```

#### Starting Phantom Worker Stub

```bash
./temporal-phantom-worker stub start -c ./config/sample.yaml
```

#### Stub Configuration

The configuration file should be in YAML format and define the workers, workflows, and activities. Each worker can have
multiple workflows and activities, along with their expected results.

| Field                | Type              | Description                                                                                                                     |
|----------------------|-------------------|---------------------------------------------------------------------------------------------------------------------------------|
| `server`             | Server            | The Temporal Server to connect to.                                                                                              |
| ├─ `host`            | String            | The server hostname. Default: `localhost`.                                                                                      |
| ├─ `port`            | Number            | The server port. Default: `7233`.                                                                                               |
| ├─ `namespace`       | String            | The namespace to connect to. Default: `default`.                                                                                |
| ├─ `mtls`            | Mtls              | The mTLS configuration if mTLS is desired. mTLS will not be used if not present.                                                |
| │   ├─ `cert_path`   | String            | Path to the client certificate file.                                                                                            |
| │   ├─ `key_path`    | String            | Path to the client key file.                                                                                                    |
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

Example configuration files can be found in the  ([config directory](config)):


### Activity

#### Trigger

Use the activity trigger command to execute an activity in isolation by wrapping it in a dynamic workflow:

```bash
./temporal-phantom-worker activity trigger -type MyTestActivity -taskqueue testQueue --input <activtyInput.yaml>
```

The input file is an optional yaml file containing the input to pass to the activity. See the command's help for full list of options.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or
encounter any problems.

