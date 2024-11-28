# Temporal Phantom Worker

Temporal Phantom Worker is a Go-based project designed to facilitate testing Temporal workflows and activities in
environments where
parts of the system are unavailable. Functionality includes:

1. Creation of Temporal Worker stubs that register workflows and activities returning predefined responses or errors
   based on provided configuration. Supports [result templating](#result-templating) to for dynamic results.
2. Testing Temporal Activities in isolation, without the need to trigger specific parent workflows

## Table of Contents

<!-- TOC -->
* [Temporal Phantom Worker](#temporal-phantom-worker)
  * [Table of Contents](#table-of-contents)
  * [Installation](#installation)
  * [Usage](#usage)
    * [Stub](#stub)
      * [Validating configuration](#validating-configuration)
      * [Starting Phantom Worker Stub](#starting-phantom-worker-stub)
      * [Stub Configuration](#stub-configuration)
      * [Result Templating](#result-templating)
    * [Activity](#activity)
      * [Trigger](#trigger)
  * [Contributing](#contributing)
<!-- TOC -->

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
./temporal-phantom-worker stub validate -c ./config/basic-success-sample.yaml
```

#### Starting Phantom Worker Stub

```bash
./temporal-phantom-worker stub start -c ./config/basic-success-sample.yaml
```

#### Stub Configuration

The configuration file should be in YAML format and define the workers, workflows, and activities. Each worker can have
multiple workflow and activity definitions, along with their expected results.

| Field                | Type              | Description                                                                                                                                                                                              |
|----------------------|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `server`             | Server            | The Temporal Server to connect to.                                                                                                                                                                       |
| ├─ `host`            | String            | The server hostname. Default: `localhost`.                                                                                                                                                               |
| ├─ `port`            | Number            | The server port. Default: `7233`.                                                                                                                                                                        |
| ├─ `namespace`       | String            | The namespace to connect to. Default: `default`.                                                                                                                                                         |
| ├─ `mtls`            | Mtls              | The mTLS configuration if mTLS is desired. mTLS will not be used if not present.                                                                                                                         |
| │   ├─ `cert_path`   | String            | Path to the client certificate file.                                                                                                                                                                     |
| │   ├─ `key_path`    | String            | Path to the client key file.                                                                                                                                                                             |
| `workers`            | Array of Worker   | A list of workers to run, each defined by a name, task queue, workflows, and activities.                                                                                                                 |
| ├─ `name`            | String            | The unique name of the worker.                                                                                                                                                                           |
| ├─ `task_queue`      | String            | The task queue to which the worker is polling for tasks on (should be unique per worker).                                                                                                                |
| ├─ `workflows`       | Array of Workflow | A list of workflows handled by the worker, defined by type, and a result or error.                                                                                                                       |
| │   ├─ `type`        | String            | The unique identifier of the workflow type (name).                                                                                                                                                       |
| │   ├─ `result`      | Any               | The expected result of the workflow, which may be a string, object, or number. Setting this denotes a successful Workflow stub. Supports [result templating](#result-templating) in field values.        |
| │   ├─ `error`       | Error             | The expected error thrown by the workflow. Setting this denotes an error Workflow stub.                                                                                                                  |
| │   │   ├─ `type`    | String            | The error type to be thrown. Example for simulating a java file not found error use: `java.io.FileNotFoundException`.                                                                                    |
| │   │   ├─ `message` | String            | A message to be included in the error. Supports [result templating](#result-templating) syntax.                                                                                                          |
| │   │   ├─ `details` | Any               | Any additional details to be included with the error (any valid yaml accepted). Supports [result templating](#result-templating) syntax in field values.                                                 |
| ├─ `activities`      | Array of Activity | A list of activities handled by the worker, defined by type, and result or error.                                                                                                                        |
| │   ├─ `type`        | String            | The unique identifier of the activity type (name).                                                                                                                                                       |
| │   ├─ `result`      | Any               | The expected result of the activity, which may be a string, object, or number. Setting this denotes a successful Activity stub. Supports [result templating](#result-templating) syntax in field values. |
| │   ├─ `error`       | Error             | The expected error thrown by the activity. Setting this denotes an error Activity stub.                                                                                                                  |
| │   │   ├─ `type`    | String            | The error type to be thrown. Example for simulating a java file not found error use: `java.io.FileNotFoundException`.                                                                                    |
| │   │   ├─ `message` | String            | A message to be included in the error. Supports [result templating](#result-templating) syntax.                                                                                                          |
| │   │   ├─ `details` | Any               | Any additional details to be included with the error (any valid yaml accepted). Supports [result templating](#result-templating) syntax in field values.                                                 |

Example configuration files can be found in the  ([config directory](config)).

#### Result Templating

Temporal Phantom Worker supports dynamic result generation for workflows and activities using Go's powerful text/template package. This allows you to create results that adapt based on input parameters or include randomized values, making your tests more flexible and robust.

Temporal Phantom Worker also supports the following additional functions to generate data:
1. `randomString(length)`
Example: `{{ randomString 10 }}`
2. `currentTime()`
Example: `{{ currentTime }}`
3. `randomUUID()`
Example: `{{ randomUUID }}`
4. `randomInt(min, max)`
Example: `{{ randomInt 5 10 }}`

For full templating  examples, refer to the [templating-samples.yaml](config/templating-sample.yaml) sample config file.

### Activity

#### Trigger

Use the activity trigger command to execute an activity in isolation by wrapping it in a dynamic workflow:

```bash
./temporal-phantom-worker activity trigger -type MyTestActivity -taskqueue testQueue --input <activtyInput.yaml>
```

The input file is an optional yaml file containing the input to pass to the activity. See the command's help for full
list of options.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or
encounter any problems.

